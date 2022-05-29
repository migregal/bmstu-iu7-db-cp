package registration

import (
	"net/http"
	"net/url"
	"neural_storage/cube/core/entities/user"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/roles"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	statCall stat.Counter
	statFail stat.Counter
	statOK   stat.Counter
)

func init() {
	statCall = stat.NewCounter("v1", "cube_registration_call", "The total number of registration attempts")
	statFail = stat.NewCounter("v1", "cube_registration_fail", "The total number of registration fails")
	statOK = stat.NewCounter("v1", "cube_registration_ok", "The total number of login attempts")
}

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}

type request struct {
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
	statCall.Inc()

	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req request
	if err := c.ShouldBind(&req); err != nil {
		statFail.Inc()
		lg.Errorf("failed to bind req: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lg.WithFields(map[string]interface{}{"req": req}).Info("req binded")

	filter := interactors.UserInfoFilter{
		Usernames: []string{req.Username},
		Emails:    []string{req.Email},
		Limit:     1,
	}
	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to find user info")
	infos, err := h.resolver.Find(c, filter)
	if err != nil {
		statFail.Inc()
		lg.Errorf("failed to fetch user info: %v", err)
		c.JSON(http.StatusBadRequest,
			jwt.Unauthorized{Message: "reginfo validation failed: " + err.Error()})
		return
	}
	if len(infos) > 0 {
		statFail.Inc()
		lg.Info("no users found")
		c.JSON(http.StatusForbidden,
			jwt.Unauthorized{Message: "email already used"})
		return
	}

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to register user")
	_, err = h.resolver.Register(
		c,
		*user.NewInfo(
			"",
			req.Username,
			req.Fullname,
			req.Email,
			req.Password,
			roles.RoleUser,
			time.Time{},
		),
	)
	if err != nil {
		statFail.Inc()
		lg.Errorf("failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError,
			jwt.Unauthorized{Message: "user registration failed: " + err.Error()})
		return
	}

	statOK.Inc()
	lg.Info("success, redirecting to login")
	location := url.URL{Path: "/api/v1/login"}
	c.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
}
