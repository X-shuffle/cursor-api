package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

// GenerateChecksumWithDefault 使用默认方式生成校验和
func GenerateChecksumWithDefault() string {
	timestamp := time.Now().Unix()
	data := fmt.Sprintf("%d", timestamp)
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// GenerateChecksumWithRepair 修复并生成校验和
func GenerateChecksumWithRepair(input string) string {
	if input == "" {
		return GenerateChecksumWithDefault()
	}
	// 如果需要，这里可以添加校验和修复逻辑
	return input
} 