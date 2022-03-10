package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		//Gender:   false,
		//Birthday: 0,
		Role: int32(user.Role),
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
