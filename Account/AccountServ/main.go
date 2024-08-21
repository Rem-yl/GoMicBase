package main

import (
	"Account/AccountServ/database"
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	logger "Account/Log"
	share "Account/Share"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	logger.Init()

	dsn := fmt.Sprintf("%s:%d", database.AccountServConfig.AccountGrpcConf.Host, database.AccountServConfig.AccountGrpcConf.Port)

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
