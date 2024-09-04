package conf

import (
	"encoding/json"
	"log"

	share "github.com/GoMicBase/Share"
)

var NacosConf share.NacosConfig
var StoreServConf StoreServConfig

func LoadDevConfig() {
	config, err := share.LoadYamlConfig("../", "dev")
	if err != nil {
		log.Panicln(err.Error())
	}

	if err := config.Unmarshal(&NacosConf); err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("NacosConfig: %v", NacosConf)
}

func LoadStoreServConfig(content string) {
	if err := json.Unmarshal([]byte(content), &StoreServConf); err != nil {
		log.Panicf(err.Error())
	}

	log.Printf("StoreServConfig: %v", StoreServConf)
}

func init() {
	LoadDevConfig()
	content, err := share.LoadConfigFromNacos(NacosConf)
	if err != nil {
		log.Panicf(err.Error())
	}

	LoadStoreServConfig(content)
	storeGrpcPort := share.GetRandomPort(StoreServConf.StoreGrpcConf.Host)
	StoreServConf.StoreGrpcConf.Id = share.GetNewUuid()
	StoreServConf.StoreGrpcConf.Port = int32(storeGrpcPort)
}
