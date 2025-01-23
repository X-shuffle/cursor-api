package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateChecksum 生成校验和
func GenerateChecksum() string {
	timestamp := time.Now().Unix()
	data := fmt.Sprintf("%d", timestamp)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// GenerateChecksumWithRepair 生成带修复的校验和
func GenerateChecksumWithRepair(oldChecksum string) string {
	if oldChecksum == "" {
		return GenerateChecksum()
	}

	// 如果旧的校验和无效，生成新的
	if len(oldChecksum) != 32 {
		return GenerateChecksum()
	}

	// 验证旧的校验和是否为有效的十六进制字符串
	if _, err := hex.DecodeString(oldChecksum); err != nil {
		return GenerateChecksum()
	}

	return oldChecksum
} 