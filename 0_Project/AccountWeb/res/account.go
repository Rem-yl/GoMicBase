package res

import (
	"Account/proto/pb"
)

type AccountRes struct {
	Id          uint32 `json:"id"`
	PhoneNumber string `json:"phone_number"`
	NickName    string `json:"nick_name"`
	Gender      uint32 `json:"gender"`
}

func Pb2model(resp *pb.AccountResponse) AccountRes {
	res := AccountRes{
		Id:          resp.Id,
		PhoneNumber: resp.PhoneNumber,
		NickName:    resp.Nickname,
		Gender:      resp.Gender,
	}

	return res
}
