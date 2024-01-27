package openapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	openAPI "github.com/lab-online/api/open-api"
	"github.com/lab-online/pkg/color"
	"github.com/lab-online/pkg/logger"
)

type Config struct {
	DOCSPrefix string // 文档地址前缀
	ServerAddr string // 服务器地址
	BasePath   string // 各 API 路径前缀
}

func Setup(engine *gin.Engine, config *Config) {
	logger.Info(
		"open api is enabled",
		color.Style(
			fmt.Sprintf("http://%s%s/index.html", config.ServerAddr, config.DOCSPrefix),
			color.ColorBlue, color.FontUnderline,
		),
	)
	openAPI.SwaggerInfo.Host = config.ServerAddr
	openAPI.SwaggerInfo.BasePath = config.BasePath
	engine.GET(
		fmt.Sprintf("%s/*any", config.DOCSPrefix),
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
}
