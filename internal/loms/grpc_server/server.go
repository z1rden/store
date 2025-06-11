package grpc_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type API interface {
	RegisterGrpcServer(s *grpc.Server)
}

type Server interface {
	Run() error
	Stop() error
	RegisterApi(api []API) error
}

type server struct {
	ctx        context.Context
	grpcServer *grpc.Server
	port       string
}

func NewServer(ctx context.Context, port string) Server {
	grpcServer := grpc.NewServer()

	return &server{
		ctx:        ctx,
		grpcServer: grpcServer,
		port:       port,
	}
}

func (s *server) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(l)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) Stop() error {
	s.grpcServer.GracefulStop()

	return nil
}

func (s *server) RegisterApi(api []API) error {
	for _, sApi := range api {
		sApi.RegisterGrpcServer(s.grpcServer)
	}

	return nil
}
