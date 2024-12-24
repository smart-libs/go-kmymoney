package reposecurities

import (
	"context"
	_ "embed"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	specSecurities "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"
)

type (
	FindSecurityBySymbolService struct {
		repo.FindOne
	}
)

//go:embed svc_find_by_symbol.sql
var sqlFindBySymbolSQL string

func (f FindSecurityBySymbolService) FindSecurityBySymbol(ctx context.Context, symbol string) specSecurities.Security {
	var (
		id                      string
		name                    string
		secType                 specSecurities.Type
		smallestAccountFraction string
		pricePrecision          int
		tradingMarket           string
		tradingCurrency         string
		roundingMethod          int
	)
	found := f.FindOne.Run(ctx, []any{symbol}, []any{
		&id,
		&name,
		&secType,
		&smallestAccountFraction,
		&pricePrecision,
		&tradingMarket,
		&tradingCurrency,
		&roundingMethod,
	})
	if found {
		return specSecurities.NewBuilder().
			WithSymbol(symbol).
			WithID(id).
			WithName(name).
			WithType(secType).
			WithSmallestAccountFraction(smallestAccountFraction).
			WithPricePrecision(pricePrecision).
			WithTradingMarket(tradingMarket).
			WithTradingCurrency(tradingCurrency).
			WithRoundingMethod(roundingMethod).
			Build()
	}
	return nil
}

func NewFindSecurityBySymbolService(provider repo.LoggerProvider) specSecurities.FindSecurityBySymbolService {
	return FindSecurityBySymbolService{
		repo.FindOne{
			ErrMsg:         "FindSecurityBySymbolService: (symbol=%s)",
			SQLCmd:         sqlFindBySymbolSQL,
			LoggerProvider: provider,
		},
	}
}
