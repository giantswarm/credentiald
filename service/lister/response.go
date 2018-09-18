package lister

// Response is the response data structure our lister service will return.
type Response struct {
	ID       string
	Provider string
	AWS      struct {
		Roles struct {
			Admin       string
			AWSOperator string
		}
	}
	Azure struct {
		ClientID       string
		SubscriptionID string
		TenantID       string
	}
}
