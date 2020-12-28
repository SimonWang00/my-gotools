# 1. 查看可用版本
docker search quay.io/coreos/etcd

# 2. 取最新版的 etcd 镜像
docker pull quay.io/coreos/etcd:latest

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run -itd --name etcd-test -p 2379:2379 quay.io/coreos/etcd

# 5. 安装成功
docker exec -it etcd-test /bin/bash