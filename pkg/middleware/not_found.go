package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin-starter/pkg/resp"
)

func NotFound(c *gin.Context) {
	resp.NotFound(c, resp.CodeNotFound, "route not found, please check your request url")
}
