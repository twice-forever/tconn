package log

import (
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.Logger

func Setup() {
	logger, _ = zap.NewDevelopment(zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(fields ...interface{}) {
	logger.Info(fmt.Sprint(fields...))
}

func Error(fields ...interface{}) {
	logger.Error(fmt.Sprint(fields...))
}
