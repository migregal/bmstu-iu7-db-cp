package adminusers

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	UserId   string `form:"user_id"`
	Username string `form:"username"`
	Email    string `form:"email"`
	Page     int    `form:"page"`
	PerPage  int    `form:"per_page"`
}

type UserInfo struct {
	Id       string `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Username string `json:"username,omitempty" example:"awesome_username"`
	Email    string `json:"email,omitempty" example:"my_awesome@email.com"`
	Fullname string `json:"fullname,omitempty" example:"Ivanov Ivan Ivanovich"`
} // @name AdminUserInfoResponse

// Registration  godoc
// @Summary      Find user info
// @Description  Find such users info as id, username, email and fullname
// @Tags         admin
// @Accept       json
// @Param        user_id query string false "UserId to search for"
// @Param        username query string false "Username to search for"
// @Param        email query string false "Email to search for"
// @Param        page query int false "Email to search for"
// @Param        per_page query int false "Email to search for"
// @Success      200 {object} []UserInfo "Users info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get user info from storage"
// @Router       /api/v1/admin/users [get]
func (h *Handler) Get(c *gin.Context) {
	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	filter := interactors.UserInfoFilter{}
	if req.Username != "" {
		filter.Usernames = []string{req.Username}
	}
	if req.Email != "" {
		filter.Emails = []string{req.Email}
	}
	if req.UserId != "" {
		filter.Ids = []string{req.UserId}
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	infos, err := h.resolver.Find(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		c.JSON(http.StatusOK, []UserInfo{})
		return
	}
	var res []UserInfo
	for _, val := range infos {
		res = append(res, UserInfo{
			Id:       *val.ID(),
			Email:    *val.Email(),
			Username: *val.Username(),
			Fullname: *val.Fullname(),
		})
	}
	c.JSON(http.StatusOK, res)
}
