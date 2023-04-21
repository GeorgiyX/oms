package grpc_server

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	RequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "homework",
		Subsystem: "grpc",
		Name:      "requests_total",
	}, []string{"method"})
	HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "homework",
		Subsystem: "grpc",
		Name:      "histogram_server_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
		[]string{"method", "status"},
	)
)

func MetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		timeStart := time.Now()
		res, err := handler(ctx, req)
		elapsed := time.Since(timeStart)

		respStatus := status.Code(err).String()
		RequestsCounter.WithLabelValues(info.FullMethod).Inc()
		HistogramResponseTime.WithLabelValues(info.FullMethod, respStatus).Observe(elapsed.Seconds())

		return res, err
	}
}
