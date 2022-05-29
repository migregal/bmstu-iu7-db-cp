package adminusers

import (
	"net/http"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	UserId string `form:"id"`
}

// Registration  godoc
// @Summary      Delete user info
// @Description  Deletes user by id
// @Tags         admin
// @Accept       json
// @Param        id query string false "UserId to delete"
// @Success      200 "User deleted"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to delete user info from storage"
// @Router       /api/v1/admin/users [delete]
func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})
	var req deleteRequest
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to delete user info")
	err := h.resolver.Delete(c, req.UserId)
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to delete model info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to delete user info")
		return
	}

	statOKDelete.Inc()
	lg.Info("status")
	c.AbortWithStatus(http.StatusOK)
}
