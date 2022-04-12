package main

import (
	"fmt"
	"seckill/sk-core/setup"
)

//sk-core秒杀核心系统，处理sk-app通过redis发来的秒杀请求，判断是否秒杀成功
func main() {
	//fmt.Println("hello,world")
	setup.InitZk()
	fmt.Println("初始化zk成功")
	setup.InitRedis()
	fmt.Println("初始化redis成功")
	setup.RunService()

}
