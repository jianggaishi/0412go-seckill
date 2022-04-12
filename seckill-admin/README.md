###使用说明  
本机构建docker
docker-compose里面四个服务，consul，mysql，zipkin，admin  
consul：服务发现  ui界面 http://localhost:8500  
mysql：用./sql文件下的sql创建数据库。 映射到本机3306了  
zipkin：链路追踪 ui界面 http://localhost:9411/zipkin/
### 启动  
    docker-compose up
###admin的API
url:http:localhost:9030/  
Get: /product/list
Get: /activity/list
Get: /health

###常用命令
