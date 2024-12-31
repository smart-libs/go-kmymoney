package transactions

type (
	TransactionSplit interface {
		GetTransaction() Transaction
		GetSplit() Split
	}

	baseTransactionSplit struct {
		Transaction
		Split
	}
)

func (b baseTransactionSplit) GetTransaction() Transaction { return b.Transaction }
func (b baseTransactionSplit) GetSplit() Split             { return b.Split }

func NewTransactionSplit(t Transaction, s Split) TransactionSplit {
	return baseTransactionSplit{
		Transaction: t,
		Split:       s,
	}
}
