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
	t              *testing.T
	Endpoint       repo.Endpoint
	LoggerProvider func(_ context.Context) *slog.Logger
}

func loggerProvider(_ context.Context) *slog.Logger { return slog.Default() }

func (env TestEnv) ExecSQL(cmd string) {
	_, err := env.Endpoint.GetDB().Exec(cmd)
	if !assert.NoError(env.t, err) {
		panic(err)
	}
}

func Run(t *testing.T, test func(env *TestEnv)) {
	ctx := context.Background()
	config := repo.NewConfig()
	factory := repo.NewEndpointFactory(config, loggerProvider)
	endpoint := factory.CreateEndpoint(ctx)
	ctx = repo.NewContextWithEndpoint(ctx, endpoint)

	env := &TestEnv{
		Context:        ctx,
		t:              t,
		Endpoint:       endpoint,
		LoggerProvider: loggerProvider,
	}

	env.ExecSQL("delete from kmmSecurities")
	env.ExecSQL("delete from kmmAccounts")
	env.ExecSQL("delete from kmmTransactions")
	env.ExecSQL("delete from kmmSplits")

	test(env)
}
