package cfg

import "github.com/spf13/viper"

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
