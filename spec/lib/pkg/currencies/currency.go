package speccurrencies

import "github.com/smart-libs/go-kmymoney/spec/lib/pkg/securities"

type (
	Currency interface {
		ISOCode() string
		Name() string
		Type() specsecurities.Type
	}

	baseCurrency struct {
		isoCode string
		name    string
		secType specsecurities.Type
	}
)

func (b baseCurrency) ISOCode() string           { return b.isoCode }
func (b baseCurrency) Name() string              { return b.name }
func (b baseCurrency) Type() specsecurities.Type { return b.secType }

func NewCurrencyFromRepo(isoCode string, name string, secType specsecurities.Type) Currency {
	return baseCurrency{isoCode: isoCode, name: name, secType: secType}
}
