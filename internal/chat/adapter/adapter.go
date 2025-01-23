package adapter

import (
	"cursor-api/internal/domain/aiserver"
	"cursor-api/internal/domain/model"
)

type ChatAdapter interface {
	Chat(req model.ChatRequest) (*model.ChatResponse, error)
	Stream(req model.ChatRequest, ch chan<- model.ChatResponse) error
}

type DefaultChatAdapter struct {
	server aiserver.AIServer
}

func NewChatAdapter(server aiserver.AIServer) *DefaultChatAdapter {
	return &DefaultChatAdapter{
		server: server,
	}
}
