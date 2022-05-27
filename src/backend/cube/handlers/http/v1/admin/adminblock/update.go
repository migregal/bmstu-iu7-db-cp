package adminblock

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	UserId string     `json:"id" form:"id" binding:"required"`
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
	var req updateRequest
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.resolver.Block(req.UserId, req.Until)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to block user")
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
