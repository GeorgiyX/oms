package cache

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	missCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "homework",
		Subsystem: "cache",
		Name:      "miss_total",
	}, []string{"name"})

	requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "homework",
		Subsystem: "cache",
		Name:      "request_total",
	}, []string{"name"})

	responseTimeCacheGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "homework",
		Subsystem: "cache",
		Name:      "rt_cache_us",
	}, []string{"name"})

	responseTimeFnGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "homework",
		Subsystem: "cache",
		Name:      "rt_fn_us",
	}, []string{"name"})
)

func countRtCache(name string) func() {
	timeStart := time.Now()
	return func() {
		elapsed := time.Since(timeStart)
		responseTimeCacheGauge.WithLabelValues(name).Set(float64(elapsed.Microseconds()))
	}
}

func countRtFn(name string) func() {
	timeStart := time.Now()
	return func() {
		elapsed := time.Since(timeStart)
		responseTimeFnGauge.WithLabelValues(name).Set(float64(elapsed.Microseconds()))
	}
}
