package service_provider

import (
	"context"
	"loms/internal/loms/closer"
	"loms/internal/loms/config"
)

type ServiceProvider struct {
	repository repository
	service    service
	closer     closer.Closer
	api        api
	cfg        *config.Config
	kafka      kafka
}

var serviceProvider *ServiceProvider

func GetServiceProvider(ctx context.Context, cfg *config.Config) *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{
			cfg: cfg,
		}
	}

	return serviceProvider
}
