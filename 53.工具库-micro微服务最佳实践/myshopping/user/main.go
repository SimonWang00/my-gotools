package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"myshopping/user/handler"
	"myshopping/user/model"
	//"myshopping/user/subscriber"
	"myshopping/user/repository"

	user "myshopping/user/proto/user"
)


// CreateDbConnection 创建mysql连接
func CreateDbConnection() (*gorm.DB, error)  {
	host := "127.0.0.1"
	username := "root"
	dbName := "myshopping"
	password := "000000"
	return gorm.Open("mysql",
		fmt.Sprintf(
			"%v:%v@tcp(%v:3306)/%v?charset=utf8&parseTime=True&loc=Local",
			username, password, host, dbName,
		),
	)
}


func main() {

	db, err := CreateDbConnection()
	defer db.Close()

	if err != nil{
		log.Fatalf("connection error : %v \n" , err)
		return
	}

	db.AutoMigrate(&model.User{})
	userDb := &repository.UserDb{db}

	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Registry(etcdReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//_ = user.RegisterUserServiceHandler(service.Server(), &handler.User{userDb})
	_ = user.RegisterUserServiceHandler(service.Server(), &handler.User{userDb})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
