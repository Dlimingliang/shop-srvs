package handler

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Dlimingliang/shop-srvs/goods-srv/global"
	"github.com/Dlimingliang/shop-srvs/goods-srv/model"
	"github.com/Dlimingliang/shop-srvs/goods-srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func ModelToGoodsRes(goods model.Goods) proto.GoodsRes {
	res := proto.GoodsRes{
		Id:              goods.ID,
		CategoryID:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		SaleNum:         goods.SaleNum,
		ClickNum:        goods.ClickNum,
		FavNum:          goods.FavNum,
		OriginPrice:     goods.OriginPrice,
		SalePrice:       goods.SalePrice,
		GoodsDesc:       goods.GoodsDes,
		ShipFree:        goods.ShipFree,
		GoodsImages:     goods.GoodsImages,
		GoodsDesImages:  goods.GoodsDesImages,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		CreateTime:      goods.CreatedAt.Unix(),
		Brand: &proto.BrandRes{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
		Category: &proto.CategoryBriefRes{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
	}
	return res
}

func (gs GoodsServer) GetGoodsPage(ctx context.Context, req *proto.GoodsPageReq) (*proto.GoodsListRes, error) {
	rsp := &proto.GoodsListRes{}

	//关键词、查询新品、热门商品、价格区间搜索、商品分类筛选
	localDb := global.DB.Model(&model.Goods{})
	if req.KeyWords != "" {
		localDb = localDb.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}
	//如何查询不是热门
	if req.IsHot != false {
		localDb = localDb.Where(&model.Goods{IsHot: true})
	}
	if req.IsNew != false {
		localDb = localDb.Where(&model.Goods{IsNew: true})
	}
	if req.Brand != 0 {
		localDb = localDb.Where(&model.Goods{BrandsID: req.Brand})
	}
	if req.PriceMax != 0 {
		localDb = localDb.Where("sale_price < ?", req.PriceMax)
	}
	if req.PriceMin != 0 {
		localDb = localDb.Where("sale_price > ?", req.PriceMin)
	}
	if req.TopCategory != 0 {
		var subSql string
		//查询是几级topcategory
		var category model.Category
		_ = global.DB.First(&category, req.TopCategory)

		if category.Level == 1 {
			subSql = fmt.Sprintf("category_id in (select id from category where parent_category_id in(select id from category where parent_category_id in(%d)))", req.TopCategory)
		} else if category.Level == 2 {
			subSql = fmt.Sprintf("category_id in (select id from category where parent_category_id in(%d))", req.TopCategory)
		} else if category.Level == 3 {
			subSql = fmt.Sprintf("category_id in (%d)", req.TopCategory)
		}
		localDb = localDb.Where(subSql)
	}

	//查询总数
	var total int64
	result := localDb.Count(&total)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp.Total = int32(total)

	//查询数据 preload待加入
	var goodsList []model.Goods
	result = localDb.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PageSize))).Find(&goodsList)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, goods := range goodsList {
		goodsRes := ModelToGoodsRes(goods)
		rsp.Data = append(rsp.Data, &goodsRes)
	}
	return rsp, nil
}

//func (UnimplementedGoodsServer) GetGoodsListByIds(context.Context, *GoodsIdsReq) (*GoodsListRes, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsListByIds not implemented")
//}
func (gs GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsReq) (*proto.GoodsRes, error) {
	var goods model.Goods
	if result := global.DB.Where(&model.Goods{Name: req.Name}).First(&goods); result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "商品已存在")
	}
	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.OriginPrice = req.OriginPrice
	goods.SalePrice = req.SalePrice
	goods.GoodsDes = req.GoodsDesc
	goods.GoodsImages = req.GoodsImages
	goods.GoodsDesImages = req.GoodsDesImages
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.CategoryID = req.CategoryID
	goods.BrandsID = req.BrandID
	result := global.DB.Create(&goods)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	goodsRes := ModelToGoodsRes(goods)
	return &goodsRes, nil
}

//func (UnimplementedGoodsServer) DeleteGoods(context.Context, *DeleteGoodsReq) (*emptypb.Empty, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
//}
//func (UnimplementedGoodsServer) UpdateGoods(context.Context, *CreateGoodsReq) (*emptypb.Empty, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
//}
//func (UnimplementedGoodsServer) GetGoods(context.Context, *GoodsReq) (*GoodsRes, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
//}
