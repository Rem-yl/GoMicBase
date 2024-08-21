package internal

import (
	conf "Account/Conf"
	"encoding/json"
)

func LoadAccountServConfig(configPath, configName string, nacosConfig *conf.NacosConfig, AccountServConf *conf.AccountServConfig) error {
	config := conf.LoadYamlConfig(configPath, configName)
	if err := config.Unmarshal(&nacosConfig); err != nil {
		return err
	}
	serverConfStr := conf.LoadConfigFromNacos(nacosConfig)

	if err := json.Unmarshal([]byte(serverConfStr), AccountServConf); err != nil {
		return err
	}

	return nil
}

func LoadAccountWebConfig(configPath, configName string, nacosConfig *conf.NacosConfig, AccountWebConf *conf.AccountWebConfig) error {
	config := conf.LoadYamlConfig(configPath, configName)
	if err := config.Unmarshal(&nacosConfig); err != nil {
		return err
	}
	serverConfStr := conf.LoadConfigFromNacos(nacosConfig)

	if err := json.Unmarshal([]byte(serverConfStr), AccountWebConf); err != nil {
		return err
	}

	return nil
}
