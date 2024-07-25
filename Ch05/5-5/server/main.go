package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Rem struct {
}

func (r *Rem) Hello(request string, resp *string) error {
	*resp = "hello, " + request
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
	}

	err = rpc.RegisterName("Rem", &Rem{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
