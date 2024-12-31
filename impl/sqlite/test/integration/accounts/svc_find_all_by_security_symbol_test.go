package currencies

import (
	repoaccounts "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/accounts"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/test/integration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindBySecurityIDService(t *testing.T) {
	integration.Run(t, func(env *integration.TestEnv) {
		env.ExecSQL(`
INSERT INTO kmmSecurities (id, name, symbol, type, typeString, smallestAccountFraction, pricePrecision, tradingMarket, tradingCurrency, roundingMethod)
VALUES ('E000063', 'CNES11', 'CNES11', 0, 'Stock', '10000', 4, '', 'BRL', 7);

INSERT INTO kmmAccounts (id, institutionId, parentId, lastReconciled, lastModified, openingDate, accountNumber, accountType, accountTypeString, isStockAccount, accountName, description, currencyId, balance, balanceFormatted, transactionCount)
VALUES ('A000260', null, 'A000136', null, '2024-12-30', '2024-04-18', null, '15', 'Stock', 'Y', 'CNES11', null, 'E000063', '460/1', '460', 10);

INSERT INTO kmmAccounts (id, institutionId, parentId, lastReconciled, lastModified, openingDate, accountNumber, accountType, accountTypeString, isStockAccount, accountName, description, currencyId, balance, balanceFormatted, transactionCount)
VALUES ('A000085', null, 'A000001', null, '2024-12-30', '2024-04-18', null, '15', 'Stock', 'Y', 'CNES11', null, 'E000063', '40/1', '40', 10);
`)
		svc := repoaccounts.NewFindAccountBySecuritySymbolService(env.LoggerProvider)
		// CNES11 is inserted into DB by migrations with 2 accounts
		cursor := svc.FindAllAccountsBySecuritySymbol(env.Context, "CNES11")
		if assert.NotNil(t, cursor) {
			if assert.True(t, cursor.Next()) {
				first := cursor.Fetch()
				if assert.NotNil(t, first) {
					assert.Equal(t, "A000260", first.ID())
				}
				if assert.True(t, cursor.Next()) {
					second := cursor.Fetch()
					if assert.NotNil(t, second) {
						assert.Equal(t, "A000085", second.ID())
					}
				}
				assert.False(t, cursor.Next())
			}
		}
	})
}
