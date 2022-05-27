package adminblock

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	UserID string `form:"id" binding:"required"`
}

type BlockInfo struct {
	ID    string     `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Until time.Time `json:"blocked_until,omitempty" example:"2025-08-09T15:00:00.053Z"`
} // @name AdminBlockUserResponse

// Registration  godoc
// @Summary      Find user block info
// @Description  Find such users info as id and block time
// @Tags         admin
// @Produce      json
// @Param        user_id query string false "UserId to search for"
// @Success      200 {object} BlockInfo "Users block info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get user info from storage"
// @Router       /api/v1/admin/users/blocked [get]
func (h *Handler) Get(c *gin.Context) {
	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	info, err := h.resolver.Get(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}

	c.JSON(http.StatusOK, BlockInfo{ID: req.UserID, Until: info.BlockedUntil()})
}
