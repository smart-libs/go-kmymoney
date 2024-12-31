package accounts

import (
	"context"
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/cursor"
)

type (
	FindAllAccountsBySecuritySymbolService interface {
		FindAllAccountsBySecuritySymbol(ctx context.Context, id string) cursor.Cursor[Account]
	}
)
