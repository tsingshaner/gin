package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/logger"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                logger.Error("recover unhandled error" ,err)
				slog.Error("recover middleware", err)
                c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": 50000,
					"msg":  "unknown error",
				})
            }
        }()

        c.Next()
    }
}
