package middleware

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"route256/libs/ratelimiter"
)

var withRateLimit map[string]struct{} = map[string]struct{}{
	"/route256.product.ProductService/GetProduct": struct{}{},
}

// RateLimiterInterceptor call method with rate limit check
func RateLimiterInterceptor() grpc.UnaryClientInterceptor {
	rl := ratelimiter.New(ratelimiter.Config{
		Interval: time.Second,
		Requests: 10,
	})
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_, ok := withRateLimit[method]
		if ok {
			err := rl.Wait(ctx)
			if err != nil {
				return errors.Wrap(err, "rate-limit fail")
			}
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
