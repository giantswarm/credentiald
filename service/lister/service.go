// Package lister offers a service to retrieve credentials for an organization.
package lister

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// the namespace in which we store credentiald secrets
	kubernetesCredentialNamespace = "giantswarm"

	// the selector we use to retrieve credentials
	kubernetesLabelSelectorMask = "app=credentiald,giantswarm.io/organization=%s"

	providerAWS   = "aws"
	providerAzure = "azure"

	// We use these data keys to detect the provider from a secret.
	providerAWSDetectionKey   = "aws.admin.arn"
	providerAzureDetectionKey = "azure.azureoperator.subscriptionid"

	defaultCredentialName = "credential-default" //nolint:gosec
)

// Config is the service configuration data structure.
type Config struct {
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger
}

// Service is our actual service.
type Service struct {
	k8sClient kubernetes.Interface
	logger    micrologger.Logger
}

// New creates a new lister service based on a configutation.
func New(config Config) (*Service, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}

	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	service := &Service{
		k8sClient: config.K8sClient,
		logger:    config.Logger,
	}

	return service, nil
}

// List returns metadata on all credentials found.
func (c *Service) List(ctx context.Context, request Request) ([]Response, error) {
	timer := prometheus.NewTimer(listTime)
	defer timer.ObserveDuration()

	c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("listing secrets for organization %#q", request.Organization))

	selector := fmt.Sprintf(kubernetesLabelSelectorMask, request.Organization)
	credentialList, err := c.k8sClient.CoreV1().Secrets(kubernetesCredentialNamespace).List(ctx, metav1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		kubernetesListErrorTotal.Inc()
		return nil, microerror.Mask(err)
	}

	resp := []Response{}

	for _, credential := range credentialList.Items {
		item := Response{}

		// We never expose the credential-default secret.
		if credential.Name == defaultCredentialName {
			continue
		}

		// get ID from name (ex: 'credential-15iv58')
		{
			parts := strings.Split(credential.Name, "-")
			if len(parts) == 2 {
				item.ID = parts[1]
			} else {
				c.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("ignoring listed secret with invalid name %#q", credential.Name))
			}
		}

		// get provider from content
		if _, ok := credential.Data[providerAWSDetectionKey]; ok {
			item.Provider = providerAWS
		} else if _, ok := credential.Data[providerAzureDetectionKey]; ok {
			item.Provider = providerAzure
		}

		// get payload
		if item.Provider == providerAWS {
			item.AWS.Roles.Admin = string(credential.Data["aws.admin.arn"])
			item.AWS.Roles.AWSOperator = string(credential.Data["aws.awsoperator.arn"])
		} else if item.Provider == providerAzure {
			item.Azure.SubscriptionID = string(credential.Data["azure.azureoperator.subscriptionid"])
			item.Azure.TenantID = string(credential.Data["azure.azureoperator.tenantid"])
			item.Azure.ClientID = string(credential.Data["azure.azureoperator.clientid"])
		}
		resp = append(resp, item)
	}

	c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("listed secrets for organization %#q", request.Organization))

	return resp, nil
}
