package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 配置
const (
	addr   = "127.0.0.1:6379"
	db	   = 0
	passwd = ""
	channel= "PUBLIC"
)

// ch 用于保证发布线程在订阅线程启动成功后才开始发布消息
var ch = make(chan int)

// ConsumeFunc consumes message at the channel.
type ConsumeFunc func(channel string, message []byte) error

// RedisClient represents a redis client with connection pool.
type RedisClient struct {
	pool *redis.Pool
}

// NewRedisClient returns a RedisClient.
func NewRedisClient(addr string, db int, passwd string) *RedisClient {
	pool := &redis.Pool{
		MaxIdle:     10,
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

// Close closes connection pool.
func (r *RedisClient) Close() error {
	err := r.pool.Close()
	return err
}

// Publish publishes message to channel.
func (r *RedisClient) Publish(channel, message string) (int, error) {
	c := r.pool.Get()
	defer c.Close()
	n, err := redis.Int(c.Do("PUBLISH", channel, message))
	if err != nil {
		return 0, fmt.Errorf("redis publish %s %s, err: %v", channel, message, err)
	}
	return n, nil
}

// Subscribe subscribes messages at the channels.
// 可支持多个channel消息的订阅
func (r *RedisClient) Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error {
	psc := redis.PubSubConn{Conn: r.pool.Get()}

	log.Printf("redis pubsub subscribe channel: %v", channel)
	if err := psc.Subscribe(redis.Args{}.AddFlat(channel)...); err != nil {
		return err
	}
	done := make(chan error, 1)
	// start a new goroutine to receive message
	go func() {
		defer psc.Close()
		for {
			// 检测接收类型
			switch msg := psc.Receive().(type) {
			case error:
				done <- fmt.Errorf("redis pubsub receive err: %v", msg)
				return
			case redis.Message:
				// 消费消息
				if err := consume(msg.Channel, msg.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				if msg.Count == 0 {
					// all channels are unsubscribed
					done <- nil
					return
				}
			}
		}
	}()

	ch <- 0

	// health check 每分钟进行定时检查
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():	// 取消
			if err := psc.Unsubscribe(); err != nil {
				return fmt.Errorf("redis pubsub unsubscribe err: %v", err)
			}
			return nil
		case err := <-done:	// 异常
			return err
		case <-tick.C:		// 定时检测
			if err := psc.Ping(""); err != nil {
				return err
			}
		}
	}

}

// myConsumer 这里是业务逻辑
func myConsumer(channel string, message []byte) error {
	log.Printf("receive message[%s] at the channel[%s]\n", string(message), channel)
	return nil
}



func main() {
	redisClient := NewRedisClient(addr, db, passwd)
	// 连续发布三个消息
	go func() {
		var subscriber int
		<-ch
		for i := 0; i < 5; i++ {
			subscriber, _ = redisClient.Publish(channel, "hello SimonWang00"+strconv.Itoa(i))
			log.Printf("there is %d subscriber.\n", subscriber)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	// 订阅消息
	err := redisClient.Subscribe(ctx,
		func(channel string, message []byte) error {
			log.Printf("receive message[%s] from the channel[%s]\n", string(message), channel)
			if string(message) == "goodbye" {
				cancel()
			}
			return nil
		},
		channel)
	if err != nil {
		fmt.Printf("get error: %v\n", err)
	}
	redisClient.Close()

	return
}

