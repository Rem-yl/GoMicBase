package database

import (
	"GoMicBase/pkg/cfg"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(config *cfg.MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.TableName)

	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return mysqlDB, nil
}
