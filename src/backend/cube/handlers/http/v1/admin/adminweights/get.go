package adminweights

import (
	"net/http"

	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
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
// @Router       /api/v1/admin/models/weights [get]
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
	lg.WithFields(map[string]interface{}{"req": req}).Info("req binded")

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

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to find weights info")
	infos, err := h.resolver.FindStructureWeights(c, filter)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to find weights info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		statOKGet.Inc()
		lg.Info("no weights found")
		c.JSON(http.StatusOK, []weights.Info{})
		return
	}
	var res []weights.Info
	for _, val := range infos {
		res = append(res, weightFromBL(*val))
	}

	statOKGet.Inc()
	lg.WithFields(map[string]interface{}{"res": res}).Info("success")
	c.JSON(http.StatusOK, res)
}
