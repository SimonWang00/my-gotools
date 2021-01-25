package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	registry "my-gotools/36.工具库-etcd服务注册与发现/registry"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/7



func main () {
	var op = registry.Options{
		Name: "svc.info",
		Ttl:  60,
		Config: clientv3.Config{
			Endpoints:   []string{"http://127.0.0.1:2379/"},
			DialTimeout: 5 * time.Second},
	}
	for i := 1; i <= 3; i++ {
		// 创建注册
		r, err := registry.NewRegistry(op)
		if err != nil {
			fmt.Errorf("注册失败：%v", err.Error())
			return
		}
		// 注册节点
		addr := fmt.Sprintf("127.0.0.1:%d%d%d%d", i, i, i, i)
		err = r.RegistryNode(registry.PutNode{Addr:addr })
		if err != nil {
			fmt.Errorf("注册失败：%v", err.Error())
			return
		}
		if i == 3 {
			go func() {
				time.Sleep(time.Second * 30)
				r.UnRegistry()
			}()
		}
	}
	time.Sleep(time.Second * 5)
}
