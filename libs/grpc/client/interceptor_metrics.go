package grpc_client

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "homework",
		Subsystem: "grpc",
		Name:      "histogram_client_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
		[]string{"method", "status"},
	)
)

func MetricsInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		timeStart := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		elapsed := time.Since(timeStart)

		respStatus := status.Code(err).String()
		HistogramResponseTime.WithLabelValues(method, respStatus).Observe(elapsed.Seconds())

		return err
	}
}
