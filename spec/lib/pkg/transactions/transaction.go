package transactions

import "time"

type (
	Transaction interface {
		ID() string
		TxType() string
		PostDate() *time.Time
		Memo() *string
		EntryDate() *time.Time
		CurrencyId() *string
		BankId() *string
	}

	baseTransaction struct {
		id         string
		txType     string
		postDate   *time.Time
		memo       *string
		entryDate  *time.Time
		currencyId *string
		bankId     *string
	}
)

func (t baseTransaction) ID() string {
	return t.id
}

func (t baseTransaction) TxType() string {
	return t.txType
}

func (t baseTransaction) PostDate() *time.Time {
	return t.postDate
}

func (t baseTransaction) Memo() *string {
	return t.memo
}

func (t baseTransaction) EntryDate() *time.Time {
	return t.entryDate
}

func (t baseTransaction) CurrencyId() *string {
	return t.currencyId
}

func (t baseTransaction) BankId() *string {
	return t.bankId
}
