package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errors "github.com/lab-online/pkg/error"
)

const (
	USER_ALREADY_EXISTS = iota + 1001
	USER_PASSWORD_EMPTY
)

var Msg = map[int]string{
	USER_ALREADY_EXISTS: "user already exists",
	USER_PASSWORD_EMPTY: "user password empty",
}

var Status = map[int]int{
	USER_ALREADY_EXISTS: http.StatusBadRequest,
	USER_PASSWORD_EMPTY: http.StatusBadRequest,
}

func New(code int) *errors.Error {
	return errors.New(code, Msg)
}

func HandleError(c *gin.Context, err error) {
	errors.HandleError(c, err, Status)
}
