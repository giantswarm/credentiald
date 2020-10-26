// Package creator offers a service to fetch a list of credentials.
package creator

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/credentiald/v2/service/creator/aws"
	"github.com/giantswarm/credentiald/v2/service/creator/azure"
)

const (
	AppLabel          = "app"
	AppValue          = "credentiald"
	ManagedByLabel    = "giantswarm.io/managed-by"
	ManagedByValue    = "credentiald"
	OrganizationLabel = "giantswarm.io/organization"
	ServiceTypeLabel  = "giantswarm.io/service-type"
	ServiceTypeValue  = "system"

	ResourceCreatedMessageFormat = "A new set of credentials has been created with ID '%s'"
)

var (
	randomIDPool = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
)

type Config struct {
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger

	IDLength         int
	NameFormat       string
	SecretsNamespace string
}

type Service struct {
	k8sClient kubernetes.Interface
	logger    micrologger.Logger

	idLength         int
	nameFormat       string
	secretsNamespace string
}

func New(config Config) (*Service, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.IDLength == 0 {
		return nil, microerror.Maskf(invalidConfigError, "%T.IDLength must not be empty", config)
	}
	if config.NameFormat == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.NameFormat must not be empty", config)
	}
	if config.SecretsNamespace == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.SecretsNamespace must not be empty", config)
	}

	service := &Service{
		k8sClient: config.K8sClient,
		logger:    config.Logger,

		idLength:         config.IDLength,
		nameFormat:       config.NameFormat,
		secretsNamespace: config.SecretsNamespace,
	}

	return service, nil
}

func (s *Service) Create(ctx context.Context, request Request) (Response, error) {
	timer := prometheus.NewTimer(createTime)
	defer timer.ObserveDuration()

	s.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("received service request: %#v", request))

	// We allow only single credential secret per organization.
	existing, err := s.existing(ctx, request.Organization)
	if err != nil {
		return Response{}, microerror.Mask(err)
	}
	if len(existing) > 0 {
		return Response{}, microerror.Maskf(alreadyExistsError, "found %d credential secrets for organization %q", len(existing), request.Organization)
	}

	credentialID := s.generateCredentialID()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf(s.nameFormat, credentialID),
			Namespace: s.secretsNamespace,
			Labels: map[string]string{
				AppLabel:          AppValue,
				ManagedByLabel:    ManagedByValue,
				OrganizationLabel: request.Organization,
				ServiceTypeLabel:  ServiceTypeValue,
			},
		},
	}

	switch request.Provider {
	case "aws":
		secret.Data = map[string][]byte{
			aws.AdminArnKey:       []byte(request.AWS.AdminARN),
			aws.AwsoperatorArnKey: []byte(request.AWS.AwsOperatorARN),
		}
	case "azure":
		secret.Data = map[string][]byte{
			azure.ClientIDKey:       []byte(request.Azure.ClientID),
			azure.ClientSecretKey:   []byte(request.Azure.SecretID),
			azure.SubscriptionIDKey: []byte(request.Azure.SubscriptionID),
			azure.TenantIDKey:       []byte(request.Azure.TenantID),
		}
	default:
		return Response{}, microerror.Maskf(invalidProviderError, "%q provider is not supported", request.Provider)
	}

	_, err = s.k8sClient.CoreV1().Secrets(secret.Namespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		kubernetesCreateErrorTotal.Inc()
		return Response{}, microerror.Mask(err)
	}

	// Check if another secret wasn't created in the meantime. If so delete
	// ours. In worst case scenario there won't be any and the request will
	// have to be replayed.
	existing, err = s.existing(ctx, request.Organization)
	if err != nil {
		return Response{}, microerror.Mask(err)
	}
	if len(existing) > 1 {
		err := s.k8sClient.CoreV1().Secrets(secret.Namespace).Delete(ctx, secret.Name, metav1.DeleteOptions{})
		if err != nil {
			kubernetesCreateErrorTotal.Inc()
			return Response{}, microerror.Mask(err)
		}
		return Response{}, microerror.Maskf(alreadyExistsError, "detected race when creating credential secret for organization %q", request.Organization)
	}

	response := Response{
		Code:    server.CodeResourceCreated,
		Message: fmt.Sprintf(ResourceCreatedMessageFormat, credentialID),

		CredentialID: credentialID,
		Organization: request.Organization,
	}

	return response, nil
}

func (s *Service) existing(ctx context.Context, organization string) ([]*corev1.Secret, error) {
	selectors := []string{
		ManagedByLabel + "=" + ManagedByValue,
		OrganizationLabel + "=" + organization,
	}

	resp, err := s.k8sClient.CoreV1().Secrets(s.secretsNamespace).List(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(selectors, ","),
	})
	if err != nil {
		kubernetesCreateErrorTotal.Inc()
		return nil, microerror.Mask(err)
	}

	// Convert to pointers avoid future type confustions.
	secrets := make([]*corev1.Secret, len(resp.Items))
	for i, s := range resp.Items {
		secrets[i] = &s
	}

	return secrets, nil
}

// generateCredentialID provides an ID suitable for credentials.
func (s *Service) generateCredentialID() string {
	pattern := regexp.MustCompile("^[a-z]+$")
	for {
		b := make([]rune, s.idLength)

		for i := range b {
			b[i] = randomIDPool[rand.Intn(len(randomIDPool))]
		}

		id := string(b)

		// Don't use an ID if it is numbers only.
		if _, err := strconv.Atoi(id); err == nil {
			continue
		}

		// Don't use an ID if its letters only.
		if pattern.MatchString(id) {
			continue
		}

		return id
	}
}
