package transactions

import "time"

type SplitBuilder struct {
	split baseSplit
}

func NewSplitBuilder() *SplitBuilder {
	return &SplitBuilder{}
}

func (b *SplitBuilder) WithTransactionId(transactionId string) *SplitBuilder {
	b.split.transactionId = transactionId
	return b
}

func (b *SplitBuilder) WithTxType(txType string) *SplitBuilder {
	b.split.txType = txType
	return b
}

func (b *SplitBuilder) WithSplitID(splitID int) *SplitBuilder {
	b.split.splitID = splitID
	return b
}

func (b *SplitBuilder) WithPayeeID(payeeID string) *SplitBuilder {
	b.split.payeeID = payeeID
	return b
}

func (b *SplitBuilder) WithReconcileDate(reconcileDate time.Time) *SplitBuilder {
	b.split.reconcileDate = &reconcileDate
	return b
}

func (b *SplitBuilder) WithAction(action string) *SplitBuilder {
	b.split.action = &action
	return b
}

func (b *SplitBuilder) WithReconcileFlag(reconcileFlag string) *SplitBuilder {
	b.split.reconcileFlag = reconcileFlag
	return b
}

func (b *SplitBuilder) WithValue(value string) *SplitBuilder {
	b.split.value = value
	return b
}

func (b *SplitBuilder) WithValueFormatted(valueFormatted string) *SplitBuilder {
	b.split.valueFormatted = valueFormatted
	return b
}

func (b *SplitBuilder) WithShares(shares string) *SplitBuilder {
	b.split.shares = shares
	return b
}

func (b *SplitBuilder) WithSharesFormatted(sharesFormatted string) *SplitBuilder {
	b.split.sharesFormatted = sharesFormatted
	return b
}

func (b *SplitBuilder) WithPrice(price string) *SplitBuilder {
	b.split.price = &price
	return b
}

func (b *SplitBuilder) WithPriceFormatted(priceFormatted string) *SplitBuilder {
	b.split.priceFormatted = &priceFormatted
	return b
}

func (b *SplitBuilder) WithMemo(memo string) *SplitBuilder {
	b.split.memo = &memo
	return b
}

func (b *SplitBuilder) WithAccountId(accountId string) *SplitBuilder {
	b.split.accountId = accountId
	return b
}

func (b *SplitBuilder) WithCostCenterID(costCenterID string) *SplitBuilder {
	b.split.costCenterID = &costCenterID
	return b
}

func (b *SplitBuilder) WithCheckNumber(checkNumber string) *SplitBuilder {
	b.split.checkNumber = &checkNumber
	return b
}

func (b *SplitBuilder) WithPostDate(postDate time.Time) *SplitBuilder {
	b.split.postDate = postDate
	return b
}

func (b *SplitBuilder) WithBankID(bankID string) *SplitBuilder {
	b.split.bankID = &bankID
	return b
}

func (b *SplitBuilder) Build() Split {
	return b.split
}
