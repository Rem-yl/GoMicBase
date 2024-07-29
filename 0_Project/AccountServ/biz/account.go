package biz

import (
	"Account/custom_error"
	"Account/internal"
	"Account/model"
	"Account/proto/pb"
	"context"
	"errors"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
}

func Paginate(pageNumber, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNumber == 0 {
			pageNumber = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageNumber - 1) * pageSize

		// Offset 指定在开始返回记录之前要跳过的记录数量; Limit 限制指定要检索的记录数
		return db.Offset(offset).Limit(pageSize)
	}
}

func AccountModel2Resp(account model.Account) *pb.AccountResponse {
	accountResp := &pb.AccountResponse{
		Id:          0,
		PhoneNumber: account.PhoneNumber,
		Password:    account.Password,
		Nickname:    account.NickName,
		Gender:      uint32(account.Gender),
		Role:        uint32(account.Role),
	}

	return accountResp
}

func (server AccountService) GetAccountList(ctx context.Context, req *pb.PageinRequest) (resp *pb.AccountListResponse, err error) {
	var accountList []model.Account // 从数据库中读到的数据结构, resp是我们要根据数据结构构造的返回消息, 这里不要弄混了
	// result := internal.DB.Find(&accountList)	// 没有分页的代码
	// Scopes 将当前数据库连接传递给参数 `func(DB) DB`, 可用于动态添加条件
	result := internal.DB.Scopes(Paginate(int(req.PageNumber), int(req.PageSize))).Find(&accountList)
	if result.Error != nil {
		return nil, result.Error
	}

	accountListResp := &pb.AccountListResponse{}
	accountListResp.Total = int32(result.RowsAffected)

	for _, account := range accountList {
		accountResp := AccountModel2Resp(account)
		accountListResp.AccountList = append(accountListResp.AccountList, accountResp)
	}

	return accountListResp, nil
}

func (server AccountService) GetAccountByName(ctx context.Context, req *pb.NameRequest) (resp *pb.AccountResponse, err error) {
	var account model.Account
	result := internal.DB.Where(&model.Account{NickName: req.Name}).First(&account)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}

	resp = AccountModel2Resp(account)
	return resp, nil
}

func (server AccountService) GetAccountByPhoneNumber(ctx context.Context, req *pb.PhoneNumberRequest) (resp *pb.AccountResponse, err error) {
	var account model.Account
	result := internal.DB.Where(&model.Account{PhoneNumber: req.PhoneNumber}).First(&account)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}

	resp = AccountModel2Resp(account)
	return resp, nil
}

func (server AccountService) GetAccountById(ctx context.Context, req *pb.IdRequest) (resp *pb.AccountResponse, err error) {
	var account model.Account
	result := internal.DB.First(&account, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}

	resp = AccountModel2Resp(account)
	return resp, nil
}

func (server AccountService) AddAccount(ctx context.Context, req *pb.AddAccountRequest) (resp *pb.AccountResponse, err error) {
	var account model.Account
	// 检查手机号和昵称是否存在
	result := internal.DB.Find(&account, "phone_number=? OR nick_name=?", req.PhoneNumber, req.Nickname)
	if result.RowsAffected != 0 {
		return nil, errors.New(custom_error.AccountExist)
	}
	salt, encoder := password.Encode(req.Password, &DefaultOptions)

	account.NickName = req.Nickname
	account.PhoneNumber = req.PhoneNumber
	account.Salt = salt
	account.Password = encoder
	account.Gender = uint32(req.Gender)
	account.Role = 1

	r := internal.DB.Create(&account)
	if r.Error != nil {
		return nil, errors.New(custom_error.AccountCreateFail)
	}

	resp = AccountModel2Resp(account)
	return resp, nil
}

func (server AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (resp *pb.AccountResponse, err error) {
	var account model.Account
	result := internal.DB.First(&account, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}

	account.NickName = req.Nickname
	account.Gender = req.Gender
	account.Password = req.Password
	account.PhoneNumber = req.PhoneNumber

	result = internal.DB.Save(&account)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.UpdateFailed)
	}

	resp = AccountModel2Resp(account)
	return resp, nil
}

func (server AccountService) CheckNamePassword(ctx context.Context, req *pb.CheckAccountRequest) (resp *pb.CheckAccountResponse, err error) {
	var account model.Account
	result := internal.DB.First(&account, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}

	resp = &pb.CheckAccountResponse{}
	resp.Check = password.Verify(req.Password, account.Salt, account.Password, &DefaultOptions)

	return resp, nil
}
