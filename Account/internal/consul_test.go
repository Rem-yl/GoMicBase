package internal

import (
	"log"
	"testing"
)

func TestReg(t *testing.T) {
	// cd Account/AccountWeb; go run main.go
	accountWebConf := AccountConf.AccountWebConf
	err := ConsulReg(accountWebConf.Host, int(accountWebConf.Port), "accountWeb", "account_web")
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("注册成功")
}

func TestGetConsulServiceList(t *testing.T) {
	// 先运行 TestReg, 注册一个web服务
	serviceList, err := GetConsulServiceList()
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Consul Service List: %v", serviceList)
}

func TestGetFilterConsulService(t *testing.T) {
	service, err := GetFilterConsulService(`Service == "accountWeb"`)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Consul Service: %v", service)
}
