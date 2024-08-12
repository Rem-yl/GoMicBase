package main

import (
	"Account/AccountServ/pb"
	"Account/AccountServ/service"
	share "Account/Share"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func loadConfig() *viper.Viper {
	config := viper.New()

	config.AddConfigPath("./conf")
	config.SetConfigName("default")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Panicf("%s : %s\n", share.ErrConfigFileNotFound, err.Error())
	}

	return config
}

func main() {
	config := loadConfig()
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
