package main

import (
	"seckill/pkg/bootstrap"
	"seckill/sk-app/setup"
)

func main() {
	//mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	setup.InitZk()
	setup.InitRedis()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
}
