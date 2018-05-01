package creator

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	AppLabel          = "app"
	AppValue          = "credentiald"
	ManagedByLabel    = "giantswarm.io/managed-by"
	ManagedByValue    = "credentiald"
	OrganizationLabel = "giantswarm.io/organization"
	ServiceTypeLabel  = "giantswarm.io/service-type"
	ServiceTypeValue  = "system"

	// AwsAdminArnKey is the key in the Secret under which the ARN for the admin role is held.
	AwsAdminArnKey = "aws.admin.arn"
	// AwsAwsoperatorArnKey is the key in the Secret under which the ARN for the aws-operator role is held.
	AwsAwsoperatorArnKey = "aws.awsoperator.arn"

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
	s.logger.Log("level", "debug", "message", fmt.Sprintf("received service request: %#v", request))

	credentialID := s.generateCredentialID()

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf(s.nameFormat, credentialID),
			Labels: map[string]string{
				AppLabel:          AppValue,
				ManagedByLabel:    ManagedByValue,
				OrganizationLabel: request.Organization,
				ServiceTypeLabel:  ServiceTypeValue,
			},
		},
		Data: map[string][]byte{
			AwsAdminArnKey:       []byte(request.AdminARN),
			AwsAwsoperatorArnKey: []byte(request.AwsOperatorARN),
		},
	}

	_, err := s.k8sClient.CoreV1().Secrets(s.secretsNamespace).Create(secret)
	if err != nil {
		return Response{}, microerror.Mask(err)
	}

	response := Response{
		Code:    server.CodeResourceCreated,
		Message: fmt.Sprintf(ResourceCreatedMessageFormat, credentialID),

		CredentialID: credentialID,
		Organization: request.Organization,
	}

	return response, nil
}

// generateCredentialID provides an ID suitable for credentials.
func (s *Service) generateCredentialID() string {
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
		matched, err := regexp.MatchString("^[a-z]+$", id)
		if err == nil && matched == true {
			continue
		}

		return id
	}
}
