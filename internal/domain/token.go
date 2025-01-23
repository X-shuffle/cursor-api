package domain

import "cursor-api/internal/domain/model"

// TokenRepository 定义令牌仓储接口
type TokenRepository interface {
	GetAll() ([]model.TokenInfo, error)
	SaveAll(tokens []model.TokenInfo) error
	Reload() ([]model.TokenInfo, error)
} 