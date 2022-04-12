module seckill

go 1.12

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/apache/thrift/lib/go/thrift v0.0.1-do-not-use // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.15+incompatible
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.12.0
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gohouse/gorose/v2 v2.1.12
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.6.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/consul/api v1.10.1
	github.com/juju/ratelimit v1.0.1
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/unknwon/com v1.0.1
	go.etcd.io/bbolt v1.3.3 // indirect
	go.etcd.io/etcd v3.3.27+incompatible
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	google.golang.org/grpc v1.40.0
	google.golang.org/grpc/examples v0.0.0-20220120193159-9cb411380883 // indirect
)

//replace golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
