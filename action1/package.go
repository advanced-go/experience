package action1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/experience/action1"
)

// GetRateLimiting - get latest rate limiting action
func GetRateLimiting(ctx context.Context, origin core.Origin) (*RateLimiting, *core.Status) {
	return &RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// AddRateLimiting - insert rate limiting action
func AddRateLimiting(ctx context.Context, origin core.Origin, inferenceId int, limit float64, burst int) *core.Status {
	return core.StatusOK()
}

// GetRouting - get latest routing action
func GetRouting(ctx context.Context, origin core.Origin) (*Routing, *core.Status) {
	return &Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// AddRouting - insert routing action
func AddRouting(ctx context.Context, origin core.Origin, inferenceId int, location string, percent int) *core.Status {
	return core.StatusOK()
}

// AddRedirect - insert redirect action
func AddRedirect(ctx context.Context, origin core.Origin, inferenceId int, location string, status string) *core.Status {
	return core.StatusOK()
}
