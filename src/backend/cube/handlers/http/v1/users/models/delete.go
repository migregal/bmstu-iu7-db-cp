package models

import (
	"net/http"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	ID string `form:"id" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
}

// Registration  godoc
// @Summary      Delete model info
// @Description  Deletes model info owned by authorized user
// @Tags         user
// @Param        id query string false "Model ID to delete"
// @Success      200 "Model info deleted"
// @Failure      400 "Invalid request"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to delete model info from storage"
// @Router       /api/v1/models [delete]
func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		statFailDelete.Inc()
		lg.Error("access token missing")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok {
		statFailDelete.Inc()
		lg.Error("invalid access token")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req DeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	defer func() {
		lg.Info("attempt to delete model from cache")
		_ = h.cache.Delete(modelStorage, req.ID)
	}()

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to delete model")
	err := h.resolver.Delete(c, usrID, req.ID)
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to delete model: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to delete model info")
		return
	}

	statOKDelete.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
