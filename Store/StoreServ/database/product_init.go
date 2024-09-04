package database

import (
	conf "Store/Conf"
	"fmt"
	"log"

	share "github.com/GoMicBase/Share"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 读取数据库以及初始化
var MysqlDB *gorm.DB
var err error

func init() {
	mysqlConf := conf.StoreConf.MysqlConf

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DdName)
	log.Printf("mysql addr : %s", dsn)

	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseConn, err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Advertise{}); err != nil {
		log.Panicf(err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Brand{}); err != nil {
		log.Panicf(err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Category{}); err != nil {
		log.Panicf(err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Product{}); err != nil {
		log.Panicf(err.Error())
	}
}
