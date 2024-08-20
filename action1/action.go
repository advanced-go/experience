package action1

import "time"

type Detail struct {
	EntryId     int       `json:"entry-id"`
	RouteName   string    `json:"route"`
	CreatedTS   time.Time `json:"created-ts"`
	InferenceId int       `json:"inference-id"`

	Limit      float64 `json:"limit"`
	Burst      int     `json:"burst"`
	Percentage int     `json:"percentage"`
	Location   string  `json:"location"`
	StatusCode string  `json:"status-code"`
}
