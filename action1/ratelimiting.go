package action1

import (
	"github.com/advanced-go/stdlib/core"
)

// RateLimiting - ingress and egress
// AgentId     string    `json:"agent-id"`
// Need to represent 2 states:
// 1. Nil or not configured - both values == -1
// 2. Configured - both values >= 0
type RateLimiting struct {
	AgentId     string      `json:"agent-id"`
	Origin      core.Origin `json:"origin"`
	InferenceId int         `json:"inference-id"`
	Limit       float64     `json:"limit"`
	Burst       int         `json:"burst"`
}

func NewRateLimiting() *RateLimiting {
	r := new(RateLimiting)
	r.Limit = -1
	r.Burst = -1
	return r
}

func (r *RateLimiting) IsActive() bool {
	return r.Limit >= 0 && r.Burst >= 0
}
