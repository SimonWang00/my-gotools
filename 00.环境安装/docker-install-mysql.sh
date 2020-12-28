# 1. 查看可用版本
docker search mysql

# 2. 取最新版的 mysql 镜像
docker pull mysql:latest

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run -itd --name mysql-test -p 6379:6379 mysql

# 5. 安装成功
docker exec -it mysql-test /bin/bash