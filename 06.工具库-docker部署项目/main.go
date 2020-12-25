package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// hello 调试请求
func hello(c *gin.Context)  {
	c.String(http.StatusOK,"Welcom to docker!")
}

func main() {
	//监听退出信号
	handleSigterm()
	// 启动web服务
	go startWebServer()

	// 阻塞等子线程，Block...
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

// handleSigterm 退出机制
func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}

// startWebServer 启动web服务
func startWebServer() {
	// 初始化配置
	g := gin.Default()
	g.GET("/hello",hello)	// 此处用于调试
	// 启动
	g.Run("0.0.0.0:8080")
}
