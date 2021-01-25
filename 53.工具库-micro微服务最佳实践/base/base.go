package base

import (
	"my-gotools/53.工具库-micro微服务最佳实践/base/config"
	"my-gotools/53.工具库-micro微服务最佳实践/base/db"
	"my-gotools/53.工具库-micro微服务最佳实践/base/tool"
)

func Init(path string) {
	config.Init(path)
	db.Init()
	tool.Init()
}
