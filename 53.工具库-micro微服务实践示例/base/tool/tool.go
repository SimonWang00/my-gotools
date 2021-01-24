package tool

import (
	"my-gotools/53.工具库-micro微服务实践示例/base/config"
)

func Init() {
	initLogger(getLoggerOptions())
}
func getLoggerOptions() *Options {
	op := &Options{}
	op.Development = config.GetToolLogConfig().GetDevelopment()
	op.LogFileDir = config.GetToolLogConfig().GetLogFileDir()
	op.AppName = config.GetToolLogConfig().GetAppName()
	op.MaxSize = config.GetToolLogConfig().GetMaxSize()
	op.MaxBackups = config.GetToolLogConfig().GetMaxBackups()
	op.MaxAge = config.GetToolLogConfig().GetMaxAge()
	return op
}
