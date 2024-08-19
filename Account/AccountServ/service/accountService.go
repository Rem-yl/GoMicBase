package service

import (
	"Account/AccountServ/database"
	"Account/AccountServ/model"
	"Account/AccountServ/pb"
	share "Account/Share"
	"context"
	"errors"
	"log"

	"github.com/anaskhan96/go-password-encoder"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
}

func (server *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (resp *pb.AccountResponse, err error) {
	db := database.MysqlDB
	var account database.Account

	result := db.Where("name=?", req.Name).First(&account)
	if result.RowsAffected != 0 {
		log.Printf("Account Name Exist: Name: %s", account.Name)
		return nil, nil
	}

	result = db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected != 0 {
		log.Printf("Account Phone Exist: Name: %s", account.Phone)
		return nil, nil
	}

	salt, hashedPassword := password.Encode(req.Password, &share.PasswordOption)
	account = database.Account{
		Name:           req.Name,
		Phone:          req.Phone,
		Password:       req.Password,
		Salt:           salt,
		HashedPassword: hashedPassword,
	}

	result = db.Create(&account)
	if result.RowsAffected == 0 {
		log.Printf("%s : %s\n", share.ErrCreateAccount, result.Error.Error())
		return nil, result.Error
	}

	log.Printf("Create Account: Name: %s, Phone: %s\n", req.Name, req.Phone)
	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) GetAccountByName(ctx context.Context, req *pb.AccountNameRequest) (resp *pb.AccountResponse, err error) {
	db := database.MysqlDB
	var account database.Account
	result := db.Where("name=?", req.Name).First(&account)
	if result.RowsAffected == 0 {
		log.Printf("Account Not Found: %s", req.Name)
		return nil, errors.New(share.ErrAccountNotFound)
	}

	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) GetAccountByPhone(ctx context.Context, req *pb.AccountPhoneRequest) (resp *pb.AccountResponse, err error) {
	db := database.MysqlDB
	var account database.Account
	result := db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected == 0 {
		log.Printf("Account Not Found: Phone: %s", req.Phone)
		return nil, errors.New(share.ErrAccountNotFound)
	}

	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) GetAccountById(ctx context.Context, req *pb.AccountIdRequest) (resp *pb.AccountResponse, err error) {
	db := database.MysqlDB
	var account database.Account
	result := db.First(&account, req.Id)
	if result.RowsAffected == 0 {
		log.Printf("Account Not Found: Id: %d", req.Id)
		return nil, errors.New(share.ErrAccountNotFound)
	}

	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) CheckNamePassword(ctx context.Context, req *pb.CheckNamePasswordRequest) (resp *pb.CheckResponse, err error) {
	db := database.MysqlDB
	var account database.Account
	result := db.Where("name=?", req.Name).First(&account)
	if result.RowsAffected == 0 {
		log.Printf("Account Not Found: Name: %s", req.Name)
		return nil, errors.New(share.ErrAccountNotFound)
	}

	check := password.Verify(req.Password, account.Salt, account.HashedPassword, &share.PasswordOption)

	resp = &pb.CheckResponse{
		Check: check,
	}

	return resp, nil
}

func (server *AccountService) CheckPhonePassword(ctx context.Context, req *pb.CheckPhonePasswordRequest) (resp *pb.CheckResponse, err error) {
	db := database.MysqlDB
	var account database.Account
	result := db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected == 0 {
		log.Printf("Account Not Found: Phone: %s", req.Phone)
		return nil, errors.New(share.ErrAccountNotFound)
	}

	check := password.Verify(req.Password, account.Salt, account.HashedPassword, &share.PasswordOption)

	resp = &pb.CheckResponse{
		Check: check,
	}

	return resp, nil
}
