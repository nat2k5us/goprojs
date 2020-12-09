package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var REQUESTS_INPROGRESS = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "go_app_requests_inprogress",
	Help: "Number of App  Requests in progress.",
})

func main() {
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
		REQUESTS_INPROGRESS.Inc()
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Happy Birthday %s :", name)
		time.Sleep(5 * time.Second)
		rw.Write([]byte(greetings))
		REQUESTS_INPROGRESS.Dec()
	}).Methods("GET")
	
	log.Println("starting the application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":8000", router)
}
