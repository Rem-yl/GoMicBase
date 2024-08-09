package Database

import (
	share "Account/Share"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func init() {
	config := viper.New()
	config.SetConfigName("db")
	config.AddConfigPath("./Conf")
	config.SetConfigType("ini")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf(share.ErrConfigNotFound + err.Error())
		} else {
			fmt.Println("读取配置文件失败" + err.Error())
			log.Printf(share.ErrConfigReadFailed + err.Error())
		}
	}

	// todo: 读取配置文件失败了
	// host := config.GetString("mysql.host")
	// port := config.GetString("mysql.port")
	// user := config.GetString("mysql.user")
	// password := config.GetString("mysql.password")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/Account?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)
	dsn := "root:123456@tcp(127.0.0.1:3307)/Account?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic("数据库连接失败")
		log.Panicln(share.ErrDatabaseConn + err.Error())
	}

	if err = MysqlDB.AutoMigrate(&share.Account{}); err != nil {
		// panic("数据库初始化失败")
		log.Panicln(share.ErrDatabaseInit + err.Error())
	}

}
