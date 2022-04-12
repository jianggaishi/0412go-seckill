###前置要求
1.zookeeper  
自行安装zookeeper，我是在本机单节点安装部署的。  
2.mysql
我已经安装完了,配置如下
>mysql:  
host: 39.99.214.230  
port: 3306  
user: root  
pwd: zhangpeng  
Db: seckill  

3.redis  
我也安装部署完了，配置如下
>redis:  
host: 39.99.214.230:6379  
password: zhangpeng  
db: 0  
Proxy2layerQueueName: proxy2layer  
Layer2proxyQueueName: Layer2proxy  
IdBlackListHash: IdBlackListHash  
IpBlackListHash: IpBlackListHash  
IdBlackListQueue: IdBlackListQueue  
IpBlackListQueue: IpBlackListQueue  
4. consul
我是本机运行的，自行去找github或者链接安装。  
###运行  
####运行admin  
1.修改项目配置文件
修改./seckill/seckill-admin/pkg/bootstrap/bootstrap_config.go  
把 initBootstrapConfig() 函数里  
viper.AddConfigPath（"path）换成你的本机文件地址  
2.执行./seckill-admin/sk-admin/main.go
3.结果
如果IDE没报错，consul成功注册了服务，且接口都可以用，说明admin启动成功。
接口部分在后面给出  
####运行app  
1.修改项目配置文件
修改./seckill/seckill-app/pkg/bootstrap/bootstrap_config.go  
把 initBootstrapConfig() 函数里  
viper.AddConfigPath（"path）换成你的本机文件地址  
2.执行./seckill-app/sk-app/main.go
3.结果
如果IDE没报错，consul成功注册了服务，且接口都可以用，说明app启动成功。
接口部分在后面给出  
####运行core  
1.修改项目配置文件  
修改./seckill/seckill-core/pkg/bootstrap/bootstrap_config.go  
把 initBootstrapConfig() 函数里  
viper.AddConfigPath（"path）换成你的本机文件地址  
2.执行./seckill-core/sk-core/main.go  
3.结果  
如果IDE没报错，且接口都可以用，说明app启动成功。
注：这里的consul不会成功注册core服务，但是不会影响功能。
接口部分在后面给出  
####接口  
这里我给的接口都是按照默认的port接口，你可以自行修改端口，后缀.yaml文档都是配置文档。  
路径：  
./seckill/接口测试文档/*

###docker启动  
默认用本地代码Dockerfile镜像，如果用dockerhub，选择对应版本2.0，把docker-compose.yaml文件里面build注释，images取消注释。  

    docker-compose up
### API接口
都在接口测试文档文件夹里，APIfox工具。
###docker 常用命令  

    sudo docker commit -m "seckill2.0" -a "zp" sk-admin zpskt/seckill_admin:2.0
    sudo docker push zpskt/seckill_admin:2.0

    