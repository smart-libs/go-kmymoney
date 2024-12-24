package specsecurities

type (
	Builder struct {
		baseSecurity baseSecurity
	}
)

func NewBuilder() *Builder { return &Builder{} }

func (b *Builder) WithSecurity(s Security) *Builder {
	b.baseSecurity = s.(baseSecurity)
	return b
}

func (b *Builder) WithID(id string) *Builder {
	b.baseSecurity.id = id
	return b
}

func (b *Builder) WithName(name string) *Builder {
	b.baseSecurity.name = name
	return b
}

func (b *Builder) WithSymbol(symbol string) *Builder {
	b.baseSecurity.symbol = symbol
	return b
}

func (b *Builder) WithType(t Type) *Builder {
	b.baseSecurity.secType = t
	return b
}

func (b *Builder) WithSmallestAccountFraction(s string) *Builder {
	b.baseSecurity.smallestAccountFraction = s
	return b
}

func (b *Builder) WithPricePrecision(p int) *Builder {
	b.baseSecurity.pricePrecision = p
	return b
}

func (b *Builder) WithTradingMarket(s string) *Builder {
	b.baseSecurity.tradingMarket = s
	return b
}

func (b *Builder) WithTradingCurrency(s string) *Builder {
	b.baseSecurity.tradingCurrency = s
	return b
}

func (b *Builder) WithRoundingMethod(i int) *Builder {
	b.baseSecurity.roundingMethod = i
	return b
}

func (b *Builder) Build() Security {
	return b.baseSecurity
}
