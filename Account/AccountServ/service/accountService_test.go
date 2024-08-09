package service

import (
	"Account/AccountServ/pb"
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func TestCreateAccount(t *testing.T) {
	// server := AccountService{}
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewAccountServiceClient(conn)

	req := &pb.CreateAccountRequest{
		Name:     "Rem",
		Phone:    "18801117212",
		Password: "123456",
	}

	resp, err := client.CreateAccount(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Create Account: Name: %s, Phone: %s\n", resp.Name, resp.Phone)

}
