package main

import (
	"cart/internal/cart/core"
	"cart/internal/cart/logger"
	"context"
)

func main() {
	ctx := context.Background()

	cartService := core.NewService(ctx)
	if err := cartService.Run(); err != nil {
		logger.Fatalf(ctx, "Failed to start cart-service: %v", err)
	}
}
