package main

import (
	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	//"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"log"
	"myshopping/task-api/common/tracer"
	pb "myshopping/task-api/proto/task"
	"myshopping/task-api/wrapper/breaker/hystrix"
)

const (
	ServerName = "go.micro.api.task"
	JaegerAddr = "127.0.0.1:6831"
	ETCDAddr = "127.0.0.1:2379"
)

// task-srv服务的restful api映射
func main() {
	g := gin.Default()

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServerName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	// 之前我们使用client.DefaultClient注入到pb.NewTaskService中
	// 现在改为标准的服务创建方式创建服务对象
	// 但这个服务并不真的运行（我们并不调用他的Init()和Run()方法）
	// 如果是task-srv这类本来就是用micro.NewService服务创建的服务
	// 则直接增加包装器，不需要再额外新增服务
	app := micro.NewService(
		micro.Name(ServerName),
		micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
		micro.WrapClient(
			// 引入hystrix包装器
			hystrix.NewClientWrapper(),
			// 配置链路追踪为jaeger
			opentracing.NewClientWrapper(jaegerTracer),
		),
	)

	// 自定义全局默认超时时间和最大并发数
	hystrixGo.DefaultSleepWindow = 200
	hystrixGo.DefaultMaxConcurrent = 3

	// 针对指定服务接口使用不同熔断配置
	// 第一个参数name=服务名.接口.方法名，这并不是固定写法，而是因为官方plugin默认用这种方式拼接命令name
	// 之后我们自定义wrapper也同样使用了这种格式
	// 如果你采用了不同的name定义方式则以你的自定义格式为准
	hystrixGo.ConfigureCommand("go.micro.service.task.TaskService.Search",
		hystrixGo.CommandConfig{
			MaxConcurrentRequests: 50,
			Timeout:               2000,
		})

	// 注册默认DefaultClient
	//cli := pb.NewTaskService("go.micro.service.task", client.DefaultClient)
	cli := pb.NewTaskService(ServerName, app.Client())

	service := web.NewService(
		web.Name(ServerName),
		web.Address(":8888"),
		web.Handler(g),
		web.Registry(etcd.NewRegistry(registry.Addrs(ETCDAddr))),
	)

	v1 := g.Group("/task")
	{
		v1.GET("/search", func(c *gin.Context) {
			req := new(pb.SearchRequest)
			if err := c.ShouldBind(req); err != nil {
				c.JSON(200, gin.H{
					"code": "500",
					"msg":  "bad param",
				})
				return
			}
			if resp, err := cli.Search(c, req); err != nil {
				c.JSON(200, gin.H{
					"code": "500",
					"msg":  err.Error(),
				})
			} else {
				c.JSON(200, gin.H{
					"code": "200",
					"data": resp,
				})
			}
		})
	}
	// 配置web路由
	router(g, cli)
	_ = service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


// web路由实现部分略
// 详情见https://gitee.com/xieyu1989/go-micro-study-notes/tree/master/go-todolist4
func router(g *gin.Engine, taskService pb.TaskService){}