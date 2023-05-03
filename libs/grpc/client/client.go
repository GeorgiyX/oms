package grpc_client

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConnection struct {
	*grpc.ClientConn
}

func NewClientConnection(target string, opts ...grpc.DialOption) (*ClientConnection, error) {
	optsAll := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			TraceInterceptor(),
			MetricsInterceptor(),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	optsAll = append(optsAll, opts...)

	conn, err := grpc.Dial(target, optsAll...)
	if err != nil {
		return nil, errors.Wrap(err, "dial to target")
	}

	return &ClientConnection{conn}, nil
}
