package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/Dlimingliang/shop_srvs/user_srv/handler"
	"github.com/Dlimingliang/shop_srvs/user_srv/proto"
)

var (
	ip   = flag.String("ip", "0.0.0.0", "IP")
	port = flag.Int("port", 9090, "端口号")
)

func main() {
	server := grpc.NewServer()
	proto.RegisterUserServer(server, handler.UserServer{})

	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		panic("Listen失败: " + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		panic("grpc启动失败: " + err.Error())
	}
}
