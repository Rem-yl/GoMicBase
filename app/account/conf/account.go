package conf

import (
	"GoMicBase/pkg/cfg"
	"GoMicBase/pkg/utils"
	"encoding/json"
)

type AccountConfig struct {
	MysqlConfig       cfg.MysqlConfig    `json:"mysql"`
	AccountServConfig cfg.GrpcServConfig `json:"account_serv"`
	ConsulConfig      cfg.ConsulConfig   `json:"consul"`
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
