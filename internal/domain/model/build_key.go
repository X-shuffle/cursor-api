package model

// BuildKeyRequest 构建密钥请求
type BuildKeyRequest struct {
	AuthToken          string           `json:"auth_token"`
	EnableStreamCheck  *bool            `json:"enable_stream_check,omitempty"`
	IncludeStopStream *bool            `json:"include_stop_stream,omitempty"`
	DisableVision     *bool            `json:"disable_vision,omitempty"`
	EnableSlowPool    *bool            `json:"enable_slow_pool,omitempty"`
	UsageCheckModels  *UsageCheckModel `json:"usage_check_models,omitempty"`
}

type UsageCheckModel struct {
	Type     string   `json:"type"`
	ModelIDs []string `json:"model_ids,omitempty"`
}

const (
	UsageCheckModelTypeDefault  = "default"
	UsageCheckModelTypeDisabled = "disabled"
	UsageCheckModelTypeAll     = "all"
	UsageCheckModelTypeCustom  = "custom"
)

// BuildKeyResponse 构建密钥响应
type BuildKeyResponse struct {
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
}

func NewKeyResponse(key string) BuildKeyResponse {
	return BuildKeyResponse{
		Type:    "key",
		Content: key,
	}
}

func NewErrorResponse(err string) BuildKeyResponse {
	return BuildKeyResponse{
		Type:    "error",
		Content: err,
	}
} 