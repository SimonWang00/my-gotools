# 1. 查看可用版本
docker search quay.io/coreos/etcd

# 2. 取最新版的 etcd 镜像
docker pull quay.io/coreos/etcd:latest

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run -d -p 2379:2379 -p 2380:2380 --name my-etcd quay.io/coreos/etcd /usr/local/bin/etcd --name s1 --data-dir /etcd-data --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380 --initial-cluster s1=http://0.0.0.0:2380 --initial-cluster-token tkn --initial-cluster-state new

# 5. 安装成功
docker exec -it etcd-test /bin/bash