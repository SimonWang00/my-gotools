package v2_endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"my-gotools/54.工具库-go-kit微服务学习进阶/v2/v2_service"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v2_service.ContextReqUUid)), zap.Any("调用 v2_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
