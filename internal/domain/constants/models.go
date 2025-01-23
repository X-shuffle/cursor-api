package constants

import "cursor-api/internal/domain/types"

const (
	// 错误信息
	ERR_UNSUPPORTED_GIF         = "不支持动态 GIF"
	ERR_UNSUPPORTED_IMAGE_FORMAT = "不支持的图片格式，仅支持 PNG、JPEG、WEBP 和非动态 GIF"
	ERR_NODATA                  = "No data"

	// 基础常量
	MODEL_OBJECT = "model"
	CREATED      = 1706659200

	// 提供商
	ANTHROPIC = "anthropic"
	CURSOR    = "cursor"
	GOOGLE    = "google"
	OPENAI    = "openai"
	DEEPSEEK  = "deepseek"

	// 模型名称
	CLAUDE_3_5_SONNET              = "claude-3.5-sonnet"
	GPT_4                          = "gpt-4"
	GPT_4O                         = "gpt-4o"
	CLAUDE_3_OPUS                  = "claude-3-opus"
	CURSOR_FAST                    = "cursor-fast"
	CURSOR_SMALL                   = "cursor-small"
	GPT_3_5_TURBO                 = "gpt-3.5-turbo"
	GPT_4_TURBO_2024_04_09        = "gpt-4-turbo-2024-04-09"
	GPT_4O_128K                   = "gpt-4o-128k"
	GEMINI_1_5_FLASH_500K         = "gemini-1.5-flash-500k"
	CLAUDE_3_HAIKU_200K           = "claude-3-haiku-200k"
	CLAUDE_3_5_SONNET_200K        = "claude-3-5-sonnet-200k"
	CLAUDE_3_5_SONNET_20241022    = "claude-3-5-sonnet-20241022"
	GPT_4O_MINI                   = "gpt-4o-mini"
	O1_MINI                       = "o1-mini"
	O1_PREVIEW                    = "o1-preview"
	O1                            = "o1"
	CLAUDE_3_5_HAIKU              = "claude-3.5-haiku"
	GEMINI_EXP_1206               = "gemini-exp-1206"
	GEMINI_2_0_FLASH_THINKING_EXP = "gemini-2.0-flash-thinking-exp"
	GEMINI_2_0_FLASH_EXP          = "gemini-2.0-flash-exp"
	DEEPSEEK_V3                   = "deepseek-v3"
	DEEPSEEK_R1                   = "deepseek-r1"
)

var AVAILABLE_MODELS = []types.Model{
	{ID: CLAUDE_3_5_SONNET, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: ANTHROPIC},
	{ID: GPT_4, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: GPT_4O, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: CLAUDE_3_OPUS, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: ANTHROPIC},
	{ID: CURSOR_FAST, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: CURSOR},
	{ID: CURSOR_SMALL, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: CURSOR},
	{ID: GPT_3_5_TURBO, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: GPT_4_TURBO_2024_04_09, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: GPT_4O_128K, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: GEMINI_1_5_FLASH_500K, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: GOOGLE},
	{ID: CLAUDE_3_HAIKU_200K, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: ANTHROPIC},
	{ID: CLAUDE_3_5_SONNET_200K, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: ANTHROPIC},
	{ID: CLAUDE_3_5_SONNET_20241022, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: ANTHROPIC},
	{ID: GPT_4O_MINI, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: O1_MINI, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: O1_PREVIEW, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: O1, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: OPENAI},
	{ID: CLAUDE_3_5_HAIKU, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: ANTHROPIC},
	{ID: GEMINI_EXP_1206, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: GOOGLE},
	{ID: GEMINI_2_0_FLASH_THINKING_EXP, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: GOOGLE},
	{ID: GEMINI_2_0_FLASH_EXP, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: GOOGLE},
	{ID: DEEPSEEK_V3, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: DEEPSEEK},
	{ID: DEEPSEEK_R1, Created: CREATED, Object: MODEL_OBJECT, OwnedBy: DEEPSEEK},
}

var USAGE_CHECK_MODELS = []string{
	CLAUDE_3_5_SONNET_20241022,
	CLAUDE_3_5_SONNET,
	GEMINI_EXP_1206,
	GPT_4,
	GPT_4_TURBO_2024_04_09,
	GPT_4O,
	CLAUDE_3_5_HAIKU,
	GPT_4O_128K,
	GEMINI_1_5_FLASH_500K,
	CLAUDE_3_HAIKU_200K,
	CLAUDE_3_5_SONNET_200K,
}

var LONG_CONTEXT_MODELS = []string{
	GPT_4O_128K,
	GEMINI_1_5_FLASH_500K,
	CLAUDE_3_HAIKU_200K,
	CLAUDE_3_5_SONNET_200K,
} 