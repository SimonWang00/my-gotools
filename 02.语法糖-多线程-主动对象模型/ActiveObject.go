package main

//File  : ActiveObject.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/25

import (
	"log"
	"time"
)


// 常量状态值
const (
	AddStatus = 1
	SubStatus = -1
)


// Service 定义类
type Service struct {
	value int
	queue chan int
}


// NewService 定义生产者模型
func NewService(buffer int) *Service {
	service := &Service{
		value:	  0,
		queue:    make(chan int, buffer),
	}
	go service.control()
	return service
}


// control 逻辑状态控制
func (service *Service) control()  {
	for{
		select {
		// 出列
		case stat := <- service.queue:
			// 执行状态逻辑
			if stat == AddStatus{
				service.value++
				log.Println("service exec value++ is ", service.value)
			} else if stat == SubStatus{
				service.value--
				log.Println("service exec value-- is ", service.value)
			}
		case  <- time.After(5 * time.Second):
			idle()
		}
	}
}


// AddFunction  AddStatus对应的执行逻辑
func (service *Service) AddFunction() {
	log.Println("入列执行Add")
	service.queue <- AddStatus
}


// SubFunction SubStatus对应的执行逻辑
func (service *Service) SubFunction() {
	log.Println("入列执行Sub")
	service.queue <- SubStatus
}


// idle 空闲状态
func idle()  {
	log.Println("service is idle 5s")
}