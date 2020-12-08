package api

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestNewClient(t *testing.T) {
	flag.Set("conf", "./../test")
	//flag.Set("f", "./../test/docker-compose.yaml")
	flag.Parse()
	//disableLich := os.Getenv("DISABLE_LICH") != ""
	//if !disableLich {
	//	if err := lich.Setup(); err != nil {
	//		panic(err)
	//	}
	//}

	conn, err := grpc.Dial("0.0.0.0:9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	c := NewDemoClient(conn)

	r, err := c.SayHelloURL(context.Background(), &HelloReq{
		Name: "11111"})
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
