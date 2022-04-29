package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"

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

func main() {
	Init()
	//TestCreateBrand()
	//TestUpdateBrand()
	//TestDeleteBrand()
	TestCreateCategory()
	_ = conn.Close()
}
