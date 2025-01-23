package service

import (
	"context"
	"cursor-api/internal/domain/aiserver"
	"cursor-api/internal/domain/errors"
	"cursor-api/internal/domain/model"
)

type ChatService struct {
	server aiserver.AIServer
}

func NewChatService(server aiserver.AIServer) *ChatService {
	return &ChatService{
		server: server,
	}
}

func (s *ChatService) Chat(ctx context.Context, req model.ChatRequest) (*model.ChatResponse, error) {
	// 验证模型
	if !s.isValidModel(req.Model) {
		return nil, errors.ErrInvalidModel
	}

	// 调用 AI 服务器
	return s.server.Chat(req)
}

func (s *ChatService) Stream(ctx context.Context, req model.ChatRequest) (<-chan model.ChatResponse, error) {
	// 验证模型
	if !s.isValidModel(req.Model) {
		return nil, errors.ErrInvalidModel
	}

	ch := make(chan model.ChatResponse)
	go func() {
		defer close(ch)
		if err := s.server.Stream(req, ch); err != nil {
			// 处理错误
		}
	}()

	return ch, nil
}

func (s *ChatService) isValidModel(model string) bool {
	// 检查模型是否在支持列表中
	for _, m := range model.AVAILABLE_MODELS {
		if m.ID == model {
			return true
		}
	}
	return false
} 