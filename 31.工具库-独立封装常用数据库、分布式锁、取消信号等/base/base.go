package base

import (
	"my-gotools/all_packaged_library/base/config"
	"my-gotools/all_packaged_library/base/db"
	"my-gotools/all_packaged_library/base/tool"
)

//配置文件的目录
func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}

