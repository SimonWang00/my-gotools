package main

import (
	"my-gotools/websocket/gateway/ws"
)

func main() {
	ws.InitWsServer()
	select {}
}
