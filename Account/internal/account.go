package internal

import (
	conf "Account/Conf"
	"encoding/json"
	"log"
)

var AccountConf conf.AccountConfig
var nacosConf conf.NacosConfig

func init() {
	config := conf.LoadYamlConfig("../", "dev")
	config.Unmarshal(&nacosConf)

	content := conf.LoadConfigFromNacos(nacosConf)
	if err := json.Unmarshal([]byte(content), &AccountConf); err != nil {
		log.Panicln(err.Error())
	}
}