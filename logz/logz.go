package logz

import (
	"context"
	"fmt"
	"moon-base/common"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogConfig() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"

	config := zap.NewProductionConfig()
	var logLevel zapcore.Level
	switch viper.GetString("LOG.LEVEL") {
	case "info":
		logLevel = zapcore.InfoLevel
	case "debug":
		logLevel = zapcore.DebugLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}
	config.Level = zap.NewAtomicLevelAt(logLevel)
	if viper.GetString("LOG.ENV") == "dev" {
		config.Encoding = "console"
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.Encoding = "json"
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	config.EncoderConfig = encoderConfig

	logger, _ := config.Build()

	return logger
}

func ExecutionTime(start time.Time, name string, l *zap.Logger) {
	elapse := time.Since(start)
	l.With(zap.Duration("duration", elapse)).Info(fmt.Sprintf("%s took %s", name, elapse))
}

func ContextLogger(ctx context.Context) *zap.Logger {
	v := ctx.Value(common.LoggerKey)
	if v == nil {
		return nil
	}
	l, ok := v.(*zap.Logger)
	if ok {
		return l
	}
	return nil
}
