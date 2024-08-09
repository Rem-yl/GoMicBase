package share

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex; not null; type: varchar(12)"`
	Phone    string `gorm:"uniqueIndex; not null; type: varchar(11)"`
	Password string `gorm:"not null; comment: hashed password"`
}
