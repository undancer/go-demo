package query

import (
	"context"
	"fmt"
	"net/http"
)

func NewQueryCtx(ctx context.Context, req *http.Request) (*QueryCtx, error) {
	q := req.FormValue("q")
	if q == "" {
		return nil, fmt.Errorf("no query supplied!")
	}
	return &QueryCtx{ctx, q}, nil
}

type QueryCtx struct {
	context.Context
	val string
}

func (ctx *QueryCtx) SetReq(req *http.Request) {
	q := req.URL.Query()
	q.Set("q", ctx.val)

	req.URL.RawQuery = q.Encode()
}