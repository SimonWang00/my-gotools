# 1. 查看可用版本
docker search elasticsearch

# 2. 取最新版的 elasticsearch 镜像
docker pull elasticsearch:7.6.2

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run --restart=always -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms512m -Xmx512m" --name elasticsearch-test --cpuset-cpus="1" -m 2G -d elasticsearch:7.6.2

# 5. 安装成功
docker exec -it elasticsearch-test /bin/bash