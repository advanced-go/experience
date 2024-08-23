package action1

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
