package azure

type AzureCredential struct {
	ClientID       string `json:"client_id"`
	SecretID       string `json:"secret_key"`
	SubscriptionID string `json:"subscription_id"`
	TenantID       string `json:"tenant_id"`
}

type Azure struct {
	Credential AzureCredential `json:"credential"`
}
