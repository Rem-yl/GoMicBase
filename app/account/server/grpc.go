package server

import (
	v1 "GoMicBase/api/account/service/v1"
	"GoMicBase/app/account/service"
	"GoMicBase/pkg/cfg"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewGrpcServer(service *service.AccountService) *grpc.Server {
	grpcServer := grpc.NewServer()
	v1.RegisterAccountServiceServer(grpcServer, service)
	// register health check
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	return grpcServer
}

func GetGrpcClient(consulConfig cfg.ConsulConfig, servConfig cfg.GrpcServConfig) (v1.AccountServiceClient, error) {
	dsn := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, servConfig.Name)
	conn, err := grpc.Dial(
		dsn,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	if err != nil {
		return nil, err
	}

	client := v1.NewAccountServiceClient(conn)
	return client, nil
}
