package utils

import (
	"go.uber.org/zap"
	"learning_tools/all_packaged_library/logtool"
)

const ContextReqUUid = "req_uuid"
var logger *zap.Logger

func NewLoggerServer() {
	logger = logtool.NewLogger(
		logtool.SetAppName("54.工具库-go-kit微服务学习进阶-v11-server"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
