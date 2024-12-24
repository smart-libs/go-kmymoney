package securities

import (
	reposecurities "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/securities"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/test/integration"
	specSecurities "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveSecurityService(t *testing.T) {
	integration.Run(t, func(env *integration.TestEnv) {
		symbol := "VALE3"
		security := specSecurities.NewBuilder().
			WithName(symbol).
			WithSymbol(symbol).
			WithType(specSecurities.Stock).
			WithPricePrecision(2).
			WithSmallestAccountFraction("100").
			WithTradingCurrency("BRL").
			WithRoundingMethod(7).
			Build()
		svc := reposecurities.NewSaveSecurityService(env.LoggerProvider)
		t.Run("I can create a new Security", func(t *testing.T) {
			savedSecurity := svc.SaveSecurity(env.Context, security)
			assert.Equal(t, "E000001", savedSecurity.ID())
		})
		t.Run("I can get a Security by ID", func(t *testing.T) {
			finder := reposecurities.NewFindSecurityBySymbolService(env.LoggerProvider)
			securityFound := finder.FindSecurityBySymbol(env.Context, symbol)
			if assert.NotNil(t, securityFound) {
				assert.Equal(t, "E000001", securityFound.ID())
			}
		})
		t.Run("I can create a second Security by Symbol", func(t *testing.T) {
			savedSecurity := svc.SaveSecurity(env.Context, security)
			assert.Equal(t, "E000002", savedSecurity.ID())
		})
	})

}
