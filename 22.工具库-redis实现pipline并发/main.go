package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: pipline并发模式
//Date  : 2021/1/4


/*
在服务端未响应时客户端继续向服务端发送请求的模式称为 Pipeline 模式。
因为减少等待网络传输的时间，Pipeline 模式可以极大的提高吞吐量，减少所需使用的 tcp 链接数。
pipeline 模式的 redis 客户端需要有两个后台协程程负责 tcp 通信，调用方通过 channel 向后台协程发送指令，
并阻塞等待直到收到响应，这是一个典型的异步编程模式。

Pipeline主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。
这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）

*/

var (
	RedisHost = "127.0.0.1:6379"
	Key		  = "pipe_test"
)

func main() {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:               RedisHost,
			Password:           "",
			DB:                 0,
			MaxRetries:         5,
			DialTimeout:        10*time.Second,
			ReadTimeout:        10*time.Second,
			WriteTimeout:       10*time.Second,
			PoolSize:           100,
			MinIdleConns:       10,
			PoolTimeout:        10,
			IdleTimeout:        10*time.Second,
		},
		)
	err := redisClient.Ping().Err()
	if nil != err {
		panic(err)
	}

	//redis乐观锁支持，可以通过watch监听一些Key, 如果这些key的值没有被其他人改变的话，才可以提交事务。
	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error{
		result, err := tx.Get(Key).Result()
		if err != nil && err != redis.Nil {
			return err
		}
		// 这里可以处理业务
		log.Println("业务逻辑：", result)
		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值，redis实际可以发多个命令
			pipe.Set(Key, "order ticket: 20210104112921", 100000000000)
			return nil
		})
		return err
	}

	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
	err = redisClient.Watch(fn, Key)
	if nil != err {
		panic(err)
	}
}
