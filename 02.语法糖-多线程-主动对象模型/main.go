package main

import (
	"fmt"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/25

// Go并发设计模式之 Active Object
func main() {
	// 生产者,状态是共享的,同个Service对象
	service := NewService(1)
	for i := 0;i <20 ;i++  {
		// 入列
		service.AddFunction()
		// 查看最后一个线程的值
		service.lastThreadValue()
	}
	// 注意最终的service.value不是20 , 而是18或者19, 队列queue中确是1-20
	// 线程与线程之间是独立的, 最后一个生产者先执行完
	fmt.Println("最终service值:", service.value)
	fmt.Println("执行完毕!")
	time.Sleep(10*time.Second)
}


// 查看生产者最后一个线程的值,不一定是20
func (info *Service) lastThreadValue() {
	fmt.Println(info.value) //不是同步返回值
}
