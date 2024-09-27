package model

import (
	v1 "GoMicBase/api/account/service/v1"
)

func AccountModel2Pb(account Account) (resp *v1.AccountResponse) {
	resp = &v1.AccountResponse{
		Id:             uint32(account.ID),
		Name:           account.Name,
		Phone:          account.Phone,
		Password:       account.Password,
		Salt:           account.Salt,
		HashedPassword: account.HashedPassword,
	}

	return resp
}

func PbResp2CustomAccount(resp *v1.AccountResponse) *CustomAccount {
	account := &CustomAccount{
		Id:             resp.Id,
		Name:           resp.Name,
		Phone:          resp.Phone,
		Password:       resp.Password,
		Salt:           resp.Salt,
		HashedPassword: resp.HashedPassword,
	}

	return account
}
