package internal

import (
	"log"
	"testing"
)

func TestConsulRegWeb(t *testing.T) {
	// cd Account/AccountWeb; go run main.go
	accountWebConf := AccountConf.AccountWebConf
	err := ConsulRegWeb(accountWebConf.Host, int(accountWebConf.Port), accountWebConf.Name, accountWebConf.Id, []string{"test"})
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("AccountWeb Register Success")
}

func TestConsulRegGrpc(t *testing.T) {
	// cd Account/AccountServ; go run main.go
	accountServConf := AccountConf.AccountServConf
	err := ConsulRegGrpc(accountServConf.Host, int(accountServConf.Port), accountServConf.Name, "AccountServ1", []string{"test"})
	if err != nil {
		log.Panicln(err.Error())
	}
	err = ConsulRegGrpc(accountServConf.Host, int(accountServConf.Port), accountServConf.Name, "AccountServ2", []string{"test"})
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("AccountServ Register Success")
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
	service, err := GetFilterConsulService(`Service == "account_serv"`)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Consul Service: %v", service)
}
