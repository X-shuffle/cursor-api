package utils

import (
	"os"
	"strconv"
	"strings"
)

// ParseStringFromEnv 从环境变量获取字符串值
func ParseStringFromEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ParseBoolFromEnv 从环境变量获取布尔值
func ParseBoolFromEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		switch strings.ToLower(value) {
		case "1", "true", "yes", "on":
			return true
		case "0", "false", "no", "off":
			return false
		}
	}
	return defaultValue
}

// ParseIntFromEnv 从环境变量获取整数值
func ParseIntFromEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

// ParseInt64FromEnv 从环境变量获取64位整数值
func ParseInt64FromEnv(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return i
		}
	}
	return defaultValue
}

// ParseFloat64FromEnv 从环境变量获取浮点数值
func ParseFloat64FromEnv(key string, defaultValue float64) float64 {
	if value := os.Getenv(key); value != "" {
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		}
	}
	return defaultValue
} 