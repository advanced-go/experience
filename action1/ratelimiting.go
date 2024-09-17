package action1

import "time"

// RateLimiting - ingress and egress
// AgentId     string    `json:"agent-id"`
// Need to represent 2 states:
// 1. Nil or not configured - both values == -1
// 2. Configured - both values >= 0
type RateLimiting struct {
	EntryId     int       `json:"entry-id"`
	Region      string    `json:"region"`
	Zone        string    `json:"zone"`
	SubZone     string    `json:"sub-zone"`
	Host        string    `json:"host"`
	Route       string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`

	Limit float64 `json:"limit"`
	Burst int     `json:"burst"`
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
