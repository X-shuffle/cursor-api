package utils

import (
	"cursor-api/internal/domain/model"
	"encoding/json"
	"os"
	"strings"
	"unicode"
)

// LoadTokens 从文件加载令牌信息
func LoadTokens() ([]model.TokenInfo, error) {
	// 默认的令牌文件路径
	tokenPath := getEnvString("TOKEN_PATH", "tokens.json")

	// 读取文件
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在，返回空切片
			return []model.TokenInfo{}, nil
		}
		return nil, err
	}

	// 解析 JSON
	var tokens []model.TokenInfo
	if err := json.Unmarshal(data, &tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

// SaveTokens 保存令牌信息到文件
func SaveTokens(tokens []model.TokenInfo) error {
	tokenPath := getEnvString("TOKEN_PATH", "tokens.json")

	// 将令牌信息转换为 JSON
	data, err := json.MarshalIndent(tokens, "", "  ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(tokenPath, data, 0644)
}

// getEnvString 获取环境变量，如果不存在则返回默认值
func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ParseToken 解析并清理令牌字符串
func ParseToken(token string) string {
	// 移除空白字符
	token = strings.TrimSpace(token)
	// 移除不可见字符
	token = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, token)
	return token
}

// ValidateToken 验证令牌是否有效
func ValidateToken(token string) bool {
	if token == "" {
		return false
	}
	// 这里可以添加更多的验证规则
	// 例如：长度限制、字符集限制等
	return true
} 