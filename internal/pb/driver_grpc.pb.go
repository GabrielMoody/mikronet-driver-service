// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: driver.proto

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
	DriverService_CreateDriver_FullMethodName      = "/dashboard.DriverService/CreateDriver"
	DriverService_GetDrivers_FullMethodName        = "/dashboard.DriverService/GetDrivers"
	DriverService_GetDriverDetails_FullMethodName  = "/dashboard.DriverService/GetDriverDetails"
	DriverService_SetStatusVerified_FullMethodName = "/dashboard.DriverService/SetStatusVerified"
)

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverServiceClient interface {
	CreateDriver(ctx context.Context, in *CreateDriverRequest, opts ...grpc.CallOption) (*Driver, error)
	GetDrivers(ctx context.Context, in *ReqDrivers, opts ...grpc.CallOption) (*Drivers, error)
	GetDriverDetails(ctx context.Context, in *ReqByID, opts ...grpc.CallOption) (*Driver, error)
	SetStatusVerified(ctx context.Context, in *ReqByID, opts ...grpc.CallOption) (*Driver, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) CreateDriver(ctx context.Context, in *CreateDriverRequest, opts ...grpc.CallOption) (*Driver, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Driver)
	err := c.cc.Invoke(ctx, DriverService_CreateDriver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetDrivers(ctx context.Context, in *ReqDrivers, opts ...grpc.CallOption) (*Drivers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Drivers)
	err := c.cc.Invoke(ctx, DriverService_GetDrivers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetDriverDetails(ctx context.Context, in *ReqByID, opts ...grpc.CallOption) (*Driver, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Driver)
	err := c.cc.Invoke(ctx, DriverService_GetDriverDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) SetStatusVerified(ctx context.Context, in *ReqByID, opts ...grpc.CallOption) (*Driver, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Driver)
	err := c.cc.Invoke(ctx, DriverService_SetStatusVerified_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
// All implementations must embed UnimplementedDriverServiceServer
// for forward compatibility.
type DriverServiceServer interface {
	CreateDriver(context.Context, *CreateDriverRequest) (*Driver, error)
	GetDrivers(context.Context, *ReqDrivers) (*Drivers, error)
	GetDriverDetails(context.Context, *ReqByID) (*Driver, error)
	SetStatusVerified(context.Context, *ReqByID) (*Driver, error)
	mustEmbedUnimplementedDriverServiceServer()
}

// UnimplementedDriverServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverServiceServer struct{}

func (UnimplementedDriverServiceServer) CreateDriver(context.Context, *CreateDriverRequest) (*Driver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDriver not implemented")
}
func (UnimplementedDriverServiceServer) GetDrivers(context.Context, *ReqDrivers) (*Drivers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrivers not implemented")
}
func (UnimplementedDriverServiceServer) GetDriverDetails(context.Context, *ReqByID) (*Driver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverDetails not implemented")
}
func (UnimplementedDriverServiceServer) SetStatusVerified(context.Context, *ReqByID) (*Driver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatusVerified not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}
func (UnimplementedDriverServiceServer) testEmbeddedByValue()                       {}

// UnsafeDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServiceServer will
// result in compilation errors.
type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv DriverServiceServer) {
	// If the following call pancis, it indicates UnimplementedDriverServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_CreateDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).CreateDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_CreateDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).CreateDriver(ctx, req.(*CreateDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDrivers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_GetDrivers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetDrivers(ctx, req.(*ReqDrivers))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetDriverDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetDriverDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_GetDriverDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetDriverDetails(ctx, req.(*ReqByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_SetStatusVerified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).SetStatusVerified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_SetStatusVerified_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).SetStatusVerified(ctx, req.(*ReqByID))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverService_ServiceDesc is the grpc.ServiceDesc for DriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDriver",
			Handler:    _DriverService_CreateDriver_Handler,
		},
		{
			MethodName: "GetDrivers",
			Handler:    _DriverService_GetDrivers_Handler,
		},
		{
			MethodName: "GetDriverDetails",
			Handler:    _DriverService_GetDriverDetails_Handler,
		},
		{
			MethodName: "SetStatusVerified",
			Handler:    _DriverService_SetStatusVerified_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}
