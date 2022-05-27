package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func (m *Handler) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	id, ok := claims[UserIdIdentityKey].(string)
	if !ok {
		return nil
	}
	flags, ok := claims[UserFlagsIdentityKey].(float64)
	if !ok {
		return nil
	}

	return &User{ID: id, Flags: uint64(flags)}
}
