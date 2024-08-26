package service

import (
	"Account/AccountServ/pb"
	"Account/internal"
	"context"
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicf("Get Account Serv Client Failed, %s\n", err.Error())
	}

	req := &pb.CreateAccountRequest{
		Name:     "ley",
		Phone:    "18801117213",
		Password: "123457",
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

func TestGetAccountList(t *testing.T) {
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	rep := &pb.AccountListRequest{
		Page:     1,
		Pagesize: 10,
	}

	resp, err := client.GetAccountList(context.Background(), rep)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Total Account Num: %d\n", resp.Total)
	for _, v := range resp.Accounts {
		log.Printf("Name: %s, Phone: %s\n", v.Name, v.Phone)
	}
}

func TestGetAccountByName(t *testing.T) {
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicf("Get Account Serv Client Failed, %s\n", err.Error())
	}

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
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicf("Get Account Serv Client Failed, %s\n", err.Error())
	}
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

func TestGetAccountById(t *testing.T) {
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicf("Get Account Serv Client Failed, %s\n", err.Error())
	}

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

func TestCheckNamePassword(t *testing.T) {
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicf("Get Account Serv Client Failed, %s\n", err.Error())
	}

	req := &pb.CheckNamePasswordRequest{
		Name:     "Ram",
		Password: "123456",
	}

	resp, err := client.CheckNamePassword(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Name: %s, Phone: %s, check: %t\n", req.Name, req.Password, resp.Check)
}

func TestModifyAccountByPhone(t *testing.T) {
	client, err := internal.GetAccountServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	req := &pb.ModifyAccountPhoneRequest{
		Phone:    "18801117213",
		Name:     "ley1",
		Password: "",
	}
	resp, err := client.ModifyAccountByPhone(context.Background(), req)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Name: %s, Phone: %s, Password: %s", resp.Name, resp.Phone, resp.Password)
}
