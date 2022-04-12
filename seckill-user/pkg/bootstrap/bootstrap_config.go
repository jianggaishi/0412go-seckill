package bootstrap

import (
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Logger kitlog.Logger

func init() {
	Logger = kitlog.NewLogfmtLogger(os.Stderr)
	Logger = kitlog.With(Logger, "ts", kitlog.DefaultTimestampUTC)
	Logger = kitlog.With(Logger, "caller", kitlog.DefaultCaller)
	viper.AutomaticEnv()
	initBootstrapConfig()
	//读取yaml文件
	//v := viper.New()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	if err := subParse("http", &HttpConfig); err != nil {
		log.Fatal("Fail to parse Http config", err)
	}
	if err := subParse("discover", &DiscoverConfig); err != nil {
		log.Fatal("Fail to parse Discover config", err)
	}
	if err := subParse("config", &ConfigServerConfig); err != nil {
		log.Fatal("Fail to parse config server", err)
	}

	if err := subParse("rpc", &RpcConfig); err != nil {
		log.Fatal("Fail to parse rpc server", err)
	}
}
func initBootstrapConfig() {
	//设置读取的配置文件
	viper.SetConfigName("bootstrap")
	//添加读取的配置文件路径
	viper.AddConfigPath("./")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	//viper.AddConfigPath("/Users/zp/mygit/go-kit/seckill-user")
	//设置配置文件类型
	viper.SetConfigType("yaml")
}

func subParse(key string, value interface{}) error {
	Logger.Log("配置文件的前缀为：", key)
	//log.Printf("配置文件的前缀为：%v", key)
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}
