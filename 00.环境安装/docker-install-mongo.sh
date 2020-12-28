# 1. 查看可用版本
docker search mongo

# 2. 取最新版的 mongo 镜像
docker pull mongo:latest

# 3. 查看本地镜像
docker images

# 4. 运行容器 --auth：需要密码才能访问容器服务。
docker run -itd --name mongo -p 27017:27017 mongo --auth

# 5. 安装成功
docker exec -it mongo mongo admin
# 创建一个名为 admin，密码为 123456 的用户。
#>  db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
# 尝试使用上面创建的用户信息进行连接。
#> db.auth('admin', '123456')