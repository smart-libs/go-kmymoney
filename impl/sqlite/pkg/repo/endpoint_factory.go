package repo

import (
	"context"
	"database/sql"
)

type (
	EndpointFactory interface {
		CreateEndpoint(ctx context.Context) Endpoint
	}

	defaultEndpointFactory struct{ db *sql.DB }
)

func (d defaultEndpointFactory) CreateEndpoint(_ context.Context) Endpoint {
	return defaultEndpoint{db: d.db}
}

func NewEndpointFactory(config Config, provider LoggerProvider) EndpointFactory {
	db := NewSQLDB(config.SQLite3Source, provider)
	return defaultEndpointFactory{
		db: db,
	}
}
