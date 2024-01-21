package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/logger"
	"github.com/lab-online/pkg/resp"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("recover unhandled error", err)
				slog.Error("recover middleware", err)
				resp.InternalServerError[any](c, resp.CodeServerError)
			}
		}()

		c.Next()
	}
}
