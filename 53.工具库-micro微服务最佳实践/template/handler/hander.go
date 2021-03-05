package handler

import (
	"my-gotools/53.工具库-micro微服务最佳实践/template/model/user"
	"sync"
)
var (
	err error
	server *Service
	m sync.Mutex
)

type Service struct {
	userServer user.Service
}

func GetService()*Service  {
	return server
}

func Init() {
	m.Lock()
	defer m.Unlock()
	server =new(Service)
	// 获取服务实例
	server.userServer, err = user.GetService()
	checkErr(err)

}

func checkErr(err error)  {
	if err!=nil {
		panic(err)
	}
}
