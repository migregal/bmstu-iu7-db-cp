package adminmodels

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	OwnerID   string `form:"ownerid"`
	ModelID   string `form:"id"`
	ModelName string `form:"name"`
	Page      int    `form:"page"`
	PerPage   int    `form:"per_page"`
}

type ModelInfo struct {
	Id        string      `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Name      string      `json:"name,omitempty" example:"awesome_username"`
	Structure interface{} `json:"structure,omitempty"`
} // @name AdminModelInfoResponse

// Registration  godoc
// @Summary      Find model info
// @Description  Find such model info as id, username, email and fullname
// @Tags         admin
// @Param        id       query string false "Model ID to search for"
// @Param        owner_id query string false "User ID that owns model to search for"
// @Param        name     query string false "Model name to search for"
// @Param        page     query int    false "Page number for pagination"
// @Param        per_page query int    false "Page size for pagination"
// @Success      200 {object} []ModelInfo "Model info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get model info from storage"
// @Router       /api/v1/admin/models [get]
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

	filter := interactors.ModelInfoFilter{}
	if req.OwnerID != "" {
		filter.Owners = append(filter.Owners, req.OwnerID)
	}
	if req.ModelID != "" {
		filter.Ids = []string{req.ModelID}
	}
	if req.ModelName != "" {
		filter.Names = []string{req.ModelName}
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to find model info")
	infos, err := h.resolver.Find(c, filter)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to fetch model info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to fetch model info")
		return
	}
	if len(infos) == 0 {
		statOKGet.Inc()
		lg.Info("no models found")
		c.JSON(http.StatusOK, []ModelInfo{})
		return
	}
	var res []ModelInfo
	for _, val := range infos {
		res = append(res, ModelInfo{
			Id:        val.ID(),
			Name:      val.Name(),
			Structure: val.Structure(),
		})
	}
	statOKGet.Inc()
	lg.WithFields(map[string]interface{}{"res": res}).Info("success")
	c.JSON(http.StatusOK, res)
}
