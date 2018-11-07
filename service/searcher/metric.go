package searcher

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "credentiald"
	subsystem = "credentials"
)

var (
	searchTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "search",
			Help:      "Time taken to search credentials.",
		},
	)

	kubernetesSearchErrorTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "kubernetes_search_error_total",
			Help:      "Counter of errors encountered searching Kubernetes credential secrets.",
		},
	)
)

func init() {
	prometheus.MustRegister(searchTime)
	prometheus.MustRegister(kubernetesSearchErrorTotal)
}
