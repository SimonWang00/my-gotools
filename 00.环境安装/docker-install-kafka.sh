# 安装Kafka之情，请先安装zookeeper
# 1. 查看可用版本
docker search kafka

# 2. 取最新版的 kafka 镜像
docker pull wurstmeister/kafka

# 3. 查看本地镜像
docker images

# 4. 运行容器, 10.8.30.55是你的本机IP
docker run -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=10.8.30.55:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://127.0.0.1:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -v /etc/localtime:/etc/localtime wurstmeister/kafka

# 5. 安装成功
docker exec -it kafka /bin/sh

#进入路径：/opt/kafka_2.11-2.0.0/bin下
#运行kafka生产者发送消息
./kafka-console-producer.sh --broker-list localhost:9092 --topic sun

#发送消息,消息内容
{"datas":[{"channel":"","metric":"temperature","producer":"SimonWang00","sn":"IJA0101-00002245","time":"1613207156000","value":"80"}],"ver":"1.0"}

#接收消息，运行kafka消费者接收消息
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sun --from-beginning
