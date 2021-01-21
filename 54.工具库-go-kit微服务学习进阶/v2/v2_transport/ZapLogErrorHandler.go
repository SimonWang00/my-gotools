package v2_transport

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"my-gotools/54.工具库-go-kit微服务学习进阶/v2/v2_service"
)

type LogErrorHandler struct {
	logger *zap.Logger
}

func NewZapLogErrorHandler(logger *zap.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(v2_service.ContextReqUUid), zap.Error(err)))
}
