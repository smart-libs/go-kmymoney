package transactions

import (
	repotransactions "github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo/transactions"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/test/integration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindBySecurityIDService(t *testing.T) {
	integration.Run(t, func(env *integration.TestEnv) {
		env.ExecSQL(`
INSERT INTO kmmTransactions (id, txType, postDate, memo, entryDate, currencyId, bankId) VALUES ('T000000000000001772', 'N', '2024-04-18', null, '2024-12-27', 'BRL', null);
INSERT INTO kmmSplits (transactionId, txType, splitId, payeeId, reconcileDate, action, reconcileFlag, value, valueFormatted, shares, sharesFormatted, price, priceFormatted, memo, accountId, costCenterId, checkNumber, postDate, bankId) VALUES ('T000000000000001772', 'N', 0, null, null, 'Add', '2', '0/1', '0', '40/1', '40.0000', null, null, null, 'A000085', null, null, '2024-04-18', '2024-04-18-4cebd50-1');

INSERT INTO kmmTransactions (id, txType, postDate, memo, entryDate, currencyId, bankId) VALUES ('T000000000000001780', 'N', '2024-04-29', null, '2024-12-27', 'BRL', null);
INSERT INTO kmmSplits (transactionId, txType, splitId, payeeId, reconcileDate, action, reconcileFlag, value, valueFormatted, shares, sharesFormatted, price, priceFormatted, memo, accountId, costCenterId, checkNumber, postDate, bankId) VALUES ('T000000000000001780', 'N', 2, null, null, 'Dividend', '2', '0/1', '0', '0/1', '0.0000', null, null, null, 'A000085', null, null, '2024-04-29', '2024-04-29-a4e6713-1');

`)
		svc := repotransactions.NewFindAllSplitTransactionsByAccountIDService(env.LoggerProvider)
		// CNES11 is inserted into DB by migrations with 2 accounts
		cursor := svc.FindAllSplitTransactionsByAccountID(env.Context, "A000085")
		if assert.NotNil(t, cursor) {
			if assert.True(t, cursor.Next()) {
				first := cursor.Fetch()
				if assert.NotNil(t, first) {
					assert.Equal(t, "T000000000000001772", first.GetTransaction().ID())
				}
				if assert.True(t, cursor.Next()) {
					second := cursor.Fetch()
					if assert.NotNil(t, second) {
						assert.Equal(t, "T000000000000001780", second.GetTransaction().ID())
					}
				}
				assert.False(t, cursor.Next())
			}
		}
	})
}
