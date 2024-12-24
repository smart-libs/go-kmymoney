package kmymoneysqlite

import (
	"context"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	repocurrencies "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/currencies"
	reposecurities "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/securities"
	"go.uber.org/fx"
	"log/slog"
)

// Use the functions below if there is no other to provide the objects returned

// NewLoggerProvider returns the default slog.Logger object
func NewLoggerProvider(_ context.Context) *slog.Logger { return slog.Default() }

var Module = fx.Module("go-kmymoney/sqlite", fx.Provide(
	repocurrencies.NewGetCurrencyByISOCodeService,
	reposecurities.NewGetNextSecurityIDService,
	reposecurities.NewSaveSecurityService,
	reposecurities.NewFindSecurityBySymbolService,
	repo.NewConfig,
	repo.NewEndpointFactory,
))
