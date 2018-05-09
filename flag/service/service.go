package service

import (
	"github.com/giantswarm/credentiald/flag/service/kubernetes"
	"github.com/giantswarm/credentiald/flag/service/secrets"
)

type Service struct {
	Kubernetes kubernetes.Kubernetes
	Secrets    secrets.Secrets
}