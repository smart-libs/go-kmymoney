package repo

import (
	"context"
	"log/slog"
)

type (
	// LoggerProvider is a type alias of a function able to returns a *slog.Logger object from the context.
	// This function shall be provided by the component user and it is defined here for supporting tests.
	// Keep in mind this is an alias not a string type, you don't have tpo use repo.LoggerProvider.
	// See: test/integration/test_env.log
	LoggerProvider = func(ctx context.Context) *slog.Logger
)
