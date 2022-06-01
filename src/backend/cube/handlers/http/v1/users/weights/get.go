package weights

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	ID          string `form:"id"`
	StructureID string `form:"structure_id"`
	Name        string `form:"name"`
	Page        int    `form:"page"`
	PerPage     int    `form:"per_page"`
}

type WeightInfo struct {
	Id      string      `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Name    string      `json:"name,omitempty" example:"awesome_username"`
	Weights interface{} `json:"offsets,omitempty"`
	Offsets interface{} `json:"weights,omitempty"`
} // @name ModelWeightsInfoResponse

// Registration  godoc
// @Summary      Find model info
// @Description  Find such model info as id, username, email and fullname
// @Tags         admin
// @Param        id           query string false "Weight ID to search for"
// @Param        structure_id query string false "Structure ID to search for"
// @Param        name         query string false "Weights name to search for"
// @Param        page         query int    false "Page number for pagination"
// @Param        per_page     query int    false "Page size for pagination"
// @Success      200 {object} []WeightInfo "Model weights info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get model weights info from storage"
// @Router       /api/v1/models/weights [get]
func (h *Handler) Get(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to get model info into cache")
	if info, err := h.cache.Get(weightStorage, req.ID); err == nil && len(info) == 2 {
		lg.Info("success to get weight info from cache")
		resp, err := unGzip(info[1].([]byte))
		if err == nil {
			statOKGet.Inc()
			c.Data(http.StatusOK, "application/json", resp)
			return
		}
		lg.Errorf("failed to ungzip weight info from cache: %v", err)
	} else {
		lg.Errorf("failed to get weight info from cache: %v", err)
	}

	filter := interactors.ModelWeightsInfoFilter{}
	if req.ID != "" {
		filter.IDs = append(filter.IDs, req.ID)
	}
	if req.StructureID != "" {
		filter.Structures = []string{req.StructureID}
	}
	if req.Name != "" {
		filter.Names = []string{req.Name}
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to get weights")
	infos, err := h.resolver.FindStructureWeights(c, filter)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to find model info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		statOKGet.Inc()
		lg.Info("no weights found")
		c.JSON(http.StatusOK, []WeightInfo{})
		return
	}
	if len(infos) == 1 {
		weight := weightFromBL(*infos[0])
		if data, err := jsonGzip(weight); err == nil {
			_ = h.cache.Update(weightStorage, req.ID, data)
		}

		statOKGet.Inc()
		lg.Info("successful get full weight info")
		c.JSON(http.StatusOK, weight)
		return
	}
	var res []WeightInfo
	for _, val := range infos {
		w := weightFromBL(*val)
		res = append(res, WeightInfo{
			Id:      w.ID,
			Name:    w.Name,
			Weights: w.Weights,
			Offsets: w.Offsets,
		})
	}
	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, res)
}
