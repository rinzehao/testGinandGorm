package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger(project, path string) *zap.Logger {
	// 方法1
	writeSyncer := getLogWriter(project, path)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	Logger = zap.New(core, zap.AddCaller())
	return Logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 所在目录必须预先创建
func getLogWriter(project, path string) zapcore.WriteSyncer {
	logDir := path + "/logger"
	if project != "" {
		logDir = "logger/" + project
	}
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(err)
	}
	consolSyncers, _, err := zap.Open(fmt.Sprintf("%s/%s.logger", logDir, time.Now().Format("2006-01-02_15:04:05")), "stdout") // stdout:打印日志在consol
	if err != nil {
		panic(err)
	}

	//defer close()
	return consolSyncers
}
