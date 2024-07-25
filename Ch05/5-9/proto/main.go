package main

import (
	"GoMicBase/Ch05/5-9/proto/pb"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	req := pb.BookRequest{Name: "rem"}
	data, err := proto.Marshal(&req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(string(data))
}
