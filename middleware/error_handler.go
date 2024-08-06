package middleware

import (
	stdErrors "errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin/middleware/requestid"
	"github.com/tsingshaner/gin/resp"
	"github.com/tsingshaner/go-pkg/errors"
	"github.com/tsingshaner/go-pkg/log"
)

func NotFoundHandler(c *gin.Context) {
	resp.NotFound(c, resp.CodeNotFound, resp.ErrNotFound.Error())
}

// NewErrorHandler is a middleware that recovers from any panics and writes a 500 if there was one.
//
// if the response status code is 200 and the response body is empty, it will call resp.Error to write a error response.
func NewErrorHandler(logger log.Slog) gin.HandlerFunc {
	logger = logger.
		WithOptions(&log.ChildLoggerOptions{AddSource: false}).
		Named("middleware.error_handler")

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				e := fmt.Errorf("%+v", err)

				logger.Error("recover a unhandled err", slog.String("err", e.Error()))

				c.Error(stdErrors.Join(resp.ErrInternal, e))
				resp.Error(c, c.Errors.Last())
			}
		}()

		c.Next()

		if c.Writer.Status() == http.StatusOK && !c.Writer.Written() {
			lastErr := c.Errors.Last()
			if restErr, ok := errors.Extract[errors.RESTError[string]](lastErr.Err); ok {
				message := strings.Split(lastErr.Error(), "\n")[0]
				resp.Failed(c, restErr.Status(), restErr.Code(), message)
			}
		}

		if errStr := c.Errors.String(); errStr != "" {
			logger.Warn(requestid.Get(c), slog.String("err", errStr))
		}
	}
}
