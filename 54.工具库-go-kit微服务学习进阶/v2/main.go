package main

import (
	"my-gotools/54.工具库-go-kit微服务学习进阶/v2/utils"
	"my-gotools/54.工具库-go-kit微服务学习进阶/v2/v2_endpoint"
	"my-gotools/54.工具库-go-kit微服务学习进阶/v2/v2_service"
	"my-gotools/54.工具库-go-kit微服务学习进阶/v2/v2_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := v2_service.NewService(utils.GetLogger())
	endpoints := v2_endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v2_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
