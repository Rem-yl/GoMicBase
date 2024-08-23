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
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	logger.Init()
	accountServConf := internal.AccountConf.AccountServConf
	dsn := fmt.Sprintf("%s:%d", accountServConf.Host, accountServConf.Port)

	grpcServer := grpc.NewServer()

	// register accountServer
	pb.RegisterAccountServiceServer(grpcServer, &service.AccountService{})
	listen, err := net.Listen("tcp", dsn)
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrListen, err.Error())
	}
	log.Printf("Start Account GRPC Service on %s\n", dsn)

	// register health check
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	// listen grpc server
	if err := grpcServer.Serve(listen); err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcServerFailed, err.Error())
	}
}
