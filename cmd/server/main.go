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
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:4936
//	@BasePath	/api/v1

//	@contact.name	Issues
//	@contact.url	http://github.com/Jon-a-than/gin-template/issues

//	@license.name	MIT
//	@license.url	https://github.com/Jon-a-than/gin-template/blob/main/LICENSE

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	ApiFox
// @externalDocs.url			https://apifox.com/apidoc/shared-3e844af7-e01f-4a3a-a44d-9b395189d4d5
func main() {
	gin.SetMode(config.Server.Mode)
	engine := gin.New()
	gin.DebugPrintRouteFunc = logger.PrintRouter

	db := database.ConnectDB(config.Database.Postgres, &gorm.Config{})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.CORS.AllowOrigins
	corsConfig.AllowMethods = config.CORS.AllowMethods
	corsConfig.AllowHeaders = config.CORS.AllowHeaders
	corsConfig.AllowCredentials = config.CORS.AllowCredentials
	corsConfig.MaxAge = time.Duration(config.CORS.MaxAge) * time.Minute

	engine.Use(
		middleware.Logger(middleware.LoggerConfig{
			Console: true,
			Level:   &config.Logger.HttpLevel,
		}),
		middleware.CORS(corsConfig),
	)

	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	openAPI.SwaggerInfo.Host = serverAddr
	openAPI.SwaggerInfo.BasePath = config.Server.Prefix

	serverApp := app.NewApp(db)
	if err := serverApp.Migrate(); err != nil {
		logger.Warn("failed to migrate database")
		logger.Warn(err.Error())
	}

	fmt.Println()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	serverApp.RoutesRegister(engine.Group(config.Server.Prefix))

	server := &http.Server{
		Addr:           serverAddr,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println()
	logger.Info(
		"server will listening on",
		color.Style(server.Addr, color.ColorBlue, color.FontUnderline),
	)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
