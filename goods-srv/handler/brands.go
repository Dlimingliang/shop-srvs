package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Dlimingliang/shop-srvs/goods-srv/proto"
)

func (gs *GoodsServer) GetBrandPage(ctx context.Context, request *proto.BrandPageReq) (*proto.BrandListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandPage not implemented")
}
func (gs *GoodsServer) CreateBrand(ctx context.Context, request *proto.BrandReq) (*proto.BrandRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (gs *GoodsServer) UpdateBrand(ctx context.Context, request *proto.BrandReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (gs *GoodsServer) DeleteBrand(ctx context.Context, request *proto.BrandReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
