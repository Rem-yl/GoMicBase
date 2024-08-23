package internal

import (
	"Account/AccountServ/pb"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
)

func GetAccountServClient() (client pb.AccountServiceClient, err error) {
	consulConf := AccountConf.ConsulConf
	servName := AccountConf.AccountServConf.Name
	dsn := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConf.Host, consulConf.Port, servName)
	conn, err := grpc.Dial(
		dsn,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	if err != nil {
		return nil, err
	}

	client = pb.NewAccountServiceClient(conn)
	return client, nil
}
