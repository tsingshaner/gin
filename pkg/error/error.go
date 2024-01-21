package error

import (
	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/resp"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(code int, msgMap map[int]string) *Error {
	return &Error{
		Code:    code,
		Message: msgMap[code],
	}
}

func HandleError(c *gin.Context, err error, statusMap map[int]int) {
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
