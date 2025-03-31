package kmymoneysqlite

import (
	"context"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	repoaccounts "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/accounts"
	repocurrencies "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/currencies"
	reposecurities "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/securities"
	repotransactions "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/transactions"
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
	reposecurities.NewBrowseSecuritiesService,
	repoaccounts.NewFindAccountBySecuritySymbolService,
	repotransactions.NewFindAllSplitTransactionsByAccountIDService,
	repo.NewConfig,
	repo.NewEndpointFactory,
))
