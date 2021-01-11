package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/11

import (
	"log"
	middleware "my-gotools/42.语法糖-使用接口实现日志中间件/middleware"
)

func main() {
	src := middleware.NewService("")
	ret := src.Add(10, 20)
	log.Println("Add result:", ret)
}
