package conf

import (
	"encoding/json"
	"log"

	share "github.com/GoMicBase/Share"
)

var NacosConf share.NacosConfig
var StoreConf StoreConfig

func LoadDevConfig() {
	config, err := share.LoadYamlConfig("../", "dev")
	if err != nil {
		log.Panicln(err.Error())
	}

	if err := config.Unmarshal(&NacosConf); err != nil {
		log.Panicln(err.Error())
	}
}

func LoadStoreServConfig(content string) {
	if err := json.Unmarshal([]byte(content), &StoreConf); err != nil {
		log.Panicf(err.Error())
	}
}

func init() {
	LoadDevConfig()
	content, err := share.LoadConfigFromNacos(NacosConf)
	if err != nil {
		log.Panicf(err.Error())
	}

	LoadStoreServConfig(content)
	storeGrpcPort := share.GetRandomPort(StoreConf.StoreServConf.Host)
	StoreConf.StoreServConf.Id = share.GetNewUuid()
	StoreConf.StoreServConf.Port = int32(storeGrpcPort)
}
