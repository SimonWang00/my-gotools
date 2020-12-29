package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/29

/*
redis实现分布式锁主要靠setnx命令

1. 当key存在时失败 , 保证互斥性

2.设置了超时 , 避免死锁

3.利用mutex保证当前程序不存在并发冲突问题
*/

func main() {
	GlobalClient := redis.NewClient(
		&redis.Options{
			Addr:         "127.0.0.1:6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     "",
			PoolSize:     10,
			DB:           0,
		},
	)
	ping, err := GlobalClient.Ping().Result()
	if nil != err {
		panic(err)
	}
	log.Println("ping", ping)
	redisLock := NewRedisLock(GlobalClient, "test", "1", time.Second*3)
	InitRedis(redisLock)
	select {}
}


func InitRedis(lock RedisLockServer) {
	go func() {
		for {
			time.Sleep(time.Second)
			if ! lock.Lock() {
				log.Println("获取锁失败")
				//获取锁失败间隔一段时间重试
			} else {
				log.Println("获取锁成功")
				key := lock.GetLockKey()
				val := lock.GetLockVal()
				log.Println("获取是key值：", key)
				log.Println("获取是val值：", val)
			}
		}
	}()
}