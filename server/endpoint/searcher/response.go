package searcher

// Response is the response format our endpoint will return.
type Response struct {
	ID       string         `json:"id"`
	Provider string         `json:"provider"`
	AWS      *ResponseAWS   `json:"aws,omitempty"`
	Azure    *ResponseAzure `json:"azure,omitempty"`
}

// ResponseAWS is a type used by above Response.
type ResponseAWS struct {
	Roles *ResponseAWSRoles `json:"roles"`
}

// ResponseAWSRoles is a type used by above AWSData struct.
type ResponseAWSRoles struct {
	Admin       string `json:"admin"`
	AWSOperator string `json:"awsoperator"`
}

// ResponseAzure is a type used by above Response struct.
type ResponseAzure struct {
	ClientID       string `json:"client_id"`
	SubscriptionID string `json:"subscription_id"`
	TenantID       string `json:"tenant_id"`
}
