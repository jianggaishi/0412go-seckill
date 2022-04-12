package setup

import (
	"github.com/go-redis/redis"
	"log"
	conf "seckill/pkg/config"
)

//初始化redis
func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       conf.Redis.Db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("Connect redis failed. Error : %v", err)
	}
	conf.Redis.RedisConn = client
}
