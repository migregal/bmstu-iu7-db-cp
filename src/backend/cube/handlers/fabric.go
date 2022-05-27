package handlers

import (
	"log"
	"net/http"

	"neural_storage/cache/adapters/hotstorage"
	"neural_storage/config/core/services/config"
	"neural_storage/cube/core/roles"
	"neural_storage/cube/core/services/interactor/model"
	"neural_storage/cube/core/services/interactor/user"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/cube/handlers/http/v1/admin/adminblock"
	"neural_storage/cube/handlers/http/v1/admin/adminmodels"
	"neural_storage/cube/handlers/http/v1/admin/adminusers"
	"neural_storage/cube/handlers/http/v1/admin/adminweights"
	"neural_storage/cube/handlers/http/v1/auth"
	"neural_storage/cube/handlers/http/v1/registration"
	"neural_storage/cube/handlers/http/v1/stat/statmodels"
	"neural_storage/cube/handlers/http/v1/stat/statusers"
	"neural_storage/cube/handlers/http/v1/stat/statweights"
	"neural_storage/cube/handlers/http/v1/users"
	"neural_storage/cube/handlers/http/v1/users/models"
	"neural_storage/cube/handlers/http/v1/users/weights"

	_ "neural_storage/cube/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           Cube API
// @version         1.0
// @description     This is cube server.

// @license.name  MIT
// @license.url   https://mit-license.org/

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  Bearer
type Server interface {
	Run(addr ...string) (err error)
}

func New(params config.Config) Server {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	initRoutes(params, engine)

	initAdminRoutes(params, engine)

	initStatRoutes(params, engine)

	initFailure(params, engine)

	return engine
}

func initRoutes(params config.Config, engine *gin.Engine) {
	authManager := auth.NewHandler(user.NewInteractor(params.UserInfo()))
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator(roles.RoleUser),
		authManager.Authorizator(roles.RoleUser),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 := engine.Group("/api/v1")
	{
		regManager := registration.New(user.NewInteractor(params.UserInfo()))
		v1.POST("/registration", regManager.Registration)
		v1.POST("/login", authMiddleware.LoginHandler)
	}

	v1Authorized := engine.Group("/api/v1").Use(authMiddleware.MiddlewareFunc())
	{
		v1Authorized.GET("/refresh_token", authMiddleware.RefreshHandler)

		usrManager := users.New(user.NewInteractor(params.UserInfo()))
		v1Authorized.GET("/users", usrManager.Get)

		modelManager := models.New(
			model.NewInteractor(params.ModelInfo()),
			hotstorage.New(params.Cache()),
		)
		v1Authorized.POST("/models", modelManager.Add)
		v1Authorized.GET("/models", modelManager.Get)
		v1Authorized.PATCH("/models", modelManager.Update)
		v1Authorized.DELETE("/models", modelManager.Delete)

		weightsManager := weights.New(model.NewInteractor(params.ModelInfo()))
		v1Authorized.POST("/models/weights", weightsManager.Add)
		v1Authorized.GET("/models/weights", weightsManager.Get)
		v1Authorized.DELETE("/models/weights", weightsManager.Delete)
		v1Authorized.PATCH("/models/weights", weightsManager.Update)
	}
}

func initAdminRoutes(params config.Config, engine *gin.Engine) {
	authManager := auth.NewHandler(user.NewInteractor(params.UserInfo()))
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator(roles.RoleAdmin),
		authManager.Authorizator(roles.RoleAdmin),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 := engine.Group("/api/v1/admin")
	{
		if params.Debug() {
			v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		v1.POST("/login", authMiddleware.LoginHandler)
	}

	v1Authorized := engine.Group("/api/v1/admin").Use(authMiddleware.MiddlewareFunc())
	{
		v1Authorized.GET("/refresh_token", authMiddleware.RefreshHandler)

		usrManager := adminusers.New(user.NewInteractor(params.AdminUserInfo()))
		v1Authorized.GET("/users", usrManager.Get)
		v1Authorized.DELETE("/users", usrManager.Delete)

		usrBlockManager := adminblock.New(user.NewInteractor(params.AdminUserInfo()))
		v1Authorized.GET("/users/blocked", usrBlockManager.Get)
		v1Authorized.DELETE("/users/blocked", usrBlockManager.Delete)
		v1Authorized.PATCH("/users/blocked", usrBlockManager.Update)

		modelManager := adminmodels.New(model.NewInteractor(params.AdminModelInfo()))
		v1Authorized.GET("/models", modelManager.Get)
		v1Authorized.DELETE("/models", modelManager.Delete)

		weightsManager := adminweights.New(model.NewInteractor(params.AdminModelInfo()))
		v1Authorized.GET("/models/weights", weightsManager.Get)
		v1Authorized.DELETE("/models/weights", weightsManager.Delete)
	}
}

func initStatRoutes(params config.Config, engine *gin.Engine) {
	authManager := auth.NewHandler(user.NewInteractor(params.UserInfo()))
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator(roles.RoleStat),
		authManager.Authorizator(roles.RoleStat),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 := engine.Group("/api/v1/stat")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
	}

	v1Authorized := engine.Group("/api/v1/stat").Use(authMiddleware.MiddlewareFunc())
	{
		v1Authorized.GET("/refresh_token", authMiddleware.RefreshHandler)

		userManager := statusers.New(user.NewInteractor(params.StatUserInfo()))
		v1Authorized.GET("/users", userManager.Get)

		modelManager := statmodels.New(model.NewInteractor(params.StatModelInfo()))
		v1Authorized.GET("/models", modelManager.Get)

		weightsManager := statweights.New(model.NewInteractor(params.StatModelInfo()))
		v1Authorized.GET("/weights", weightsManager.Get)
	}
}

func initFailure(params config.Config, engine *gin.Engine) {
	authManager := auth.NewHandler(user.NewInteractor(params.UserInfo()))
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator(roles.RoleUser),
		authManager.Authorizator(roles.RoleUser),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	engine.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
