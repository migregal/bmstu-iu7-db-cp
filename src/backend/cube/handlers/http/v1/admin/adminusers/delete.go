package adminusers

import (
	"net/http"

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
	var req deleteRequest
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.resolver.Delete(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to delete user info")
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
