package grpc_server

import (
	"context"
	"github.com/hashicorp/go-multierror"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(serverName string, optIn ...grpc.ServerOption) (*Server, error) {
	err := initTracing(serverName)
	if err != nil {
		return nil, errors.Wrap(err, "init tracing")
	}

	chainedInterceptor := grpc.ChainUnaryInterceptor(
		TraceInterceptor(),
		MetricsInterceptor(),
	)
	opts := append(optIn, chainedInterceptor)

	server := grpc.NewServer(opts...)
	reflection.Register(server)
	return &Server{server}, nil
}

func (s *Server) Serve(grpcAddr, httpAddr string) error {

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return errors.Wrap(err, "create tcp listener")
	}

	router := httprouter.New()
	router.Handler(http.MethodGet, "/metrics", promhttp.Handler())
	httpServer := &http.Server{Addr: httpAddr, Handler: router}

	var errsGRPC, errsHTTP error
	wg := sync.WaitGroup{}

	// run grpc server
	wg.Add(1)
	go func() {
		defer wg.Done()
		errIn := s.grpcServer.Serve(lis)
		if errIn != nil {
			errsGRPC = multierror.Append(errsGRPC, errors.Wrap(errIn, "serve grpc server"))
		}
		errIn = httpServer.Shutdown(context.Background())
		if errIn != nil {
			errsGRPC = multierror.Append(errsGRPC, errors.Wrap(errIn, "stop http server"))
		}
	}()

	// run http server for serving metrics page
	wg.Add(1)
	go func() {
		defer wg.Done()
		errIn := httpServer.ListenAndServe()
		if errIn != nil {
			errsHTTP = multierror.Append(errsHTTP, errors.Wrap(errIn, "serve http server"))
		}
		s.grpcServer.GracefulStop()
	}()

	wg.Wait()
	return multierror.Append(errsGRPC, errsHTTP)
}

func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.grpcServer.RegisterService(desc, impl)
}
