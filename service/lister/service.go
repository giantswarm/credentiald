// Package lister offers a service to retrieve credentials for an organization.
package lister

import (
	"fmt"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// the namespace in which we store credentiald secrets
	kubernetesCredentialNamespace = "giantswarm"

	// the selector we use to retrieve credentials
	kubernetesLabelSelectorMask = "app=credentiald,giantswarm.io/organization=%s"

	gaugeValue = float64(1)
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
func (c *Service) List(request Request) ([]*Response, error) {
	c.logger.Log("level", "debug", "message", fmt.Sprintf("listing secrets for organization %s", request.Organization))

	selector := fmt.Sprintf(kubernetesLabelSelectorMask, request.Organization)
	credentialList, err := c.k8sClient.CoreV1().Secrets(kubernetesCredentialNamespace).List(metav1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		c.logger.Log("level", "error", "message", "could not list secrets", "stack", fmt.Sprintf("%#v", err))
	}

	resp := []*Response{}

	for _, credential := range credentialList.Items {
		// get ID from name (ex: 'credential-15iv58')
		parts := strings.Split(credential.Name, "-")
		if len(parts) == 2 {
			resp = append(resp, &Response{ID: parts[1]})
		} else {
			c.logger.Log("level", "error", "message", fmt.Sprintf("Invalid secret name found: %q", credential.Name))
		}
	}

	c.logger.Log("level", "debug", "message", "finished listing secrets")

	return resp, nil
}
