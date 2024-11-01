package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

/* 	//logger, _ := zap.NewProduction()
	logger := zap.NewExample()
	logger.Info("Hello, World")

	//development logger

	logger, _ = zap.NewDevelopment()
	logger.Info("Hello, World")

	//production logger
	logger, _ = zap.NewProduction()
	logger.Info("Hello, World") */

	//custom logger
	formatLog := formatLog()
	writeLogSync := writeLogSync()
	coreLogger := zapcore.NewCore(formatLog, writeLogSync, zapcore.InfoLevel)
	logger := zap.New(coreLogger, zap.AddCaller())
	logger.Info("Hello, World", zap.String("key", "value"))



}

func formatLog() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	//time encoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "Time"

	//level encoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//caller encoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func writeLogSync() zapcore.WriteSyncer{
	os.MkdirAll(".log", os.ModePerm)
	file, _ := os.OpenFile(".log/log.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}

