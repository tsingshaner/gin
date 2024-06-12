package middleware

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tsingshaner/gin-starter/pkg/resp"
	"github.com/tsingshaner/gin-starter/pkg/shared"
	"github.com/tsingshaner/go-pkg/color"
)

// ValidatorOptions Todo interface{} 约束为结构体指针
type ValidatorOptions struct {
	BodyKey string
	Body    any

	ParamsKey string
	Params    any

	QueryKey string
	Query    any

	HeaderKey string
	Header    any
}

const (
	KeyBody   = "body"
	KeyParams = "params"
	KeyQuery  = "query"
	KeyHeader = "header"
)

type validatorMap map[string]interface{}
type validatorKeyMap map[string]string

func Validator(opts *ValidatorOptions) gin.HandlerFunc {
	options := validatorMap{}
	optionKey := validatorKeyMap{}

	registerValidator(options, optionKey, KeyBody, opts.BodyKey, opts.Body)
	registerValidator(options, optionKey, KeyParams, opts.ParamsKey, opts.Params)
	registerValidator(options, optionKey, KeyQuery, opts.QueryKey, opts.Query)
	registerValidator(options, optionKey, KeyHeader, opts.HeaderKey, opts.Header)

	checkIsStructPointer(options)

	return func(c *gin.Context) {
		for key, obj := range options {
			var err error
			switch key {
			case KeyBody:
				err = c.ShouldBindJSON(obj)
			case KeyParams:
				err = c.ShouldBindUri(obj)
			case KeyQuery:
				err = c.ShouldBindQuery(obj)
			case KeyHeader:
				err = c.ShouldBindHeader(obj)
			}

			if err != nil {
				handleValidatorError(c, err)
				return
			} else {
				c.Set(optionKey[key], obj)
			}
		}

		c.Next()
	}
}

func checkIsStructPointer(options validatorMap) {
	for _, obj := range options {
		val := reflect.ValueOf(obj)
		if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
			panic(color.UnsafeBold(color.UnsafeRed("fatal:")) +
				"Validator middleware options must be struct pointer")
		}
	}
}

func registerValidator(options validatorMap, optionKey validatorKeyMap, key string, customKey string, obj any) {
	if obj != nil {
		options[key] = obj
		if customKey == "" {
			optionKey[key] = key
		} else {
			optionKey[key] = customKey
		}
	}
}

func handleValidatorError(c *gin.Context, err error) {
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		errMsg := map[string]string{}
		for _, e := range validationErrs {
			key := namespaceToLowerCase(e.StructNamespace())
			errMsg[key] = fmt.Sprintf("validate failed: %s", e.ActualTag())
		}
		resp.BadRequest(c, resp.CodeValidateError, errMsg)
	} else {
		resp.BadRequest(c, resp.CodeValidateError, err.Error())
	}
}

func namespaceToLowerCase(namespace string) string {
	tokens := strings.Split(namespace, ".")
	for i, token := range tokens {
		tokens[i] = shared.UpperCamelCaseToLowerCameCase(token)
	}
	return strings.Join(tokens, ".")
}
