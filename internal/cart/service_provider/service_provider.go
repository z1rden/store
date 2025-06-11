package service_provider

import (
	"cart/internal/cart/closer"
	"context"
)

// Паттерн Singleton
var serviceProvider *ServiceProvider

// ServiceProvider - это компонент, который используется для централизованного управления зависимостями и
// сервисами приложения.
type ServiceProvider struct {
	ctx        context.Context
	repository repository
	service    service
	closer     closer.Closer
	client     client
	api        api
}

func GetServiceProvider(ctx context.Context) *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{
			ctx: ctx,
		}
	}
	return serviceProvider
}
