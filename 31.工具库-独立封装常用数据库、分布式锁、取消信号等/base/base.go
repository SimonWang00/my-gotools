package base

import (
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/config"
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/db"
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/tool"
)

//配置文件的目录
func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}

