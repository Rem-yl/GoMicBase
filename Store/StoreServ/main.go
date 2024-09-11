package main

import (
	conf "Store/Conf"
	"Store/StoreServ/pb"
	"Store/StoreServ/service"
	"fmt"
	"log"
	"net"

	register "github.com/GoMicBase/Register"
	share "github.com/GoMicBase/Share"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// grpc注册服务
	storeGrpcConf := conf.StoreConf.StoreServConf
	consulConf := conf.StoreConf.ConsulConf

	dsn := fmt.Sprintf("%s:%d", storeGrpcConf.Host, storeGrpcConf.Port)

	grpcServer := grpc.NewServer()

	// register storeServer
	pb.RegisterStoreServiceServer(grpcServer, &service.StoreService{})
	listen, err := net.Listen("tcp", dsn)
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrListen, err.Error())
	}
	log.Printf("Start Store GRPC Service on %s\n", dsn)

	// register health check
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	consulRegistery := &register.ConsulRegistery{
		Config: &register.ConsulConfig{
			Host: consulConf.Host,
			Port: consulConf.Port,
		},
	}

	err = consulRegistery.NewClient()
	if err != nil {
		log.Println(err.Error())
	}

	err = consulRegistery.RegisterGrpcServ(storeGrpcConf.Host, int(storeGrpcConf.Port), storeGrpcConf.Name, storeGrpcConf.Id, []string{"test"})
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcRegister, err.Error())
	}

	// listen grpc server
	if err := grpcServer.Serve(listen); err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcServerFailed, err.Error())
	}
}
