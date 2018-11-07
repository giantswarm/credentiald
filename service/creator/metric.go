package creator

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "credentiald"
	subsystem = "credentials"
)

var (
	createTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "create",
			Help:      "Time taken to create credentials.",
		},
	)

	kubernetesCreateErrorTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "kubernetes_create_error_total",
			Help:      "Counter of errors encountered creating Kubernetes credential secrets.",
		},
	)
)

func init() {
	prometheus.MustRegister(createTime)
	prometheus.MustRegister(kubernetesCreateErrorTotal)
}
