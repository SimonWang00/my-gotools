module my-gotools

go 1.14

require (
	cloud.google.com/go/firestore v1.4.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/PuerkitoBio/goquery v1.6.0
	github.com/RedisBloom/redisbloom-go v0.9.1
	github.com/Shopify/sarama v1.27.2
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/apache/rocketmq-client-go v1.2.4
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.44.0
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-sqlite/sqlite3 v0.0.0-20180313105335-53dd8e640ee7 // indirect
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.2
	github.com/gonuts/binary v0.2.0 // indirect
	github.com/google/uuid v1.1.2
	github.com/google/wire v0.4.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/hpcloud/tail v1.0.0
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/keybase/go-keychain v0.0.0-20201121013009-976c83ec27a6 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nats-io/stan.go v0.7.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic/v7 v7.0.22
	github.com/onyas/go-browsercookie v0.0.0-20190726085653-a0ba54f39260
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.9.0
	github.com/robfig/cron v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.2.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.5.1
	github.com/tsuna/gohbase v0.0.0-20201125011725-348991136365
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/urfave/cli v1.22.5
	github.com/urfave/cli/v2 v2.3.0
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	go.mongodb.org/mongo-driver v1.4.4
	go.uber.org/ratelimit v0.1.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	google.golang.org/api v0.36.0
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gopkg.in/go-oauth2/redis.v3 v3.2.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/oauth2.v3 v3.12.0
	github.com/apache/rocketmq-client-go/v2 v2.1.0-rc5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
