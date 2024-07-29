package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	PhoneNumber string `gorm:"uniqueIndex:idx_phoneNumber; not null; type : varchar(11)"`
	Password    string `gorm:"not null; type:varchar(64)"`
	Salt        string `gorm:"type:varchar(16)"`
	NickName    string `gorm:"uniqueIndex:idx_userName; not null; type : varchar(16)"`
	Gender      uint32 `gorm:"comment : 0-men,1-women,2-unknown"`
	Role        uint32 `gorm:"not null; comment: 0-manager, 1-normal"`
}
