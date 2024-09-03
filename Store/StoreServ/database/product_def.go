package database

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

// Category : 类别标签
type Category struct {
	gorm.Model
	// https://gorm.io/zh_CN/docs/belongs_to.html
	Name             string `gorm:"type:varchar(32); not null; index"` // index代表这个字段是一个索引(可能是外键)
	ParentCategoryID uint64
	ParentCategory   *Category
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID; references: ID"`
}

type Brand struct {
	gorm.Model
	Name string `gorm:"type:varchar(32); not null; index"` // index代表这个字段是一个索引(可能是外键)
	Logo string `gorm:"type:varchar(256); not null; default:''"`
}

type Advertise struct {
	gorm.Model
	Index int32  `gorm:"type:int; not null; default: 1"`
	Image string `gorm:"type:varchar(256); not null"`
	Url   string `gorm:"type:varchar(256); not null"`
}

// Product : 商品数据库
type Product struct {
	gorm.Model
	// 如果没有显式地指定 foreignKey 和 references，GORM 会按照约定优于配置的原则自动推断外键
	// Product 结构体中有一个字段 CategoryID，而你的 Category 结构体中有一个字段 ID，GORM 会自动推断 CategoryID 是指向 Category 表中 ID 字段的外键。
	CategoryName string   `gorm:"not null"`
	Category     Category `gorm:"foreignKey:CategoryName; references:Name"`

	BrandName string `gorm:"not null"`
	Brand     Brand  `gorm:"foreignKey:BrandName; references:Name"`

	Selling  bool `gorm:"default:false"`
	ShipFree bool `gorm:"default:false"`  // 是否包邮
	IsPop    bool `gorm:"default:false"`  // 是否是热卖商品
	IsNew    bool `gorm:"deafault:false"` //是否包邮

	Name       string          `gorm:"type:varchar(64); not null"`
	SN         string          `gorm:"type:varchar(64); not null"` //商品条码
	FavNum     int32           `gorm:"type:int; default:0"`        //商品收藏数
	SoldNum    int32           `gorm:"type:int; default:0"`        // 销售额
	Price      float32         `gorm:"not null"`
	RealPrice  float32         `gorm:"not null"`
	ShoreDesc  string          `gorm:"type:varchar(256); not null"` //商品介绍
	DescImages JSONStringSlice `gorm:"type:json; not null"`         // 使用 []string 作为类型可能需要自定义类型映射或处理，GORM 默认不支持直接映射 []string 到 varchar，建议使用 JSON 类型或字符串类型存储，并在代码中进行序列化和反序列化
	CoverImage string          `gorm:"type:varchar(256); not null"` // 封面
}

type JSONStringSlice []string

func (j *JSONStringSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j JSONStringSlice) Value() (driver.Value, error) {
	return json.Marshal(j)
}
