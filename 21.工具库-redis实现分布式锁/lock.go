package main

//File  : lock.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/29

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

var (
	lock = sync.Mutex{}
)

// redis锁
type RedisLock struct {
	conn    *redis.Client
	timeout time.Duration
	key     string
	val     string
}

// 初始化
func NewRedisLock(conn *redis.Client, key, val string, timeout time.Duration) *RedisLock {
	return &RedisLock{conn: conn, timeout: timeout, key: key, val: val}
}


//setnx 实现当key存在时失败 , 保证互斥性
func (redisLock *RedisLock) Lock() bool  {
	lock.Lock()
	defer lock.Unlock()
	result, err := redisLock.conn.SetNX(redisLock.key, 1, 10*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}


func (redisLock *RedisLock) Unlock() int64 {
	result, err := redisLock.conn.Del(redisLock.key).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}


func (redisLock *RedisLock) GetLockKey() string {
	return redisLock.key
}

func (redisLock *RedisLock) GetLockVal() string {
	return redisLock.val
}
