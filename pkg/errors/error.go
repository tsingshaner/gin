package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin-starter/pkg/resp"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func New(code int, msg string) *Error {
	return &Error{code, msg}
}

func HandleError(c *gin.Context, err error, statusMap map[int]int) {
	c.Errors = append(c.Errors, &gin.Error{Err: err})
	if customErr, ok := err.(*Error); ok {
		c.AbortWithStatusJSON(
			statusMap[customErr.Code],
			resp.NewFailedRespBody(customErr.Code, &customErr.Message),
		)
	} else {
		c.Error(err)
		resp.InternalServerError(c, resp.CodeServerError, "server unknown error")
	}
}
