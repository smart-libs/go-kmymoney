package securities

import (
	reposecurities "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/securities"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/test/integration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextSecurityIDService(t *testing.T) {
	integration.Run(t, func(env *integration.TestEnv) {
		svc := reposecurities.NewGetNextSecurityIDService(env.LoggerProvider)
		assert.Equal(t, "E000001", svc.NextSecurityID(env.Context))
	})
}
