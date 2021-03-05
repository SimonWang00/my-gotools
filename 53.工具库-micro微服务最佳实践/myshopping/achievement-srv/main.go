package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"myshopping/achievement-srv/repository"
	"myshopping/achievement-srv/subscriber"
)

// 这里是我内网的mongo地址，请根据你得实际情况配置，推荐使用dockers部署
const MONGO_URI = "mongodb://admin:000000@127.0.0.1:27017"

// 连接到MongoDB
func connectMongo(uri string, timeout time.Duration) (*mongo.Client, error) {
	// context是go的标准库包，不是很清楚这个包的作用，文档上面也没有写很清楚，知道的朋友希望指点
	//一下
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	// 连接uri
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

// task-srv服务
func main() {
	// 在日志中打印文件路径，便于调试代码
	log.SetFlags(log.Llongfile)

	conn, err := connectMongo(MONGO_URI, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Disconnect(context.Background())

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.achievement"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
		micro.Broker(nats.NewBroker(broker.Addrs("nats://127.0.0.1:4222"))),
	)

	// Initialise service
	service.Init()

	// Register Handler
	handler := &subscriber.AchievementSub{
		Repo: &repository.AchievementRepoImpl{
			Conn: conn,
		},
	}
	// 这里的topic注意与task-srv注册的要一致
	if err := micro.RegisterSubscriber("go.micro.service.task.finished", service.Server(), handler); err != nil {
		log.Fatal(errors.WithMessage(err, "subscribe"))
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
