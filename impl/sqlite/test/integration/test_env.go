package integration

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"

	"context"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	"log/slog"
)

type TestEnv struct {
	context.Context
	LoggerProvider func(_ context.Context) *slog.Logger
}

func loggerProvider(_ context.Context) *slog.Logger { return slog.Default() }

func Run(t *testing.T, test func(env *TestEnv)) {
	ctx := context.Background()
	config := repo.NewConfig()
	factory := repo.NewEndpointFactory(config, loggerProvider)
	endpoint := factory.CreateEndpoint(ctx)
	ctx = repo.NewContextWithEndpoint(ctx, endpoint)

	_, err := endpoint.GetDB().Exec("delete from kmmSecurities")
	if assert.NoError(t, err) {
		test(&TestEnv{
			Context:        ctx,
			LoggerProvider: loggerProvider,
		})
	}
}
