package validator

import (
	"errors"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tsingshaner/gin/resp"
)

const (
	keyBody   = "@@body"
	keyParams = "@@params"
	keyQuery  = "@@query"
	keyHeader = "@@header"
)

type Binder func(*gin.Context) func(obj any) error

// handleValidatorError 处理验证错误
var ErrorHandler func(*gin.Context, error) = func(c *gin.Context, err error) {
	var validationErrs validator.ValidationErrors

	if errors.As(err, &validationErrs) {
		translator, _ := Translators.FindTranslator(
			strings.Split(c.Request.Header.Get("Accept-Language"), ",")...,
		)
		resp.ValidateError(c, validationErrs.Translate(translator))
		return
	} else if errors.Is(err, io.EOF) {
		resp.ValidateError(c, "request args is empty")
		return
	}

	c.Error(errors.Join(errors.New("unknown args validate error"), resp.ErrValidate, err))
	c.Abort()
}

func Validator[T any](binder Binder, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := new(T)

		if err := binder(c)(body); err != nil {
			ErrorHandler(c, err)
		}

		c.Set(key, body)
	}
}

func Query[T any]() gin.HandlerFunc {
	return Validator[T](func(ctx *gin.Context) func(obj any) error {
		return ctx.ShouldBindQuery
	}, keyQuery)
}

func Body[T any]() gin.HandlerFunc {
	return Validator[T](func(ctx *gin.Context) func(obj any) error {
		return ctx.ShouldBindJSON
	}, keyBody)
}

func Params[T any]() gin.HandlerFunc {
	return Validator[T](func(ctx *gin.Context) func(obj any) error {
		return ctx.ShouldBindUri
	}, keyParams)
}

func Header[T any]() gin.HandlerFunc {
	return Validator[T](func(ctx *gin.Context) func(obj any) error {
		return ctx.ShouldBindHeader
	}, keyHeader)
}

func GetBody[T any](c *gin.Context, key ...string) *T {
	return c.MustGet(getKey(keyBody, key)).(*T)
}

func GetQuery[T any](c *gin.Context, key ...string) *T {
	return c.MustGet(getKey(keyQuery, key)).(*T)
}

func GetParams[T any](c *gin.Context, key ...string) *T {
	return c.MustGet(getKey(keyParams, key)).(*T)
}

func GetHeader[T any](c *gin.Context, key ...string) *T {
	return c.MustGet(getKey(keyHeader, key)).(*T)
}

func getKey(fallback string, keys []string) string {
	if len(keys) > 0 {
		return keys[0]
	}

	return fallback
}
