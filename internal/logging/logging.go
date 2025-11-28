package logging

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func Init(production bool) {
	if production {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(fmt.Errorf("failed to initialize production logger: %w", err))
		}
		Logger = logger.Sugar()
	} else {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(fmt.Errorf("failed to initialize development logger: %w", err))
		}
		Logger = logger.Sugar()
	}
}
