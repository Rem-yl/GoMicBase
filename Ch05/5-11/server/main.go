package main

import (
	"GoMicBase/Ch05/5-11/pb"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type Rem struct {
	pb.UnimplementedRemServiceServer
}

func (r *Rem) HelloName(request *pb.RemStreamRequest, server pb.RemService_HelloNameServer) error {
	fmt.Println("收到客户端请求")
	server.Send(&pb.RemStreamResponse{Msg: "hello, " + request.Name})
	return nil
}

func (r *Rem) PostName(server pb.RemService_PostNameServer) error {
	for {
		recv, err := server.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv)
	}
	return nil
}

func (r *Rem) FullStream(server pb.RemService_FullStreamServer) error {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan string, 5)
	go func() {
		defer wg.Done()
		for {
			recv, err := server.Recv()
			if err != nil {
				fmt.Println(err)
			}
			c <- recv.Name
			fmt.Printf("get name: %s\n", recv.Name)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			name := <-c
			err := server.Send(&pb.RemStreamResponse{Msg: "hello, " + name})
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterRemServiceServer(server, &Rem{})

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
