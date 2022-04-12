### 简介
配套《Go语言高并发与微服务实战》使用，具体讲解见图书13章
### 简介2
gateway 网关服务
oauth-service认证服务
admin 管理端服务
app 

### 依赖基础组件
- redis
- zookeeper
- git仓库
- consul

#### 部署
- 1 部署 consul 
参考书籍第六章6.3小节内容，安装部署 consul
- 2 部署 Redis,Zookeeper,MySQL。
参考对应组件的官方部署文档，安装完MySQL后，可以导入主目录下的seckill.sql
- 3 新建git repo
可以参考 https://gitee.com/cloud-source/config-repo 创建对应项目的文件，修改Redis，MySQL，Zookeeper等组件的配置
- 4 部署 Config-Service
参考书籍第八章8.3.1小节 在ch8-config文件夹下有 config-service项目，
在yml文件中配置对应的git项目地址和consul地址，构建并运行Java程序，将config-service注册到consul上
- 5 修改bootstrap文件
修改各个项目中的bootstrap.yml文件discover相关的consul地址和config-service的相关配置

## 安装注意事项
### zookeeper
在启动时可能会遇到这个错误  
>Error contacting service. It is probably not running.  

这说明你的zookeeper没启动，可以看安装目录/logs/***.out文件看日志 
1.会默认启用8080端口
修改zk配置文件，添加下面的配置，端口自定义  

    vim zookeeper/conf/zoo.cfg 
>admin.serverPort=8001

客户端链接
    ./zkCli.sh -server localhost:2181
#### zookeeper可视化工具
    git clone https://github.com/DeemOpen/zkui.git
    cd zkui
    mvn clean package -DskipTests=true
    cp config.cfg ./target/config.cfg
    java -jar ./target/zkui-2.0-SNAPSHOT-jar-with-dependencies.jar
访问http://localhost:9090，如能正常访问并看到如下界面，则运行正常
>角色为ADMIN
username: admin
password: manager
角色为USER
username: appconfig
password: appconfig

### 使用方法
#### sk-admin
打开consul，zookeeper，mysql,[zipkin-server](https://github.com/openzipkin/zipkin)


    go run ./sk-admin/main.go 