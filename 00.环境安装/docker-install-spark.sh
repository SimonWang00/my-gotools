# 1. 取最新版的 spark 镜像
docker pull p7hb/docker-spark:2.1.0

# 2. 查看本地镜像
docker images

# 3. 运行容器
docker run -it -p 4040:4040 -p 8080:8080 -p 8081:8081 -h spark --name=spark p7hb/docker-spark:2.1.0

# 4. 安装成功
docker exec -it spark /bin/bash