package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"my-gotools/all_packaged_library/base/config"
	"my-gotools/all_packaged_library/base/tool"
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
