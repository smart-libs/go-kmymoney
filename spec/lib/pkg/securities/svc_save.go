package specsecurities

import "context"

type (
	SaveSecurityService interface {
		SaveSecurity(ctx context.Context, security Security) Security
	}
)
