package mytest

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "39.99.214.230:6379",
		Password: "zhangpeng",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("Connect redis failed. Error : %v", err)
	}
	log.Printf("init redis success")
}
