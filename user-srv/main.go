package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Dlimingliang/shop-srvs/user-srv/global"
	"github.com/Dlimingliang/shop-srvs/user-srv/handler"
	"github.com/Dlimingliang/shop-srvs/user-srv/initialize"
	"github.com/Dlimingliang/shop-srvs/user-srv/proto"
)

const (
	address = "127.0.0.1"
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

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serviceID
	registration.Port = *port
	registration.Address = address
	//本地开发，暂时注释健康检查
	//check := &api.AgentServiceCheck{
	//	GRPC:                           fmt.Sprintf("%s:%d", address, *port),
	//	Timeout:                        "5s",
	//	Interval:                       "5s",
	//	DeregisterCriticalServiceAfter: "10s",
	//}
	//registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Panic("user-srv注册consul失败: " + err.Error())
	}
	go func() {
		err = server.Serve(listen)
		if err != nil {
			zap.S().Panic("grpc启动失败: " + err.Error())
		}
	}()

	//优雅终止
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c
		err = client.Agent().ServiceDeregister(serviceID)
		if err != nil {
			zap.S().Error("user-srv注销失败")
		}
		zap.S().Info("user-srv注销成功")
		close(processed)
	}()
	zap.S().Infof("user-srv启动成功,%s:%d", *ip, *port)
	<-processed
}
