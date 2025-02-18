package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

// Info logs a message with some additional context
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
	log.Sync()
}

// Error logs a message with some additional context
func Error(message string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", err))
	log.Error(message, fields...)
	log.Sync()
}

// Debug logs a message with some additional context
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
	log.Sync()
}
