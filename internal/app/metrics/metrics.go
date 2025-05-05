package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsService struct {
	httpRequestsTotal   *prometheus.CounterVec
	httpRequestDuration *prometheus.HistogramVec
	httpRequestInFlight prometheus.Gauge
}

func NewMetricsService() *MetricsService {
	return &MetricsService{
		httpRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		httpRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "Duration of HTTP requests",
				Buckets: []float64{0.1, 0.5, 1, 2, 5},
			},
			[]string{"method", "path"},
		),
		httpRequestInFlight: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "http_requests_in_flight",
				Help: "Number of currently processing HTTP requests",
			},
		),
	}
}

func (m *MetricsService) InstrumentHandler(method, path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.httpRequestInFlight.Inc()
		defer m.httpRequestInFlight.Dec()

		rw := &responseWriter{ResponseWriter: w}
		handler(rw, r)

		duration := time.Since(start).Seconds()
		m.httpRequestDuration.WithLabelValues(method, path).Observe(duration)
		m.httpRequestsTotal.WithLabelValues(method, path, strconv.Itoa(rw.status)).Inc()
	}
}

func (m *MetricsService) MetricsHandler() http.Handler {
	return promhttp.Handler()
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
