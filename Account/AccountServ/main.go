package main

import (
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	share "Account/Share"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, &service.AccountService{})
	listen, err := net.Listen("tcp", "127.0.0.1:9095")
	if err != nil {
		log.Fatalln(share.ErrListen + err.Error())
	}

	log.Println("Start Account GRPC Service on 127.0.0.1:9095")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf(share.ErrGrpcServerFailed + err.Error())
	}
}
