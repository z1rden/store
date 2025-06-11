package main

import (
	"context"
	"loms/internal/loms/core"
	"loms/internal/loms/logger"
)

func main() {
	ctx := context.Background()
	service := core.NewService(ctx)

	if err := service.Run(); err != nil {
		logger.Fatal(ctx, "can't run service")
	}
}
