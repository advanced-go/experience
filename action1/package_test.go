package action1

import "github.com/advanced-go/stdlib/core"

func ExampleEgressCurrentRateLimiting() {
	Ingress.CurrentRateLimiting(nil, core.Origin{})
}
