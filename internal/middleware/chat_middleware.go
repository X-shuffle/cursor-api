package middleware

import (
	"context"
	"cursor-api/internal/domain/errors"
	"cursor-api/internal/domain/model"
	"time"
)

type ChatMiddleware struct {
	rateLimiter *RateLimiter
}

func NewChatMiddleware() *ChatMiddleware {
	return &ChatMiddleware{
		rateLimiter: NewRateLimiter(10, time.Second), // 每秒10个请求
	}
}

func (m *ChatMiddleware) Process(ctx context.Context, req model.ChatRequest, next func(context.Context, model.ChatRequest) (*model.ChatResponse, error)) (*model.ChatResponse, error) {
	// 速率限制
	if !m.rateLimiter.Allow() {
		return nil, errors.ErrRateLimitExceeded
	}

	// 请求验证
	if err := m.validateRequest(req); err != nil {
		return nil, err
	}

	// 调用下一个处理器
	return next(ctx, req)
}

func (m *ChatMiddleware) validateRequest(req model.ChatRequest) error {
	if req.Model == "" {
		return errors.ErrInvalidModel
	}
	if len(req.Messages) == 0 {
		return &errors.ChatError{Code: "invalid_request", Message: "Messages cannot be empty"}
	}
	return nil
} 