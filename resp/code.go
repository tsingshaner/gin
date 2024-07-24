package resp

import "github.com/tsingshaner/go-pkg/errors"

const (
	CodeSuccess = "20000"

	CodeBadRequest = "40000"
	CodeValidateError
	CodeAuthError      = "40001"
	CodeForbidden      = "40003"
	CodeNotFound       = "40004"
	CodeBearerTokenErr = "40005"

	CodeInternalError    = "50000"
	CodeSetDeadlineError = "50001"
)

var (
	// ErrBadRequest 客户端请求错误
	ErrBadRequest = errors.NewREST(400, CodeBadRequest, "客户端请求错误")
	// ErrValidate 参数校验失败
	ErrValidate = errors.NewREST(400, CodeValidateError, "参数校验失败")
	// ErrAuth 身份验证失败
	ErrAuth = errors.NewREST(401, CodeAuthError, "身份验证失败")
	// ErrForbidden 无权限访问
	ErrForbidden = errors.NewREST(403, CodeForbidden, "无权限访问")
	// ErrNotFound 未找到资源
	ErrNotFound = errors.NewREST(404, CodeNotFound, "未找到资源")
	// ErrBearerToken Bearer Token无效
	ErrBearerToken = errors.NewREST(400, CodeBearerTokenErr, "Bearer Token 无效")
	// ErrInternal 内部服务错误
	ErrInternal = errors.NewREST(500, CodeInternalError, "内部服务错误")
	// ErrSetDeadline 服务端超时设置失败
	ErrSetDeadline = errors.NewREST(500, CodeSetDeadlineError, "服务端超时设置失败")
)
