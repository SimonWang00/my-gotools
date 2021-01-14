package utils

import (
	"go.uber.org/zap"
	"my-gotools/all_packaged_library/logtool"
)

var logger *zap.Logger

func NewLoggerServer() {
	logger = logtool.NewLogger(
		logtool.SetAppName("54.工具库-go-kit微服务学习进阶"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
