package conf

import (
	share "Account/Share"
	"log"

	"github.com/spf13/viper"
)

type AccountConfig struct {
	AccountServConf AccountServConfig `json:"account_serv"`
	AccountWebConf  AccountWebConfig  `json:"account_web"`
	MysqlConf       MysqlConfig       `json:"mysql"`
}

type AccountServConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type AccountWebConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type MysqlConfig struct {
	TableName string `json:"tableName"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

func LoadYamlConfig(configPath, configName string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(configPath)
	config.SetConfigName(configName)
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Panicf("%s : %s\n", share.ErrConfigFileNotFound, err.Error())
	}

	return config
}
