package conf

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/log"
)

var (
	// Conf config
	Conf = &Config{}
)

// Config .
type Config struct {
	Log        *log.Config
	Server     *Server
	Redis      *redis.Config
	MySQL      *sql.Config
	App        *App
	TestClient *warden.ClientConfig
}

type Server struct {
	*warden.ServerConfig
}

// App .
type App struct {
	AppName     string
	CorsDomains string
	AppCode     int
}

func init() {
	flag.Parse()
}

// Init init conf
func Init() error {
	if err := paladin.Init(); err != nil {
		panic(err)
	}

	return local()
}

func local() (err error) {
	initLogConfig()
	initMysqlConfig()
	initRedisConfig()
	initServerConfig()
	initTestClientConfig()
	//if err := paladin.Get("etcd.toml").UnmarshalTOML(&Conf); err != nil {
	//	// 不存在时，将会为nil使用默认配置
	//	if err != paladin.ErrNotExist {
	//		panic(err)
	//	}
	//}
	//
	//
	//
	//if err := paladin.Get("http.toml").UnmarshalTOML(&Conf); err != nil {
	//	// 不存在时，将会为nil使用默认配置
	//	if err != paladin.ErrNotExist {
	//		panic(err)
	//	}
	//}
	//
	//if err := paladin.Get("jwt.toml").UnmarshalTOML(&Conf); err != nil {
	//	// 不存在时，将会为nil使用默认配置
	//	if err != paladin.ErrNotExist {
	//		panic(err)
	//	}
	//}
	//
	//if err := paladin.Get("app.toml").UnmarshalTOML(&Conf); err != nil {
	//	// 不存在时，将会为nil使用默认配置
	//	if err != paladin.ErrNotExist {
	//		panic(err)
	//	}
	//}
	//
	//fmt.Printf("%+v", Conf.App)

	return
}

func initServerConfig() {
	var (
		ct paladin.TOML
	)
	if err := paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		panic(err)
	}
	if err := ct.Get("Server").UnmarshalTOML(&Conf.Server); err != nil {
		panic(err)
	}

	fmt.Printf("initServerConfig %+v \n", Conf.Server)
}

func initTestClientConfig() {
	var (
		ct paladin.TOML
	)
	if err := paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		panic(err)
	}
	if err := ct.Get("Client").UnmarshalTOML(&Conf.TestClient); err != nil {
		panic(err)
	}
}

func initMysqlConfig() {

}

func initRedisConfig() {
	if err := paladin.Get("redis.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
}

func initLogConfig() {
	if err := paladin.Get("log.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
}
