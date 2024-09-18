package action1

import "time"

// Actions - no updates as history with corresponding inference is needed

type Actions struct {
	RouteName        string  `json:"route"`
	RedirectLocation string  `json:"redirect-location"`
	StatusCode       string  `json:"status-code"`
	Limit            float64 `json:"limit"`
	Burst            int     `json:"burst"`
	RouteLocation    string  `json:"route-location"`
	Percentage       int     `json:"percentage"`
}

// Action - host
type Action struct {
	// Origin + route for uniqueness
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	RouteName string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	Action     string  `json:"action"`
	Timeout    int     `json:"timeout"`
	RateLimit  float64 `json:"rate-limit"`
	RateBurst  int     `json:"rate-burst"`
	Primary    string  `json:"primary"`
	Secondary  string  `json:"secondary"`
	Percentage int     `json:"percentage"`
	Status     string  `json:"status"`
}
