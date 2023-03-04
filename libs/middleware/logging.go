package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("%s %s: %s --- %v\n", color.GreenString("[gRPC]"), time.Now().Format(time.RFC850), info.FullMethod, req)

	res, err := handler(ctx, req)
	if err != nil {
		fmt.Printf("%s %s: %s --- %v\n", color.RedString("[gRPC]"), time.Now().Format(time.RFC850), info.FullMethod, err)
		return nil, err
	}

	fmt.Printf("%s %s: %s --- %v\n", color.GreenString("[gRPC]"), time.Now().Format(time.RFC850), info.FullMethod, res)
	return res, nil
}
