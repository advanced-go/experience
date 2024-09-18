package action1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/experience/action1"
)

// TODO : need a way to retrieve all current actions for a given host and route

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

// IngressCurrentRateLimiting - get latest rate limiting action
func IngressCurrentRateLimiting(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// IngressAddRateLimiting - insert rate limiting action
func IngressAddRateLimiting(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status {
	return core.StatusOK()
}

// IngressCurrentRouting - get latest routing action
func IngressCurrentRouting(ctx context.Context, origin core.Origin) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// IngressAddRouting - insert routing action
func IngressAddRouting(ctx context.Context, origin core.Origin, action Routing) *core.Status {
	return core.StatusOK()
}

// IngressAddRedirect - insert redirect action
func IngressAddRedirect(ctx context.Context, origin core.Origin, action Redirect) *core.Status {
	return core.StatusOK()
}

// Egress

// EgressCurrentRateLimiting - get latest rate limiting action
func EgressCurrentRateLimiting(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// EgressAddRateLimiting - insert rate limiting action
func EgressAddRateLimiting(ctx context.Context, origin core.Origin, action RateLimiting) *core.Status {
	return core.StatusOK()
}

// EgressCurrentRouting - get latest routing action
func EgressCurrentRouting(ctx context.Context, origin core.Origin) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// EgressAddRouting - insert routing action
func EgressAddRouting(ctx context.Context, origin core.Origin, action Routing) *core.Status {
	return core.StatusOK()
}

// EgressAddRedirect - insert redirect action
func EgressAddRedirect(ctx context.Context, origin core.Origin, action Redirect) *core.Status {
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
