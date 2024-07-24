// Package resp
//
// All resp helper will call c.Abort()
package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/go-pkg/errors"
)

func NewBody[T errors.Coder](code T) *Body[T] {
	return &Body[T]{code}
}

func NewSuccessBody[C errors.Coder, Data any](code C, data Data) *SuccessBody[Data, C] {
	return &SuccessBody[Data, C]{
		data,
		Body[C]{code},
	}
}

func NewFailedBody[C errors.Coder, Err any](code C, err Err) *FailedBody[Err, C] {
	return &FailedBody[Err, C]{
		err,
		Body[C]{code},
	}
}

func Success[C errors.Coder, Data any](c *gin.Context, status int, code C, data ...Data) {
	if len(data) == 0 {
		c.AbortWithStatusJSON(status, NewBody(code))
	} else {
		c.AbortWithStatusJSON(status, NewSuccessBody(code, data[0]))
	}
}

func OK[C errors.Coder, Data any](c *gin.Context, code C, data ...Data) {
	Success(c, http.StatusOK, code, data...)
}

func Created[C errors.Coder, Data any](c *gin.Context, code C, data ...Data) {
	Success(c, http.StatusCreated, code, data...)
}

func NoContent[T errors.Coder](c *gin.Context) {
	c.AbortWithStatus(http.StatusNoContent)
}

func Failed[C errors.Coder, Err any](c *gin.Context, status int, code C, err ...Err) {
	if len(err) == 0 {
		c.AbortWithStatusJSON(status, NewBody(code))
	} else {
		c.AbortWithStatusJSON(status, NewFailedBody(code, err[0]))
	}
}

func BadRequest[C errors.Coder, Err any](c *gin.Context, code C, err ...Err) {
	Failed(c, http.StatusBadRequest, code, err...)
}

func NotFound[C errors.Coder, Err any](c *gin.Context, code C, err ...Err) {
	Failed(c, http.StatusNotFound, code, err...)
}

func InternalServerError[C errors.Coder, Err any](c *gin.Context, code C, err ...Err) {
	Failed(c, http.StatusInternalServerError, code, err...)
}

func ValidateError[T any](c *gin.Context, message T) {
	BadRequest(c, CodeBadRequest, message)
}
