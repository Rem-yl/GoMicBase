package database

import (
	share "Account/Share"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

type Account struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex; not null; type: varchar(12)"`
	Phone    string `gorm:"uniqueIndex; not null; type: varchar(11)"`
	Password string `gorm:"not null; comment: hashed password"`
}

func loadConfig() *viper.Viper {
	config := viper.New()

	config.AddConfigPath("./conf")
	config.SetConfigName("default")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Panicf("%s : %s\n", share.ErrConfigFileNotFound, err.Error())
	}

	return config
}

func init() {
	config := loadConfig()
	user := config.GetString("db.user")
	password := config.GetString("db.password")
	host := config.GetString("db.host")
	port := config.GetString("db.port")
	tableName := config.GetString("db.tableName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, tableName)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseConn, err.Error())
	}

	if err = MysqlDB.AutoMigrate(&Account{}); err != nil {
		log.Panicf("%s:%s\n", share.ErrDatabaseInit, err.Error())
	}
}
