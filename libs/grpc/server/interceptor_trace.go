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

const TraceHeader = "x-trace-id"
const SpanHeader = "x-span-id"

func initTracing(serviceName string, reporter string) error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: reporter,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return errors.Wrap(err, "cannot init tracing")
	}

	return nil
}

func newRootSpan(ctx context.Context, method string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, method)
	spanContext, ok := span.Context().(jaeger.SpanContext)
	if !ok {
		logger.Warn("can't cast span context to jaeger.SpanContext")
		return span, ctx
	}

	ctx = metadata.AppendToOutgoingContext(ctx, TraceHeader, spanContext.TraceID().String())

	md := metadata.New(map[string]string{
		TraceHeader: spanContext.TraceID().String(),
	})

	err := grpc.SetHeader(ctx, md)
	if err != nil {
		logger.Warn("can't set trace id metadata header", zap.Error(err))
	}

	return span, ctx
}

func attachToRootSpan(ctx context.Context, method, trace, span string) (opentracing.Span, context.Context) {
	ctxString := trace + ":" + span + ":0:0"
	spanContext, err := jaeger.ContextFromString(ctxString)
	if err != nil {
		logger.Warn("can't create span context", zap.Error(err))
		return newRootSpan(ctx, method)
	}

	ctx = metadata.AppendToOutgoingContext(ctx, TraceHeader, spanContext.TraceID().String())
	return opentracing.StartSpanFromContext(ctx, method, opentracing.ChildOf(spanContext))
}

// newServiceSpan search traceId and spanId in incoming metadata and if its found create span
// linked with external span. Otherwise, create completely new span.
// In both cases write traceId to outgoing context.
func newServiceSpan(ctx context.Context, method string) (opentracing.Span, context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get(TraceHeader)) == 0 || len(md.Get(SpanHeader)) == 0 {
		return newRootSpan(ctx, method)
	}
	traceID := md.Get(TraceHeader)[0]
	spanID := md.Get(SpanHeader)[0]
	return attachToRootSpan(ctx, method, traceID, spanID)
}

func TraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		span, ctx := newServiceSpan(ctx, info.FullMethod)
		defer span.Finish()

		span.SetTag("operation", info.FullMethod)

		res, err := handler(ctx, req)
		if status.Code(err) != codes.OK {
			ext.Error.Set(span, true)
		}

		span.SetTag("status_code", status.Code(err).String())

		return res, err
	}
}
