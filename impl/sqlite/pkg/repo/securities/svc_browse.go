package reposecurities

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	specSecurities "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"
)

type (
	BrowseSecuritiesService struct {
		repo.FindAll
	}
)

const (
	BrowseSecuritiesServiceName = "BrowseSecuritiesService"
)

//go:embed svc_browse.sql
var sqlBrowseSQL string

func (f BrowseSecuritiesService) BrowseSecurities(ctx context.Context, req specSecurities.BrowseSecuritiesRequest) {
	rows := f.Run(ctx)
	if rows == nil {
		panic(fmt.Errorf("%s: receives a nil rows object", BrowseSecuritiesServiceName))
	}
	if rows.Err() != nil {
		panic(errors.Join(rows.Err(), fmt.Errorf("%s: receives error from query DB", BrowseSecuritiesServiceName)))
	}
	defer func() { _ = rows.Close() }()

	var (
		id                      string
		name                    string
		notStop                 bool
		pricePrecision          int
		roundingMethod          int
		secType                 specSecurities.Type
		smallestAccountFraction string
		symbol                  string
		tradingMarket           string
		tradingCurrency         string
	)

	for rows.Next() {
		err := rows.Scan(
			&id, &symbol, &name, &secType, &smallestAccountFraction, &pricePrecision,
			&tradingMarket, &tradingCurrency, &roundingMethod,
		)
		if err != nil {
			panic(errors.Join(rows.Err(), fmt.Errorf("%s: receives error from Scan", BrowseSecuritiesServiceName)))
		}

		security := specSecurities.NewBuilder().
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
		notStop, err = req.OnSecurity(ctx, security)
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("%s: stopped by error", BrowseSecuritiesServiceName)))
		}
		if notStop {
			continue
		}
		break
	}
}

func NewBrowseSecuritiesService(provider repo.LoggerProvider) specSecurities.BrowseSecuritiesService {
	return BrowseSecuritiesService{
		repo.FindAll{
			ErrMsg:         BrowseSecuritiesServiceName + ": failed",
			SQLCmd:         sqlBrowseSQL,
			LoggerProvider: provider,
		},
	}
}
