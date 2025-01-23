package model

type TokenAddRequest struct {
	Token    string  `json:"token"`
	Checksum *string `json:"checksum,omitempty"`
}

type TokensDeleteRequest struct {
	Tokens      []string        `json:"tokens"`
	Expectation TokenExpectation `json:"expectation"`
}

type TokenExpectation string

const (
	TokenExpectationNone     TokenExpectation = "none"
	TokenExpectationDetailed TokenExpectation = "detailed"
	TokenExpectationUpdated  TokenExpectation = "updated"
	TokenExpectationFailed   TokenExpectation = "failed"
)

func (e TokenExpectation) NeedsUpdatedTokens() bool {
	return e == TokenExpectationDetailed || e == TokenExpectationUpdated
}

func (e TokenExpectation) NeedsFailedTokens() bool {
	return e == TokenExpectationDetailed || e == TokenExpectationFailed
} 