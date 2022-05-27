package auth

import (
	"github.com/gin-gonic/gin"
)

func (m *Handler) Authorizator(role uint64) func(c *gin.Context, data interface{}) bool {
	return func(_ *gin.Context, data interface{}) bool {
		usr, ok := data.(*User)
		return ok && (usr.Flags&role == role)
	}
}
