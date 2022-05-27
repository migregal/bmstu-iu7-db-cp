package auth

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"neural_storage/cube/core/ports/interactors"
	. "neural_storage/cube/handlers/http/jwt"
)

type request struct {
	Email    string `form:"email" json:"email" binding:"required" example:"my_awesome@email.com"`
	Password string `form:"password" json:"password" binding:"required" example:"Really, you're waiting for example?"`
} // @name LoginRequest

// Registration  godoc
// @Summary      User login
// @Description  Login to existing account
// @Tags         auth,user
// @Accept       json
// @Param        Body body request true "The body to create a thing"
// @Success      200 {object} LoginResponse "Login was successfull"
// @Failure      401 {object} Unauthorized "Login data is invalid or missing, check request"
// @Router       /api/v1/login [post]
func (m *Handler) Authenticator(role uint64) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var req request
		if err := c.ShouldBind(&req); err != nil {
			return "", ErrMissingCreds
		}

		infos, err := m.resolver.Find(interactors.UserInfoFilter{
			Emails: []string{req.Email},
			Limit:  1,
		})
		if err != nil {
			return nil, ErrFailedAuth
		}
		if len(infos) == 0 {
			return nil, ErrFailedAuth
		}

		creds := infos[0]
		if !creds.BlockedUntil().IsZero() && creds.BlockedUntil().After(time.Now()) {
			return nil, fmt.Errorf("blocked")
		}

		if creds.Flags()&role != role {
			return nil, fmt.Errorf("not permitted")
		}

		if *creds.Pwd() != hex.EncodeToString(getPasswordHash(req.Password)) {
			return nil, ErrFailedAuth
		}

		return &User{
			ID:       *creds.ID(),
			Email:    *creds.Email(),
			Username: *creds.Username(),
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
