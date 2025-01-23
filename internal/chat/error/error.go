package error

type ChatError struct {
	Code    string
	Message string
}

func (e *ChatError) Error() string {
	return e.Message
}

var (
	ErrInvalidModel = &ChatError{Code: "invalid_model", Message: "Invalid model specified"}
	ErrNoToken = &ChatError{Code: "no_token", Message: "No available token"}
	ErrRateLimit = &ChatError{Code: "rate_limit", Message: "Rate limit exceeded"}
) 