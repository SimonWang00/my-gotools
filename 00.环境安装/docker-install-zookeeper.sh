# 1. 查看可用版本
docker search zookeeper

# 2. 取最新版的 zookeeper 镜像
docker pull wurstmeister/zookeeper

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run -itd --name zookeeper  -p 2181:2181 wurstmeister/zookeeper

# 5. 安装成功
docker exec -it zookeeper /bin/bash
