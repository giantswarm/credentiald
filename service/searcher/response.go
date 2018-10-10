package searcher

// Response is the response data structure our searcher service will return.
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
		Credential struct {
			ClientID       string
			SubscriptionID string
			TenantID       string
		}
	}
}
