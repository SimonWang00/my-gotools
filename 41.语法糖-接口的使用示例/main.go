package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/11

import (
	"fmt"
	"my-gotools/interface/src"
)

type CallObject struct {
	caller src.Agent
}

func (a *CallObject) TestAA() {
	fmt.Println("接口测试AA")
}

func main() {
	var A = &CallObject{}
	A.caller = src.NewMyClass()
	A.TestAA()
	A.caller.Name()
	A.caller.Run()
}
