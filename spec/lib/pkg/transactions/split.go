package transactions

import "time"

type (
	Split interface {
		TransactionId() string
		TxType() string
		SplitID() int
		PayeeID() string
		ReconcileDate() *time.Time
		Action() *string
		ReconcileFlag() string
		Value() string
		ValueFormatted() string
		Shares() string
		SharesFormatted() string
		Price() *string
		PriceFormatted() *string
		Memo() *string
		AccountId() string
		CostCenterID() *string
		CheckNumber() *string
		PostDate() time.Time
		BankID() *string
	}

	baseSplit struct {
		transactionId   string
		txType          string
		splitID         int
		payeeID         string
		reconcileDate   *time.Time
		action          *string
		reconcileFlag   string
		value           string
		valueFormatted  string
		shares          string
		sharesFormatted string
		price           *string
		priceFormatted  *string
		memo            *string
		accountId       string
		costCenterID    *string
		checkNumber     *string
		postDate        time.Time
		bankID          *string
	}
)

func (s baseSplit) TransactionId() string {
	return s.transactionId
}

func (s baseSplit) TxType() string {
	return s.txType
}

func (s baseSplit) SplitID() int {
	return s.splitID
}

func (s baseSplit) PayeeID() string {
	return s.payeeID
}

func (s baseSplit) ReconcileDate() *time.Time {
	return s.reconcileDate
}

func (s baseSplit) Action() *string {
	return s.action
}

func (s baseSplit) ReconcileFlag() string {
	return s.reconcileFlag
}

func (s baseSplit) Value() string {
	return s.value
}

func (s baseSplit) ValueFormatted() string {
	return s.valueFormatted
}

func (s baseSplit) Shares() string {
	return s.shares
}

func (s baseSplit) SharesFormatted() string {
	return s.sharesFormatted
}

func (s baseSplit) Price() *string {
	return s.price
}

func (s baseSplit) PriceFormatted() *string {
	return s.priceFormatted
}

func (s baseSplit) Memo() *string {
	return s.memo
}

func (s baseSplit) AccountId() string {
	return s.accountId
}

func (s baseSplit) CostCenterID() *string {
	return s.costCenterID
}

func (s baseSplit) CheckNumber() *string {
	return s.checkNumber
}

func (s baseSplit) PostDate() time.Time {
	return s.postDate
}

func (s baseSplit) BankID() *string {
	return s.bankID
}
