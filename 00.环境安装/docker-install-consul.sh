# 1. 拉取consul镜像
docker pull consul:latest

# 2. 启动consul
docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul agent -server -bootstrap-expect=1 -ui -bind=0.0.0.0 -client=0.0.0.0

# 3. 查看信息
我们可以打开浏览器：http://127.0.0.1:8500 来查看整个集群的信息
