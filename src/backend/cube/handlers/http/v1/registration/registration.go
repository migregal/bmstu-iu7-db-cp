package registration

import (
	"net/http"
	"net/url"
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/roles"
	"neural_storage/cube/handlers/http/jwt"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	resolver interactors.UserInfoInteractor
}

func New(resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver}
}

type Request struct {
	Username string `binding:"required" example:"my_awesome_nickname"`
	Email    string `binding:"required" example:"my_awesome@email.com"`
	Fullname string `binding:"required" example:"John Smith"`
	Password string `binding:"required" example:"Really, you're waiting for example?"`
} // @name RegistrationRequest

// Registration  godoc
// @Summary      User registration
// @Description  register new user
// @Tags         auth,user
// @Accept       json
// @Param        Body body Request true "The body to create a thing"
// @Success      307 "Registration was successfull, redirect request to login (/api/v1/login)"
// @Failure      400 {object} jwt.Unauthorized "Registration data is invalid or missing, check request"
// @Failure      500 {object} jwt.Unauthorized "Failed to register user due to some reasons. For example: user already exists"
// @Router       /api/v1/registration [post]
func (h *Handler) Registration(c *gin.Context) {
	var req Request

	if c.ShouldBind(&req) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	infos, err := h.resolver.Find(interactors.UserInfoFilter{
		Usernames: []string{req.Username},
		Emails:    []string{req.Email},
		Limit:     1,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest,
			jwt.Unauthorized{Message: "reginfo validation failed: " + err.Error()})
		return
	}
	if len(infos) > 0 {
		c.JSON(http.StatusForbidden,
			jwt.Unauthorized{Message: "email already used"})
		return
	}

	_, err = h.resolver.Register(
		*user.NewInfo(
			nil,
			&req.Username,
			&req.Fullname,
			&req.Email,
			&req.Password,
			roles.RoleUser,
			time.Time{},
		),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			jwt.Unauthorized{Message: "user registration failed: " + err.Error()})
		return
	}

	location := url.URL{Path: "/api/v1/login"}
	c.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
}
