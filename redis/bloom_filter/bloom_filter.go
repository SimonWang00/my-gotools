package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//git clone https://github.com/RedisBloom/RedisBloom.git
//cd RedisBloom
//make //编译 生成so文件
//redis-server --loadmodule /path/to/rebloom.so

func main() {
	GlobalClient := redis.NewClient(
		&redis.Options{
			Addr:         "127.0.0.1:6379",		// redis conn url
			DialTimeout:  10 * time.Second,		// 连接超时时间
			ReadTimeout:  30 * time.Second,		// 读超时
			WriteTimeout: 30 * time.Second,		// 写超时
			Password:     "",
			PoolSize:     10,
			DB:           0,
		},
	)
	err := GlobalClient.Ping().Err()
	if nil != err {
		panic(err)
	}
	fmt.Println("connect redis success!")
	info:=redis.NewStatusCmd("bf.add", "bl", "1")
	_ = GlobalClient.Process(info)
	if err := info.Err(); err != nil {
		print(err)
	}
	info1:=redis.NewStatusCmd("bf.add", "bl", "2")
	_ = GlobalClient.Process(info1)
	if err := info1.Err(); err != nil {
		print(err)
	}
	info3:=redis.NewStatusCmd("bf.add", "bl", "3")
	_ = GlobalClient.Process(info3)
	if err := info3.Err(); err != nil {
		print(err)
	}
	info4:=redis.NewIntCmd("bf.exists", "bl", "6")
	_ = GlobalClient.Process(info4)
	if err := info4.Err(); err != nil {
		print(err)
	}
	v,err:=info4.Result()
	fmt.Println("err",err)
	fmt.Println("v",v)//存在 v==1  不存在 ==0
	//fmt.Println(GlobalClient.Get("mykey").String())
}
