package repocurrencies

import (
	"context"
	_ "embed"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	specCurrencies "github.com/smart-libs/go-kmymoney/spec/lib/pkg/currencies"
	specSecurities "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"
)

type (
	GetCurrencyByISOCodeService struct {
		repo.FindOne
	}
)

//go:embed svc_get_by_iso_code.sql
var sqlCmd string

func (g GetCurrencyByISOCodeService) GetCurrencyByISOCode(ctx context.Context, isoCode string) specCurrencies.Currency {
	var (
		name         string
		securityType specSecurities.Type
	)
	if g.FindOne.Run(ctx, []any{isoCode}, []any{&name, &securityType}) {
		return specCurrencies.NewCurrencyFromRepo(isoCode, name, securityType)
	}
	return nil
}

func NewGetCurrencyByISOCodeService(loggerProvider repo.LoggerProvider) specCurrencies.GetCurrencyByISOCodeService {
	return GetCurrencyByISOCodeService{
		repo.FindOne{
			ErrMsg:         "GetCurrencyByISOCodeService: (isoCode=%s)",
			SQLCmd:         sqlCmd,
			LoggerProvider: loggerProvider,
		},
	}
}
