package transactions

import "time"

type TransactionBuilder struct {
	transaction baseTransaction
}

func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{}
}

func (b *TransactionBuilder) WithID(id string) *TransactionBuilder {
	b.transaction.id = id
	return b
}

func (b *TransactionBuilder) WithTxType(txType string) *TransactionBuilder {
	b.transaction.txType = txType
	return b
}

func (b *TransactionBuilder) WithPostDate(postDate time.Time) *TransactionBuilder {
	b.transaction.postDate = &postDate
	return b
}

func (b *TransactionBuilder) WithMemo(memo string) *TransactionBuilder {
	b.transaction.memo = &memo
	return b
}

func (b *TransactionBuilder) WithEntryDate(entryDate time.Time) *TransactionBuilder {
	b.transaction.entryDate = &entryDate
	return b
}

func (b *TransactionBuilder) WithCurrencyId(currencyId string) *TransactionBuilder {
	b.transaction.currencyId = &currencyId
	return b
}

func (b *TransactionBuilder) WithBankId(bankId string) *TransactionBuilder {
	b.transaction.bankId = &bankId
	return b
}

func (b *TransactionBuilder) Build() Transaction {
	return b.transaction
}
