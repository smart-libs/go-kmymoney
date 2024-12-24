package specsecurities

type (
	Type int
)

const (
	Stock      Type = 0
	MutualFund Type = 1
	Bond       Type = 2
	Currency   Type = 3
	None       Type = 4
)

func (t Type) String() string {
	switch t {
	case Stock:
		return "Stock"
	case MutualFund:
		return "MutualFund"
	case Bond:
		return "Bond"
	case Currency:
		return "Currency"
	case None:
		return "None"
	default:
		return "Unknown"
	}
}
