// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: store.proto

package pb

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
	StoreService_CreateNewBrand_FullMethodName  = "/StoreService/CreateNewBrand"
	StoreService_GetBrandByName_FullMethodName  = "/StoreService/GetBrandByName"
	StoreService_GetBrandById_FullMethodName    = "/StoreService/GetBrandById"
	StoreService_UpdateBrandById_FullMethodName = "/StoreService/UpdateBrandById"
	StoreService_DeleteBrandById_FullMethodName = "/StoreService/DeleteBrandById"
)

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreServiceClient interface {
	// 品牌表服务
	// 增
	CreateNewBrand(ctx context.Context, in *CreateNewBrandRequest, opts ...grpc.CallOption) (*BrandResponse, error)
	// 查
	GetBrandByName(ctx context.Context, in *BrandNameRequest, opts ...grpc.CallOption) (*BrandResponse, error)
	GetBrandById(ctx context.Context, in *BrandIdRequest, opts ...grpc.CallOption) (*BrandResponse, error)
	// 改
	UpdateBrandById(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*BrandResponse, error)
	// 删
	DeleteBrandById(ctx context.Context, in *BrandIdRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) CreateNewBrand(ctx context.Context, in *CreateNewBrandRequest, opts ...grpc.CallOption) (*BrandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, StoreService_CreateNewBrand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) GetBrandByName(ctx context.Context, in *BrandNameRequest, opts ...grpc.CallOption) (*BrandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, StoreService_GetBrandByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) GetBrandById(ctx context.Context, in *BrandIdRequest, opts ...grpc.CallOption) (*BrandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, StoreService_GetBrandById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) UpdateBrandById(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*BrandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BrandResponse)
	err := c.cc.Invoke(ctx, StoreService_UpdateBrandById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) DeleteBrandById(ctx context.Context, in *BrandIdRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, StoreService_DeleteBrandById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
// All implementations must embed UnimplementedStoreServiceServer
// for forward compatibility.
type StoreServiceServer interface {
	// 品牌表服务
	// 增
	CreateNewBrand(context.Context, *CreateNewBrandRequest) (*BrandResponse, error)
	// 查
	GetBrandByName(context.Context, *BrandNameRequest) (*BrandResponse, error)
	GetBrandById(context.Context, *BrandIdRequest) (*BrandResponse, error)
	// 改
	UpdateBrandById(context.Context, *UpdateBrandRequest) (*BrandResponse, error)
	// 删
	DeleteBrandById(context.Context, *BrandIdRequest) (*CheckResponse, error)
	mustEmbedUnimplementedStoreServiceServer()
}

// UnimplementedStoreServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStoreServiceServer struct{}

func (UnimplementedStoreServiceServer) CreateNewBrand(context.Context, *CreateNewBrandRequest) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewBrand not implemented")
}
func (UnimplementedStoreServiceServer) GetBrandByName(context.Context, *BrandNameRequest) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandByName not implemented")
}
func (UnimplementedStoreServiceServer) GetBrandById(context.Context, *BrandIdRequest) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandById not implemented")
}
func (UnimplementedStoreServiceServer) UpdateBrandById(context.Context, *UpdateBrandRequest) (*BrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrandById not implemented")
}
func (UnimplementedStoreServiceServer) DeleteBrandById(context.Context, *BrandIdRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrandById not implemented")
}
func (UnimplementedStoreServiceServer) mustEmbedUnimplementedStoreServiceServer() {}
func (UnimplementedStoreServiceServer) testEmbeddedByValue()                      {}

// UnsafeStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServiceServer will
// result in compilation errors.
type UnsafeStoreServiceServer interface {
	mustEmbedUnimplementedStoreServiceServer()
}

func RegisterStoreServiceServer(s grpc.ServiceRegistrar, srv StoreServiceServer) {
	// If the following call pancis, it indicates UnimplementedStoreServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StoreService_ServiceDesc, srv)
}

func _StoreService_CreateNewBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).CreateNewBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_CreateNewBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).CreateNewBrand(ctx, req.(*CreateNewBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_GetBrandByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetBrandByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_GetBrandByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetBrandByName(ctx, req.(*BrandNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_GetBrandById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetBrandById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_GetBrandById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetBrandById(ctx, req.(*BrandIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_UpdateBrandById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).UpdateBrandById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_UpdateBrandById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).UpdateBrandById(ctx, req.(*UpdateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_DeleteBrandById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).DeleteBrandById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_DeleteBrandById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).DeleteBrandById(ctx, req.(*BrandIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreService_ServiceDesc is the grpc.ServiceDesc for StoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewBrand",
			Handler:    _StoreService_CreateNewBrand_Handler,
		},
		{
			MethodName: "GetBrandByName",
			Handler:    _StoreService_GetBrandByName_Handler,
		},
		{
			MethodName: "GetBrandById",
			Handler:    _StoreService_GetBrandById_Handler,
		},
		{
			MethodName: "UpdateBrandById",
			Handler:    _StoreService_UpdateBrandById_Handler,
		},
		{
			MethodName: "DeleteBrandById",
			Handler:    _StoreService_DeleteBrandById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}
