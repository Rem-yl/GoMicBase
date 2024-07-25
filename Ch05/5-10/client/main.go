package main

import (
	"GoMicBase/Ch05/5-10/server/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewStudyClient(c)
	resp, err := client.Study(context.Background(), &pb.BookRequest{Name: "rem"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Msg)

}
