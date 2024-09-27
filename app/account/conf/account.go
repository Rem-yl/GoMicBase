package conf

import (
	"GoMicBase/pkg/cfg"
	"GoMicBase/pkg/utils"
	"encoding/json"
)

type AccountConfig struct {
	MysqlConfig       MysqlConfig       `json:"mysql"`
	AccountServConfig AccountServConfig `json:"account_serv"`
	ConsulConfig      ConsulConfig      `json:"consul"`
}

type MysqlConfig struct {
	TableName string `json:"tableName"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

type AccountServConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

var (
	accountConfig *AccountConfig
	nacosConfig   *cfg.NacosConfig
)

func NewAccountConfig(path, name string) (*AccountConfig, error) {
	config, err := cfg.LoadYamlConfig(path, name)
	if err != nil {
		return nil, err
	}

	if err = config.Unmarshal(&nacosConfig); err != nil {
		return nil, err
	}

	content, err := cfg.GetContentFromNacos(*nacosConfig)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(content), &accountConfig); err != nil {
		return nil, err
	}

	accountConfig.AccountServConfig.Port = int32(utils.GetRandomPort(accountConfig.AccountServConfig.Host))
	accountConfig.AccountServConfig.Id = utils.GetNewUuid()

	return accountConfig, nil
}
