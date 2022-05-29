package users

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"

	"github.com/gin-gonic/gin"
)

var (
	statCall stat.Counter
	statFail stat.Counter
	statOK   stat.Counter
)

func init() {
	statCall = stat.NewCounter("v1", "cube_users_call", "The total number of getting user info attempts")
	statFail = stat.NewCounter("v1", "cube_users_fail", "The total number of getting user info fails")
	statOK = stat.NewCounter("v1", "cube_users_ok", "The total number of login attempts")
}

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{lg: lg, resolver: resolver}
}

type request struct {
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
} // @name UserInfoResponse

// Registration  godoc
// @Summary      Find user info
// @Description  Find such users info as id, username, email and fullname
// @Tags         user
// @Accept       json
// @Param        user_id  query string false "UserId to search for"
// @Param        username query string false "Username to search for"
// @Param        email    query string false "Email to search for"
// @Param        page     query int false "Page number for pagination"
// @Param        per_page query int false "Page size for pagination"
// @Success      200 {object} []UserInfo "Users info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get user info from storage"
// @Router       /api/v1/users [get]
func (h *Handler) Get(c *gin.Context) {
	statCall.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req request
	if err := c.ShouldBind(&req); err != nil {
		statFail.Inc()
		lg.Error("failed to bind request")
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

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to find user info")
	infos, err := h.resolver.Find(c, filter)
	if err != nil {
		statFail.Inc()
		lg.Error("failed to fetch user info")
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		statOK.Inc()
		lg.Info("no users found")
		c.JSON(http.StatusOK, []UserInfo{})
		return
	}
	var res []UserInfo
	for _, val := range infos {
		res = append(res, UserInfo{
			Id:       val.ID(),
			Email:    val.Email(),
			Username: val.Username(),
			Fullname: val.Fullname(),
		})
	}

	statOK.Inc()
	lg.WithFields(map[string]interface{}{"res": res}).Info("success")
	c.JSON(http.StatusOK, res)
}
