package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func NewConfig(opts ...ConfigOption) *zap.Config {
	config := zap.NewProductionConfig()

	// Название «атомный» связано с тем, что операции по установке и получению уровня автоматически синхронизированы:
	// даже если несколько потоков пытаются читать или записывать уровень одновременно, они не будут влиять друг на
	// друга.
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.DisableCaller = true
	config.DisableStacktrace = true

	for _, opt := range opts {
		opt(&config)
	}

	return &config
}
