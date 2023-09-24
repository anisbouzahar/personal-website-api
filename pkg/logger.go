package pkg

import (
	"fmt"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, func()) {
	logger, err := zap.NewProduction(zap.WithCaller(false))
	if err != nil {
		panic(fmt.Sprintf("Couldn't initialize zap logger: %v", err))
	}

	return logger, func() {
		_ = logger.Sync()
	}
}
