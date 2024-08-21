package conf

import (
	"fmt"
	"testing"
)

func TestLoadNacosConf(t *testing.T) {
	addr := "127.0.0.1"
	port := 8848
	namespaceId := "3eff1bb2-2e38-40e6-8cb4-04b3ab983da3"
	dataId := "account_serv.json"
	group := "account"

	config := NacosConfig{
		Addr:        addr,
		Port:        int32(port),
		NamespaceId: namespaceId,
		DataId:      dataId,
		Group:       group,
	}
	content := LoadConfigFromNacos(&config)
	fmt.Println(content)
}
