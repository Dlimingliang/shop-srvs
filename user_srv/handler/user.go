package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	"github.com/Dlimingliang/shop_srvs/user_srv/global"
	"github.com/Dlimingliang/shop_srvs/user_srv/model"
	"github.com/Dlimingliang/shop_srvs/user_srv/proto"
)

type UserServer struct{}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ModelToUserResponse(user model.User) proto.UserResponse {
	//proto生成的对象是有默认值的,我们不可以随便将数据库查询的东西赋值给proto,如果为nil,那么grpc可能报错
	//所以我们查询出来的东西，如果是可以为空的，那么我们要进行判断
	userRs := proto.UserResponse{
		Id:       user.ID,
		UserName: user.UserName,
		Mobile:   user.Mobile,
		Password: user.Password,
		//Gender:   "",
		//Birthday: 0,
		Role: int32(user.Role),
	}

	if user.Gender != "" {
		userRs.Gender = user.Gender
	}
	if user.Birthday != nil {
		userRs.Birthday = uint64(user.Birthday.Unix())
	}

	return userRs
}

func (us UserServer) GetUserPage(ctx context.Context, request *proto.UserPageRequest) (*proto.UserPageResponse, error) {
	//查询全部用户
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	//分页查询用户数据
	result = global.DB.Scopes(Paginate(int(request.Page), int(request.PageSize))).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	//构建返回数据
	rsp := &proto.UserPageResponse{}
	rsp.Total = int32(result.RowsAffected)
	for _, user := range users {
		userRs := ModelToUserResponse(user)
		rsp.Data = append(rsp.Data, &userRs)
	}
	return rsp, nil
}

func (us UserServer) GetUserByID(ctx context.Context, request *proto.IDRequest) (*proto.UserResponse, error) {
	//通过id查询用户
	var user model.User
	result := global.DB.First(&user, request.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	userRs := ModelToUserResponse(user)
	return &userRs, nil
}

func (us UserServer) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserResponse, error) {
	//通过电话查询用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile: request.Mobile}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	userRs := ModelToUserResponse(user)
	return &userRs, nil
}

func (us UserServer) CreateUser(ctx context.Context, request *proto.CreateUserRequest) (*proto.UserResponse, error) {
	//新建用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile: request.Mobile}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.UserName = request.UserName
	user.Mobile = request.Mobile
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(request.Password, options)
	user.Password = fmt.Sprintf("%s$%s", salt, encodedPwd)
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userRs := ModelToUserResponse(user)
	return &userRs, nil
}

func (us UserServer) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*emptypb.Empty, error) {
	//更新用户
	var user model.User
	result := global.DB.First(&user, request.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	user.UserName = request.UserName
	user.Gender = request.Gender
	birthDay := time.Unix(int64(request.Birthday), 0)
	user.Birthday = &birthDay
	//会更新所有字段即使是非零的值,因为之前查询了所有字段出来，并且进行更新，所以不会有问题
	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (us UserServer) CheckPassword(ctx context.Context,
	request *proto.PasswordCheckRequest) (*proto.PasswordCheckResponse, error) {
	//校验密码
	passwords := strings.Split(request.EncryptedPassword, "$")
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	check := password.Verify(request.Password, passwords[0], passwords[1], options)
	return &proto.PasswordCheckResponse{Success: check}, nil
}
