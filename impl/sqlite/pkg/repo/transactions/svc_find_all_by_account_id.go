package repotransactions

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/cursor"
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/transactions"
	"time"
)

type (
	FindAllSplitTransactionsByAccountIDService struct {
		repo.FindAll
	}
)

//go:embed svc_find_all_by_account_id.sql
var sqlFindAllByAccountIDCmd string

func NewFindAllSplitTransactionsByAccountIDService(provider repo.LoggerProvider) transactions.FindAllSplitTransactionsByAccountIDService {
	return &FindAllSplitTransactionsByAccountIDService{
		repo.FindAll{
			ErrMsg:         "FindAllSplitTransactionsByAccountIDService failed using accountID=[%s]",
			SQLCmd:         sqlFindAllByAccountIDCmd,
			LoggerProvider: provider,
		},
	}
}

func (f FindAllSplitTransactionsByAccountIDService) FindAllSplitTransactionsByAccountID(ctx context.Context, accountID string) cursor.Cursor[transactions.TransactionSplit] {
	rows := f.FindAll.Run(ctx, accountID)
	return repo.NewSQLCursor(rows, func(givenRows repo.Rows) transactions.TransactionSplit {
		return f.fetch(givenRows, accountID)
	})
}

func (f FindAllSplitTransactionsByAccountIDService) fetch(rows repo.Rows, accountID string) transactions.TransactionSplit {
	var (
		id              string
		txType          string
		postDate        time.Time
		memo            sql.NullString
		entryDate       sql.NullTime
		currencyID      sql.NullString
		bankID          sql.NullString
		splitID         int
		payeeID         sql.NullString
		reconcileDate   sql.NullTime
		action          sql.NullString
		reconcileFlag   string
		value           string
		valueFormatted  string
		shares          string
		sharesFormatted string
		price           sql.NullString
		priceFormatted  sql.NullString
		costCenterID    sql.NullString
		checkNumber     sql.NullString
		splitPostDate   time.Time
		splitBankID     sql.NullString
		splitMemo       sql.NullString
	)

	err := rows.Scan(&id, &txType, &postDate, &memo, &entryDate, &currencyID, &bankID,
		&splitID, &payeeID, &reconcileDate, &action, &reconcileFlag, &value, &valueFormatted,
		&shares, &sharesFormatted, &price, &priceFormatted, &costCenterID, &checkNumber, &splitPostDate,
		&splitBankID, &splitMemo)
	if err != nil {
		panic(errors.Join(err, fmt.Errorf(f.ErrMsg, accountID)))
	}

	tBuilder := transactions.NewTransactionBuilder().
		WithID(id).
		WithTxType(txType).
		WithPostDate(postDate)

	if memo.Valid {
		tBuilder.WithMemo(memo.String)
	}
	if entryDate.Valid {
		tBuilder.WithEntryDate(entryDate.Time)
	}
	if currencyID.Valid {
		tBuilder.WithCurrencyId(currencyID.String)
	}
	if bankID.Valid {
		tBuilder.WithBankId(bankID.String)
	}

	sBuilder := transactions.NewSplitBuilder().
		WithSplitID(splitID).
		WithReconcileFlag(reconcileFlag).
		WithValue(value).
		WithValueFormatted(valueFormatted).
		WithShares(shares).
		WithSharesFormatted(sharesFormatted).
		WithAccountId(accountID).
		WithPostDate(splitPostDate)
	if payeeID.Valid {
		sBuilder.WithPayeeID(payeeID.String)
	}
	if reconcileDate.Valid {
		sBuilder.WithReconcileDate(reconcileDate.Time)
	}
	if action.Valid {
		sBuilder.WithAction(action.String)
	}
	if price.Valid {
		sBuilder.WithPrice(price.String)
	}
	if priceFormatted.Valid {
		sBuilder.WithPriceFormatted(priceFormatted.String)
	}
	if costCenterID.Valid {
		sBuilder.WithCostCenterID(costCenterID.String)
	}
	if checkNumber.Valid {
		sBuilder.WithCheckNumber(checkNumber.String)
	}
	if splitBankID.Valid {
		sBuilder.WithBankID(splitBankID.String)
	}
	if splitMemo.Valid {
		sBuilder.WithMemo(splitMemo.String)
	}
	return transactions.NewTransactionSplit(tBuilder.Build(), sBuilder.Build())
}
