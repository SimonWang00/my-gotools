package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"myshopping/hello/handler"
	"myshopping/hello/subscriber"

	hello "myshopping/hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
