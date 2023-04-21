package grpc_server

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"route256/libs/logger"
)

func initTracing(serviceName string) error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return errors.Wrap(err, "cannot init tracing")
	}

	return nil
}

func TraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		span, ctx := opentracing.StartSpanFromContext(ctx, info.FullMethod)
		defer span.Finish()

		span.SetTag("method", info.FullMethod)

		// jaeger.SpanContext != context.Context. Set by jaeger lib. Opentracing lib nothing know about jaeger.SpanContext.
		spanContext, ok := span.Context().(jaeger.SpanContext)
		if ok {
			md := metadata.New(map[string]string{
				"x-trace-id": spanContext.TraceID().String(),
			})

			err := grpc.SetHeader(ctx, md)
			if err != nil {
				logger.Warn("can't set trace id metadata header", zap.Error(err))
			}
		}

		res, err := handler(ctx, req)
		if status.Code(err) != codes.OK {
			ext.Error.Set(span, true)
		}
		span.SetTag("status_code", status.Code(err).String())

		return res, err
	}
}
