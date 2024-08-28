package share

import (
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func LoadYamlConfig(path, name string) (config *viper.Viper, err error) {
	config = viper.New()
	config.AddConfigPath(path)
	config.SetConfigFile("yaml")
	config.SetConfigName(name)

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	return config, err
}

func GetNewUuid() string {
	uid := uuid.New().String()

	return uid
}

func GetRandomPort(host string) int {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:0", host))
	if err != nil {
		log.Panicf(err.Error())
	}
	defer listen.Close()

	addr := listen.Addr().(*net.TCPAddr)
	return addr.Port
}
