package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	hystrix "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"myshopping/task-srv/common/tracer"
	"myshopping/task-srv/handler"
	pb "myshopping/task-srv/proto/task"
	"myshopping/task-srv/repository"

	"time"
)

const (
	ServerName = "go.micro.service.task"
	MONGO_URI = "mongodb://admin:000000@localhost:27017/"
	ETCD_URI = "localhost:2379"
	JaegerAddr = "localhost:6831"
)

// 连接到MongoDB
func connectMongo(uri string, timeout time.Duration) (*mongo.Client, error) {
	// context是go的标准库包，不是很清楚这个包的作用，文档上面也没有写很清楚，知道的朋友希望指点一下
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	// 构建mongo连接可选属性配置
	opt := new(options.ClientOptions)
	// 设置最大连接的数量
	opt = opt.SetMaxPoolSize(uint64(10))
	// 设置连接超时时间 5000 毫秒
	du, _ := time.ParseDuration("5000")
	opt = opt.SetConnectTimeout(du)
	// 设置连接的空闲时间 毫秒
	mt, _ := time.ParseDuration("5000")
	opt = opt.SetMaxConnIdleTime(mt)
	// 开启驱动
	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), opt)
	if err != nil {
		return nil, errors.WithMessage(err, "create mongo connection session")
	}
	// 注意，在这一步才开始正式连接mongo
	_ = MongoClient.Ping(ctx, readpref.Primary())
	return MongoClient, nil
}


func main() {
	// 在日志中打印文件路径，便于调试代码
	log.SetFlags(log.Llongfile)

	conn, err := connectMongo(MONGO_URI, 10 *time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Disconnect(context.Background())

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServerName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Version("latest"),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379
		micro.Registry(etcd.NewRegistry(
			// 地址是我本地etcd服务器地址，不要照抄
			registry.Addrs(ETCD_URI),
		)),
		micro.Broker(nats.NewBroker(broker.Addrs("nats://127.0.0.1:4222"))),	// nats消息组件
		//micro.Broker(http.NewBroker(broker.Addrs("nats://127.0.0.1:4222"))),	// http 消息组件
		// 熔断器
		micro.WrapClient(hystrix.NewClientWrapper()),
		// 配置链路追踪为jaeger
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
	)

	// Initialise service
	service.Init()

	// Register Handler
	taskHandler := &handler.TaskHandler{
		TaskRepository: &repository.TaskRepositoryImpl{
			Conn: conn,
		},
		// 注入消息发送实例,为避免消息名冲突,这里的topic我们用服务名+自定义消息名拼出
		TaskFinishedPubEvent: micro.NewEvent("go.micro.service."+handler.TaskFinishedTopic, service.Client()),
	}

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.task", service.Server(), new(subscriber.TaskSrv))
	if err := pb.RegisterTaskServiceHandler(service.Server(), taskHandler); err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
