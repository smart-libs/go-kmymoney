package accounts

import "time"

type (
	Account interface {
		ID() string
		InstitutionId() string
		ParentId() string
		LastReconciled() time.Time
		LastModified() time.Time
		OpeningDate() time.Time
		AccountNumber() string
		AccountType() Type
		IsStockAccount() bool
		AccountName() string
		Description() string
		CurrencyId() string
		Balance() string
		BalanceFormatted() string
		TransactionCount() uint64
	}

	// baseAccount is the struct that implements the Account interface
	baseAccount struct {
		id               string
		institutionId    string
		parentId         string
		lastReconciled   time.Time
		lastModified     time.Time
		openingDate      time.Time
		accountNumber    string
		accountType      Type
		accountName      string
		description      string
		currencyId       string
		balance          string
		transactionCount uint64
	}
)

func (a baseAccount) ID() string {
	return a.id
}

func (a baseAccount) InstitutionId() string {
	return a.institutionId
}

func (a baseAccount) ParentId() string {
	return a.parentId
}

func (a baseAccount) LastReconciled() time.Time {
	return a.lastReconciled
}

func (a baseAccount) LastModified() time.Time {
	return a.lastModified
}

func (a baseAccount) OpeningDate() time.Time {
	return a.openingDate
}

func (a baseAccount) AccountNumber() string {
	return a.accountNumber
}

func (a baseAccount) AccountType() Type {
	return a.accountType
}

func (a baseAccount) IsStockAccount() bool {
	return a.accountType == Stock
}

func (a baseAccount) AccountName() string {
	return a.accountName
}

func (a baseAccount) Description() string {
	return a.description
}

func (a baseAccount) CurrencyId() string {
	return a.currencyId
}

func (a baseAccount) Balance() string {
	return a.balance
}

func (a baseAccount) BalanceFormatted() string {
	// Example of simple formatting, adapt as necessary
	return a.balance + " USD"
}

func (a baseAccount) TransactionCount() uint64 {
	return a.transactionCount
}
