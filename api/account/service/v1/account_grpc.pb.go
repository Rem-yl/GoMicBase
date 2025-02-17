// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: account.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountService_CreateAccount_FullMethodName        = "/account.v1.AccountService/CreateAccount"
	AccountService_GetAccountList_FullMethodName       = "/account.v1.AccountService/GetAccountList"
	AccountService_GetAccountByName_FullMethodName     = "/account.v1.AccountService/GetAccountByName"
	AccountService_GetAccountByPhone_FullMethodName    = "/account.v1.AccountService/GetAccountByPhone"
	AccountService_GetAccountById_FullMethodName       = "/account.v1.AccountService/GetAccountById"
	AccountService_CheckNamePassword_FullMethodName    = "/account.v1.AccountService/CheckNamePassword"
	AccountService_CheckPhonePassword_FullMethodName   = "/account.v1.AccountService/CheckPhonePassword"
	AccountService_DeleteAccountByName_FullMethodName  = "/account.v1.AccountService/DeleteAccountByName"
	AccountService_DeleteAccountByPhone_FullMethodName = "/account.v1.AccountService/DeleteAccountByPhone"
	AccountService_ModifyAccountByPhone_FullMethodName = "/account.v1.AccountService/ModifyAccountByPhone"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	// create
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	// get
	GetAccountList(ctx context.Context, in *AccountListRequest, opts ...grpc.CallOption) (*AccountListResponse, error)
	GetAccountByName(ctx context.Context, in *AccountNameRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetAccountByPhone(ctx context.Context, in *AccountPhoneRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetAccountById(ctx context.Context, in *AccountIdRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	CheckNamePassword(ctx context.Context, in *CheckNamePasswordRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	CheckPhonePassword(ctx context.Context, in *CheckPhonePasswordRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// delete
	DeleteAccountByName(ctx context.Context, in *AccountNameRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	DeleteAccountByPhone(ctx context.Context, in *AccountPhoneRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	// modify
	ModifyAccountByPhone(ctx context.Context, in *ModifyAccountPhoneRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountList(ctx context.Context, in *AccountListRequest, opts ...grpc.CallOption) (*AccountListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountListResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccountList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByName(ctx context.Context, in *AccountNameRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccountByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByPhone(ctx context.Context, in *AccountPhoneRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccountByPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountById(ctx context.Context, in *AccountIdRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_GetAccountById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CheckNamePassword(ctx context.Context, in *CheckNamePasswordRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, AccountService_CheckNamePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CheckPhonePassword(ctx context.Context, in *CheckPhonePasswordRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, AccountService_CheckPhonePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccountByName(ctx context.Context, in *AccountNameRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteAccountByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccountByPhone(ctx context.Context, in *AccountPhoneRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteAccountByPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ModifyAccountByPhone(ctx context.Context, in *ModifyAccountPhoneRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_ModifyAccountByPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility.
type AccountServiceServer interface {
	// create
	CreateAccount(context.Context, *CreateAccountRequest) (*AccountResponse, error)
	// get
	GetAccountList(context.Context, *AccountListRequest) (*AccountListResponse, error)
	GetAccountByName(context.Context, *AccountNameRequest) (*AccountResponse, error)
	GetAccountByPhone(context.Context, *AccountPhoneRequest) (*AccountResponse, error)
	GetAccountById(context.Context, *AccountIdRequest) (*AccountResponse, error)
	CheckNamePassword(context.Context, *CheckNamePasswordRequest) (*CheckResponse, error)
	CheckPhonePassword(context.Context, *CheckPhonePasswordRequest) (*CheckResponse, error)
	// delete
	DeleteAccountByName(context.Context, *AccountNameRequest) (*AccountResponse, error)
	DeleteAccountByPhone(context.Context, *AccountPhoneRequest) (*AccountResponse, error)
	// modify
	ModifyAccountByPhone(context.Context, *ModifyAccountPhoneRequest) (*AccountResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountServiceServer struct{}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountList(context.Context, *AccountListRequest) (*AccountListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountList not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByName(context.Context, *AccountNameRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByName not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByPhone(context.Context, *AccountPhoneRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByPhone not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountById(context.Context, *AccountIdRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}
func (UnimplementedAccountServiceServer) CheckNamePassword(context.Context, *CheckNamePasswordRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNamePassword not implemented")
}
func (UnimplementedAccountServiceServer) CheckPhonePassword(context.Context, *CheckPhonePasswordRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPhonePassword not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccountByName(context.Context, *AccountNameRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccountByName not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccountByPhone(context.Context, *AccountPhoneRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccountByPhone not implemented")
}
func (UnimplementedAccountServiceServer) ModifyAccountByPhone(context.Context, *ModifyAccountPhoneRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyAccountByPhone not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}
func (UnimplementedAccountServiceServer) testEmbeddedByValue()                        {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccountList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountList(ctx, req.(*AccountListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccountByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByName(ctx, req.(*AccountNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccountByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByPhone(ctx, req.(*AccountPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_GetAccountById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountById(ctx, req.(*AccountIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CheckNamePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckNamePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CheckNamePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CheckNamePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CheckNamePassword(ctx, req.(*CheckNamePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CheckPhonePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPhonePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CheckPhonePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CheckPhonePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CheckPhonePassword(ctx, req.(*CheckPhonePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccountByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccountByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteAccountByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccountByName(ctx, req.(*AccountNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccountByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccountByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteAccountByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccountByPhone(ctx, req.(*AccountPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ModifyAccountByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyAccountPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ModifyAccountByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_ModifyAccountByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ModifyAccountByPhone(ctx, req.(*ModifyAccountPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.v1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccountList",
			Handler:    _AccountService_GetAccountList_Handler,
		},
		{
			MethodName: "GetAccountByName",
			Handler:    _AccountService_GetAccountByName_Handler,
		},
		{
			MethodName: "GetAccountByPhone",
			Handler:    _AccountService_GetAccountByPhone_Handler,
		},
		{
			MethodName: "GetAccountById",
			Handler:    _AccountService_GetAccountById_Handler,
		},
		{
			MethodName: "CheckNamePassword",
			Handler:    _AccountService_CheckNamePassword_Handler,
		},
		{
			MethodName: "CheckPhonePassword",
			Handler:    _AccountService_CheckPhonePassword_Handler,
		},
		{
			MethodName: "DeleteAccountByName",
			Handler:    _AccountService_DeleteAccountByName_Handler,
		},
		{
			MethodName: "DeleteAccountByPhone",
			Handler:    _AccountService_DeleteAccountByPhone_Handler,
		},
		{
			MethodName: "ModifyAccountByPhone",
			Handler:    _AccountService_ModifyAccountByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
