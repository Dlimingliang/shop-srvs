package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/Dlimingliang/shop-srvs/user-srv/handler"
	"github.com/Dlimingliang/shop-srvs/user-srv/initialize"
	"github.com/Dlimingliang/shop-srvs/user-srv/proto"
)

var (
	ip   = flag.String("ip", "0.0.0.0", "IP")
	port = flag.Int("port", 8090, "端口号")
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, handler.UserServer{})

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		panic(any("Listen失败: " + err.Error()))
	}
	err = server.Serve(listen)
	if err != nil {
		panic(any("grpc启动失败: " + err.Error()))
	}
}
