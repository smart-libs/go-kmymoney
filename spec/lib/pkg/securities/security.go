package specsecurities

type (
	Security interface {
		ID() string
		Name() string
		Symbol() string
		Type() Type
		SmallestAccountFraction() string
		PricePrecision() int
		TradingMarket() string
		TradingCurrency() string
		RoundingMethod() int
	}

	baseSecurity struct {
		id                      string
		name                    string
		symbol                  string
		secType                 Type
		smallestAccountFraction string
		pricePrecision          int
		tradingMarket           string
		tradingCurrency         string
		roundingMethod          int
	}
)

func (b baseSecurity) ID() string                      { return b.id }
func (b baseSecurity) Name() string                    { return b.name }
func (b baseSecurity) Symbol() string                  { return b.symbol }
func (b baseSecurity) Type() Type                      { return b.secType }
func (b baseSecurity) SmallestAccountFraction() string { return b.smallestAccountFraction }
func (b baseSecurity) PricePrecision() int             { return b.pricePrecision }
func (b baseSecurity) TradingMarket() string           { return b.tradingMarket }
func (b baseSecurity) TradingCurrency() string         { return b.tradingCurrency }
func (b baseSecurity) RoundingMethod() int             { return b.roundingMethod }
