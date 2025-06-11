package logger

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"os"
	"syscall"
)

type Logger struct {
	logger *zap.Logger
}

var logger = NewLogger(WithDebugLevel(), WithOutputStdout())

func NewLogger(opts ...ConfigOption) *Logger {
	l, err := NewConfig().Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		logger: l,
	}
}

func Close() error {
	if logger != nil && logger.logger != nil {
		err := logger.logger.Sync()

		// TODO Разобраться подробнее: иначе будет выкидывать ошибку.
		var pathErr *fs.PathError
		if errors.Is(err, syscall.ENOTTY) || errors.As(err, &pathErr) {
			return nil
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to close logger: %v", err)
			return err
		}
	}

	return nil
}

func WithNameApp(ctx context.Context, appName string) {
	logger.logger = logger.logger.With(zap.String("app", appName))
}

func (l *Logger) Info(ctx context.Context, m string) {
	l.logger.Info(m)
}

func (l *Logger) Infof(ctx context.Context, m string, args ...any) {
	l.logger.Info(fmt.Sprintf(m, args...))
}

func (l *Logger) Warn(ctx context.Context, m string) {
	l.logger.Warn(m)
}

func (l *Logger) Warnf(ctx context.Context, m string, args ...any) {
	l.logger.Warn(fmt.Sprintf(m, args...))
}

func (l *Logger) Error(ctx context.Context, m string) {
	l.logger.Error(m)
}

func (l *Logger) Errorf(ctx context.Context, m string, args ...any) {
	l.logger.Error(fmt.Sprintf(m, args...))
}

func (l *Logger) Panic(ctx context.Context, m string) {
	l.logger.Panic(m)
}

func (l *Logger) Panicf(ctx context.Context, m string, args ...any) {
	l.logger.Panic(fmt.Sprintf(m, args...))
}

func (l *Logger) Fatal(ctx context.Context, m string) {
	l.logger.Fatal(m)
}

func (l *Logger) Fatalf(ctx context.Context, m string, args ...any) {
	l.logger.Fatal(fmt.Sprintf(m, args...))
}

func Info(ctx context.Context, m string) {
	logger.Info(ctx, m)
}

func Infof(ctx context.Context, m string, args ...any) {
	logger.Infof(ctx, m, args...)
}

func Warn(ctx context.Context, m string) {
	logger.Warn(ctx, m)
}

func Warnf(ctx context.Context, m string, args ...any) {
	logger.Warnf(ctx, m, args...)
}

func Error(ctx context.Context, m string) {
	logger.Error(ctx, m)
}

func Errorf(ctx context.Context, m string, args ...any) {
	logger.Errorf(ctx, m, args...)
}

func Panic(ctx context.Context, m string) {
	logger.Panic(ctx, m)
}

func Panicf(ctx context.Context, m string, args ...any) {
	logger.Panicf(ctx, m, args...)
}

func Fatal(ctx context.Context, m string) {
	logger.Fatal(ctx, m)
}

func Fatalf(ctx context.Context, m string, args ...any) {
	logger.Fatalf(ctx, m, args...)
}
