package transactions

import (
	"context"
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/cursor"
)

type (
	FindAllSplitTransactionsByAccountIDService interface {
		FindAllSplitTransactionsByAccountID(ctx context.Context, accountID string) cursor.Cursor[TransactionSplit]
	}
)
