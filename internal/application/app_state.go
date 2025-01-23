package application

import (
	"sync"
	"time"
	"cursor-api/internal/domain/model"
)

type AppState struct {
	mu             sync.RWMutex
	tokenInfos     []model.TokenInfo
	totalRequests  int64
	activeRequests int64
	errorRequests  int64
	requestLogs    []model.RequestLog
}

func NewAppState(tokens []model.TokenInfo) *AppState {
	return &AppState{
		tokenInfos:  tokens,
		requestLogs: make([]model.RequestLog, 0, 100), // 预分配100条日志的空间
	}
}

// GetTotalRequests 获取总请求数
func (s *AppState) GetTotalRequests() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.totalRequests
}

// GetActiveRequests 获取活跃请求数
func (s *AppState) GetActiveRequests() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.activeRequests
}

// GetErrorRequests 获取错误请求数
func (s *AppState) GetErrorRequests() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.errorRequests
}

// GetLogs 获取请求日志
func (s *AppState) GetLogs() []model.RequestLog {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.requestLogs
}

// AddRequestLog 添加请求日志
func (s *AppState) AddRequestLog(log model.RequestLog) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requestLogs = append(s.requestLogs, log)
	s.totalRequests++
}

// IncrementActiveRequests 增加活跃请求数
func (s *AppState) IncrementActiveRequests() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.activeRequests++
}

// DecrementActiveRequests 减少活跃请求数
func (s *AppState) DecrementActiveRequests() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.activeRequests > 0 {
		s.activeRequests--
	}
}

// IncrementErrorRequests 增加错误请求数
func (s *AppState) IncrementErrorRequests() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errorRequests++
}

// GetTokenInfos 获取所有令牌信息
func (s *AppState) GetTokenInfos() []model.TokenInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tokenInfos
}

// UpdateTokenInfos 更新令牌信息
func (s *AppState) UpdateTokenInfos(tokens []model.TokenInfo) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokenInfos = tokens
}

// GetChecksum 获取当前校验和
func (s *AppState) GetChecksum() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return time.Now().Format("20060102150405")
} 