package main

import (
	"GoMicBase/Ch05/5-10/server/pb"
	"context"
	"net"

	"google.golang.org/grpc"
)

type BookInfo struct {
	pb.UnimplementedStudyServer
}

func (b *BookInfo) Study(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{Msg: "hello, " + req.Name}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStudyServer(server, &BookInfo{})

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
