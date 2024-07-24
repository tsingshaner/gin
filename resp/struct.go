package resp

import "github.com/tsingshaner/go-pkg/errors"

// Body response body with a custom resp code
type Body[T errors.Coder] struct {
	Code T `json:"code" binding:"required" example:"20000"` // 业务状态码
}

// FailedBody status `4xx` | `5xx`
type FailedBody[Err any, C errors.Coder] struct {
	Err Err `json:"err" binding:"required"` // 错误信息
	Body[C]
}

// SuccessBody status `2xx`
type SuccessBody[Data any, C errors.Coder] struct {
	Data Data `json:"data" binding:"required"` // 响应数据
	Body[C]
}

// FailedBodyWithString a alias of FailedBody with string error message
type FailedBodyWithString[C errors.Coder] struct {
	Err string `json:"err" binding:"required" example:"请求参数错误"` // 错误信息
	Body[C]
}
