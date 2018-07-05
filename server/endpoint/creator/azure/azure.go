package azure

type Credential struct {
	ClientID       string `json:"clientID"`
	SecretID       string `json:"secretID"`
	SubscriptionID string `json:"subscriptionID"`
	TenantID       string `json:"tenantID"`
}

type Azure struct {
	Credential Credential `json:"credential"`
}
