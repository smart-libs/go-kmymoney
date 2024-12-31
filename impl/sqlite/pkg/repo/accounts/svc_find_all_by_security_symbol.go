package repoaccounts

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/accounts"
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/cursor"
)

type (
	FindAllAccountsBySecuritySymbolService struct {
		repo.FindAll
	}
)

//go:embed svc_find_all_by_security_symbol.sql
var sqlFindAllBySecuritySymbolCmd string

func NewFindAccountBySecuritySymbolService(provider repo.LoggerProvider) accounts.FindAllAccountsBySecuritySymbolService {
	return &FindAllAccountsBySecuritySymbolService{
		FindAll: repo.FindAll{
			ErrMsg:         "FindAllAccountsBySecuritySymbolService: failed to get accounts using securityID=[%s]",
			SQLCmd:         sqlFindAllBySecuritySymbolCmd,
			LoggerProvider: provider,
		},
	}
}

func (s FindAllAccountsBySecuritySymbolService) FindAllAccountsBySecuritySymbol(ctx context.Context, id string) cursor.Cursor[accounts.Account] {
	rows := s.FindAll.Run(ctx, id)
	return repo.NewSQLCursor(rows, func(givenRows repo.Rows) accounts.Account {
		return s.fetch(givenRows, id)
	})
}

func (s FindAllAccountsBySecuritySymbolService) fetch(rows repo.Rows, securityId string) accounts.Account {
	var (
		id               string
		institutionId    sql.NullString
		parentId         string
		lastReconciled   sql.NullTime
		lastModified     sql.NullTime
		openingDate      sql.NullTime
		accountNumber    sql.NullString
		accountType      accounts.Type
		accountName      sql.NullString
		description      sql.NullString
		currencyId       sql.NullString
		balance          sql.NullString
		transactionCount sql.NullInt64
	)

	err := rows.Scan(&id, &institutionId, &parentId, &lastReconciled, &lastModified, &openingDate, &accountNumber,
		&accountType, &accountName, &description, &currencyId, &balance, &transactionCount)
	if err != nil {
		panic(errors.Join(err, fmt.Errorf(s.ErrMsg, securityId)))
	}
	return accounts.NewBuilder().
		WithID(id).
		WithInstitutionId(institutionId.String).
		WithParentId(parentId).
		WithLastReconciled(lastReconciled.Time).
		WithLastModified(lastModified.Time).
		WithOpeningDate(openingDate.Time).
		WithAccountNumber(accountNumber.String).
		WithAccountType(accountType).
		WithAccountName(accountName.String).
		WithDescription(description.String).
		WithCurrencyId(currencyId.String).
		WithBalance(balance.String).
		WithTransactionCount(uint64(transactionCount.Int64)).
		Build()
}
