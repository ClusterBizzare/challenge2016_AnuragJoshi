// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: distributor/v1/distributor.proto

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
	CreateDistributorService_CreateDistributor_FullMethodName = "/distributor.v1.CreateDistributorService/CreateDistributor"
)

// CreateDistributorServiceClient is the client API for CreateDistributorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateDistributorServiceClient interface {
	CreateDistributor(ctx context.Context, in *CreateDistributorRequest, opts ...grpc.CallOption) (*CreateDistributorResponse, error)
}

type createDistributorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateDistributorServiceClient(cc grpc.ClientConnInterface) CreateDistributorServiceClient {
	return &createDistributorServiceClient{cc}
}

func (c *createDistributorServiceClient) CreateDistributor(ctx context.Context, in *CreateDistributorRequest, opts ...grpc.CallOption) (*CreateDistributorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDistributorResponse)
	err := c.cc.Invoke(ctx, CreateDistributorService_CreateDistributor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateDistributorServiceServer is the server API for CreateDistributorService service.
// All implementations must embed UnimplementedCreateDistributorServiceServer
// for forward compatibility.
type CreateDistributorServiceServer interface {
	CreateDistributor(context.Context, *CreateDistributorRequest) (*CreateDistributorResponse, error)
	mustEmbedUnimplementedCreateDistributorServiceServer()
}

// UnimplementedCreateDistributorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCreateDistributorServiceServer struct{}

func (UnimplementedCreateDistributorServiceServer) CreateDistributor(context.Context, *CreateDistributorRequest) (*CreateDistributorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDistributor not implemented")
}
func (UnimplementedCreateDistributorServiceServer) mustEmbedUnimplementedCreateDistributorServiceServer() {
}
func (UnimplementedCreateDistributorServiceServer) testEmbeddedByValue() {}

// UnsafeCreateDistributorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateDistributorServiceServer will
// result in compilation errors.
type UnsafeCreateDistributorServiceServer interface {
	mustEmbedUnimplementedCreateDistributorServiceServer()
}

func RegisterCreateDistributorServiceServer(s grpc.ServiceRegistrar, srv CreateDistributorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCreateDistributorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CreateDistributorService_ServiceDesc, srv)
}

func _CreateDistributorService_CreateDistributor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDistributorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateDistributorServiceServer).CreateDistributor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateDistributorService_CreateDistributor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateDistributorServiceServer).CreateDistributor(ctx, req.(*CreateDistributorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateDistributorService_ServiceDesc is the grpc.ServiceDesc for CreateDistributorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateDistributorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "distributor.v1.CreateDistributorService",
	HandlerType: (*CreateDistributorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDistributor",
			Handler:    _CreateDistributorService_CreateDistributor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributor/v1/distributor.proto",
}

const (
	CheckPermissionService_CheckPermission_FullMethodName = "/distributor.v1.CheckPermissionService/CheckPermission"
)

// CheckPermissionServiceClient is the client API for CheckPermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckPermissionServiceClient interface {
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
}

type checkPermissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckPermissionServiceClient(cc grpc.ClientConnInterface) CheckPermissionServiceClient {
	return &checkPermissionServiceClient{cc}
}

func (c *checkPermissionServiceClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, CheckPermissionService_CheckPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckPermissionServiceServer is the server API for CheckPermissionService service.
// All implementations must embed UnimplementedCheckPermissionServiceServer
// for forward compatibility.
type CheckPermissionServiceServer interface {
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	mustEmbedUnimplementedCheckPermissionServiceServer()
}

// UnimplementedCheckPermissionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCheckPermissionServiceServer struct{}

func (UnimplementedCheckPermissionServiceServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedCheckPermissionServiceServer) mustEmbedUnimplementedCheckPermissionServiceServer() {
}
func (UnimplementedCheckPermissionServiceServer) testEmbeddedByValue() {}

// UnsafeCheckPermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckPermissionServiceServer will
// result in compilation errors.
type UnsafeCheckPermissionServiceServer interface {
	mustEmbedUnimplementedCheckPermissionServiceServer()
}

func RegisterCheckPermissionServiceServer(s grpc.ServiceRegistrar, srv CheckPermissionServiceServer) {
	// If the following call pancis, it indicates UnimplementedCheckPermissionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CheckPermissionService_ServiceDesc, srv)
}

func _CheckPermissionService_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckPermissionServiceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckPermissionService_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckPermissionServiceServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckPermissionService_ServiceDesc is the grpc.ServiceDesc for CheckPermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckPermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "distributor.v1.CheckPermissionService",
	HandlerType: (*CheckPermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPermission",
			Handler:    _CheckPermissionService_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributor/v1/distributor.proto",
}