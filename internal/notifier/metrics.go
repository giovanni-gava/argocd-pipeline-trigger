package notifier

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	SyncTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "argocd_sync_total",
			Help: "Total number of ArgoCD sync requests",
		},
		[]string{"app", "status"},
	)

	SyncDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "argocd_sync_duration_seconds",
			Help:    "Duration of ArgoCD sync operations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"app"},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(SyncTotal, SyncDuration)
}
