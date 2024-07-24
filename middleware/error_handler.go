package middleware

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin/resp"
	iErrors "github.com/tsingshaner/go-pkg/errors"
	"github.com/tsingshaner/go-pkg/log"
)

func NotFoundHandler(c *gin.Context) {
	resp.NotFound(c, resp.CodeNotFound, resp.ErrNotFound.Error())
}

func NewErrorHandler(logger log.Slog) gin.HandlerFunc {
	logger = logger.
		WithOptions(&log.ChildLoggerOptions{AddSource: false}).
		Named("middleware.error_handler")

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				e := fmt.Errorf("%+v", err)

				logger.Error("recover a unhandled err", slog.String("error", e.Error()))
				c.Error(errors.Join(resp.ErrInternal, e))
			}
		}()

		c.Next()

		if c.Writer.Status() == http.StatusOK && !c.Writer.Written() {
			lastErr := c.Errors.Last()
			var restErr iErrors.RESTError[string]
			if errors.As(lastErr.Err, &restErr) {
				message := strings.Split(lastErr.Error(), "\n")[0]
				resp.Failed(c, restErr.Status(), restErr.Code(), message)
			}
		}
	}
}
