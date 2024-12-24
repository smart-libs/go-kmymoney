package speccurrencies

import "context"

type (
	GetCurrencyByISOCodeService interface {
		GetCurrencyByISOCode(ctx context.Context, isoCode string) Currency
	}
)
