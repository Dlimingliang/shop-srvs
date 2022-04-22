package model

type Category struct {
	BaseModel
	Name             string `gorm:"column:name;type:varchar(20);not null;comment:名称"`
	Level            int32  `gorm:"column:level;type:int;not null;comment:级别"`
	IsTab            bool   `gorm:"column:is_table;default:false;not null;comment:是否展示在搜索处"`
	ParentCategoryId int32  `gorm:"column:parent_category_id;type:int;not null;comment:父类id"`
	ParentCategory   *Category
}

type Brands struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(20);not null;comment:名称"`
	Logo string `gorm:"column:logo;type:varchar(200);default:'',not null;comment:logo图"`
}

type CategoryBrandRelation struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:index_category_brand,unique"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:index_category_brand,unique"`
	Brands     Brands
}

type Banner struct {
	BaseModel
	Image string `gorm:"column:image;type:varchar(200);not null;comment:图片地址"`
	Url   string `gorm:"column:url;type:varchar(200);not null;comment:跳转url"`
	Order int32  `gorm:"column:order;type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	Name            string   `gorm:"type:varchar(50);not null;comment:名称"`
	GoodsSn         string   `gorm:"type:varchar(50);not null;comment:编号"`
	SaleNum         int32    `gorm:"type:int;default:0;not null;comment:售卖数量"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:点击数"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:收藏数"`
	OriginPrice     float32  `gorm:"not null;comment:原价"`
	SalePrice       float32  `gorm:"not null;comment:卖价"`
	GoodsDes        string   `gorm:"type:varchar(100);not null;comment:描述"`
	GoodsImages     GormList `gorm:"type:varchar(1000);not null;comment:图片列表"`
	GoodsDesImages  GormList `gorm:"type:varchar(1000);not null;comment:描述图片列表"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:前端图"`

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`
}
