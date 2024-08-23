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
