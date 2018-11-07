package lister

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "credentiald"
	subsystem = "credentials"
)

var (
	listTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "list",
			Help:      "Time taken to list credentials.",
		},
	)

	kubernetesListErrorTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "kubernetes_list_error_total",
			Help:      "Counter of errors encountered listing Kubernetes credential secrets.",
		},
	)
)

func init() {
	prometheus.MustRegister(listTime)
	prometheus.MustRegister(kubernetesListErrorTotal)
}
