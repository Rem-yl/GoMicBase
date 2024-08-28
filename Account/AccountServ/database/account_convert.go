package database

import (
	"Account/AccountServ/pb"
)

func AccountModel2Pb(account Account) (resp *pb.AccountResponse) {
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

func PbResp2CustomAccount(resp *pb.AccountResponse) (account CustomAccount) {
	account = CustomAccount{
		Id:             resp.Id,
		Name:           resp.Name,
		Phone:          resp.Phone,
		Password:       resp.Password,
		Salt:           resp.Salt,
		HashedPassword: resp.HashedPassword,
	}

	return account
}
