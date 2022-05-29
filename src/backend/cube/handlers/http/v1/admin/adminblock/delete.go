package adminblock

import (
	"net/http"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	UserId string `form:"id"`
}

// Registration  godoc
// @Summary      Delete user block info
// @Description  Deletes user block info by user id
// @Tags         admin
// @Param        id query string false "User ID to unblock"
// @Success      200 "User unblocked"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to delete user block info from storage"
// @Router       /api/v1/admin/users/blocked [delete]
func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req deleteRequest
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to get user block info")
	err := h.resolver.Block(c, req.UserId, time.Now().Add(5*time.Minute))
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to unblock user %v", err)
		c.JSON(http.StatusInternalServerError, "failed to unblock user")
		return
	}

	statOKDelete.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
