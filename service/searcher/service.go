// Package searcher provides functionality to retrieve a single credential.
package searcher

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// the namespace in which we store credentiald secrets
	kubernetesCredentialNamespace = "giantswarm"

	// a label we expect to be set to the owner org
	kubernetesOrganizationLabel = "giantswarm.io/organization"

	providerAWS   = "aws"
	providerAzure = "azure"

	// We use these data keys to detect the provider from a secret.
	providerAWSDetectionKey   = "aws.admin.arn"
	providerAzureDetectionKey = "azure.azureoperator.subscriptionid"

	// we name our default credential "credential-default". This is the second part of it.
	defaultCredentialNameIDPart = "default"
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

// Search returns metadata about one credential.
func (c *Service) Search(request Request) (*Response, error) {
	ctx := context.Background()
	timer := prometheus.NewTimer(searchTime)
	defer timer.ObserveDuration()

	c.logger.Log("level", "debug", "message", fmt.Sprintf("searching secret for organization %s, ID %s", request.Organization, request.ID))

	// We never expose the credential-default secret. From the outside, this does not exist.
	if request.ID == defaultCredentialNameIDPart {
		return nil, microerror.Mask(secretNotFoundError)
	}

	name := "credential-" + request.ID
	credential, err := c.k8sClient.CoreV1().Secrets(kubernetesCredentialNamespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		kubernetesSearchErrorTotal.Inc()
		return nil, microerror.Mask(err)
	}

	// make sure the found credential really belongs to the organization indicated
	if ownerOrganization, ok := credential.Labels[kubernetesOrganizationLabel]; ok {
		if ownerOrganization != request.Organization {
			return nil, microerror.Mask(wrongOwnerOrganizationError)
		}
	} else {
		// no organization set
		return nil, microerror.Mask(wrongOwnerOrganizationError)
	}

	resp := &Response{
		ID: request.ID,
	}

	// get provider from content
	if _, ok := credential.Data[providerAWSDetectionKey]; ok {
		resp.Provider = providerAWS
	} else if _, ok := credential.Data[providerAzureDetectionKey]; ok {
		resp.Provider = providerAzure
	} else {
		return nil, microerror.Mask(secretInUnexpectedFormatError)
	}

	// get payload
	switch resp.Provider {
	case providerAWS:
		resp.AWS.Roles.Admin = string(credential.Data["aws.admin.arn"])
		resp.AWS.Roles.AWSOperator = string(credential.Data["aws.awsoperator.arn"])
	case providerAzure:
		resp.Azure.Credential.SubscriptionID = string(credential.Data["azure.azureoperator.subscriptionid"])
		resp.Azure.Credential.TenantID = string(credential.Data["azure.azureoperator.tenantid"])
		resp.Azure.Credential.ClientID = string(credential.Data["azure.azureoperator.clientid"])
	default:
		return nil, microerror.Mask(secretInUnexpectedFormatError)
	}

	c.logger.Log("level", "debug", "message", "finished getting secret")

	return resp, nil
}
