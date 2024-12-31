package accounts

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	t.Run("Build complete baseAccount", func(t *testing.T) {
		builder := NewBuilder()
		now := time.Now()
		account := builder.
			WithID("123").
			WithInstitutionId("456").
			WithParentId("789").
			WithLastReconciled(now).
			WithLastModified(now).
			WithOpeningDate(now.AddDate(-1, 0, 0)).
			WithAccountNumber("ACC123").
			WithAccountType(Savings).
			WithAccountName("Test Account").
			WithDescription("A test account description").
			WithCurrencyId("USD").
			WithBalance("1000.50").
			WithTransactionCount(42).
			Build()

		assert.Equal(t, "123", account.ID())
		assert.Equal(t, "456", account.InstitutionId())
		assert.Equal(t, "789", account.ParentId())
		assert.Equal(t, now, account.LastReconciled())
		assert.Equal(t, now, account.LastModified())
		assert.Equal(t, now.AddDate(-1, 0, 0), account.OpeningDate())
		assert.Equal(t, "ACC123", account.AccountNumber())
		assert.Equal(t, "Savings", account.AccountType().String())
		assert.Equal(t, false, account.IsStockAccount())
		assert.Equal(t, "Test Account", account.AccountName())
		assert.Equal(t, "A test account description", account.Description())
		assert.Equal(t, "USD", account.CurrencyId())
		assert.Equal(t, "1000.50", account.Balance())
		assert.Equal(t, uint64(42), account.TransactionCount())
	})

	t.Run("Build partial baseAccount", func(t *testing.T) {
		builder := NewBuilder()
		account := builder.
			WithID("partial").
			WithAccountName("Partial Account").
			Build()

		assert.Equal(t, "partial", account.ID())
		assert.Equal(t, "Partial Account", account.AccountName())
		assert.Equal(t, "", account.InstitutionId())
		assert.Equal(t, "", account.ParentId())
		assert.Equal(t, time.Time{}, account.LastReconciled())
		assert.Equal(t, time.Time{}, account.LastModified())
		assert.Equal(t, time.Time{}, account.OpeningDate())
		assert.Equal(t, "", account.AccountNumber())
		assert.Equal(t, Unknown, account.AccountType())
		assert.Equal(t, false, account.IsStockAccount())
		assert.Equal(t, "", account.Description())
		assert.Equal(t, "", account.CurrencyId())
		assert.Equal(t, "", account.Balance())
		assert.Equal(t, uint64(0), account.TransactionCount())
	})
}
