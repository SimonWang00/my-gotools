package src

import "fmt"

type MyClass struct {
	Addr string
	Data string
}

func NewMyClass()*MyClass  {
	var cls =new(MyClass)
	cls.Addr="127.0.0.1"
	cls.Data="我是测试信息i"
	return cls
}

func (cls *MyClass) Name() string  {
	fmt.Println("调用Name")
	return cls.Data
}

func (cls *MyClass) Run()  {
	fmt.Println("调用RUN")
}


