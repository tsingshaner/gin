package resp

// status `2xx` | `4xx` | `5xx`
type BaseRespBody struct {
	Code int `json:"code" binding:"required" example:"20000"` // 业务状态码
}

// status `4xx` | `5xx`
type FailedRespBody[T any] struct {
	BaseRespBody
	Err T `json:"err" binding:"required"` // 错误信息
}

// status `2xx`
type SuccessRespBody[T any] struct {
	BaseRespBody
	Data T `json:"data" binding:"required"` // 响应数据
}
