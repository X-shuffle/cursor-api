package model

// TokenInfo 表示令牌信息的数据模型
type TokenInfo struct {
	Token     string       `json:"token"`
	Checksum  string       `json:"checksum"`
	Profile   *UserProfile `json:"profile,omitempty"`
	Used      int64       `json:"used"`
	Available bool        `json:"available"`
}

type UserProfile struct {
	User   UserInfo   `json:"user"`
	Stripe StripeInfo `json:"stripe"`
	Usage  UsageInfo  `json:"usage"`
}

type UserInfo struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type StripeInfo struct {
	MembershipType       string `json:"membership_type"`
	PaymentID           string `json:"payment_id,omitempty"`
	DaysRemainingOnTrial int    `json:"days_remaining_on_trial"`
}

type UsageInfo struct {
	Premium  *ModelUsage `json:"premium,omitempty"`
	Standard *ModelUsage `json:"standard,omitempty"`
	Unknown  *ModelUsage `json:"unknown,omitempty"`
}

type ModelUsage struct {
	Requests    int64 `json:"requests"`
	Tokens      int64 `json:"tokens"`
	MaxRequests int64 `json:"max_requests,omitempty"`
} 