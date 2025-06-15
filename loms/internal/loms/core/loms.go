package core

import (
	"context"
	"loms/internal/loms/config"
	"loms/internal/loms/grpc_server"
	"loms/internal/loms/http_server"
	"loms/internal/loms/kafka_admin"
	"loms/internal/loms/logger"
	"loms/internal/loms/service_provider"
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
	cfg := config.MustLoad()
	logger.WithNameApp(ctx, cfg.AppName)

	ctx, cancel := context.WithCancel(context.Background())
	serviceProvider := service_provider.GetServiceProvider(ctx, cfg)
	serviceProvider.GetCloser(ctx).Add(func() error {
		cancel()
		return nil
	})

	return &service{
		ctx:             ctx,
		serviceProvider: serviceProvider,
		cfg:             cfg,
	}
}

func (s *service) Run() error {
	logger.Infof(s.ctx, "Starting service")
	defer logger.Infof(s.ctx, "Stopping service")

	closer := s.serviceProvider.GetCloser(s.ctx)
	defer closer.Wait()

	kafkaAdmin := s.serviceProvider.GetKafkaAdmin(s.ctx)
	if err := kafkaAdmin.CreateTopic(s.cfg.Kafka.Topic.Name,
		kafka_admin.WithNumPartitions(int32(s.cfg.Kafka.Topic.NumPartitions)),
		kafka_admin.WithReplicationFactor(int16(s.cfg.Kafka.Topic.ReplicationFactor)),
		kafka_admin.WithRetentionMSMinute(s.cfg.Kafka.Topic.RetentionsMS)); err != nil {

		return err
	}
	if err := kafkaAdmin.Close(); err != nil {
		logger.Fatalf(s.ctx, "Failed to close kafka admin: %v", err)
	}

	orderApi := s.serviceProvider.GetOrderAPI(s.ctx)
	stockApi := s.serviceProvider.GetStockAPI(s.ctx)

	grpcServer := grpc_server.NewServer(s.ctx, s.cfg.GrpcPort)
	err := grpcServer.RegisterApi([]grpc_server.API{orderApi, stockApi})
	if err != nil {
		return err
	}
	closer.Add(grpcServer.Stop)

	go func() {
		logger.Infof(s.ctx, "Starting grpc server on port %s", s.cfg.GrpcPort)
		err := grpcServer.Run()
		if err != nil {
			logger.Fatalf(s.ctx, "grpc server run failed: %s", err)
			closer.CloseAll()
		}
	}()

	httpServer, err := http_server.NewServer(s.ctx, s.cfg.HttpPort, s.cfg.GrpcPort)
	if err != nil {
		return err
	}
	err = httpServer.RegisterApi([]http_server.API{orderApi, stockApi})
	if err != nil {
		return err
	}
	closer.Add(httpServer.Stop)

	go func() {
		logger.Infof(s.ctx, "Starting http server on port %s", s.cfg.HttpPort)
		err := httpServer.Run()
		if err != nil {
			logger.Fatalf(s.ctx, "http server run failed: %s", err)
			closer.CloseAll()
		}
	}()

	closer.Add(logger.Close)

	kafkaService := s.serviceProvider.GetKafkaService(s.ctx)
	kafkaService.SendMessages(s.ctx)
	closer.Add(kafkaService.StopSendMessages)

	return nil
}
