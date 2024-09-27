package service

import (
	v1 "GoMicBase/api/account/service/v1"
	"GoMicBase/app/account/model"
	"GoMicBase/pkg/zlog"
	"GoMicBase/share"
	"context"
	"errors"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
)

type AccountService struct {
	v1.UnimplementedAccountServiceServer

	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	if err := db.AutoMigrate(&model.Account{}); err != nil {
		zlog.Panicln(err.Error())
	}

	return &AccountService{
		db: db,
	}
}

func (server *AccountService) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account
	result := server.db.Where("name=?", req.Name).First(&account)

	if result.RowsAffected != 0 {
		zlog.Infof("Account Name Exist: Name: %s", account.Name)
		return nil, nil
	}

	result = server.db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected != 0 {
		zlog.Infof("Account Phone Exist: Name: %s", account.Phone)
		return nil, nil
	}
	salt, hashedPassword := password.Encode(req.Password, share.PasswordOption)

	newAccount := model.Account{
		Name:           req.Name,
		Phone:          req.Phone,
		Password:       req.Password,
		Salt:           salt,
		HashedPassword: hashedPassword,
	}
	result = server.db.Create(&newAccount)

	if result.RowsAffected == 0 {
		zlog.Info(result.Error.Error())
		return nil, result.Error
	}

	zlog.Infof("Create Account: Name: %s, Phone: %s\n", req.Name, req.Phone)
	resp = model.AccountModel2Pb(newAccount)
	return resp, nil
}

func (server *AccountService) GetAccountList(ctx context.Context, req *v1.AccountListRequest) (resp *v1.AccountListResponse, err error) {
	var accounts []model.Account
	page := int(req.Page)
	pageSize := int(req.Pagesize)

	result := server.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&accounts)
	resp = &v1.AccountListResponse{
		Total:    0,
		Accounts: []*v1.AccountResponse{},
	}
	total := result.RowsAffected
	if total == 0 {
		return resp, nil
	}

	resp.Total = int32(total)
	for _, account := range accounts {
		accountResp := &v1.AccountResponse{
			Id:             uint32(account.ID),
			Name:           account.Name,
			Phone:          account.Phone,
			Password:       account.Password,
			Salt:           account.Salt,
			HashedPassword: account.HashedPassword,
		}
		resp.Accounts = append(resp.Accounts, accountResp)
	}
	return resp, nil
}

func (server *AccountService) GetAccountByName(ctx context.Context, req *v1.AccountNameRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account
	result := server.db.Where("name=?", req.Name).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: %s", req.Name)
		return nil, result.Error
	}

	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) GetAccountByPhone(ctx context.Context, req *v1.AccountPhoneRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account
	result := server.db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Phone: %s", req.Phone)
		return nil, result.Error
	}

	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) GetAccountById(ctx context.Context, req *v1.AccountIdRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account
	result := server.db.First(&account, req.Id)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Id: %d", req.Id)
		return nil, result.Error
	}

	resp = model.AccountModel2Pb(account)
	return resp, nil
}

func (server *AccountService) CheckNamePassword(ctx context.Context, req *v1.CheckNamePasswordRequest) (resp *v1.CheckResponse, err error) {
	var account model.Account
	result := server.db.Where("name=?", req.Name).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Name: %s", req.Name)
		return nil, result.Error
	}

	check := password.Verify(req.Password, account.Salt, account.HashedPassword, share.PasswordOption)

	resp = &v1.CheckResponse{
		Check: check,
	}

	return resp, nil
}

func (server *AccountService) CheckPhonePassword(ctx context.Context, req *v1.CheckPhonePasswordRequest) (resp *v1.CheckResponse, err error) {
	var account model.Account
	result := server.db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Phone: %s", req.Phone)
		return nil, result.Error
	}

	check := password.Verify(req.Password, account.Salt, account.HashedPassword, share.PasswordOption)

	resp = &v1.CheckResponse{
		Check: check,
	}

	return resp, nil
}

func (server *AccountService) ModifyAccountByPhone(ctx context.Context, req *v1.ModifyAccountPhoneRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account

	result := server.db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Phone: %s", req.Phone)
		return nil, result.Error
	}

	// 更新用户信息
	if req.Name == "" {
		req.Name = account.Name
	} else if req.Name == account.Name {
		zlog.Infof("Name: %s already exists!", req.Name)
		return nil, errors.New("Name exists.")
	}

	if req.Password == "" {
		req.Password = account.Password
	}
	salt, hashedPassword := password.Encode(req.Password, share.PasswordOption) // 更新用户密码

	update := map[string]interface{}{
		"name":            req.Name,
		"password":        req.Password,
		"hashed_password": hashedPassword,
		"salt":            salt,
	}

	result.Updates(update)
	resp = &v1.AccountResponse{
		Id:             uint32(account.ID),
		Name:           account.Name,
		Phone:          account.Phone,
		Password:       account.Password,
		Salt:           account.Salt,
		HashedPassword: account.HashedPassword,
	}

	return resp, nil
}

func (server *AccountService) DeleteAccountByName(ctx context.Context, req *v1.AccountNameRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account

	result := server.db.Where("name=?", req.Name).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Name: %s", req.Name)
		return nil, result.Error
	}

	result.Delete(account.ID)

	resp = &v1.AccountResponse{
		Id:             uint32(account.ID),
		Name:           account.Name,
		Phone:          account.Phone,
		Password:       account.Phone,
		Salt:           account.Salt,
		HashedPassword: account.HashedPassword,
	}
	return resp, nil
}

func (server *AccountService) DeleteAccountByPhone(ctx context.Context, req *v1.AccountPhoneRequest) (resp *v1.AccountResponse, err error) {
	var account model.Account

	result := server.db.Where("phone=?", req.Phone).First(&account)
	if result.RowsAffected == 0 {
		zlog.Infof("Account Not Found: Phone: %s", req.Phone)
		return nil, result.Error
	}

	result.Delete(account.ID)

	resp = &v1.AccountResponse{
		Id:             uint32(account.ID),
		Name:           account.Name,
		Phone:          account.Phone,
		Password:       account.Phone,
		Salt:           account.Salt,
		HashedPassword: account.HashedPassword,
	}
	return resp, nil
}
