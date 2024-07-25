package main

import (
	"GoMicBase/Ch05/5-11/pb"
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewRemServiceClient(conn)
	//! 双向流模式
	names := []string{"rem", "ram", "fish"}
	fullStreamClient, err := client.FullStream(context.Background())
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			recv, err := fullStreamClient.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("recv %s\n", recv.Msg)
		}
	}()

	go func(names []string) {
		defer wg.Done()
		for _, name := range names {
			err := fullStreamClient.Send(&pb.RemStreamRequest{Name: name})
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}(names)

	wg.Wait()

	//! 客户端流模式
	// ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second) // 客户端10s退出
	// defer cancelFunc()

	// PostNameClient, err := client.PostName(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// Names := []string{"rem", "ram", "fish"}
	// for _, name := range Names {
	// 	err := PostNameClient.Send(&pb.RemStreamRequest{Name: name})
	// 	fmt.Printf("send %s\n", name)
	// 	time.Sleep(time.Second * 1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// }

	//! 服务端流模式
	// res, err := client.HelloName(context.Background(), &pb.RemStreamRequest{Name: "rem"})
	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	recv, err := res.Recv()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}

	// 	fmt.Println(recv.Msg)
	// }

}
