package di

import (
	"fmt"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("di.NewLogger: %w", err)
	}
	defer logger.Sync()
	return logger, nil
}

func NewFxLogger(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{
		Logger: logger,
	}
}
