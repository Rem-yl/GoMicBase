package conf

import (
	share "Account/Share"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./conf")
	config.SetConfigName("default")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Panicf("%s : %s\n", share.ErrConfigFileNotFound, err.Error())
	}

	return config
}
