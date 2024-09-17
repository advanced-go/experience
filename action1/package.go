package action1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/experience/action1"
)

// TODO : need a way to retrieve all current actions for a given host and route

// RateLimiting

// GetCurrentRateLimiting - get latest rate limiting action
func GetCurrentRateLimiting(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// AddRateLimiting - insert rate limiting action
func AddRateLimiting(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status {
	return core.StatusOK()
}

// Routing

// GetCurrentRouting - get latest routing action
func GetCurrentRouting(ctx context.Context, origin core.Origin) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// AddRouting - insert routing action
func AddRouting(ctx context.Context, origin core.Origin, action Routing) *core.Status {
	return core.StatusOK()
}

// ResetRouting - if the current routing acton is configured, then add a nil routing action,
// with a corresponding inference.
func ResetRouting(ctx context.Context, origin core.Origin, agentId string) *core.Status {
	// If the current routing action is configured, then add an action to reset.
	// Also add an inference

	return core.StatusOK()
}

// AddRedirect - insert redirect action
func AddRedirect(ctx context.Context, origin core.Origin, action Redirect) *core.Status {
	return core.StatusOK()
}
