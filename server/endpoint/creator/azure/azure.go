package azure

type Credential struct {
	ClientID       string `json:"client_id"`
	SecretID       string `json:"secret_key"`
	SubscriptionID string `json:"subscription_id"`
	TenantID       string `json:"tenant_id"`
}

type Azure struct {
	Credential Credential `json:"credential"`
}
