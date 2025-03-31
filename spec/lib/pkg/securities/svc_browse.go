package specsecurities

import "context"

type (
	BrowseSecuritiesRequest struct {
		// OnSecurity provides the lambda that will be invoked by each Security found in DB
		// The lambda must return true to continue the retrieval or false to stop
		// If the lambda returns error, then the retrieval is interrupted and the error is returned by
		// the service.
		OnSecurity func(context.Context, Security) (bool, error)
	}

	// BrowseSecuritiesService allows the caller to access all Security instances in DB.
	// The service panics if it fails
	BrowseSecuritiesService interface {
		BrowseSecurities(ctx context.Context, req BrowseSecuritiesRequest)
	}
)
