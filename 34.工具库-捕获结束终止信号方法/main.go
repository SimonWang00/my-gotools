package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/7


import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//syscall.SIGQUIT 用户发送QUIT字符(Ctrl+/)触发
//syscall.SIGTERM 结束程序(可以被捕获、阻塞或忽略)
//syscall.SIGINT 用户发送INTR字符(Ctrl+C)触发
func QuitSignal(quitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Printf("server start success pid:%d\n", os.Getpid())
	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			quitFunc()
			return
		default:
			return
		}
	}
}

// handleSigterm 退出机制
// 不带方法的
func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}


func main() {
	handleSigterm()
	QuitSignal(func() {
		fmt.Println("终止信号来了！")
	})
	select {
	}
}
