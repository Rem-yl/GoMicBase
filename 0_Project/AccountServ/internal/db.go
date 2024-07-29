package internal

import (
	"Account/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error // 这里有个坑, 如果不把error定义为全局变量, 则会把DB赋值为局部变量无法导出

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent, // logger.Info
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// connect mysql
	dsn := "rem:123456@tcp(localhost:3306)/accountServ?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	err = DB.AutoMigrate(&model.Account{})
	if err != nil {
		fmt.Println(err)
	}

}
