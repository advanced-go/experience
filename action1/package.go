package action1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/experience/action1"
)

// TODO : need a way to retrieve all current actions for a given host and route

// ActiveRateLimiting - get latest rate limiting action
func ActiveRateLimiting(ctx context.Context, origin core.Origin, ingress bool) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// ActiveRouting - get latest routing action
func ActiveRouting(ctx context.Context, origin core.Origin, ingress bool) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// AddRateLimiting - insert rate limiting action
func AddRateLimiting(ctx context.Context, origin core.Origin, action *RateLimiting, ingress bool) *core.Status {
	if action == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: action is nil"))
	}
	return core.StatusOK()
}

// AddRouting - insert routing action
func AddRouting(ctx context.Context, origin core.Origin, action *Routing, ingress bool) *core.Status {
	if action == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: action is nil"))
	}
	return core.StatusOK()
}

// AddRedirect - insert redirect action
func AddRedirect(ctx context.Context, origin core.Origin, action *Redirect, ingress bool) *core.Status {
	if action == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: action is nil"))
	}
	return core.StatusOK()
}

/*
// Egress

// GetEgressActiveRateLimiting - get latest rate limiting action
func GetEgressActiveRateLimiting(ctx context.Context, origin core.Origin) (RateLimiting, *core.Status) {
	return RateLimiting{Limit: -1, Burst: -1}, core.StatusOK()
}

// GetEgressActiveRouting - get latest routing action
func GetEgressActiveRouting(ctx context.Context, origin core.Origin) (Routing, *core.Status) {
	return Routing{Location: "", Percentage: -1}, core.StatusOK()
}

// AddEgressRateLimiting - insert rate limiting action
func AddEgressRateLimiting(ctx context.Context, origin core.Origin, action *RateLimiting) *core.Status {
	if action == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: action is nil"))
	}
	return core.StatusOK()
}

// AddEgressRouting - insert routing action
func AddEgressRouting(ctx context.Context, origin core.Origin, action *Routing) *core.Status {
	if action == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: action is nil"))
	}
	return core.StatusOK()
}

// AddEgressRedirect - insert redirect action
func AddEgressRedirect(ctx context.Context, origin core.Origin, action *Redirect) *core.Status {
	if action == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: action is nil"))
	}
	return core.StatusOK()
}


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
