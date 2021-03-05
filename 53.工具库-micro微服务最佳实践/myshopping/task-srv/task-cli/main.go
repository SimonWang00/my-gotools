package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/4

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	pb "myshopping/task-srv/proto/task"
	"myshopping/task-srv/repository"
	"time"
)

const (
	ServerName = "go.micro.client.task"
	// 服务地址
	JaegerAddr = "127.0.0.1:6831"
)

// 插入测试任务
func insertTask(taskService pb.TaskService, body string, start, end int64) {
	_, err := taskService.Create(context.Background(), &pb.Task{
		UserId: "100",
		Body:      body,
		StartTime: start,
		EndTime:   end,
	})
	if err != nil {
		log.Fatal("create", err)
	}
	log.Println("create task success! ")
}

// 模拟client调用task-srv服务
func main() {
	// 在日志中打印文件路径，便于调试代码
	log.SetFlags(log.Llongfile)

	// 客户端也注册为服务
	server := micro.NewService(micro.Name(ServerName),
								micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379")), ),
		)
	server.Init()
	taskService := pb.NewTaskService(ServerName, server.Client())

	// 调用服务生成三条任务
	now := time.Now()
	insertTask(taskService, "完成学习笔记（一）", now.Unix(), now.Add(time.Hour*24).Unix())
	insertTask(taskService, "完成学习笔记（二）", now.Add(time.Hour*24).Unix(), now.Add(time.Hour*48).Unix())
	insertTask(taskService, "完成学习笔记（三）", now.Add(time.Hour*48).Unix(), now.Add(time.Hour*72).Unix())

	// 分页查询任务列表
	page, err := taskService.Search(context.Background(), &pb.SearchRequest{
		PageCode: 1,
		PageSize: 20,
	})
	if err != nil {
		log.Fatal("search1", err)
	}
	log.Println(page)

	// 更新第一条记录为完成
	row := page.Rows[0]
	if _, err = taskService.Finished(context.Background(), &pb.Task{
		Id:         row.Id,
		IsFinished: repository.Finished,
	}); err != nil {
		log.Fatal("finished", row.Id, err)
	}

	// 再次分页查询，校验修改结果
	page, err = taskService.Search(context.Background(), &pb.SearchRequest{})
	if err != nil {
		log.Fatal("search2", err)
	}
	log.Println(page)
}
