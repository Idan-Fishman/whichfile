package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Level zap.AtomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel) // Current log level.
	L     *zap.Logger     = InitializeLogger()                  // Global logger instance.
)

// InitializeLogger initializes a logger with custom encoding and returns it.
func InitializeLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             Level,
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	logger, _ := config.Build()
	return logger
}

// SetLevel sets the log level.
func SetLevel(level string) {
	l, _ := zapcore.ParseLevel(level)
	Level.SetLevel(l)
}

// Flush flushes any buffered log entries.
func Flush() {
	L.Sync()
}
