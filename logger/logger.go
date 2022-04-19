package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	var err error
	Log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string) {
	Log.Info(message)
}

func Warn(message string) {
	Log.Warn(message)
}

func Fatal(message string) {
	Log.Fatal(message)
}
