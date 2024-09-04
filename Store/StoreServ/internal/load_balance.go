package internal

import (
	conf "Store/StoreServ/Conf"
	"Store/StoreServ/pb"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" //! It's important

	"google.golang.org/grpc"
)

// load StoreGrpcServ client from consul
func GetStoreServClient() (pb.StoreServiceClient, error) {
	consulConf := conf.StoreServConf.ConsulConf
	servName := conf.StoreServConf.StoreGrpcConf.Name

	dsn := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConf.Host, consulConf.Port, servName)

	conn, err := grpc.Dial(
		dsn,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	if err != nil {
		return nil, err
	}

	client := pb.NewStoreServiceClient(conn)
	return client, nil
}
