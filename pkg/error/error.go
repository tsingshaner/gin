package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lab-online/pkg/constant"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) ToJSON() map[string]any {
	return gin.H{
		"code": e.Code,
		"msg":  e.Message,
	}
}

func New(code int, msgMap map[int]string) *Error {
	return &Error{
		Code:    code,
		Message: msgMap[code],
	}
}

func HandleError(c *gin.Context, err error, statusMap map[int]int) {
	if customErr, ok := err.(*Error); ok {
		c.AbortWithStatusJSON(statusMap[customErr.Code], customErr.ToJSON())
	} else {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": constant.INTERNAL_SERVER_ERROR,
			"msg":  constant.CodeMsg[constant.INTERNAL_SERVER_ERROR],
		})
	}
}
