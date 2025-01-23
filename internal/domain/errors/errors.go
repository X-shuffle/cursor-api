package errors

import "errors"

var (
	ErrTokenNotFound     = errors.New("token not found")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenUnavailable  = errors.New("token unavailable")
	ErrInvalidRequest    = errors.New("invalid request")
	ErrUnauthorized      = errors.New("unauthorized")
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

type ChatError struct {
	Code    string
	Message string
}

func (e *ChatError) Error() string {
	return e.Message
}

var (
	// 通用错误
	ErrInvalidModel       = &ChatError{Code: "invalid_model", Message: "Invalid model specified"}
	ErrNoToken           = &ChatError{Code: "no_token", Message: "No available token"}
	ErrRateLimitExceeded = &ChatError{Code: "rate_limit", Message: "Rate limit exceeded"}
	ErrServerError       = &ChatError{Code: "server_error", Message: "Internal server error"}
	
	// 图片相关错误
	ErrUnsupportedGIF         = &ChatError{Code: "unsupported_gif", Message: "不支持动态 GIF"}
	ErrUnsupportedImageFormat = &ChatError{Code: "unsupported_format", Message: "不支持的图片格式，仅支持 PNG、JPEG、WEBP 和非动态 GIF"}
	
	// 数据相关错误
	ErrNoData = &ChatError{Code: "no_data", Message: "No data"}
) 