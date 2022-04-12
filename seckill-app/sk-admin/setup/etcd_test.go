package setup

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
	"time"
)

func TestInitEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"39.99.214.230:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err1 := cli.Put(ctx, "sec_kill_product", "sample")
	if err1 != nil {
		log.Printf("Get falied. Error : %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	nresp, err := cli.Get(ctx, "sec_kill_product")

	if err != nil {
		log.Printf("Get falied. Error : %v", err)
	}

	for _, ev := range nresp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
