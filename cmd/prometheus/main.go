package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	service := &http.Server{
		Handler: router(),
		Addr:    ":8484",
	}

	logrus.
		WithError(service.ListenAndServe()).
		Panic()
}

func router() http.Handler {
	r := mux.NewRouter()

	r.Handle("/metrics", promhttp.Handler())
	metric, err := newMetric()
	if nil != err {
		panic(err)
	}

	r.HandleFunc("/buggy", getBuggy(metric))

	return r
}

func newMetric() (prometheus.Counter, error) {
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace:   "at",
			Subsystem:   "learn",
			Name:        "buggy",
			Help:        "just for learning",
			ConstLabels: nil,
		},
	)

	if err := prometheus.Register(counter); nil != err {
		return nil, err
	}

	return counter, nil
}

func getBuggy(counter prometheus.Counter) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer counter.Inc()

		w.WriteHeader(200)
		_, _ = w.Write([]byte(`OK`))
	}
}
