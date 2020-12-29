package main

//File  : lock_server.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/29

// 分布式锁的接口
type RedisLockServer interface {
	Lock()		 bool
	Unlock()	 int64
	GetLockKey() string
	GetLockVal() string
}
