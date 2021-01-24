package main

import (
	"my-gotools/35.工具库-基于go实现的websocket/gateway/ws"
)

func main() {
	ws.InitWsServer()
	select {}
}
