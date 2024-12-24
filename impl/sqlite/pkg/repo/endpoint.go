package repo

import (
	"context"
	"database/sql"
	"errors"
)

type (
	Endpoint interface {
		GetDB() *sql.DB
	}

	EndpointCtxKey int

	defaultEndpoint struct{ db *sql.DB }
)

func (d defaultEndpoint) GetDB() *sql.DB { return d.db }

func NewContextWithEndpoint(ctx context.Context, endpoint Endpoint) context.Context {
	return context.WithValue(ctx, EndpointCtxKey(1), endpoint)
}

func EndpointFromContext(ctx context.Context) Endpoint {
	endpoint, found := ctx.Value(EndpointCtxKey(1)).(Endpoint)
	if !found {
		panic(errors.New("db.Endpoint not found in context"))
	}
	return endpoint
}
