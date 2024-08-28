package conf

import (
	"encoding/json"
	"log"

	share "github.com/GoMicBase/Share"
)

var AccountConf AccountConfig
var nacosConf share.NacosConfig

func init() {
	config, err := share.LoadYamlConfig("../", "dev")
	if err != nil {
		log.Panicln(err.Error())
	}

	if err := config.Unmarshal(&nacosConf); err != nil {
		log.Panicln(err.Error())
	}
	content, _ := share.LoadConfigFromNacos(nacosConf)
	if err := json.Unmarshal([]byte(content), &AccountConf); err != nil {
		log.Panicln(err.Error())
	}

	accountServPort := share.GetRandomPort(AccountConf.AccountServConf.Host)
	accountServUuid := share.GetNewUuid()
	AccountConf.AccountServConf.Port = int32(accountServPort)
	AccountConf.AccountServConf.Id = accountServUuid
}
