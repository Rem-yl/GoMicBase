package database

import (
	conf "Account/Conf"
	logger "Account/Log"
	"fmt"
	"log"

	share "github.com/GoMicBase/Share"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var nacosConfig conf.NacosConfig
var AccountServConfig conf.AccountServConfig
var err error

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

func init() {
	logger.Init()

	mysqlConf := conf.AccountConf.MysqlConf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.TableName)
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseConn, err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Account{}); err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseInit, err.Error())
	}
}
