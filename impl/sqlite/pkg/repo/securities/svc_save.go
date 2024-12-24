package reposecurities

import (
	"context"
	_ "embed"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	specSecurities "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"
)

type (
	SaveSecurityService struct {
		repo.Save
		GetNextSecurityIDService
	}
)

//go:embed svc_save.sql
var sqlSaveCmd string

func (s SaveSecurityService) SaveSecurity(ctx context.Context, security specSecurities.Security) specSecurities.Security {
	if security.ID() == "" {
		id := s.NextSecurityID(ctx)
		newSecurity := specSecurities.NewBuilder().WithSecurity(security).WithID(id).Build()
		return s.SaveSecurity(ctx, newSecurity)
	}
	s.Save.Run(ctx,
		security.ID(),
		security.Name(),
		security.Symbol(),
		security.Type(),
		security.Type().String(),
		security.SmallestAccountFraction(),
		security.PricePrecision(),
		security.TradingMarket(),
		security.TradingCurrency(),
		security.RoundingMethod(),
	)
	return security
}

func NewSaveSecurityService(loggerProvider repo.LoggerProvider) specSecurities.SaveSecurityService {
	return SaveSecurityService{
		Save: repo.Save{
			ErrMsg:         "SaveSecurityService: (id=%s,name=%s,symbol=%s,type=%d,smallestAccountFraction=%s,pricePrecision=%d,tradingMarket=%s,tradingCurrency=%s,roundingMethod=%d)",
			SQLCmd:         sqlSaveCmd,
			LoggerProvider: nil,
		},
		GetNextSecurityIDService: NewGetNextSecurityIDService(loggerProvider),
	}
}
