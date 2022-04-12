package setup

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	conf "seckill/pkg/config"
	"time"
)

//初始化zookeeper
func InitZk() {
	var hosts = []string{"192.168.145.151:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println("InitZk有错 err：", err)
		return
	}
	conf.Zk.ZkConn = conn
	//设置关键字
	conf.Zk.SecProductKey = "/product"
}
