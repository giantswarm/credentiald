package azure

const (
	ClientIDKey       = "azure.azureoperator.clientid"
	ClientSecretKey   = "azure.azureoperator.clientsecret"
	SubscriptionIDKey = "azure.azureoperator.subscriptionid"
	TenantIDKey       = "azure.azureoperator.tenantid"
)

type Request struct {
	ClientID       string
	SecretID       string
	SubscriptionID string
	TenantID       string
}
