package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/4

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	grpcSvc "myshopping/client/proto/user"
	webLib "myshopping/client/webLib"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error  {
	fmt.Println("调用接口")
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp, opts...)
}

func NewLogWrapper(c client.Client) client.Client  {
	return &logWrapper{c}
}

func main()  {
	//使用etcd作为注册中心
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	mySvc := micro.NewService(
		micro.Name("go.micro.srv.user.client"),
		micro.WrapClient(NewLogWrapper), //包装log函数
	)
	//后端grpc服务的client
	prodSvc := grpcSvc.NewUserService("go.micro.srv.user", mySvc.Client())
	httpSvc := web.NewService(
		web.Name("go.micro.srv.user.client"), //服务名称
		web.Address(":8002"), //监听端口
		//将gin引入, 里面包含gin注册的http路由，参数prodSvc用于调用后端grpc服务
		web.Handler(webLib.NewGinRouter(prodSvc)),
		web.Registry(etcdReg), //将服务注册中心引入
	)

	httpSvc.Init()
	httpSvc.Run() //启动
}
