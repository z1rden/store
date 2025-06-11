package http_server

import (
	"cart/internal/cart/http_server/middleware"
	"cart/internal/cart/logger"
	"cart/internal/cart/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Run() error
	Stop() error
	AddHandlers(handlers []model.HttpAPIHandler)
}

type server struct {
	ctx        context.Context
	mux        *http.ServeMux
	httpServer *http.Server
}

func NewServer(ctx context.Context, port string) Server {
	s := &server{
		ctx: ctx,
		mux: http.NewServeMux(),
	}

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		Handler:           s.mux,
	}

	return s
}

func (s *server) Run() error {
	const operation = "server.Run"

	if err := s.httpServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("%s: failed to start http server: %w", operation, err)
		}
	}

	return nil
}

func (s *server) Stop() error {
	const operation = "server.Stop"

	err := s.httpServer.Shutdown(s.ctx)
	if err != nil {
		logger.Errorf(s.ctx, "%s: failed to stop http server: %v", operation, err)
		return fmt.Errorf("failed to stop http server: %w", err)
	}

	logger.Infof(s.ctx, "%s: http server is stopped successfully", operation)
	return nil
}

func (s *server) AddHandlers(handlers []model.HttpAPIHandler) {
	for _, handler := range handlers {
		s.mux.HandleFunc(
			handler.Pattern,
			middleware.Logger(
				handler.Handler))
	}
}
