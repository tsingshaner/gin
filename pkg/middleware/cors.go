package middleware

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/go-pkg/color"
)

func CORS(config cors.Config) gin.HandlerFunc {
	corsTag := color.UnsafeGreen(color.UnsafeBold(" cors:"))
	fmt.Println(corsTag, "origins:", config.AllowOrigins)
	fmt.Println(corsTag, "methods:", config.AllowMethods)
	fmt.Println(corsTag, "headers:", config.AllowHeaders)
	fmt.Println(corsTag, "credentials:", config.AllowCredentials)
	fmt.Println(corsTag, "max age:", config.MaxAge)

	return cors.New(config)
}
