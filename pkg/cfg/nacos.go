package cfg

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosConfig struct {
	Host        string `mapstructure:"host"`
	Port        int32  `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespaceId"`
	DataId      string `mapstructure:"dataId"`
	GroupId     string `mapstructure:"group"`
}

// https://github.com/nacos-group/nacos-sdk-go
func GetContentFromNacos(conf NacosConfig) (content string, err error) {
	// create client
	clientConfig := constant.ClientConfig{
		NamespaceId:         conf.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      conf.Host,
			Port:        uint64(conf.Port),
			ContextPath: "/nacos",
			Scheme:      "http",
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return "", err
	}

	content, err = configClient.GetConfig(vo.ConfigParam{
		DataId: conf.DataId,
		Group:  conf.GroupId,
	})
	if err != nil {
		return "", nil
	}

	return content, nil
}
