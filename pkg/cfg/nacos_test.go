package cfg

import (
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/stretchr/testify/assert"
)

func TestPublishNacos(t *testing.T) {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("public"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	success, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "test.json",
		Group:   "test",
		Content: "hello, world",
	})

	if err != nil {
		t.Error(err.Error())
	}

	assert.True(t, success)
}

func TestGetNacosConfig(t *testing.T) {
	// create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	// create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("public"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	content, err := client.GetConfig(vo.ConfigParam{
		Group:  "test",
		DataId: "test.json",
	})

	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "hello, world", content)
}

func TestGetContentFromNacos(t *testing.T) {
	var nacosConfig NacosConfig
	config, err := LoadYamlConfig(".", "dev")
	if err != nil {
		t.Error(err.Error())
	}

	if err := config.Unmarshal(&nacosConfig); err != nil {
		t.Error(err.Error())
	}

	content, err := GetContentFromNacos(nacosConfig)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "hello, world", content)
}
