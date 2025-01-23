package model

type ApiStatus string

const (
	ApiStatusSuccess ApiStatus = "success"
	ApiStatusFailed  ApiStatus = "failed"
	ApiStatusPending ApiStatus = "pending"
)

type ApiResponse struct {
	Status  ApiStatus   `json:"status"`
	Code    *int       `json:"code,omitempty"`
	Error   *string    `json:"error,omitempty"`
	Message *string    `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data interface{}, message string) ApiResponse {
	return ApiResponse{
		Status:  ApiStatusSuccess,
		Data:    data,
		Message: &message,
	}
}

func NewErrorResponse(code int, err string) ApiResponse {
	return ApiResponse{
		Status: ApiStatusFailed,
		Code:   &code,
		Error:  &err,
	}
} 