package specsecurities

import (
	"context"
)

type (
	FindSecurityBySymbolService interface {
		FindSecurityBySymbol(ctx context.Context, symbol string) Security
	}
)
