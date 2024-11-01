package logger

import (
	"os"

	"go-ecommerce-be/pkg/settings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config settings.LoggerSettings) *LoggerZap {
	logLever := config.Level
	// debug, info, warn, error, dpanic, panic, fatal

	var level zapcore.Level
	switch logLever {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	formatLog := formatLog()

	hook := lumberjack.Logger{
		Filename:   config.File,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
		Compress:   config.Compress,
	}

	coreLogger := zapcore.NewCore(
		formatLog,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &LoggerZap{zap.New(coreLogger, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func formatLog() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	// time encoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "Time"

	// level encoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller encoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
