# 1. 查看可用版本
docker search hbase

# 2. 取 hbase 镜像
docker pull harisekhon/hbase:1.3

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run -d -h myhbase -p 2181:2181 -p 8080:8080 -p 8085:8085 -p 9090:9090 -p 9095:9095 -p 16000:16000 -p 16010:16010 -p 16201:16201 -p 16301:16301 --name hbase1.3 harisekhon/hbase:1.3

# 5. 安装成功
访问 http://127.0.0.1:16010/master-status