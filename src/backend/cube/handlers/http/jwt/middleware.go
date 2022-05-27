package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

const IdentityKey = "user_id"

var ErrMissingCreds = jwt.ErrMissingLoginValues
var ErrFailedAuth = jwt.ErrFailedAuthentication

type LoginResponse struct {
	Token  string `json:"token" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc..."`
	Expire string `json:"expire" example:"2022-03-20T17:00:01Z"`
} // @name LoginResponse

type Unauthorized struct {
	Message string `json:"message" example:"user not found"`
} // @name Unauthorized

func NewJWTMiddleware(
	authenticator func(c *gin.Context) (interface{}, error),
	authorizator func(c *gin.Context, data interface{}) bool,
	payload func(data interface{}) jwt.MapClaims,
	identityHandler func(c *gin.Context) interface{},
	privKeyFile string,
	pubKeyFile string,
) *jwt.GinJWTMiddleware {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "Archmage",
		SigningAlgorithm: "RS256",
		Key:              x509.MarshalPKCS1PrivateKey(key),
		Timeout:          time.Hour,
		MaxRefresh:       24 * time.Hour,
		IdentityKey:      IdentityKey,
		PayloadFunc:      payload,
		IdentityHandler:  identityHandler,
		Authenticator:    authenticator,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return authorizator(c, data)
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if code != http.StatusOK {
				c.AbortWithStatus(code)
			}
			c.JSON(http.StatusOK, LoginResponse{token, expire.Format(time.RFC3339)})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if code != http.StatusOK {
				c.AbortWithStatus(code)
			}
			c.JSON(http.StatusOK, LoginResponse{token, expire.Format(time.RFC3339)})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, Unauthorized{message})
		},
		PrivKeyFile: privKeyFile,
		PubKeyFile:  pubKeyFile,
		TimeFunc:    time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	return jwt.ExtractClaims(c)
}
