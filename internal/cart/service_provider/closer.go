package service_provider

import (
	"cart/internal/cart/closer"
	"context"
	"syscall"
)

func (s *ServiceProvider) GetCloser(ctx context.Context) closer.Closer {
	if s.closer == nil {
		s.closer = closer.NewCloser(syscall.SIGINT, syscall.SIGTERM)
	}

	return s.closer
}
