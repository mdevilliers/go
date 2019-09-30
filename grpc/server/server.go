package server

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	*grpc.Server
}

// DefaultIptions returns a reasonaly set of grpc.ServerOptions
func DefaultOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		// clients should reconnect regularly so we can loadbalace effectively
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
	}
}

// New returns a configured GRPC service
func NewGRPCServer(options ...grpc.ServerOption) *server {
	return &server{
		Server: grpc.NewServer(
			options...,
		),
	}
}

// NewDefaultGRPCServer returns a configured GRPC service using the DefaultOptions
func NewDefaultGRPCServer() *server {
	return NewGRPCServer(DefaultOptions()...)
}

// Start successfully starts the server or returns an error
func (s *server) Start(address string) error {

	// Create a new listener on the defined address
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrap(err, "error creating listener")
	}

	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "error running server")
	}
	return nil
}

// Shutdown attempts a graceful shutdown of the server
func (s *server) Shutdown(ctx context.Context) {
	s.GracefulStop()
}
