package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/config"
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/tool"
	"time"
)

func initRedis() {

	redisDb = redis.NewClient(
		&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()),
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     config.GetRedisConfig().GetPass(),
			PoolSize:     config.GetRedisConfig().GetMaxOpen(),
		},
	)
	err = redisDb.Ping().Err()
	if nil != err {
		tool.GetLogger().Error("ping redis err:", zap.Error(err))
		panic(err)
	}
	tool.GetLogger().Debug("redis success : " + fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()))

}

func closeRedis() {
	if redisDb != nil {
		_ = redisDb.Close()
	}
}
