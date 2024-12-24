package currencies

import (
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/currencies"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/test/integration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByAliasService(t *testing.T) {
	integration.Run(t, func(env *integration.TestEnv) {
		svc := repocurrencies.NewGetCurrencyByISOCodeService(env.LoggerProvider)
		assert.NotNil(t, svc.GetCurrencyByISOCode(env.Context, "BRL"))
	})
}
