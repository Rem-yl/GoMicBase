package model

import (
	"Account/AccountServ/database"
	"Account/AccountServ/pb"
	share "Account/Share"
)

func AccountModel2Pb(account database.Account) (resp *pb.AccountResponse) {
	resp = &pb.AccountResponse{
		Id:             uint32(account.ID),
		Name:           account.Name,
		Phone:          account.Phone,
		Password:       account.Password,
		Salt:           account.Salt,
		HashedPassword: account.HashedPassword,
	}

	return resp
}

func PbResp2CustomAccount(resp *pb.AccountResponse) (account share.CustomAccount) {
	account = share.CustomAccount{
		Id:             resp.Id,
		Name:           resp.Name,
		Phone:          resp.Phone,
		Password:       resp.Password,
		Salt:           resp.Salt,
		HashedPassword: resp.HashedPassword,
	}

	return account
}
