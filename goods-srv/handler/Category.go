package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Dlimingliang/shop-srvs/goods-srv/global"
	"github.com/Dlimingliang/shop-srvs/goods-srv/model"
	"github.com/Dlimingliang/shop-srvs/goods-srv/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func ModelToCategoryRes(category model.Category) proto.CategoryRes {
	res := proto.CategoryRes{
		Id:    category.ID,
		Name:  category.Name,
		Level: category.Level,
		IsTab: category.IsTab,
	}
	return res
}

func (gs GoodsServer) GetAllCategoryList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategoryList not implemented")
}
func (gs GoodsServer) GetSubCategoryList(ctx context.Context, req *proto.SubCategoryReq) (*proto.SubCategoryListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubCategoryList not implemented")
}
func (gs GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryReq) (*proto.CategoryRes, error) {
	var category model.Category
	if result := global.DB.Where(&model.Category{Name: req.Name}).First(&category); result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "分类已存在")
	}

	category.Name = req.Name
	category.Level = req.Level
	category.IsTab = req.IsTab
	category.ParentCategoryId = req.ParentCategoryId
	result := global.DB.Create(&category)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	categoryRes := ModelToCategoryRes(category)
	return &categoryRes, nil
}
func (gs GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryReq) (*emptypb.Empty, error) {
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.ParentCategoryId != 0 {
		category.ParentCategoryId = req.ParentCategoryId
	}
	category.IsTab = req.IsTab
	if result := global.DB.Save(category); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
func (gs GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryReq) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
