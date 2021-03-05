package hystrix

//File  : hystrix.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/5

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"log"
	pb "myshopping/task-api/proto/task"
)

type clientWrapper struct {
	client.Client
}

// 这是我们真正需要实现的方法，实现一下熔断器的降级
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	// 命令名的写法参考官方插件，服务名和方法名拼接
	name := req.Service() + "." + req.Endpoint()
	// 自定义当前命令的熔断配置，除了超时时间还有很多其他配置请自行研究
	// 这些配置在wrapper调用时才执行，因此具有最高的优先级
	// ---如果打算使用全局参数配置，请注释掉下面几行---
	config := hystrix.CommandConfig{
		Timeout: 2000,		// 2秒的超时
	}
	hystrix.ConfigureCommand(name, config)
	// ---如果打算使用全局参数配置，请注释掉上面几行---

	return hystrix.Do(name,
		func() error {
			// 这里调用了真正的服务
			return c.Client.Call(ctx, req, rsp, opts...)
		},


		// 这个叫做降级函数，用来自定义调用失败后的处理
		// 一般我们可以选择返回特定错误信息，或者返回预设默认值
		// 这个方法尽量简单，尽量不要加入额外的服务调用和IO操作，避免降级函数自身异常
		func(err error) error {
			// 因为是示例程序，只处理请求超时这一种错误的降级，其他错误仍抛给上级调用函数
			if err != hystrix.ErrTimeout {
				return err
			}
			// 直接返回默认的查询结果并记录日志
			switch r:=rsp.(type) {
			// 这个服务我只实现了search一个接口的调用*pb.SearchResponse
			case *pb.SearchResponse:
				log.Print("search task fail: ", err)
				*r = pb.SearchResponse{
					PageSize: 20,
					PageCode: 1,
					SortBy:   "createTime",
					Order:    -1,
					Rows: []*pb.Task{
						{Body: "超时熔断后，自定义的结果返回"},
					},
				}
			default:
				log.Print("unknown err: ", err)
			}
			return nil
		})
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}