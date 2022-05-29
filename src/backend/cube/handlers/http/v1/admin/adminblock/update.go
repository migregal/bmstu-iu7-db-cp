package adminblock

import (
	"net/http"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	UserId string    `json:"id" form:"id" binding:"required"`
	Until  time.Time `json:"until" form:"until" binding:"required"`
}

// Registration  godoc
// @Summary      Block user
// @Description  Blocks user until specified moment
// @Tags         stat
// @Param        id    query string false "User ID to block"
// @Param        until query string false "Time to block until"
// @Success      200 "User blocked"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to block user "
// @Router       /api/v1/stat/users/blocked [patch]
func (h *Handler) Update(c *gin.Context) {
	statCallUpdate.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req updateRequest
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to block user")
	err := h.resolver.Block(c, req.UserId, req.Until)
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to block user: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to block user")
		return
	}

	statOKUpdate.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
