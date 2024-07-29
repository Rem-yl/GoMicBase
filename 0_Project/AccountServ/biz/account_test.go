package biz

import (
	"Account/custom_error"
	"Account/internal"
	"Account/proto/pb"
	"context"
	"fmt"
	"testing"
)

func init() {
	internal.InitDB()
}

func TestAccountServer_AddAccount(t *testing.T) {
	accountServer := AccountService{}
	repList := []*pb.AddAccountRequest{}
	account1 := &pb.AddAccountRequest{
		PhoneNumber: "18801117212",
		Password:    "123456",
		Nickname:    "Rem",
		Gender:      0,
		Role:        0,
	}
	account2 := &pb.AddAccountRequest{
		PhoneNumber: "18810992127",
		Password:    "1234567",
		Nickname:    "Ram",
		Gender:      1,
		Role:        1,
	}
	account3 := &pb.AddAccountRequest{
		PhoneNumber: "15357240810",
		Password:    "12345678",
		Nickname:    "Fish",
		Gender:      1,
		Role:        0,
	}
	repList = append(repList, account1, account2, account3)

	for _, rep := range repList {
		resp, err := accountServer.AddAccount(context.Background(), rep)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Id: %d, Name: %s\n", resp.Id, resp.Nickname)
	}
}

func TestAccountServer_GetAccountById(t *testing.T) {
	accountServer := AccountService{}
	req := &pb.IdRequest{
		Id: 1,
	}
	resp, err := accountServer.GetAccountById(context.Background(), req)
	if err != nil {
		fmt.Println(custom_error.AccountNotFound)
		return
	}

	fmt.Printf("ID: %d, Name: %s, phone: %s\n", resp.Id, resp.Nickname, resp.PhoneNumber)
}

func TestAccountServer_GetAccountByName(t *testing.T) {
	accountServer := AccountService{}
	rep := &pb.NameRequest{
		Name: "Rem",
	}
	resp, err := accountServer.GetAccountByName(context.Background(), rep)
	if err != nil {
		fmt.Println(custom_error.AccountNotFound)
		return
	}
	fmt.Printf("ID: %d, Name: %s, phone: %s\n", resp.Id, resp.Nickname, resp.PhoneNumber)
}

func TestAccountServer_GetAccountByPhoneNumber(t *testing.T) {
	accountServer := AccountService{}
	req := &pb.PhoneNumberRequest{
		PhoneNumber: "18801117212",
	}
	resp, err := accountServer.GetAccountByPhoneNumber(context.Background(), req)
	if err != nil {
		fmt.Println(custom_error.AccountNotFound)
		return
	}
	fmt.Printf("ID: %d, Name: %s, phone: %s\n", resp.Id, resp.Nickname, resp.PhoneNumber)
}

func TestAccountServer_GetAccountList(t *testing.T) {
	accountServer := AccountService{}
	req := &pb.PageinRequest{
		PageNumber: 0,
		PageSize:   10,
	}
	resp, err := accountServer.GetAccountList(context.Background(), req)
	if err != nil {
		fmt.Println(custom_error.AccountNotFound)
		return
	}

	fmt.Println(resp)
}

func TestAccountServer_CheckNamePassword(t *testing.T) {
	accountServer := AccountService{}
	req := &pb.CheckAccountRequest{
		Password: "123456",
		Id:       1,
	}
	resp, err := accountServer.CheckNamePassword(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	if resp.Check {
		fmt.Println("Password check ok!")
	} else {
		fmt.Println("Password check failed!")
	}
}
