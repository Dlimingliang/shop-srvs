package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Dlimingliang/shop-srvs/goods-srv/proto"
)

var conn *grpc.ClientConn
var goodsClient proto.GoodsClient

func Init() {
	addr := flag.String("addr", "localhost:8010", "the address to connect to")
	var err error
	conn, err = grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		panic(any(err.Error()))
	}
	goodsClient = proto.NewGoodsClient(conn)
}

func TestCreateBrand() {
	for i := 0; i < 20; i++ {
		brand, err := goodsClient.CreateBrand(context.Background(), &proto.BrandReq{
			Name: fmt.Sprintf("品牌%d", i),
			Logo: "https://www.tukuppt.com/muban/xpdejwgw.html",
		})
		if err != nil {
			panic(any(err))
		}
		fmt.Println(brand.Id)
	}
}

func TestUpdateBrand() {
	rsp, err := goodsClient.GetBrandPage(context.Background(), &proto.BrandPageReq{
		Pages:    1,
		PageSize: 2,
	})
	if err != nil {
		panic(any(err))
	}

	brand := rsp.Data[0]
	updateBrand := proto.BrandReq{
		Id:   brand.Id,
		Name: "update-品牌",
	}
	_, err = goodsClient.UpdateBrand(context.Background(), &updateBrand)
	if err != nil {
		panic(any(err))
	}
}

func TestDeleteBrand() {
	rsp, err := goodsClient.GetBrandPage(context.Background(), &proto.BrandPageReq{
		Pages:    1,
		PageSize: 2,
	})
	if err != nil {
		panic(any(err))
	}
	_, err = goodsClient.DeleteBrand(context.Background(), &proto.BrandReq{Id: rsp.Data[0].Id})
	if err != nil {
		panic(any(err))
	}
}

func TestCreateCategory() {
	category, err := goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:  "家用电器",
		Level: 1,
		IsTab: true,
	})
	if err != nil {
		panic(any(err))
	}

	air, err := goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:             "空调",
		Level:            2,
		IsTab:            true,
		ParentCategoryId: category.Id,
	})
	if err != nil {
		panic(any(err))
	}
	_, err = goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:             "新风空调",
		Level:            3,
		IsTab:            true,
		ParentCategoryId: air.Id,
	})
	if err != nil {
		panic(any(err))
	}
	_, err = goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:             "空调挂机",
		Level:            3,
		IsTab:            true,
		ParentCategoryId: air.Id,
	})
	if err != nil {
		panic(any(err))
	}

	xiyiji, err := goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:             "洗衣机",
		Level:            2,
		IsTab:            true,
		ParentCategoryId: category.Id,
	})
	if err != nil {
		panic(any(err))
	}
	_, err = goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:             "滚筒洗衣机",
		Level:            3,
		IsTab:            true,
		ParentCategoryId: xiyiji.Id,
	})
	if err != nil {
		panic(any(err))
	}
	_, err = goodsClient.CreateCategory(context.Background(), &proto.CategoryReq{
		Name:             "洗烘一体",
		Level:            3,
		IsTab:            true,
		ParentCategoryId: xiyiji.Id,
	})
	if err != nil {
		panic(any(err))
	}
	fmt.Println("category初始化完成")
}

func TestGetAllCategoryList() {
	list, err := goodsClient.GetAllCategoryList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(any(err))
	}

	for _, category := range list.Data {
		fmt.Println(category.Name)
	}
}

func TestCreateGoods() {
	_, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsReq{
		Name:            "商品6",
		GoodsSn:         "123",
		OriginPrice:     100,
		SalePrice:       25.99,
		GoodsDesc:       "我是描绘",
		ShipFree:        false,
		GoodsImages:     []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsDesImages:  []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsFrontImage: "https://www.tukuppt.com/muban/xpdejwgw.html",
		IsNew:           false,
		IsHot:           false,
		OnSale:          false,
		CategoryID:      6,
		BrandID:         21,
	})
	if err != nil {
		panic(any(err))
	}

	_, err = goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsReq{
		Name:            "商品7",
		GoodsSn:         "123",
		OriginPrice:     100,
		SalePrice:       25.99,
		GoodsDesc:       "我是描绘",
		ShipFree:        false,
		GoodsImages:     []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsDesImages:  []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsFrontImage: "https://www.tukuppt.com/muban/xpdejwgw.html",
		IsNew:           false,
		IsHot:           false,
		OnSale:          false,
		CategoryID:      7,
		BrandID:         21,
	})
	if err != nil {
		panic(any(err))
	}

	_, err = goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsReq{
		Name:            "商品9",
		GoodsSn:         "123",
		OriginPrice:     100,
		SalePrice:       25.99,
		GoodsDesc:       "我是描绘",
		ShipFree:        false,
		GoodsImages:     []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsDesImages:  []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsFrontImage: "https://www.tukuppt.com/muban/xpdejwgw.html",
		IsNew:           false,
		IsHot:           false,
		OnSale:          false,
		CategoryID:      9,
		BrandID:         21,
	})
	if err != nil {
		panic(any(err))
	}

	_, err = goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsReq{
		Name:            "商品10",
		GoodsSn:         "123",
		OriginPrice:     100,
		SalePrice:       25.99,
		GoodsDesc:       "我是描绘",
		ShipFree:        false,
		GoodsImages:     []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsDesImages:  []string{"https://www.tukuppt.com/muban/xpdejwgw.html"},
		GoodsFrontImage: "https://www.tukuppt.com/muban/xpdejwgw.html",
		IsNew:           false,
		IsHot:           false,
		OnSale:          false,
		CategoryID:      10,
		BrandID:         21,
	})
	if err != nil {
		panic(any(err))
	}
}

func TestGetGoodsPage() {
	res, err := goodsClient.GetGoodsPage(context.Background(), &proto.GoodsPageReq{
		//KeyWords: "商品",
		//IsHot:    false,
		//Brand:    21,
		//PriceMax: 100,
		//PriceMin: 50,
		TopCategory: 4,
		Pages:       1,
		PageSize:    10,
	})
	if err != nil {
		panic(any(err))
	}
	for _, goods := range res.Data {
		fmt.Println(goods.Name)
		if goods.Brand != nil {
			fmt.Println(goods.Brand.Name)
		}
		if goods.Category != nil {
			fmt.Println(goods.Category.Name)
		}
	}
}

func main() {
	Init()
	//TestCreateBrand()
	//TestUpdateBrand()
	//TestDeleteBrand()
	//TestCreateCategory()
	//TestGetAllCategoryList()
	//TestCreateGoods()
	TestGetGoodsPage()
	_ = conn.Close()
}
