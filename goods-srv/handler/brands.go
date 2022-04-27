package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Dlimingliang/shop-srvs/goods-srv/global"
	"github.com/Dlimingliang/shop-srvs/goods-srv/model"
	"github.com/Dlimingliang/shop-srvs/goods-srv/proto"
)

func ModelToBrandRes(brand model.Brands) proto.BrandRes {
	res := proto.BrandRes{
		Id:   brand.ID,
		Name: brand.Name,
		Logo: brand.Logo,
	}
	return res
}

func (gs GoodsServer) GetBrandPage(ctx context.Context, request *proto.BrandPageReq) (*proto.BrandListRes, error) {
	rsp := &proto.BrandListRes{}

	//查询总数
	var total int64
	result := global.DB.Model(&model.Brands{}).Count(&total)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp.Total = int32(total)

	//查询数据
	var brandList []model.Brands
	result = global.DB.Scopes(Paginate(int(request.Pages), int(request.PageSize))).Find(&brandList)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, brand := range brandList {
		brandRes := ModelToBrandRes(brand)
		rsp.Data = append(rsp.Data, &brandRes)
	}
	return rsp, nil
}
func (gs GoodsServer) CreateBrand(ctx context.Context, request *proto.BrandReq) (*proto.BrandRes, error) {
	var brand model.Brands
	if result := global.DB.Where(&model.Brands{Name: request.Name}).First(&brand); result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "品牌已存在")
	}

	brand.Name = request.Name
	brand.Logo = request.Logo
	result := global.DB.Create(&brand)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	brandRes := ModelToBrandRes(brand)
	return &brandRes, nil
}
func (gs GoodsServer) UpdateBrand(ctx context.Context, request *proto.BrandReq) (*emptypb.Empty, error) {
	var brand model.Brands
	if result := global.DB.First(&brand, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}

	if request.Name != "" {
		brand.Name = request.Name
	}
	if request.Logo != "" {
		brand.Logo = request.Logo
	}
	if result := global.DB.Save(brand); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
func (gs GoodsServer) DeleteBrand(ctx context.Context, request *proto.BrandReq) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, request.Id); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
