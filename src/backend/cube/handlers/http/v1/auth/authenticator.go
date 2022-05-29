package auth

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"neural_storage/cube/core/ports/interactors"
	. "neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
)

type request struct {
	Email    string `form:"email" json:"email" binding:"required" example:"my_awesome@email.com"`
	Password string `form:"password" json:"password" binding:"required" example:"Really, you're waiting for example?"`
} // @name LoginRequest

var (
	statCall         stat.Counter
	statFail         stat.Counter
	statBlocked      stat.Counter
	statNotPermitted stat.Counter
	statOK           stat.Counter
)

func init() {
	statCall = stat.NewCounter("v1", "cube_login_call", "The total number of login attempts")
	statFail = stat.NewCounter("v1", "cube_login_fail", "The total number of login fails")
	statBlocked = stat.NewCounter("v1", "cube_login_blocked", "The total number of login attempts for blocked users")
	statNotPermitted = stat.NewCounter("v1", "cube_login_not_permitted", "The total number of login attempts for not permitted accounts")
	statOK = stat.NewCounter("v1", "cube_login_ok", "The total number of successful login attempts")
}

// Registration  godoc
// @Summary      User login
// @Description  Login to existing account
// @Tags         auth,user
// @Accept       json
// @Param        Body body request true "The body to create a thing"
// @Success      200 {object} LoginResponse "Login was successfull"
// @Failure      401 {object} Unauthorized "Login data is invalid or missing, check request"
// @Router       /api/v1/login [post]
func (h *Handler) Authenticator(role uint64) func(c *gin.Context) (interface{}, error) {

	return func(c *gin.Context) (interface{}, error) {
		statCall.Inc()

		lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

		var req request
		if err := c.ShouldBind(&req); err != nil {
			statFail.Inc()
			lg.Errorf("failed to bind request: %v", err)
			return "", ErrMissingCreds
		}
		lg.WithFields(map[string]interface{}{"req": req}).Info("req binded")

		filter := interactors.UserInfoFilter{Emails: []string{req.Email}, Limit: 1}
		lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to find weights info")
		infos, err := h.resolver.Find(c, filter)
		if err != nil {
			statFail.Inc()
			lg.Errorf("failed to find user info: %v", err)
			return nil, ErrFailedAuth
		}
		if len(infos) == 0 {
			statFail.Inc()
			lg.Info("no users found")
			return nil, ErrFailedAuth
		}

		creds := infos[0]
		if !creds.BlockedUntil().IsZero() && creds.BlockedUntil().After(time.Now()) {
			statBlocked.Inc()
			lg.Info("user blocked")
			return nil, fmt.Errorf("blocked")
		}

		if creds.Flags()&role != role {
			statNotPermitted.Inc()
			lg.Info("not permitted")
			return nil, fmt.Errorf("not permitted")
		}

		if creds.Pwd() != hex.EncodeToString(getPasswordHash(req.Password)) {
			statFail.Inc()
			lg.Info("invalid creds")
			return nil, ErrFailedAuth
		}

		statOK.Inc()
		lg.Info("success")
		return &User{
			ID:       creds.ID(),
			Email:    creds.Email(),
			Username: creds.Username(),
			Flags:    creds.Flags(),
		}, nil
	}
}

// Registration  godoc
// @Summary      Admin login
// @Description  login to existing account
// @Tags         auth,admin
// @Accept       json
// @Param        Body body request true "The body to create a thing"
// @Success      200 {object} LoginResponse "Login was successfull"
// @Failure      401 {object} Unauthorized "Login data is invalid or missing, check request"
// @Router       /api/v1/admin/login [post]
func _() {}

// Registration  godoc
// @Summary      Stat login
// @Description  login to existing account
// @Tags         auth,stat
// @Accept       json
// @Param        Body body request true "The body to create a thing"
// @Success      200 {object} LoginResponse "Login was successfull"
// @Failure      401 {object} Unauthorized "Login data is invalid or missing, check request"
// @Router       /api/v1/stat/login [post]
func _() {}
