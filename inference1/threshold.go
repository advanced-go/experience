package inference1

const (
	PercentilePercent = 95
	PercentileValue   = 3000 // milliseconds
	PercentileMinimum = 0

	StatusCodePercent = 10
	StatusCodeValue   = 0
	StatusCodeMinimum = 100
)

type Threshold struct {
	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Value   int `json:"value"`   // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}

func InitPercentileThreshold(t *Threshold) {
	if t != nil {
		t.Minimum = PercentileMinimum
		t.Percent = PercentilePercent
		t.Value = PercentileValue
	}
}

func InitStatusCodeThreshold(t *Threshold) {
	if t != nil {
		t.Minimum = StatusCodeMinimum
		t.Percent = StatusCodePercent
		t.Value = StatusCodeValue
	}
}
