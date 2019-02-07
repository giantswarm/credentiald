package lister

// Response is the response data structure our lister service will return.
type Response struct {
	ID       string
	Provider string
	AWS      ResponseAWS
	Azure    ResponseAzure
}

type ResponseAWS struct {
	Roles ResponseAWSRoles
}

type ResponseAWSRoles struct {
	Admin       string
	AWSOperator string
}

type ResponseAzure struct {
	ClientID       string
	SubscriptionID string
	TenantID       string
}
