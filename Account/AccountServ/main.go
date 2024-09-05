package main

import (
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	conf "Account/Conf"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	register "github.com/GoMicBase/Register"
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
	consulRegistery := &register.ConsulRegistery{
		Config: &register.ConsulConfig{
			Host: consulConf.Host,
			Port: consulConf.Port,
		},
	}

	if err = consulRegistery.NewClient(); err != nil {
		log.Panicln(err.Error())
	}

	if err = consulRegistery.RegisterGrpcServ(accountServConf.Host, int(accountServConf.Port), accountServConf.Name, accountServConf.Id, []string{"test"}); err != nil {
		log.Panicln(err.Error())
	}

	if err != nil {
		log.Panicf("%s:%s\n", share.ErrGrpcRegister, err.Error())
	}

	// listen grpc server
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Panicf("%s:%s\n", share.ErrGrpcServerFailed, err.Error())
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	log.Println("Received interrupt signal, shutting down gracefully...")
	// 注销Grpc服务
	if err = consulRegistery.Deregister(accountServConf.Id); err != nil {
		log.Println(err.Error())
	}

	// 停止Grpc服务
	grpcServer.GracefulStop()
	log.Println("Grpc server shut down gracefully")
}
