package action1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/experience/action1"
)

// TODO : need a way to retrieve all current actions for a given host and route

// GetIngressActiveRateLimiting - get latest rate limiting action
func GetIngressActiveRateLimiting(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// AddIngressRateLimiting - insert rate limiting action
func AddIngressRateLimiting(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status {
	return core.StatusOK()
}

// GetIngressActiveRouting - get latest routing action
func GetIngressActiveRouting(ctx context.Context, origin core.Origin) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// AddIngressRouting - insert routing action
func AddIngressRouting(ctx context.Context, origin core.Origin, action Routing) *core.Status {
	return core.StatusOK()
}

// AddIngressRedirect - insert redirect action
func AddIngressRedirect(ctx context.Context, origin core.Origin, action Redirect) *core.Status {
	return core.StatusOK()
}

// Egress

// GetEgressActiveRateLimiting - get latest rate limiting action
func GetEgressActiveRateLimiting(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// AddEgressRateLimiting - insert rate limiting action
func AddEgressRateLimiting(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status {
	return core.StatusOK()
}

// GetEgressActiveRouting - get latest routing action
func GetEgressActiveRouting(ctx context.Context, origin core.Origin) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// AddEgressRouting - insert routing action
func AddEgressRouting(ctx context.Context, origin core.Origin, action Routing) *core.Status {
	return core.StatusOK()
}

// AddEgressRedirect - insert redirect action
func AddEgressRedirect(ctx context.Context, origin core.Origin, action Redirect) *core.Status {
	return core.StatusOK()
}

/*
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


*/

/*
// ingress - interface
type ingress struct {
	CurrentRateLimiting func(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status)
	AddRateLimiting     func(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status
}

var Ingress = func() *ingress {
	return &ingress{
		CurrentRateLimiting: func(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
			return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
		},
		AddRateLimiting: func(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status {
			return core.StatusOK()
		},
	}
}()
*/
