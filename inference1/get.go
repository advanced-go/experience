package inference1

import (
	"context"
	"github.com/advanced-go/experience/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, resource, template string) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}

	h2 = httpx.Forward(h2, h)
	h2.Set(core.XFrom, module.Authority)
	entries, status = pgxsql.QueryT[T](ctx, h, resource, template, values)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
	}
	return
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	if !core.OriginMatch(e.Origin, filter) {
		return false
	}
	// Additional filtering
	return true
}
