package internal

import (
	conf "Account/Conf"
	"encoding/json"
)

func LoadAccountServConfig(configPath, configName string, nacosConfig *conf.NacosConfig, AccountservConf *conf.AccountServConfig) error {
	config := conf.LoadYamlConfig(configPath, configName)
	if err := config.Unmarshal(&nacosConfig); err != nil {
		return err
	}
	serverConfStr := conf.LoadConfigFromNacos(nacosConfig)

	if err := json.Unmarshal([]byte(serverConfStr), AccountservConf); err != nil {
		return err
	}

	return nil
}
