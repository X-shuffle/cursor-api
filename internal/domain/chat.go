package domain

// Message 表示聊天消息的领域模型
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest 表示聊天请求的领域模型
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// ChatService 定义聊天服务接口
type ChatService interface {
	HandleChat(req ChatRequest) (interface{}, error)
	GetModels() ([]string, error)
} 