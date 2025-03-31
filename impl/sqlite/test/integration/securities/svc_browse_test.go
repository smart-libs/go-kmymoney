package securities

import (
	"context"
	"errors"
	reposecurities "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/securities"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/test/integration"
	specsecurities "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrowseSecuritiesService(t *testing.T) {
	integration.Run(t, func(env *integration.TestEnv) {
		env.ExecSQL(`
INSERT INTO kmmSecurities(id, name, symbol, type, typeString, smallestAccountFraction, pricePrecision, tradingMarket, tradingCurrency, roundingMethod)
VALUES ('E000001', 'CNES11', 'CNES11', 0, 'Stock', '10000', 4, 'XX', 'BRL', 7);
INSERT INTO kmmSecurities(id, name, symbol, type, typeString, smallestAccountFraction, pricePrecision, tradingMarket, tradingCurrency, roundingMethod)
VALUES ('E000002', 'CNES12', 'CNES12', 0, 'Stock', '10000', 4, 'XX', 'BRL', 7);
`)
		svc := reposecurities.NewBrowseSecuritiesService(env.LoggerProvider)

		t.Run("I can browse Securities", func(t *testing.T) {
			invoked := 0
			req := specsecurities.BrowseSecuritiesRequest{
				OnSecurity: func(_ context.Context, security specsecurities.Security) (bool, error) {
					invoked++
					if invoked == 1 {
						assert.Equal(t, "E000001", security.ID())
						assert.Equal(t, "CNES11", security.Symbol())
						assert.Equal(t, "CNES11", security.Name())
					} else {
						assert.Equal(t, "E000002", security.ID())
						assert.Equal(t, "CNES12", security.Symbol())
						assert.Equal(t, "CNES12", security.Name())
					}
					assert.Equal(t, specsecurities.Stock, security.Type())
					assert.Equal(t, "10000", security.SmallestAccountFraction())
					assert.Equal(t, 4, security.PricePrecision())
					assert.Equal(t, "XX", security.TradingMarket())
					assert.Equal(t, "BRL", security.TradingCurrency())
					assert.Equal(t, 7, security.RoundingMethod())
					return true, nil
				},
			}
			svc.BrowseSecurities(env.Context, req)
			assert.Equal(t, 2, invoked, "OnSecurity lambda should have been invoked twice")
		})
		t.Run("I can stop browsing Securities", func(t *testing.T) {
			invoked := 0
			req := specsecurities.BrowseSecuritiesRequest{
				OnSecurity: func(_ context.Context, security specsecurities.Security) (bool, error) {
					invoked++
					assert.Equal(t, "E000001", security.ID())
					assert.Equal(t, "CNES11", security.Symbol())
					assert.Equal(t, "CNES11", security.Name())
					assert.Equal(t, specsecurities.Stock, security.Type())
					assert.Equal(t, "10000", security.SmallestAccountFraction())
					assert.Equal(t, 4, security.PricePrecision())
					assert.Equal(t, "XX", security.TradingMarket())
					assert.Equal(t, "BRL", security.TradingCurrency())
					assert.Equal(t, 7, security.RoundingMethod())
					return false, nil
				},
			}
			svc.BrowseSecurities(env.Context, req)
			assert.Equal(t, 1, invoked, "OnSecurity lambda should have been invoked once")
		})
		t.Run("Browse panics if the OnSecurity lambda returns error", func(t *testing.T) {
			err := errors.New("some error")
			req := specsecurities.BrowseSecuritiesRequest{
				OnSecurity: func(_ context.Context, security specsecurities.Security) (bool, error) {
					return true, err
				},
			}
			assert.PanicsWithError(t, "some error\nBrowseSecuritiesService: stopped by error", func() {
				svc.BrowseSecurities(env.Context, req)
			})
		})
	})
}
