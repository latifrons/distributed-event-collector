// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: protos/report.proto

package dec

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DecService_Report_FullMethodName             = "/dec.DecService/Report"
	DecService_GetEventFlow_FullMethodName       = "/dec.DecService/GetEventFlow"
	DecService_GetEventStatistics_FullMethodName = "/dec.DecService/GetEventStatistics"
)

// DecServiceClient is the client API for DecService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecServiceClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	GetEventFlow(ctx context.Context, in *GetEventFlowRequest, opts ...grpc.CallOption) (*GetEventFlowResponse, error)
	GetEventStatistics(ctx context.Context, in *GetEventStatisticsRequest, opts ...grpc.CallOption) (*GetEventStatisticsResponse, error)
}

type decServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDecServiceClient(cc grpc.ClientConnInterface) DecServiceClient {
	return &decServiceClient{cc}
}

func (c *decServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, DecService_Report_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decServiceClient) GetEventFlow(ctx context.Context, in *GetEventFlowRequest, opts ...grpc.CallOption) (*GetEventFlowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventFlowResponse)
	err := c.cc.Invoke(ctx, DecService_GetEventFlow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decServiceClient) GetEventStatistics(ctx context.Context, in *GetEventStatisticsRequest, opts ...grpc.CallOption) (*GetEventStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventStatisticsResponse)
	err := c.cc.Invoke(ctx, DecService_GetEventStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecServiceServer is the server API for DecService service.
// All implementations must embed UnimplementedDecServiceServer
// for forward compatibility
type DecServiceServer interface {
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	GetEventFlow(context.Context, *GetEventFlowRequest) (*GetEventFlowResponse, error)
	GetEventStatistics(context.Context, *GetEventStatisticsRequest) (*GetEventStatisticsResponse, error)
	mustEmbedUnimplementedDecServiceServer()
}

// UnimplementedDecServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDecServiceServer struct {
}

func (UnimplementedDecServiceServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedDecServiceServer) GetEventFlow(context.Context, *GetEventFlowRequest) (*GetEventFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventFlow not implemented")
}
func (UnimplementedDecServiceServer) GetEventStatistics(context.Context, *GetEventStatisticsRequest) (*GetEventStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStatistics not implemented")
}
func (UnimplementedDecServiceServer) mustEmbedUnimplementedDecServiceServer() {}

// UnsafeDecServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecServiceServer will
// result in compilation errors.
type UnsafeDecServiceServer interface {
	mustEmbedUnimplementedDecServiceServer()
}

func RegisterDecServiceServer(s grpc.ServiceRegistrar, srv DecServiceServer) {
	s.RegisterService(&DecService_ServiceDesc, srv)
}

func _DecService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecService_Report_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecService_GetEventFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecServiceServer).GetEventFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecService_GetEventFlow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecServiceServer).GetEventFlow(ctx, req.(*GetEventFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecService_GetEventStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecServiceServer).GetEventStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecService_GetEventStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecServiceServer).GetEventStatistics(ctx, req.(*GetEventStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DecService_ServiceDesc is the grpc.ServiceDesc for DecService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DecService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dec.DecService",
	HandlerType: (*DecServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _DecService_Report_Handler,
		},
		{
			MethodName: "GetEventFlow",
			Handler:    _DecService_GetEventFlow_Handler,
		},
		{
			MethodName: "GetEventStatistics",
			Handler:    _DecService_GetEventStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/report.proto",
}
