package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/28

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var (
	GlobalClient *redis.Client		// redis连接客户端接口
)

func initRedis()  error{
	GlobalClient = redis.NewClient(
		&redis.Options{
			Addr:         "127.0.0.1:6379",		// redis conn url
			DialTimeout:  10 * time.Second,		// 连接超时时间
			ReadTimeout:  30 * time.Second,		// 读超时
			WriteTimeout: 30 * time.Second,		// 写超时
			Password:     "",					// 登录密码
			PoolSize:     10,					// 线程池
			DB:           0,					// db
		},
	)
	err := GlobalClient.Ping().Err()
	if nil != err {
		log.Fatalf("connect redis failed! error(%v)", err)
		return err
	}
	log.Println("connect redis success!")
	GlobalClient.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			log.Printf("starting process:<%s>\n", cmd)
			err := old(cmd)
			log.Printf("finished process:<%s>\n", cmd)
			return err
		}
	})
	return nil
}


//get 查set
func Get(key string) *redis.StringCmd {
	cmd := redis.NewStringCmd("get", key)
	err := GlobalClient.Process(cmd)
	if err != nil{
		return nil
	}
	return cmd
}


//set 赋值
func Set( key string, value string) *redis.StringCmd {
	cmd := redis.NewStringCmd("set", key, value)
	err := GlobalClient.Process(cmd)
	if err != nil{
		return nil
	}
	return cmd
}

// pop
func Pop(key string) *redis.StringCmd {
	cmd := redis.NewStringCmd("pop", key)
	err := GlobalClient.Process(cmd)
	if err != nil{
		return nil
	}
	return cmd
}



func main() {
	initRedis()
	_, errSet := Set( "myKey", "myValue2").Result()
	if errSet != nil{
		log.Fatalf("exec redis Set function failed! error(%v)", errSet)
	}
	cmd := Get("myKey")
	log.Println(cmd)
	// 过期时间为10s， 后面自动删除
	cmd = redis.NewStringCmd("set", "myKey1", "123", "ex", "10")
	errSet = GlobalClient.Process(cmd)
	log.Println(errSet)
	//get expire
	cmd1 := redis.NewIntCmd("ttl", "myKey1")
	GlobalClient.Process(cmd1)
	expire, errEx := cmd1.Result()
	if errEx != nil {
		log.Println("ttl failed.", errEx)
	} else {
		log.Println("expire of key is", expire)
	}
}
