package accounts

import "time"

type (
	Builder struct {
		baseAccount baseAccount
	}
)

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithID(id string) *Builder {
	b.baseAccount.id = id
	return b
}

func (b *Builder) WithInstitutionId(institutionId string) *Builder {
	b.baseAccount.institutionId = institutionId
	return b
}

func (b *Builder) WithParentId(parentId string) *Builder {
	b.baseAccount.parentId = parentId
	return b
}

func (b *Builder) WithLastReconciled(lastReconciled time.Time) *Builder {
	b.baseAccount.lastReconciled = lastReconciled
	return b
}

func (b *Builder) WithLastModified(lastModified time.Time) *Builder {
	b.baseAccount.lastModified = lastModified
	return b
}

func (b *Builder) WithOpeningDate(openingDate time.Time) *Builder {
	b.baseAccount.openingDate = openingDate
	return b
}

func (b *Builder) WithAccountNumber(accountNumber string) *Builder {
	b.baseAccount.accountNumber = accountNumber
	return b
}

func (b *Builder) WithAccountType(accountType Type) *Builder {
	b.baseAccount.accountType = accountType
	return b
}

func (b *Builder) WithAccountName(accountName string) *Builder {
	b.baseAccount.accountName = accountName
	return b
}

func (b *Builder) WithDescription(description string) *Builder {
	b.baseAccount.description = description
	return b
}

func (b *Builder) WithCurrencyId(currencyId string) *Builder {
	b.baseAccount.currencyId = currencyId
	return b
}

func (b *Builder) WithBalance(balance string) *Builder {
	b.baseAccount.balance = balance
	return b
}

func (b *Builder) WithTransactionCount(transactionCount uint64) *Builder {
	b.baseAccount.transactionCount = transactionCount
	return b
}

func (b *Builder) Build() Account {
	return b.baseAccount
}
