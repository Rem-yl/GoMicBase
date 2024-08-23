package internal

import (
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
)

func GetNewUuid() string {
	uid := uuid.New().String()

	return uid
}

func GetRandomPort(host string) int {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:0", host))
	if err != nil {
		log.Panicf(err.Error())
	}
	defer listen.Close()

	addr := listen.Addr().(*net.TCPAddr)
	return addr.Port
}
