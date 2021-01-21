package base

import (
	"my-gotools/53.工具库-micro微服务实践示例/base/config"
	"my-gotools/53.工具库-micro微服务实践示例/base/db"
	"my-gotools/53.工具库-micro微服务实践示例/base/tool"
)

func Init(path string) {
	config.Init(path)
	db.Init()
	tool.Init()
}
