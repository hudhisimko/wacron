package messenger

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	messageReceivedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "message_received_total",
		Help: "Total message fetched by cron",
	}, []string{"tenant_id", "tenant_branch"})

	messageSentTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "message_sent_total",
		Help: "Total message successfully sent",
	}, []string{"tenant_id", "tenant_branch", "provider"})

	messageFailedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "message_failed_total",
		Help: "Total message failed to sent",
	}, []string{"tenant_id", "tenant_branch", "reason"})
)

func ListenMetric() error {
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":3541", h)
}
