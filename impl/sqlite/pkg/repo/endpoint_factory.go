package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
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
	logger := provider(nil)
	logger.Info("NewEndpointFactory", "sqlite3_db", config.SQLite3Source)
	dir, err := os.Getwd()
	logger.Info("NewEndpointFactory", "PWD", dir)
	db, err := sql.Open("sqlite3", config.SQLite3Source)
	if err != nil {
		panic(errors.Join(err, fmt.Errorf("NewEndpointFactory: failed to open database %s", config.SQLite3Source)))
	}
	return defaultEndpointFactory{
		db: db,
	}
}
