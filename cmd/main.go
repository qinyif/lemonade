package main

import (
	"github.com/byteconv/lemonade/internal/conf"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/byteconv/lemonade/internal/di"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	//flag.Parse()
	err := conf.Init()

	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("github.com/byteconv/lemonade start")
	//paladin.Init()

	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("github.com/byteconv/lemonade exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
