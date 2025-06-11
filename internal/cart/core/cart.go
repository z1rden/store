package core

import (
	"cart/internal/cart/config"
	"cart/internal/cart/http_server"
	"cart/internal/cart/logger"
	"cart/internal/cart/service_provider"
	"context"
)

type Service interface {
	Run() error
}

type service struct {
	ctx             context.Context
	serviceProvider *service_provider.ServiceProvider
	cfg             *config.Config
}

func NewService(ctx context.Context) Service {
	ctx, cancel := context.WithCancel(ctx)

	sp := service_provider.GetServiceProvider(ctx)
	sp.GetCloser(ctx).Add(func() error {
		cancel()
		return nil
	})
	cfg := config.MustLoad()

	logger.WithNameApp(ctx, cfg.AppName)

	return &service{
		ctx:             ctx,
		serviceProvider: sp,
		cfg:             cfg,
	}
}

func (s *service) Run() error {
	logger.Info(s.ctx, "cartService is starting")
	defer logger.Info(s.ctx, "cartService finished")

	closer := s.serviceProvider.GetCloser(s.ctx)
	defer closer.Wait()

	api := s.serviceProvider.GetApi(s.cfg.LomsServiceGrpcHost)

	httpPort := s.cfg.HttpPort
	httpServer := http_server.NewServer(s.ctx, httpPort)
	httpServer.AddHandlers(api.GetHandlers())
	closer.Add(httpServer.Stop)

	go func() {
		logger.Infof(s.ctx, "http cartService server is starting on localhost:%s", httpPort)
		err := httpServer.Run()
		if err != nil {
			logger.Errorf(s.ctx, "failed to start http cartService server on localhost:%s : %v",
				httpPort, err)
			closer.CloseAll()
		}
	}()

	closer.Add(logger.Close)

	return nil
}
