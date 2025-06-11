package http_server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"os"
)

type API interface {
	RegisterHttpHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type Server interface {
	Run() error
	Stop() error
	RegisterApi(api []API) error
}

type server struct {
	ctx        context.Context
	gatewayMux *runtime.ServeMux
	mux        *http.ServeMux
	conn       *grpc.ClientConn
	httpServer *http.Server
}

func NewServer(ctx context.Context, httpPort string, grpcPort string) (Server, error) {
	s := &server{
		ctx:        ctx,
		gatewayMux: runtime.NewServeMux(),
		mux:        http.NewServeMux(),
	}

	var err error
	s.conn, err = grpc.NewClient(fmt.Sprintf(":%s", grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: s.mux,
	}

	s.mux.HandleFunc("/swagger.json", s.handleSwagger)

	fs := http.FileServer(http.Dir("pkg/swagger-ui"))
	s.mux.Handle("/docs/", http.StripPrefix("/docs/", fs))

	s.mux.Handle("/", s.gatewayMux)

	return s, nil
}

func (s *server) Run() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

func (s *server) Stop() error {
	err := s.httpServer.Shutdown(s.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) RegisterApi(api []API) error {
	for _, sApi := range api {
		err := sApi.RegisterHttpHandlers(s.ctx, s.gatewayMux, s.conn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) handleSwagger(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("pkg/swagger/swagger.json")
	if err != nil {
		http.Error(w, "swagger not found", http.StatusNotFound)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	io.Copy(w, reader)
}
