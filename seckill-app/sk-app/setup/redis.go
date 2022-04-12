package setup

import (
	"github.com/go-redis/redis"
	"github.com/unknwon/com"
	"log"
	conf "seckill/pkg/config"
	"seckill/sk-app/service/srv_redis"
	"time"
)

//初始化Redis
func InitRedis() {
	log.Printf("连接 redis 数据库： %s", conf.Redis.Db)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       conf.Redis.Db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("Connect redis failed. Error : %v", err)
	}
	log.Printf("connect redis %s success", conf.Redis.Db)
	conf.Redis.RedisConn = client
	//加载黑名单ip列表
	loadBlackList(client)
	initRedisProcess()
}

//加载黑名单列表
func loadBlackList(conn *redis.Client) {
	conf.SecKill.IPBlackMap = make(map[string]bool, 10000)
	conf.SecKill.IDBlackMap = make(map[int]bool, 10000)

	//用户Id
	idList, err := conn.HGetAll(conf.Redis.IdBlackListHash).Result()

	if err != nil {
		log.Printf("hget all failed. Error : %v", err)
		return
	}

	for _, v := range idList {
		id, err := com.StrTo(v).Int()
		if err != nil {
			log.Printf("invalid user id [%v]", id)
			continue
		}
		conf.SecKill.IDBlackMap[id] = true
	}

	//用户Ip
	ipList, err := conn.HGetAll(conf.Redis.IpBlackListHash).Result()
	if err != nil {
		log.Printf("hget all failed. Error : %v", err)
		return
	}
	log.Printf("用户ip为： ", ipList)
	for _, v := range ipList {
		conf.SecKill.IPBlackMap[v] = true
	}

	//go syncIpBlackList(conn)
	//go syncIdBlackList(conn)
	return
}

//同步用户ID黑名单
func syncIdBlackList(conn *redis.Client) {
	for {
		idArr, err := conn.BRPop(time.Minute, conf.Redis.IdBlackListQueue).Result()
		if err != nil {
			log.Printf("brpop id failed, err : %v", err)
			continue
		}
		id, _ := com.StrTo(idArr[1]).Int()
		conf.SecKill.RWBlackLock.Lock()
		{
			conf.SecKill.IDBlackMap[id] = true
		}
		conf.SecKill.RWBlackLock.Unlock()
	}
}

//同步用户IP黑名单
func syncIpBlackList(conn *redis.Client) {
	var ipList []string
	lastTime := time.Now().Unix()

	for {
		ipArr, err := conn.BRPop(time.Minute, conf.Redis.IpBlackListQueue).Result()
		if err != nil {
			log.Printf("brpop ip failed, err : %v", err)
			continue
		}

		ip := ipArr[1]
		curTime := time.Now().Unix()
		ipList = append(ipList, ip)

		if len(ipList) > 100 || curTime-lastTime > 5 {
			conf.SecKill.RWBlackLock.Lock()
			{
				for _, v := range ipList {
					conf.SecKill.IPBlackMap[v] = true
				}
			}
			conf.SecKill.RWBlackLock.Unlock()

			lastTime = curTime
			log.Printf("sync ip list from redis success, ip[%v]", ipList)
		}
	}
}

//初始化redis进程
func initRedisProcess() {
	//根据系统业务配置数据中心中配置的属性值，启动对应数量的协程执行 WriteHandle和ReadHandle 。，
	//把SecReqChan中的请求发送到Redis对应（**）队列，使用BPpop监听Redis的对应队列，获取秒杀系统返回值，将其发送到resultChan
	log.Printf("initRedisProcess %d %d", conf.SecKill.AppWriteToHandleGoroutineNum, conf.SecKill.AppReadFromHandleGoroutineNum)
	for i := 0; i < conf.SecKill.AppWriteToHandleGoroutineNum; i++ {
		//写数据到redis进程
		go srv_redis.WriteHandle()
	}

	for i := 0; i < conf.SecKill.AppReadFromHandleGoroutineNum; i++ {
		go srv_redis.ReadHandle()
	}

}
