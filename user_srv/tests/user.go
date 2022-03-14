package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"

	"github.com/Dlimingliang/shop_srvs/user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	addr := flag.String("addr", "localhost:9090", "the address to connect to")
	var err error
	conn, err = grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserPage() {
	rsp, err := userClient.GetUserPage(context.Background(), &proto.UserPageRequest{
		Page:     1,
		PageSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.UserName, user.Mobile, user.Password)
		checkResponse, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckRequest{
			Password:          "123456",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkResponse.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		user, err := userClient.CreateUser(context.Background(), &proto.CreateUserRequest{
			UserName: fmt.Sprintf("lml%d", i),
			Mobile:   fmt.Sprintf("1388961430%d", i),
			Password: "123456",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(user.Id)
	}
}

func main() {
	Init()
	//TestGetUserPage()
	TestCreateUser()
	conn.Close()
}
