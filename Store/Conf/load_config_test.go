package conf

import (
	"log"
	"testing"

	share "github.com/GoMicBase/Share"
)

func TestLoadDevConfig(t *testing.T) {
	LoadDevConfig()

	content, _ := share.LoadConfigFromNacos(NacosConf)
	log.Println(content)
}

func TestLoadStoreServConfig(t *testing.T) {
	LoadDevConfig()

	content, _ := share.LoadConfigFromNacos(NacosConf)
	log.Println(content)
	LoadStoreServConfig(content)
}
