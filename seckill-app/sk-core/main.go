package main

import (
	"log"
	"seckill/sk-core/setup"
)

//sk-core秒杀核心系统，处理sk-app通过redis发来的秒杀请求，判断是否秒杀成功
func main() {
	//fmt.Println("hello,world")
	log.Println("开始zookeeper初始化")
	setup.InitZk()
	log.Println("初始化zk成功")
	log.Println("初始化redis")
	setup.InitRedis()
	//下面这个有问题

	setup.RunService()

}
