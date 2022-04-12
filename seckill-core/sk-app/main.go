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
	setup.InitZk()
	fmt.Println("初始化zk成功")
	setup.InitRedis()
	fmt.Println("初始化redis成功")

	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
	fmt.Println("初始化服务成功")

}
