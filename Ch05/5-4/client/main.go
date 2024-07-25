package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	reply := ""
	err = client.Call("RemService.HelloName", "rem", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

}
