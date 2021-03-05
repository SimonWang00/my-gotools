package main

//File  : hellocli.go.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/4


import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"log"
	pb "myshopping/hello/proto/hello"
)

func main() {
	// 这里以HelloService默认提供的Call接口调用为例示范服务的调用
	// 可以看到他的调用就像调用本地方法一样，go-micro为我们隐藏了背后的服务注册、发现、负载均衡以及网络操作
	testCallFunc()

	// 这里示范消息的发送
	testSendMessage()
}
func testCallFunc(){
	// 获取hello服务
	// 这里第一个参数"go.micro.service.hello"必须与hello-service注册信息一致
	// 一般由micro生成的项目默认服务名为：{namespace 默认[go.micro]}.{type 默认[service]}.{项目名}组成
	// 如果要修改默认值，在生成项目时可以这样： micro --namespace=XXX --type=YYYY ZZZZ
	// 当然也可以直接修改main.go中micro.Name("go.micro.service.hello")的内容
	helloService := pb.NewHelloService("go.micro.service.hello", client.DefaultClient)

	// 默认生成的hello服务中自带三个接口: Call() Stream() PingPong(),分别对应参数调用、流传输和心跳
	resp, err := helloService.Call(context.Background(), &pb.Request{
		Name: "my first go micro~~",
	})
	if err != nil {
		log.Panic("call func", err)
	}
	log.Println("call func success!", resp.Msg)
}

func testSendMessage(){
	// 消息主题，定义规则与服务一致
	// 同样，也可以修改main.go的micro.RegisterSubscriber("go.micro.service.hello", service.Server(), new(subscriber.Hello))
	const topic = "go.micro.service.hello"
	// 获取消息发送接口，这里我一直使用的时micro.NewPublisher()
	// 但在写文时发现NewPublisher()已经被废止，改为NewEvent()，二者参数和返回值一致
	event := micro.NewEvent(topic, client.DefaultClient)
	if err := event.Publish(context.Background(), &pb.Message{
		Say: "hello server!",
	}); err != nil {
		log.Panic("send msg", err)
	}
	log.Println("send msg success!")
}
