package middleware

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info(
			"[gRPC] Accept request",
			zap.String("method", info.FullMethod),
			zap.String("request", fmt.Sprintf("%v", req)),
		)

		res, err := handler(ctx, req)
		if err != nil {
			logger.Error(
				"[gRPC] Request finished with error",
				zap.String("method", info.FullMethod),
				zap.Error(err),
			)
			return nil, err
		}

		logger.Info(
			"[gRPC] Request finished OK",
			zap.String("method", info.FullMethod),
			zap.String("request", fmt.Sprintf("%v", res)),
		)
		return res, nil
	}
}
