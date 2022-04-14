package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Dlimingliang/shop-srvs/user-srv/global"
	"github.com/Dlimingliang/shop-srvs/user-srv/handler"
	"github.com/Dlimingliang/shop-srvs/user-srv/initialize"
	"github.com/Dlimingliang/shop-srvs/user-srv/proto"
)

const (
	address = "172.0.0.1"
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

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		zap.S().Panic("Listen失败: " + err.Error())
	}

	server := grpc.NewServer()
	proto.RegisterUserServer(server, handler.UserServer{})
	//注册健康检查服务
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//注册grpc服务
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Panic("生成consulclient失败: ", err.Error())
	}
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, *port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *port
	registration.Address = address
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Panic("user-srv注册consul失败: " + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		zap.S().Panic("grpc启动失败: " + err.Error())
	}
}
