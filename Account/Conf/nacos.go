package conf

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosConfig struct {
	Addr        string `mapstructure:"addr" json:"addr"`
	Port        int32  `mapstructure:"port" json:"port"`
	NamespaceId string `mapstructure:"namespaceId" json:"namespaceId"`
	DataId      string `mapstructure:"dataId" json:"dataId"`
	Group       string `mapstructure:"group" json:"group"`
}

func LoadConfigFromNacos(config *NacosConfig) string {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.Addr,
			Port:   uint64(config.Port),
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		log.Panicln(err.Error())
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
	})
	if err != nil {
		log.Panicln(err.Error())
	}

	return content
}
