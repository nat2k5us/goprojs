package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "myrequests", Help: "requests being made"})
var (
	reqGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_random_value",
		Help: "randomly generated Go random value",
	})
)

var (
	// apiRequestDuration tracks the duration separate for each HTTP status
	// class (1xx, 2xx, ...). This creates a fair amount of time series on
	// the Prometheus server. Usually, you would track the duration of
	// serving HTTP request without partitioning by outcome. Do something
	// like this only if needed. Also note how only status classes are
	// tracked, not every single status code. The latter would create an
	// even larger amount of time series. Request counters partitioned by
	// status code are usually OK as each counter only creates one time
	// series. Histograms are way more expensive, so partition with care and
	// only where you really need separate latency tracking. Partitioning by
	// status class is only an example. In concrete cases, other partitions
	// might make more sense.
	apiRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram for the request duration of the public API, partitioned by status class.",
			Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
		},
		[]string{"status_class"},
	)
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/testCounter", func(w http.ResponseWriter, r *http.Request) {
		reqCounter.Add(1)
		fmt.Println("Incremented counter")
	})
	r.HandleFunc("/testGauge", func(w http.ResponseWriter, r *http.Request) {
		randNumber := rand.Float64()
		reqGauge.Set(randNumber)
		fmt.Println("Random number generation")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		// The ObserverFunc gets called by the deferred ObserveDuration and
		// decides which Histogram's Observe method is called.
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			switch {
			case status >= 500: // Server error.
				apiRequestDuration.WithLabelValues("5xx").Observe(v)
			case status >= 400: // Client error.
				apiRequestDuration.WithLabelValues("4xx").Observe(v)
			case status >= 300: // Redirection.
				apiRequestDuration.WithLabelValues("3xx").Observe(v)
			case status >= 200: // Success.
				apiRequestDuration.WithLabelValues("2xx").Observe(v)
			default: // Informational.
				apiRequestDuration.WithLabelValues("1xx").Observe(v)
			}
		}))
		defer timer.ObserveDuration()

		// Handle the request. Set status accordingly.
		// ...
	})
	r.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8002", r))
}
