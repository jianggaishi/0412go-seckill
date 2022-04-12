package main

import (
	"fmt"
	"seckill/pkg/bootstrap"
	"seckill/sk-app/setup"
)

//sk-app:秒杀业务系统，主要接收用户的秒杀请求，处理用户黑名单，然后把秒杀请求通过redis发给sk-core
func main() {
	//？为什么注释？在admin有
	//mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	//从zookeeper中加载秒杀活动数据到内存中
	setup.InitZk()
	fmt.Println("初始化zk成功，并且加载完秒杀商品信息")

	//从redis中拉取用户id以及用户IP的黑名单设置
	setup.InitRedis()
	fmt.Println("初始化redis成功")
	//fmt.Println("zk商品数据是：", conf.SecKill.SecProductInfoMap[1])
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)

}
