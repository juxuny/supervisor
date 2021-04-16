// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package supervisor

import (
	context "context"
	proxy "github.com/juxuny/supervisor/proxy"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SupervisorClient is the client API for Supervisor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupervisorClient interface {
	ProxyStatus(ctx context.Context, in *ProxyStatusReq, opts ...grpc.CallOption) (*proxy.StatusResp, error)
	Apply(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*StopResp, error)
}

type supervisorClient struct {
	cc grpc.ClientConnInterface
}

func NewSupervisorClient(cc grpc.ClientConnInterface) SupervisorClient {
	return &supervisorClient{cc}
}

func (c *supervisorClient) ProxyStatus(ctx context.Context, in *ProxyStatusReq, opts ...grpc.CallOption) (*proxy.StatusResp, error) {
	out := new(proxy.StatusResp)
	err := c.cc.Invoke(ctx, "/supervisor.Supervisor/ProxyStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Apply(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error) {
	out := new(ApplyResp)
	err := c.cc.Invoke(ctx, "/supervisor.Supervisor/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/supervisor.Supervisor/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*StopResp, error) {
	out := new(StopResp)
	err := c.cc.Invoke(ctx, "/supervisor.Supervisor/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupervisorServer is the server API for Supervisor service.
// All implementations must embed UnimplementedSupervisorServer
// for forward compatibility
type SupervisorServer interface {
	ProxyStatus(context.Context, *ProxyStatusReq) (*proxy.StatusResp, error)
	Apply(context.Context, *ApplyReq) (*ApplyResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
	Stop(context.Context, *StopReq) (*StopResp, error)
	mustEmbedUnimplementedSupervisorServer()
}

// UnimplementedSupervisorServer must be embedded to have forward compatible implementations.
type UnimplementedSupervisorServer struct {
}

func (UnimplementedSupervisorServer) ProxyStatus(context.Context, *ProxyStatusReq) (*proxy.StatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyStatus not implemented")
}
func (UnimplementedSupervisorServer) Apply(context.Context, *ApplyReq) (*ApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedSupervisorServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSupervisorServer) Stop(context.Context, *StopReq) (*StopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSupervisorServer) mustEmbedUnimplementedSupervisorServer() {}

// UnsafeSupervisorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupervisorServer will
// result in compilation errors.
type UnsafeSupervisorServer interface {
	mustEmbedUnimplementedSupervisorServer()
}

func RegisterSupervisorServer(s grpc.ServiceRegistrar, srv SupervisorServer) {
	s.RegisterService(&Supervisor_ServiceDesc, srv)
}

func _Supervisor_ProxyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ProxyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Supervisor/ProxyStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ProxyStatus(ctx, req.(*ProxyStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Supervisor/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Apply(ctx, req.(*ApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Supervisor/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Supervisor/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Stop(ctx, req.(*StopReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Supervisor_ServiceDesc is the grpc.ServiceDesc for Supervisor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Supervisor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.Supervisor",
	HandlerType: (*SupervisorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProxyStatus",
			Handler:    _Supervisor_ProxyStatus_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Supervisor_Apply_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Supervisor_Get_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Supervisor_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor.proto",
}
