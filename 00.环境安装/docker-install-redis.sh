# 1. 查看可用版本
docker search redis

# 2. 取最新版的 Redis 镜像
docker pull redis:latest

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run -itd --name redis-test -p 6379:6379 redis

# 5. 安装成功
docker exec -it redis-test /bin/bash