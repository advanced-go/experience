package action1

import (
	"github.com/advanced-go/stdlib/core"
)

// Routing - ingress and egress
// Need to determine how to represent 2 states:
// 1. Nil or not configured - location empty, percentage = -1
// 2. Re-routing in progress - location not empty, percentage >= 0
type Routing struct {
	AgentId     string      `json:"agent-id"`
	Origin      core.Origin `json:"origin"`
	InferenceId int         `json:"inference-id"`
	Location    string      `json:"location"`
	Percentage  int         `json:"percentage"`
}

func NewRouting() *Routing {
	r := new(Routing)
	r.Percentage = -1
	return r
}

func (r *Routing) IsActive() bool {
	return r.Location != "" && r.Percentage >= 0
}

// RoutingAction - ingress and egress
// Need to determine how to represent 2 states:
// 1. Nil or not configured - location empty, percentage = -1
// 2. Re-routing in progress - location not empty, percentage >= 0
/*
type RoutingAction struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`
	Location    string    `json:"location"`
	Percentage  int       `json:"percentage"`
}


*/
