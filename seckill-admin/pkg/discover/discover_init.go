package discover

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
	"seckill/pkg/bootstrap"
	"seckill/pkg/common"
	"seckill/pkg/loadbalance"
)

var ConsulService DiscoveryClient
var LoadBalance loadbalance.LoadBalance
var Logger *log.Logger
var NoInstanceExistedErr = errors.New("no available client")

func init() {
	// 1.实例化一个 Consul 客户端，此处实例化了原生态实现版本
	ConsulService = New(bootstrap.DiscoverConfig.Host, bootstrap.DiscoverConfig.Port)
	LoadBalance = new(loadbalance.RandomLoadBalance)
	Logger = log.New(os.Stderr, "", log.LstdFlags)

}

//
func CheckHealth(writer http.ResponseWriter, reader *http.Request) {
	Logger.Println("Health check!")
	_, err := fmt.Fprintln(writer, "Server is OK!")
	if err != nil {
		Logger.Println(err)
	}
}

func DiscoveryService(serviceName string) (*common.ServiceInstance, error) {
	instances := ConsulService.DiscoverServices(serviceName, Logger)

	if len(instances) < 1 {
		Logger.Printf("no available client for %s.", serviceName)
		return nil, NoInstanceExistedErr
	}
	return LoadBalance.SelectService(instances)

}

func Register() {
	//// 实例失败，停止服务
	if ConsulService == nil {
		panic(0)
	}
	fmt.Println("-begin test--------------------")
	svclist := ConsulService.DiscoverServices("config-service", Logger)
	fmt.Println("我发现的服务有： ", svclist)
	//fmt.Println("consul客户端好用，且也能发现服务，测试通过")
	fmt.Println("--test down-------------------")

	//判空 instanceId,通过 go.uuid 获取一个服务实例ID
	instanceId := bootstrap.DiscoverConfig.InstanceId

	if instanceId == "" {
		instanceId = bootstrap.DiscoverConfig.ServiceName + uuid.NewV4().String()
	}
	log.Println("bootstrap.HttpConfig是", bootstrap.HttpConfig)
	log.Println("bootstrap.DiscoverConfig ：", bootstrap.DiscoverConfig)
	//设置weight

	//我要注册consul服务
	ConsulService.Register(instanceId, "sk-admin", "/health",
		bootstrap.HttpConfig.Port, bootstrap.DiscoverConfig.ServiceName,
		bootstrap.DiscoverConfig.Weight,
		map[string]string{
			"rpcPort": bootstrap.RpcConfig.Port,
		}, nil, Logger)
	//fmt.Println("我注册完服务了")
	//这里应该是服务的应用ip：port，
	if !ConsulService.Register(instanceId, "sk-admin", "/health",
		bootstrap.HttpConfig.Port, bootstrap.DiscoverConfig.ServiceName,
		bootstrap.DiscoverConfig.Weight,
		map[string]string{
			"rpcPort": bootstrap.RpcConfig.Port,
		}, nil, Logger) {
		Logger.Printf("register service %s failed.", bootstrap.DiscoverConfig.ServiceName)
		// 注册失败，服务启动失败
		panic(0)
	}

	Logger.Printf(bootstrap.DiscoverConfig.ServiceName+"-service for service %s success.", bootstrap.DiscoverConfig.ServiceName)

}

func Deregister() {
	//// 实例失败，停止服务
	if ConsulService == nil {
		panic(0)
	}
	//判空 instanceId,通过 go.uuid 获取一个服务实例ID
	instanceId := bootstrap.DiscoverConfig.InstanceId

	if instanceId == "" {
		instanceId = bootstrap.DiscoverConfig.ServiceName + "-" + uuid.NewV4().String()
	}
	if !ConsulService.DeRegister(instanceId, Logger) {
		Logger.Printf("deregister for service %s failed.", bootstrap.DiscoverConfig.ServiceName)
		panic(0)
	}
}
