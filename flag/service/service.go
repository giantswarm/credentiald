package service

import (
	"github.com/giantswarm/operatorkit/v2/pkg/flag/service/kubernetes"

	"github.com/giantswarm/credentiald/v2/flag/service/secrets"
)

type Service struct {
	Kubernetes kubernetes.Kubernetes
	Secrets    secrets.Secrets
}
