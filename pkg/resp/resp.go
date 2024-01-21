package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBaseRespBody(code int) *BaseRespBody {
	return &BaseRespBody{code}
}

func NewSuccessRespBody[T any](code int, data T) *SuccessRespBody[T] {
	return &SuccessRespBody[T]{
		BaseRespBody{code},
		data,
	}
}

func NewFailedRespBody[T any](code int, err T) *FailedRespBody[T] {
	return &FailedRespBody[T]{
		BaseRespBody{code},
		err,
	}
}

func Success[T any](c *gin.Context, code int, data ...T) {
	if len(data) == 0 {
		c.JSON(http.StatusOK, NewBaseRespBody(code))
	} else {
		c.JSON(http.StatusOK, NewSuccessRespBody(code, data[0]))
	}
}

func Created[T any](c *gin.Context, code int, data ...T) {
	if len(data) == 0 {
		c.AbortWithStatusJSON(http.StatusCreated, NewBaseRespBody(code))
	} else {
		c.AbortWithStatusJSON(http.StatusCreated, NewSuccessRespBody(code, data[0]))
	}
}

func Deleted(c *gin.Context, code int) {
	c.JSON(http.StatusNoContent, NewBaseRespBody(code))
}

func BadRequest[T any](c *gin.Context, code int, err ...T) {
	if len(err) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewBaseRespBody(code))
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewFailedRespBody(code, err[0]))
	}
}

func NotFound[T any](c *gin.Context, code int, err ...T) {
	if len(err) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, NewBaseRespBody(code))
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, NewFailedRespBody(code, err[0]))
	}
}

func InternalServerError[T any](c *gin.Context, code int, err ...T) {
	if len(err) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewBaseRespBody(code))
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewFailedRespBody(code, err[0]))
	}
}
