package database

import (
	"Account/AccountServ/internal"
	conf "Account/Conf"
	logger "Account/Log"
	share "Account/Share"
	"fmt"
	"log"

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

func init() {
	logger.Init()

	if err = internal.LoadAccountServConfig("./conf", "dev", &nacosConfig, &AccountServConfig); err != nil {
		log.Panicln(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", AccountServConfig.MysqlConf.User, AccountServConfig.MysqlConf.Password, AccountServConfig.MysqlConf.Host, AccountServConfig.MysqlConf.Port, AccountServConfig.MysqlConf.TableName)
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseConn, err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Account{}); err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseInit, err.Error())
	}
}
