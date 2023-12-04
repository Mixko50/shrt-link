package response

type BaseResponse[T any] struct {
	Success bool           `json:"success"`
	Data    *T             `json:"data"`
	Error   *ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	Detail       string `json:"detail"`
}
