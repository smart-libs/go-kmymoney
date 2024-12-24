package kmymoneysqlite

import (
	repocurrencies "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/currencies"
	reposecurities "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/securities"
	"go.uber.org/fx"
)

var Module = fx.Module("go-kmymoney/sqlite", fx.Provide(
	repocurrencies.NewGetCurrencyByISOCodeService,
	reposecurities.NewGetNextSecurityIDService,
	reposecurities.NewSaveSecurityService,
))
