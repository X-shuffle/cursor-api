package constants

import (
	"os"
	"strconv"
	"time"
)

var (
	// 版本信息
	PKG_VERSION = getEnvString("PKG_VERSION", "1.0.0")

	// 路由前缀
	ROUTE_PREFIX = getEnvString("ROUTE_PREFIX", "")

	// 服务超时时间
	SERVICE_TIMEOUT = getEnvDuration("SERVICE_TIMEOUT", 30*time.Second)

	// OpenAI API 配置
	OPENAI_API_HOST = getEnvString("OPENAI_API_HOST", "api.openai.com")
	OPENAI_API_URL  = "https://" + OPENAI_API_HOST + "/v1/chat/completions"

	// Cursor API 配置
	CURSOR_API2_HOST = getEnvString("CURSOR_API2_HOST", "api2.cursor.sh")
	CURSOR_HOST      = getEnvString("CURSOR_HOST", "cursor.sh")

	// Cursor API URLs
	CURSOR_API2_CHAT_URL   = getCursorApiUrl(CURSOR_API2_HOST, "/aiserver.v1.AiService/StreamChat")
	CURSOR_API2_STRIPE_URL = getCursorApiUrl(CURSOR_API2_HOST, "/auth/full_stripe_profile")
	CURSOR_USAGE_API_URL   = getCursorApiUrl(CURSOR_HOST, "/api/usage")
	CURSOR_USER_API_URL    = getCursorApiUrl(CURSOR_HOST, "/api/auth/me")

	// 反向代理配置
	REVERSE_PROXY_HOST = getEnvString("REVERSE_PROXY_HOST", "")
	USE_REVERSE_PROXY  = REVERSE_PROXY_HOST != ""

	// 调试配置
	DEBUG          = getEnvBool("DEBUG", false)
	DEBUG_LOG_FILE = getEnvString("DEBUG_LOG_FILE", "debug.log")

	// API 路由
	ROUTE_MODELS_PATH = ROUTE_PREFIX + "/v1/models"
	ROUTE_CHAT_PATH   = ROUTE_PREFIX + "/v1/chat/completions"
)

// 其他路由常量
const (
	ROUTE_ROOT_PATH              = "/"
	ROUTE_HEALTH_PATH           = "/health"
	ROUTE_ABOUT_PATH            = "/about"
	ROUTE_API_PATH              = "/api"
	ROUTE_BASIC_CALIBRATION_PATH = "/calibration"
	ROUTE_BUILD_KEY_PATH        = "/build-key"
	ROUTE_CONFIG_PATH           = "/config"
	ROUTE_ENV_EXAMPLE_PATH      = "/env-example"
	ROUTE_GET_CHECKSUM          = "/checksum"
	ROUTE_GET_HASH              = "/hash"
	ROUTE_GET_TIMESTAMP_HEADER  = "/timestamp"
	ROUTE_LOGS_PATH            = "/logs"
	ROUTE_README_PATH          = "/readme"
	ROUTE_STATIC_PATH          = "/static/*filepath"
	ROUTE_TOKENS_PATH          = "/tokens"
	ROUTE_TOKENS_ADD_PATH      = "/tokens/add"
	ROUTE_TOKENS_DELETE_PATH   = "/tokens/delete"
	ROUTE_TOKENS_GET_PATH      = "/tokens/get"
	ROUTE_TOKENS_RELOAD_PATH   = "/tokens/reload"
	ROUTE_TOKENS_UPDATE_PATH   = "/tokens/update"
	ROUTE_USER_INFO_PATH       = "/user-info"
)

func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}

// getCursorApiUrl 生成 Cursor API URL
func getCursorApiUrl(host, path string) string {
	if USE_REVERSE_PROXY {
		return "https://" + REVERSE_PROXY_HOST + path
	}
	return "https://" + host + path
} 