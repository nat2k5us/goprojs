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

var reqCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "Requests"})
var (
	reqGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_random_value",
		Help: "radomly generated Go random value",
	})
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
	r.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", r))
}
