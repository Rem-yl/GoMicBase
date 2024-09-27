package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex; not null; type: varchar(12)"`
	Phone          string `gorm:"uniqueIndex; not null; type: varchar(11)"`
	Password       string `gorm:"not null; comment: origin password"`
	HashedPassword string `gorm:"not null; comment: hashed password"`
	Salt           string `gorm:"not null; comment: use to hash password"`
}

type CustomAccount struct {
	Id             uint32 `json:"id"`
	Name           string `json:"string"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	Salt           string `json:"salt"`
	HashedPassword string `json:"hashed_password"`
}
