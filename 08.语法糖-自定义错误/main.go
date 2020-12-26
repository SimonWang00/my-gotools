package main

import (
	"errors"
	"fmt"
	"time"
)

/**
 * @Author: SimonWang00
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/12/26 14:10
 */


// 自定义错误
type MyError struct{
	Msg string
	Err error
	time.Time
}

// 自己实现Error
func (myerror MyError) Error() string  {
	return fmt.Sprintf("%v %v %v", myerror.Msg, myerror.Time, myerror.Err)
}

// 初始化
func NewMyError() error {
	return MyError{
		Msg:  "MyError",
		Time: time.Now(),
	}
}

// 默认错误类
func (myerror MyError) Unwrap() error {
	return myerror.Err
}

func TestMyError()  error{
	err := NewMyError()
	err = fmt.Errorf("自定义错误%w", err)
	return err
}


func main() {
	err := TestMyError()
	fmt.Println("TestMyError 的错误：", err)
	var myerror MyError
	// 查询myerror是否自定义了错误，并解析数据
	fmt.Println(errors.As(err, &myerror))
	fmt.Println(myerror)
	// 是否包含了该错误
	fmt.Println(errors.Is(err, myerror))
	// 去掉最上的错误
	fmt.Println(errors.Unwrap(myerror))
}
