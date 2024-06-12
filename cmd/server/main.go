package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tsingshaner/gin-starter/config"
	app "github.com/tsingshaner/gin-starter/internal"
	"github.com/tsingshaner/gin-starter/pkg/database"
	"github.com/tsingshaner/gin-starter/pkg/jwt"
	"github.com/tsingshaner/gin-starter/pkg/logger"
	"github.com/tsingshaner/gin-starter/pkg/middleware"
	"github.com/tsingshaner/gin-starter/pkg/openapi"
	"github.com/tsingshaner/go-pkg/color"
)

//	@title			在线实验平台
//	@version		1.0
//	@description	在线实验平台 API 文档
//	@description.markdown
//	@termsOfService	http://swagger.io/terms/

//	@securityDefinitions.apikey	bearer
//	@in							header
//	@name						Authorization

//	@contact.name	Issues
//	@contact.url	http://github.com/Jon-a-than/gin-template/issues

//	@license.name	MIT
//	@license.url	https://github.com/Jon-a-than/gin-template/blob/main/LICENSE

// @externalDocs.description	ApiFox
// @externalDocs.url			https://apifox.com/apidoc/shared-3e844af7-e01f-4a3a-a44d-9b395189d4d5
func main() {
	engine := initEngine()
	db := database.ConnectDB(config.Database.Postgres, &gorm.Config{})
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	bindGlobalMiddleware(engine)
	jsonWebToken, _ := jwt.New(config.JWT)
	serverApp := app.New(db, jsonWebToken)

	if err := serverApp.Migrate(); err != nil {
		logger.Warn("failed to migrate database")
		logger.Warn(err.Error())
	}

	if config.Server.EnableOpenAPI {
		openapi.Setup(engine, &openapi.Config{
			DOCSPrefix: config.Server.OpenAPIRoute,
			ServerAddr: serverAddr,
			BasePath:   config.Server.Prefix,
		})
	}
	serverApp.RoutesRegister(engine.Group(config.Server.Prefix))

	server := &http.Server{
		Addr:           serverAddr,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Info(
		"server will listening on",
		color.UnsafeUnderline(color.UnsafeBlue(server.Addr)),
	)
	if err := server.ListenAndServe(); err != nil {
		logger.Error(err.Error())
	}
}

func initEngine() *gin.Engine {
	gin.SetMode(config.Server.Mode)
	engine := gin.New()
	gin.DebugPrintRouteFunc = logger.PrintRouter
	engine.NoRoute(middleware.NotFound)
	return engine
}

func bindGlobalMiddleware(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.CORS.AllowOrigins
	corsConfig.AllowMethods = config.CORS.AllowMethods
	corsConfig.AllowHeaders = config.CORS.AllowHeaders
	corsConfig.AllowCredentials = config.CORS.AllowCredentials
	corsConfig.MaxAge = time.Duration(config.CORS.MaxAge) * time.Minute

	engine.Use(
		middleware.Logger(middleware.LoggerConfig{
			Console: &config.Logger.ConsoleHttp,
			Level:   &config.Logger.HttpLevel,
		}),
		middleware.CORS(corsConfig),
		middleware.Recover(),
	)
}
