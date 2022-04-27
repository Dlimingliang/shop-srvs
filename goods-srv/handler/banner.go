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

func ModelToBannerRes(banner model.Banner) proto.BannerRes {
	res := proto.BannerRes{
		Id:    banner.ID,
		Image: banner.Image,
		Url:   banner.Url,
		Order: banner.Order,
	}
	return res
}

func (gs GoodsServer) GetBannerList(ctx context.Context, empty *emptypb.Empty) (*proto.BannerListRes, error) {
	var bannerList []model.Banner
	result := global.DB.Find(&bannerList)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.BannerListRes{}
	for _, banner := range bannerList {
		bannerRes := ModelToBannerRes(banner)
		rsp.Data = append(rsp.Data, &bannerRes)
	}
	return rsp, nil
}
func (gs GoodsServer) CreateBanner(ctx context.Context, request *proto.BannerReq) (*proto.BannerRes, error) {
	banner := model.Banner{
		Image: request.Image,
		Url:   request.Url,
		Order: request.Order,
	}
	result := global.DB.Create(&banner)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	bannerRes := ModelToBannerRes(banner)
	return &bannerRes, nil
}
func (gs GoodsServer) UpdateBanner(ctx context.Context, request *proto.BannerReq) (*emptypb.Empty, error) {
	var banner model.Banner
	if result := global.DB.First(&banner, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}

	if request.Url != "" {
		banner.Url = request.Url
	}
	if request.Image != "" {
		banner.Image = request.Image
	}
	if request.Order != 0 {
		banner.Order = request.Order

	}
	if result := global.DB.Save(banner); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
func (gs GoodsServer) DeleteBanner(ctx context.Context, request *proto.BannerReq) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, request.Id); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
