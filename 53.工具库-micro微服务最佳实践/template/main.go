package main

import (
	"flag"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"my-gotools/53.工具库-micro微服务最佳实践/template/base"
	"my-gotools/53.工具库-micro微服务最佳实践/template/base/config"
	"my-gotools/53.工具库-micro微服务最佳实践/template/base/tool"
	"my-gotools/53.工具库-micro微服务最佳实践/template/handler"
	"my-gotools/53.工具库-micro微服务最佳实践/template/model"
	user_agent "my-gotools/53.工具库-micro微服务最佳实践/template/proto/user"
	"strings"
	"time"
)

var conf = flag.String("conf", "C:\\workon\\Go\\src\\my-gotools\\53.工具库-micro微服务最佳实践\\conf", "conf path")

func main() {
	base.Init(*conf)
	// 注册中心-go-plugins好强大!
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Timeout = time.Second * 5
		options.Addrs = strings.Split(config.GetServerConfig().GetEtcdAddr(), ",")
	})
	// 服务注册
	service := micro.NewService(
		micro.Name(config.GetServerConfig().GetServerName()),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()   // 注册模型对象
			handler.Init() // 初始化服务实例
		}),
	)
	tool.GetLogger().Info("start service " + config.GetServerConfig().GetServerName() + " success")
	// handle RegisterUserHandler
	_ = user_agent.RegisterUserHandler(service.Server(), handler.GetService())
	// 启动服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
