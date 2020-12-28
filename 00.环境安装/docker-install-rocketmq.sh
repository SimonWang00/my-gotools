# 1. 查看可用版本
docker search rocketmq

# 2. 取最新版的 rocketmq 镜像
docker pull rocketmqinc/rocketmq:4.4.0

# 3. 查看本地镜像
docker images

# 4. 安装 Namesrv
docker run -d -p 9876:9876 --name rmqnamesrv -e "MAX_POSSIBLE_HEAP=100000000" rocketmqinc/rocketmq:4.4.0 sh mqnamesrv


# 5. 安装 broker
docker run -d -p 10911:10911 -p 10909:10909 --name rmqbroker --link rmqnamesrv:namesrv -e "NAMESRV_ADDR=namesrv:9876" -e "MAX_POSSIBLE_HEAP=200000000" rocketmqinc/rocketmq:4.4.0 sh mqbroker -c /opt/rocketmq-4.4.0/conf/broker.conf


# 6. 安装 rocketmq 控制台
docker pull styletang/rocketmq-console-ng
docker run -d -e "JAVA_OPTS=-Drocketmq.namesrv.addr=10.8.30.55:9876 -Dcom.rocketmq.sendMessageWithVIPChannel=false" -p 8080:8080 -t styletang/rocketmq-console-ng


# 7. 安装成功
访问 http://127.0.0.1:8080/