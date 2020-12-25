package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: 上下文管理
//Date  : 2020/12/25

var siganChannel = make(chan os.Signal, 1)

func main() {
	// 取消信号
	//ContextWithCancel()
	// 超时信号
	//ContextWithTimeout()
	// 截至日期信号
	ContextWithDeadline()
}


// 上下文取消信号
func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println(ctx.Value("SimonWang00"))
				return
			}
		}
	}()
	log.Println("上下文开始的地方")
	time.AfterFunc(5*time.Second, func() {
		ctx = context.WithValue(ctx, "SimonWang00", "5秒后调用执行cancel()方法")
		cancel()
		log.Println("上下文结束啦~~")
	})
	// 退出程序
	Exit()
}


// 上下文超时信号
func ContextWithTimeout() {
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println(time.Since(start))
				log.Println(time.Now().Format("2006-01-02 15:04:05"))
				return
			}
		}
	}()
	Exit()
}

// 上下文截至日期
func ContextWithDeadline() {
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	start := time.Now()
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			log.Println(time.Since(start))
			log.Println(time.Now().Format("2006-01-02 15:04:05"))
			return
		}
	}()
	Exit()
	req, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	req = req.WithContext(ctx)
	client := &http.Client{}
	res, err := client.Do(req)
	log.Println(err, res )
}

func Exit() {
	signal.Notify(siganChannel, os.Kill, os.Interrupt)
	<-siganChannel
}