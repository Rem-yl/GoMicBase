package main

import (
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	conf "Account/Conf"
	logger "Account/Log"
	share "Account/Share"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	logger.Init()
	config := conf.LoadConfig()
	host := config.GetString("grpc.host")
	port := config.GetString("grpc.port")
	dsn := fmt.Sprintf("%s:%s", host, port)

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
