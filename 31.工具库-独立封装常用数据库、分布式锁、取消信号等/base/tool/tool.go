package tool

import (
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/config"
	"sync"
)

var (
	one sync.Once
	err error
)

func Init() {
	one.Do(func() {
		initLogger(getLoggerOptions())
		//NewWorker ID 填入 分布式的服务唯一ID从1到1024
		if err = NewWorker(config.GetToolLogConfig().GetDcId()); err != nil {
			panic(err)
		}
	})
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
