package main

import (
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	conf "Account/Conf"
	"fmt"
	"log"
	"net"

	share "github.com/GoMicBase/Share"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	accountServConf := conf.AccountConf.AccountServConf
	consulConf := conf.AccountConf.ConsulConf
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
	consulClient, err := share.GetConsulClient(share.ConsulConfig(consulConf))
	if err != nil {
		log.Panicln(err.Error())
	}
	err = share.ConsulRegGrpc(consulClient, accountServConf.Host, int(accountServConf.Port), accountServConf.Name, accountServConf.Id, []string{"test"})
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcRegister, err.Error())
	}

	// listen grpc server
	if err := grpcServer.Serve(listen); err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcServerFailed, err.Error())
	}
}
