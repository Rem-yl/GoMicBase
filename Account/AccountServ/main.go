package main

import (
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	logger "Account/Log"
	share "Account/Share"
	"Account/internal"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	logger.Init()
	accountServConf := internal.AccountConf.AccountServConf
	dsn := fmt.Sprintf("%s:%d", accountServConf.Host, accountServConf.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, &service.AccountService{})
	listen, err := net.Listen("tcp", dsn)
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrListen, err.Error())
	}

	log.Printf("Start Account GRPC Service on %s\n", dsn)

	if err := grpcServer.Serve(listen); err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcServerFailed, err.Error())
	}
}
