package action1

import (
	"github.com/advanced-go/stdlib/core"
)

// Redirect - ingress and egress
type Redirect struct {
	AgentId     string      `json:"agent-id"`
	Origin      core.Origin `json:"origin"`
	InferenceId int         `json:"inference-id"`
	Location    string      `json:"location"`
	StatusCode  string      `json:"status-code"` // Only for ingress
}

func NewRedirect(location, statusCode string, inferenceId int) *Redirect {
	r := new(Redirect)
	r.Location = location
	r.StatusCode = statusCode
	r.InferenceId = inferenceId
	return r
}
