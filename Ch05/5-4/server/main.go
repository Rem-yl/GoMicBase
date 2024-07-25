package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type RemService struct {
}

func (r *RemService) HelloName(request string, resp *string) error {
	// 符合go rpc格式的函数签名
	*resp = "Hello, " + request
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err) // 一般不在服务端panic
	}

	err = rpc.RegisterName("RemService", &RemService{})
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	rpc.ServeConn(conn)

}
