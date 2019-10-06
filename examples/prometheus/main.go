package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gebn/go-stamp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	addr = ":8080"
)

var (
	buildInfo = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "example_build_info",
			Help: "The version and commit of the example code. Constant 1.",
		},
		// the runtime version is already exposed by the default Go collector
		[]string{"version", "commit"},
	)
	buildTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "example_build_time",
		Help: "When the example code was build, as seconds since the Unix Epoch.",
	})
)

func init() {
	buildInfo.WithLabelValues(stamp.Version, stamp.Commit).Set(1)
	buildTime.Set(float64(stamp.Time().UnixNano()) / float64(time.Second))
}

func main() {
	log.Printf("about to listen on %v; visit localhost%[1]v/metrics", addr)
	http.ListenAndServe(addr, promhttp.Handler())
}
