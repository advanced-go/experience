package inference1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath           = "github/advanced-go/experience/inference1"
	inferenceResource = "inference"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	return get[core.Log, Entry](ctx, core.AddRequestId(h), values, inferenceResource, "")
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), inferenceResource, "", body)
}

// Ingress
/*
// IngressQuery - query ingress inference
func IngressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

*/

// AddIngress - insert ingress inference
func AddIngress(ctx context.Context, h http.Header, e Entry) (int, *core.Status) {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e})
	// Need to return this from database
	inferenceId := 0
	return inferenceId, status
}

/*
// IngressInsertInterval - insert ingress interval inference
func IngressInsertInterval(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e}, nil)
	return status
}
*/

// Egress

/*
// EgressQuery - query egress inference
func EgressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}


*/

// AddEgress - insert egress inference
func AddEgress(ctx context.Context, h http.Header, e Entry) (int, *core.Status) {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e})
	// Need to return from database
	inferenceId := 0
	return inferenceId, status
}

/*
// EgressInsertInterval - insert egress interval inference
func EgressInsertInterval(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e}, nil)
	return status
}


*/
