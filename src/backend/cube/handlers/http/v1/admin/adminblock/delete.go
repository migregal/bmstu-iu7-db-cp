package adminblock

import (
	"net/http"
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
	var req deleteRequest
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.resolver.Block(req.UserId, time.Now().Add(5*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to unblock user")
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
