package auth

import (
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

func (m *Handler) Authorizator(role uint64) func(c *gin.Context, data interface{}) bool {
	return func(c *gin.Context, data interface{}) bool {
		usr, ok := data.(*User)

		c.Set(jwt.IdentityKey, usr.ID)
		return ok && (usr.Flags&role == role)
	}
}
