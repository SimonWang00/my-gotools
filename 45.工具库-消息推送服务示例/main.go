package main

import (
	"my-gotools/45.工具库-消息推送服务示例/gateway"
	"my-gotools/45.工具库-消息推送服务示例/logic"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go gateway.InitWsServer()
	go logic.InitHttpServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
