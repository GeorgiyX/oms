package grpc_client

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	grpc_server "route256/libs/grpc/server"
	"route256/libs/logger"
)

func spanIDToOutgoingContext(ctx context.Context, span opentracing.Span) context.Context {
	spanContext, ok := span.Context().(jaeger.SpanContext)
	if !ok {
		logger.Warn("can't cast span context to jaeger.SpanContext")
		return ctx
	}

	return metadata.AppendToOutgoingContext(ctx, grpc_server.SpanHeader, spanContext.SpanID().String())
}

func TraceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		span, ctx := opentracing.StartSpanFromContext(ctx, "call: "+method)
		defer span.Finish()

		span.SetTag("method", method)

		ctx = spanIDToOutgoingContext(ctx, span)
		err := invoker(ctx, method, req, reply, cc, opts...)
		if status.Code(err) != codes.OK {
			ext.Error.Set(span, true)
		}

		span.SetTag("status_code", status.Code(err).String())
		return err
	}
}
