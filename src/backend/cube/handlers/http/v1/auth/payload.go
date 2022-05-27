package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
)

func (m *Handler) Payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{
			UserIdIdentityKey:    v.ID,
			UserFlagsIdentityKey: v.Flags,
		}
	}

	return jwt.MapClaims{}
}
