package transactions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionBuilder(t *testing.T) {
	t.Run("Build complete transaction", func(t *testing.T) {
		postDate := time.Now()
		entryDate := time.Now().Add(-24 * time.Hour)
		transaction := NewTransactionBuilder().
			WithID("TX123").
			WithTxType("Credit").
			WithPostDate(postDate).
			WithMemo("Sample transaction").
			WithEntryDate(entryDate).
			WithCurrencyId("USD").
			WithBankId("BANK123").
			Build()

		assert.Equal(t, "TX123", transaction.ID())
		assert.Equal(t, "Credit", transaction.TxType())
		assert.NotNil(t, transaction.PostDate())
		assert.Equal(t, postDate, *transaction.PostDate())
		assert.NotNil(t, transaction.Memo())
		assert.Equal(t, "Sample transaction", *transaction.Memo())
		assert.NotNil(t, transaction.EntryDate())
		assert.Equal(t, entryDate, *transaction.EntryDate())
		assert.NotNil(t, transaction.CurrencyId())
		assert.Equal(t, "USD", *transaction.CurrencyId())
		assert.NotNil(t, transaction.BankId())
		assert.Equal(t, "BANK123", *transaction.BankId())
	})

	t.Run("Build partial transaction", func(t *testing.T) {
		transaction := NewTransactionBuilder().
			WithID("TX124").
			WithTxType("Debit").
			Build()

		assert.Equal(t, "TX124", transaction.ID())
		assert.Equal(t, "Debit", transaction.TxType())
		assert.Nil(t, transaction.PostDate())
		assert.Nil(t, transaction.Memo())
		assert.Nil(t, transaction.EntryDate())
		assert.Nil(t, transaction.CurrencyId())
		assert.Nil(t, transaction.BankId())
	})
}
