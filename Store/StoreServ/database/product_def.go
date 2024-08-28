package database

import (
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"` // gorm用于软删除
}

// Category : 类别标签
type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(32); not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID; references: ID"`
}

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(32); not null"`
	Logo string `gorm:"type:varchar(256); not null; default:''"`
}

type Advertise struct {
	BaseModel
	Index int32  `gorm:"type:int; not null; default: 1"`
	Image string `gorm:"type:varchar(256); not null"`
	Url   string `gorm:"type:varchar(256); not null"`
	Sort  int32  `gorm:"type:int; not null; default: 1"`
}

// Product : 商品数据库
type Product struct {
	BaseModel
	CategoryID int32    `gorm:"type:int; not null"`
	Category   Category // Product访问Category中的字段需要通过Product.Category来访问

	BrandID int32 `gorm:"type:int; not null"`
	Brand         // Brand中的字段会直接成为Product中的字段, 不需要通过Product.Brand来访问

	Selling  bool `gorm:"deafault:false"`
	ShipFree bool `gorm:"deafault:false"` // 是否包邮
	IsPop    bool `gorm:"default:false"`  // 是否是热卖商品
	IsNew    bool `gorm:"deafault:false"` //是否包邮

	Name       string   `gorm:"type:varchar(64); not null"`
	SN         string   `gorm:"type:varchar(64); not null"` //商品条码
	FavNum     int32    `gorm:"type:int; default:0"`        //商品收藏数
	SoldNum    int32    `gorm:"type:int; default:0"`        // 销售额
	Price      float32  `gorm:"not null"`
	RealPrice  float32  `gorm:"not null"`
	ShoreDesc  string   `gorm:"type:varchar(256); not null"`   //商品介绍
	DescImages []string `gorm:"type:varchar(1024); notl null"` // 商品图片
	CoverImage string   `gorm:"type:varchar(256); not null"`   // 封面
}
