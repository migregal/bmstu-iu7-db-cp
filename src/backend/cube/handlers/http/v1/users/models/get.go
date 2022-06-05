package models

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/v1/entities/model"
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

// Registration  godoc
// @Summary      Find model info
// @Description  Find such model info as id, username, email and fullname
// @Tags         user
// @Param        id       query string false "Model ID to search for"
// @Param        owner_id query string false "User ID that owns model to search for"
// @Param        name     query string false "Model name to search for"
// @Param        page     query int    false "Page number for pagination"
// @Param        per_page query int    false "Page size for pagination"
// @Success      200 {object} []model.Info "Model info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get model info from storage"
// @Router       /api/v1/models [get]
func (h *Handler) Get(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to get model info into cache")
	if info, err := h.cache.Get(modelStorage, req.ModelID); err == nil && len(info) >= 2 {
		lg.Info("success to get model info from cache")
		resp, err := unGzip(info[1].([]byte))
		if err == nil {
			statOKGet.Inc()
			c.Data(http.StatusOK, "application/json", resp)
			return
		}
		lg.Errorf("failed to ungzip model info from cache: %v", err)
	} else {
		lg.Errorf("failed to get model info from cache: %v", err)
	}

	filter := interactors.ModelInfoFilter{}
	if len(req.ModelID) > 0 {
		filter.Ids = []string{req.ModelID}
	}
	if req.OwnerID != "" {
		filter.Owners = append(filter.Owners, req.OwnerID)
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
		lg.Errorf("failed to find model info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		statOKGet.Inc()
		lg.Info("no models found")
		c.JSON(http.StatusOK, []model.Info{})
		return
	}

	var res []model.Info
	for _, val := range infos {
		res = append(res, modelFromBL(val))
	}

	if len(req.ModelID) == 1 && req.ModelID != "" {
		if data, err := jsonGzip(res); err == nil {
			_ = h.cache.Update(modelStorage, req.ModelID, data)
		}
	}
	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, res)

}
