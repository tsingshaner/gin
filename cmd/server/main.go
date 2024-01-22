package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	openAPI "github.com/lab-online/api/open-api"
	"github.com/lab-online/config"
	app "github.com/lab-online/internal"
	"github.com/lab-online/pkg/color"
	"github.com/lab-online/pkg/database"
	"github.com/lab-online/pkg/logger"
	"github.com/lab-online/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

//	@title			在线实验平台
//	@version		1.0
//	@description	在线实验平台 API 文档
//	@description.markdown
//	@termsOfService	http://swagger.io/terms/

//	@securityDefinitions.apikey	BearerToken
//	@in							header
//	@name						Authorization
//	@Security					BearerToken

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
	serverApp := app.NewApp(db)

	if err := serverApp.Migrate(); err != nil {
		logger.Warn("failed to migrate database")
		logger.Warn(err.Error())
	}

	setupOpenAPI(engine, serverAddr, config.Server.Prefix)
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
		color.Style(server.Addr, color.ColorBlue, color.FontUnderline),
	)
	err := server.ListenAndServe()
	if err != nil {
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

func setupOpenAPI(engine *gin.Engine, serverAddr string, BasePath string) {
	if !config.Server.EnableOpenAPI {
		return
	}

	logger.Info(
		"open api is enabled",
		color.Style(
			fmt.Sprintf("http://%s%s/index.html", serverAddr, config.Server.OpenAPIRoute),
			color.ColorBlue, color.FontUnderline,
		),
	)
	openAPI.SwaggerInfo.Host = serverAddr
	openAPI.SwaggerInfo.BasePath = config.Server.Prefix
	engine.GET(
		fmt.Sprintf("%s/*any", config.Server.OpenAPIRoute),
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
}
