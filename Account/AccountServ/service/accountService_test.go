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
		log.Panicln(err)
	}

	client := pb.NewAccountServiceClient(conn)

	req := &pb.CreateAccountRequest{
		Name:     "Ram",
		Phone:    "18801117212",
		Password: "123456",
	}

	resp, err := client.CreateAccount(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("xxxxxxx: %v", resp)
	if resp == nil {
		log.Printf("Account Exist: Name: %s, Phone: %s\n", req.Name, req.Phone)
	} else {
		log.Printf("Create Account: Name: %s, Phone: %s\n", resp.Name, resp.Phone)
	}
}

func TestGetAccountByName(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	client := pb.NewAccountServiceClient(conn)

	req := &pb.AccountNameRequest{
		Name: "ley",
	}

	resp, err := client.GetAccountByName(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Name: %s, Phone: %s", resp.Name, resp.Phone)
}

func TestGetAccountByPhone(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	client := pb.NewAccountServiceClient(conn)
	req := pb.AccountPhoneRequest{
		Phone: "18801117212",
	}

	resp, err := client.GetAccountByPhone(context.Background(), &req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Name: %s, Phone: %s", resp.Name, resp.Phone)
}

func TestGGetAccountById(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	client := pb.NewAccountServiceClient(conn)
	req := pb.AccountIdRequest{
		Id: 123,
	}

	resp, err := client.GetAccountById(context.Background(), &req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Name: %s, Phone: %s", resp.Name, resp.Phone)
}
