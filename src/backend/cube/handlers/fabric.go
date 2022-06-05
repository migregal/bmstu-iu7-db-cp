package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"

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
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat/statmiddleware"

	_ "neural_storage/cube/docs"
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

func New(params config.Config, lg *logger.Logger) Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Use(logger.RequestIDSetter())
	engine.Use(logger.RequestLogger(lg))
	engine.Use(statmiddleware.MeasureResponseDuration())

	initRoutes(params, lg, engine)

	initAdminRoutes(params, lg, engine)

	initStatRoutes(params, lg, engine)

	initFailure(params, lg, engine)

	return engine
}

func initRoutes(params config.Config, lg *logger.Logger, engine *gin.Engine) {
	engine.GET("/prometheus", gin.WrapH(promhttp.Handler()))

	authManager := auth.NewHandler(lg, user.NewInteractor(lg, params.UserInfo()))
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator(roles.RoleUser),
		authManager.Authorizator(roles.RoleUser),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	login := func(c *gin.Context) {
		authMiddleware.LoginHandler(c)
	}

	v1 := engine.Group("/api/v1")
	{
		regManager := registration.New(lg, user.NewInteractor(lg, params.UserInfo()))
		v1.POST("/registration", regManager.Registration)
		v1.POST("/login", login)
	}

	v1Authorized := engine.
		Group("/api/v1").
		Use(authMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		v1Authorized.GET("/refresh", authMiddleware.RefreshHandler)
		v1Authorized.GET("/logout", authMiddleware.LogoutHandler)

		usrManager := users.New(lg, user.NewInteractor(lg, params.UserInfo()))
		v1Authorized.GET("/users", usrManager.Get)

		modelManager := models.New(
			lg,
			model.NewInteractor(lg, params.ModelInfo()),
			hotstorage.New(params.Cache()),
		)
		v1Authorized.POST("/models", modelManager.Add)
		v1Authorized.GET("/models", modelManager.Get)
		v1Authorized.PATCH("/models", modelManager.Update)
		v1Authorized.DELETE("/models", modelManager.Delete)

		weightsManager := weights.New(
			lg,
			model.NewInteractor(lg, params.ModelInfo()),
			hotstorage.New(params.Cache()))
		v1Authorized.POST("/models/weights", weightsManager.Add)
		v1Authorized.GET("/models/weights", weightsManager.Get)
		v1Authorized.DELETE("/models/weights", weightsManager.Delete)
		v1Authorized.PATCH("/models/weights", weightsManager.Update)
	}
}

func initAdminRoutes(params config.Config, lg *logger.Logger, engine *gin.Engine) {
	authManager := auth.NewHandler(lg, user.NewInteractor(lg, params.UserInfo()))
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

	v1Authorized := engine.
		Group("/api/v1/admin").
		Use(authMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		v1Authorized.GET("/refresh", authMiddleware.RefreshHandler)
		v1Authorized.GET("/logout", authMiddleware.LogoutHandler)

		usrManager := adminusers.New(lg, user.NewInteractor(lg, params.AdminUserInfo()))
		v1Authorized.GET("/users", usrManager.Get)
		v1Authorized.DELETE("/users", usrManager.Delete)

		usrBlockManager := adminblock.New(lg, user.NewInteractor(lg, params.AdminUserInfo()))
		v1Authorized.GET("/users/blocked", usrBlockManager.Get)
		v1Authorized.DELETE("/users/blocked", usrBlockManager.Delete)
		v1Authorized.PATCH("/users/blocked", usrBlockManager.Update)

		modelManager := adminmodels.New(lg, model.NewInteractor(lg, params.AdminModelInfo()))
		v1Authorized.GET("/models", modelManager.Get)
		v1Authorized.DELETE("/models", modelManager.Delete)

		weightsManager := adminweights.New(lg, model.NewInteractor(lg, params.AdminModelInfo()))
		v1Authorized.GET("/models/weights", weightsManager.Get)
		v1Authorized.DELETE("/models/weights", weightsManager.Delete)
	}
}

func initStatRoutes(params config.Config, lg *logger.Logger, engine *gin.Engine) {
	authManager := auth.NewHandler(lg, user.NewInteractor(lg, params.UserInfo()))
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

	v1Authorized := engine.
		Group("/api/v1/stat").
		Use(authMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		v1Authorized.GET("/refresh", authMiddleware.RefreshHandler)
		v1Authorized.GET("/logout", authMiddleware.LogoutHandler)

		userManager := statusers.New(lg, user.NewInteractor(lg, params.StatUserInfo()))
		v1Authorized.GET("/users", userManager.Get)

		modelManager := statmodels.New(lg, model.NewInteractor(lg, params.StatModelInfo()))
		v1Authorized.GET("/models", modelManager.Get)

		weightsManager := statweights.New(lg, model.NewInteractor(lg, params.StatModelInfo()))
		v1Authorized.GET("/weights", weightsManager.Get)
	}
}

func initFailure(params config.Config, lg *logger.Logger, engine *gin.Engine) {
	authManager := auth.NewHandler(lg, user.NewInteractor(lg, params.UserInfo()))
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
		lg.WithFields(logrus.Fields{"claims": claims}).Info("no route")
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
