package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/4

const (
	Addr   = "127.0.0.1:6379"
	Db     = 0
	Passwd = ""
)


type RedisClient struct {
	pool *redis.Pool
}

func NewRedisClient(addr string, db int, passwd string) *RedisClient {
	pool := &redis.Pool{
		MaxIdle:  10,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialPassword(passwd), redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	log.Printf("new redis pool at %s", addr)
	client := &RedisClient{
		pool: pool,
	}
	return client
}


// publish
func (r *RedisClient) publish() (int, error) {
	c := r.pool.Get()
	defer c.Close()
	n, err := redis.Int(c.Do("PUBLISH", "SimonWang00", "hello"))
	if err != nil {
		return 0, err
	}
	return n, nil
}


// CheckErr 错误监听
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main()  {
	redisClient := NewRedisClient(Addr,Db,Passwd)
	psc := redis.PubSubConn{redisClient.pool.Get()}
	// 订阅channel
	psc.Subscribe("PUBLISH")
	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Println("每3秒钟发布一次")
			redisClient.publish()
		}
	}()

	//
	for {
		switch v := psc.Receive().(type) {
		case redis.Subscription:
			log.Printf("1  %s: %s %d\n", v.Channel, v.Kind, v.Count)
			break
		case redis.Message: //单个订阅subscribe
			log.Printf("2  %s: message: %s\n", v.Channel, v.Data)
			break
		case error:
			log.Println(v)
			break
		}
	}
}

