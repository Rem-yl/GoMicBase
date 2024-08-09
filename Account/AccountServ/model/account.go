package model

import (
	"Account/AccountServ/pb"
	share "Account/Share"
)

func AccountModel2Pb(account share.Account) (resp *pb.AccountResponse) {
	resp = &pb.AccountResponse{
		Id:       uint32(account.ID),
		Name:     account.Name,
		Phone:    account.Phone,
		Password: account.Password,
	}

	return resp
}
