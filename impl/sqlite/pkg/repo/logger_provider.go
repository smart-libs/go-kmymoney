package repo

import (
	"context"
	"log/slog"
)

type (
	LoggerProvider = func(ctx context.Context) *slog.Logger
)
