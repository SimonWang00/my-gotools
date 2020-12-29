package main

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
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

var (
	locker 	  = sync.Mutex{}
)

// redis锁
type RedisLock struct {
	conn    *redis.Client		// redis连接
	timeout time.Duration		// 连接超时时间
	key     string
	val     string
}


// 初始化
func NewRedisLock(conn *redis.Client, key, val string, timeout time.Duration) *RedisLock {
	return &RedisLock{conn: conn, timeout: timeout, key: key, val: val}
}


//setnx 实现当key存在时失败 , 保证互斥性
func (redisLock *RedisLock) Lock() bool  {
	locker.Lock()
	defer locker.Unlock()
	// 独占期3秒钟，每个线程占用资源的时间
	result, err := redisLock.conn.SetNX(redisLock.key, 1, 1*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

// 释放资源
func (redisLock *RedisLock) Unlock() int64 {
	result, err := redisLock.conn.Del(redisLock.key).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}


func (redisLock *RedisLock) GetLockKey() string {
	// dosmoething
	return redisLock.key
}

func (redisLock *RedisLock) GetLockVal() string {
	// dosmoething
	return redisLock.val
}


// 分布式锁的接口
type RedisLockServer interface {
	Lock()		 bool
	Unlock()	 int64
	GetLockKey() string
	GetLockVal() string
}


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
	GrapLock(redisLock)
	select {}
}

// 抢锁
func GrapLock(lock RedisLockServer) {
	go func() {
		for {
			// 模拟每1秒钟竞争一次资源
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