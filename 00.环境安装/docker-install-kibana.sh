# 1. 查看可用版本
docker search kibana

# 2. 取最新版的 kibana 镜像
docker pull kibana:7.6.2

# 3. 查看本地镜像
docker images

# 4. 运行容器
docker run --name kibana -e ELASTICSEARCH_URL=http://127.0.0.1:9200 -p 5601:5601 -d f29a1ee41030

# 5. 浏览器测试访问
访问 http://127.0.0.1:5601