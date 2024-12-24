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
		security := specSecurities.NewBuilder().
			WithName("VALE3").
			WithSymbol("VALE3").
			WithType(specSecurities.Stock).
			WithPricePrecision(2).
			WithSmallestAccountFraction("100").
			WithTradingCurrency("BRL").
			WithRoundingMethod(7).
			Build()
		svc := reposecurities.NewSaveSecurityService(env.LoggerProvider)
		savedSecurity := svc.SaveSecurity(env.Context, security)
		assert.Equal(t, "E000001", savedSecurity.ID())

		savedSecurity = svc.SaveSecurity(env.Context, security)
		assert.Equal(t, "E000002", savedSecurity.ID())
	})

}
