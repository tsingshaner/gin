package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/constant"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    constant.NOT_FOUND,
		"message": constant.CodeMsg[constant.NOT_FOUND],
	})
}
