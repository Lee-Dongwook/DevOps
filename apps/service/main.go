package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of received requests",
		},
		[]string {"path"},
	)
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Request duration",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func handler(w http.ResponseWriter, r *http.Request){
	start := time.Now()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from kube-blueprint with metrics!")

	duration := time.Since(start).Seconds()
	reqCounter.WithLabelValues(r.URL.Path).Inc()
	reqDuration.WithLabelValues(r.URL.Path).Observe(duration)
}

func main() {
	prometheus.MustRegister(reqCounter, reqDuration)

	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}