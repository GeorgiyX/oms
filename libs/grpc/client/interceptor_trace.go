package grpc_client

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TraceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		span, ctx := opentracing.StartSpanFromContext(ctx, "call: "+method)
		defer span.Finish()

		span.SetTag("method", method)

		err := invoker(ctx, method, req, reply, cc, opts...)
		if status.Code(err) != codes.OK {
			ext.Error.Set(span, true)
		}

		span.SetTag("status_code", status.Code(err).String())

		return err
	}
}
