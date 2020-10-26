// Package collector offers a service to retrieve metrics about credentials.
package collector

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
	PrometheusNamespace = "credentiald"

	PrometheusNameLabel         = "name"
	PrometheusOrganizationLabel = "organization"

	KubernetesOrganizationLabel = "giantswarm.io/organization"

	KubernetesCredentialNamespace = "giantswarm"
	KubernetesLabelSelector       = "app=credentiald"

	gaugeValue = float64(1)
)

var (
	credentials = prometheus.NewDesc(
		prometheus.BuildFQName(PrometheusNamespace, "", "credential_info"),
		"Credential info.",
		[]string{
			PrometheusNameLabel,
			PrometheusOrganizationLabel,
		},
		nil,
	)
)

type Config struct {
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger
}

type Collector struct {
	k8sClient kubernetes.Interface
	logger    micrologger.Logger
}

// New creates a new Collector based on a configutation.
func New(config Config) (*Collector, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}

	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	c := &Collector{
		k8sClient: config.K8sClient,
		logger:    config.Logger,
	}

	return c, nil
}

// Describe collects metadata for a prometheus metric.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- credentials
}

// Collect collects metrics for prometheus
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	ctx := context.Background()
	c.logger.Log("level", "debug", "message", "collecting metrics")

	credentialList, err := c.k8sClient.CoreV1().Secrets(KubernetesCredentialNamespace).List(ctx, metav1.ListOptions{
		LabelSelector: KubernetesLabelSelector,
	})
	if err != nil {
		c.logger.Log("level", "error", "message", "could not list secrets", "stack", fmt.Sprintf("%#v", err))
	}

	for _, credential := range credentialList.Items {
		ch <- prometheus.MustNewConstMetric(
			credentials,
			prometheus.GaugeValue,
			gaugeValue,
			credential.Name,
			credential.Labels[KubernetesOrganizationLabel],
		)
	}

	c.logger.Log("level", "debug", "message", "finished collecting metrics")
}
