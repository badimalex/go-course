package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HttpDuration      *prometheus.HistogramVec
	HttpRequestsTotal *prometheus.CounterVec
}

func New(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		HttpDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_response_duration_seconds",
			Help: "Duration of HTTP requests.",
		}, []string{"path"}),
		HttpRequestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests.",
		}, []string{"method", "path"}),
	}

	reg.MustRegister(m.HttpDuration)
	reg.MustRegister(m.HttpRequestsTotal)

	return m
}
