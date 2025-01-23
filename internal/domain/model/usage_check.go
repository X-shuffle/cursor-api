package model

import (
	"strings"
	"cursor-api/internal/domain/constants"
)

// UsageCheck 使用量检查类型
type UsageCheck struct {
	Type    string   `json:"type"`
	Models  []string `json:"models,omitempty"`
}

const (
	UsageCheckTypeNone    = "none"
	UsageCheckTypeDefault = "default"
	UsageCheckTypeAll     = "all"
	UsageCheckTypeList    = "list"
)

// NewUsageCheck 创建使用量检查
func NewUsageCheck(checkType string, models []string) UsageCheck {
	return UsageCheck{
		Type:   checkType,
		Models: filterValidModels(models),
	}
}

// FromString 从字符串创建使用量检查
func (u *UsageCheck) FromString(s string) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "none", "disabled":
		u.Type = UsageCheckTypeNone
	case "default":
		u.Type = UsageCheckTypeDefault
	case "all", "everything":
		u.Type = UsageCheckTypeAll
	default:
		if s == "" {
			u.Type = UsageCheckTypeDefault
			return
		}
		models := strings.Split(s, ",")
		validModels := filterValidModels(models)
		if len(validModels) == 0 {
			u.Type = UsageCheckTypeDefault
		} else {
			u.Type = UsageCheckTypeList
			u.Models = validModels
		}
	}
}

func filterValidModels(models []string) []string {
	validModels := make([]string, 0)
	for _, model := range models {
		model = strings.TrimSpace(model)
		for _, availableModel := range constants.AVAILABLE_MODELS {
			if availableModel.ID == model {
				validModels = append(validModels, model)
				break
			}
		}
	}
	return validModels
} 