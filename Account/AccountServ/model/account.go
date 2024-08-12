package model

import (
	"Account/AccountServ/database"
	"Account/AccountServ/pb"
)

func AccountModel2Pb(account database.Account) (resp *pb.AccountResponse) {
	resp = &pb.AccountResponse{
		Id:       uint32(account.ID),
		Name:     account.Name,
		Phone:    account.Phone,
		Password: account.Password,
	}

	return resp
}
