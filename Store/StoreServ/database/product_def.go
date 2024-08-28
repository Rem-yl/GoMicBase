package database

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

// Category : 类别标签
type Category struct {
	// 构建的表结构中有一个parent_category_id列是外键列, 它必须是ID中的值
	gorm.Model
	Name             string `gorm:"type:varchar(32); not null"`
	ParentCategoryID uint64
	ParentCategory   *Category
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID; references: ID"`
}

type Brand struct {
	gorm.Model
	Name string `gorm:"type:varchar(32); not null"`
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
	CategoryID uint64   `gorm:"not null"` // 根据字段名来推断外键
	Category   Category // Product访问Category中的字段需要通过Product.Category来访问

	BrandID uint64 `gorm:"not null"`
	Brand   Brand  // Brand中的字段会直接成为Product中的字段, 不需要通过Product.Brand来访问

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
