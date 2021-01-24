package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/7

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MyContext struct {
	context.Context
	Gin *gin.Context
}

type MyHandleFunc func (c *MyContext)

func WithMyContext(myHandle MyHandleFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		fmt.Println(<- chan struct{}(nil))

		// 可以在gin.Context中设置key-value
		c.Set("trace", "假设这是一个调用链追踪sdk")

		// 全局超时控制
		timeoutCtx, cancelFunc := context.WithTimeout(c, 5 * time.Second)
		defer cancelFunc()

		// ZDM上下文
		myCtx := MyContext{Context: timeoutCtx, Gin: c}

		// 回调接口
		myHandle(&myCtx)
	}
}

// 模拟一个MYSQL查询
func dbQuery(ctx context.Context, sql string) {
	// 模拟调用链埋点
	trace := ctx.Value("trace").(string)

	// 模拟长时间逻辑阻塞, 被context的5秒超时中断
	//<- ctx.Done()

	fmt.Println(trace)
}

// 有bug
func main() {
	r := gin.New()

	r.GET("/test", WithMyContext(func(c *MyContext) {
		// 业务层处理
		dbQuery(c, "select * from user")
		// 调用gin应答
		c.Gin.String(200, "请求完成!")
	}))

	r.Run("127.0.0.1:8090")
}