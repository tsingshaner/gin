package resp

// status `2xx` | `4xx` | `5xx`
type BaseRespBody struct {
	Code int `json:"code" binding:"required" example:"20000"`
}

// status `4xx` | `5xx`
type FailedRespBody[T any] struct {
	BaseRespBody
	Err T `json:"err" binding:"required"`
}

// status `2xx`
type SuccessRespBody[T any] struct {
	BaseRespBody
	Data T `json:"data" binding:"required"`
}
