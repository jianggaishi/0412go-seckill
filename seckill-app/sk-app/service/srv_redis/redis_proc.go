package srv_redis

import (
	"encoding/json"
	"fmt"
	"log"
	conf "seckill/pkg/config"
	"seckill/sk-app/config"
	"seckill/sk-app/model"
	"time"
)

//写数据到Redis
func WriteHandle() {
	log.Printf("我到WriteHandle了，我要开始写数据到redis里了")
	for {
		log.Println("wirter data to redis.")
		req := <-config.SkAppContext.SecReqChan
		fmt.Println("accessTime : ", req.AccessTime)
		conn := conf.Redis.RedisConn

		data, err := json.Marshal(req)
		if err != nil {
			log.Printf("json.Marshal req failed. Error : %v, req : %v", err, req)
			continue
		}

		err = conn.LPush(conf.Redis.Proxy2layerQueueName, string(data)).Err()
		if err != nil {
			log.Printf("lpush req failed. Error : %v, req : %v", err, req)
			continue
		}
		log.Printf("lpush req success. req : %v", string(data))
	}
}

//从redis读取数据
func ReadHandle() {
	log.Printf("我到ReadHandle函数了，我要从redis中读取数据了")
	for {
		conn := conf.Redis.RedisConn
		//阻塞弹出
		//阻塞弹出队列数据（core处理完的数据：放在Layer2proxy队列中）
		data, err := conn.BRPop(time.Second, conf.Redis.Layer2proxyQueueName).Result()
		if err != nil {
			continue
		}
		//把取出的data包装成SecResult形式
		var result *model.SecResult
		err = json.Unmarshal([]byte(data[1]), &result)
		if err != nil {
			log.Printf("json.Unmarshal failed. Error : %v", err)
			continue
		}

		userKey := fmt.Sprintf("%d_%d", result.UserId, result.ProductId)
		log.Println("userKey : ", userKey)
		config.SkAppContext.UserConnMapLock.Lock()
		resultChan, ok := config.SkAppContext.UserConnMap[userKey]
		config.SkAppContext.UserConnMapLock.Unlock()
		if !ok {
			log.Printf("user not found : %v", userKey)
			continue
		}
		log.Printf("request result 开始发送到 resultChan")

		resultChan <- result
		log.Printf("request result 发送到 resultChan 成功, userKey是 : %v", userKey)
	}
}
