package main

import (
	"Account/biz"
	"Account/internal"
	"Account/proto/pb"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func init() {
	internal.InitDB()
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "默认ip")
	port := flag.Int("port", 9095, "默认端口")
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	fmt.Println("Start New Grpc Server ... ")
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountService{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
