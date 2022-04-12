package main

import (
	"seckill/sk-core/setup"
)

func main() {

	setup.InitZk()
	setup.InitRedis()
	setup.RunService()

}
