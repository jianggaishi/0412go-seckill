package main

import (
	"fmt"
	"seckill/pkg/bootstrap"
	conf "seckill/pkg/config"
	"seckill/pkg/mysql"
	"seckill/sk-admin/setup"
)

//秒杀管理系统，创建删除秒杀活动，商品配置
func main() {
	//fmt.Println("开始延时2：")
	//time.Sleep(50 * time.Second)
	//初始化mysql
	//fmt.Println("mysqlConfig是：", conf.MysqlConfig)
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db

	fmt.Println("\nmysql初始化成功")

	//setup.InitEtcd()

	setup.InitZk()
	fmt.Println("\nzk初始化成功")

	//传入两个参数，服务的ip和端口
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
	//fmt.Println("\n服务初始化成功")

}
