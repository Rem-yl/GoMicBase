package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	c, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	reply := ""
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))

	err = client.Call("Rem.Hello", "rem", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

}
