// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: perf.proto

package gopb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CedarPerformanceMetricsClient is the client API for CedarPerformanceMetrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CedarPerformanceMetricsClient interface {
	CreateMetricSeries(ctx context.Context, in *ResultData, opts ...grpc.CallOption) (*MetricsResponse, error)
	AttachArtifacts(ctx context.Context, in *ArtifactData, opts ...grpc.CallOption) (*MetricsResponse, error)
	AttachRollups(ctx context.Context, in *RollupData, opts ...grpc.CallOption) (*MetricsResponse, error)
	SendMetrics(ctx context.Context, opts ...grpc.CallOption) (CedarPerformanceMetrics_SendMetricsClient, error)
	CloseMetrics(ctx context.Context, in *MetricsSeriesEnd, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type cedarPerformanceMetricsClient struct {
	cc grpc.ClientConnInterface
}

func NewCedarPerformanceMetricsClient(cc grpc.ClientConnInterface) CedarPerformanceMetricsClient {
	return &cedarPerformanceMetricsClient{cc}
}

func (c *cedarPerformanceMetricsClient) CreateMetricSeries(ctx context.Context, in *ResultData, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarPerformanceMetrics/CreateMetricSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarPerformanceMetricsClient) AttachArtifacts(ctx context.Context, in *ArtifactData, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarPerformanceMetrics/AttachArtifacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarPerformanceMetricsClient) AttachRollups(ctx context.Context, in *RollupData, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarPerformanceMetrics/AttachRollups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarPerformanceMetricsClient) SendMetrics(ctx context.Context, opts ...grpc.CallOption) (CedarPerformanceMetrics_SendMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CedarPerformanceMetrics_ServiceDesc.Streams[0], "/cedar.CedarPerformanceMetrics/SendMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &cedarPerformanceMetricsSendMetricsClient{stream}
	return x, nil
}

type CedarPerformanceMetrics_SendMetricsClient interface {
	Send(*MetricsEvent) error
	CloseAndRecv() (*SendResponse, error)
	grpc.ClientStream
}

type cedarPerformanceMetricsSendMetricsClient struct {
	grpc.ClientStream
}

func (x *cedarPerformanceMetricsSendMetricsClient) Send(m *MetricsEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cedarPerformanceMetricsSendMetricsClient) CloseAndRecv() (*SendResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cedarPerformanceMetricsClient) CloseMetrics(ctx context.Context, in *MetricsSeriesEnd, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarPerformanceMetrics/CloseMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CedarPerformanceMetricsServer is the server API for CedarPerformanceMetrics service.
// All implementations must embed UnimplementedCedarPerformanceMetricsServer
// for forward compatibility
type CedarPerformanceMetricsServer interface {
	CreateMetricSeries(context.Context, *ResultData) (*MetricsResponse, error)
	AttachArtifacts(context.Context, *ArtifactData) (*MetricsResponse, error)
	AttachRollups(context.Context, *RollupData) (*MetricsResponse, error)
	SendMetrics(CedarPerformanceMetrics_SendMetricsServer) error
	CloseMetrics(context.Context, *MetricsSeriesEnd) (*MetricsResponse, error)
	mustEmbedUnimplementedCedarPerformanceMetricsServer()
}

// UnimplementedCedarPerformanceMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedCedarPerformanceMetricsServer struct {
}

func (UnimplementedCedarPerformanceMetricsServer) CreateMetricSeries(context.Context, *ResultData) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMetricSeries not implemented")
}
func (UnimplementedCedarPerformanceMetricsServer) AttachArtifacts(context.Context, *ArtifactData) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachArtifacts not implemented")
}
func (UnimplementedCedarPerformanceMetricsServer) AttachRollups(context.Context, *RollupData) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachRollups not implemented")
}
func (UnimplementedCedarPerformanceMetricsServer) SendMetrics(CedarPerformanceMetrics_SendMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMetrics not implemented")
}
func (UnimplementedCedarPerformanceMetricsServer) CloseMetrics(context.Context, *MetricsSeriesEnd) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseMetrics not implemented")
}
func (UnimplementedCedarPerformanceMetricsServer) mustEmbedUnimplementedCedarPerformanceMetricsServer() {
}

// UnsafeCedarPerformanceMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CedarPerformanceMetricsServer will
// result in compilation errors.
type UnsafeCedarPerformanceMetricsServer interface {
	mustEmbedUnimplementedCedarPerformanceMetricsServer()
}

func RegisterCedarPerformanceMetricsServer(s grpc.ServiceRegistrar, srv CedarPerformanceMetricsServer) {
	s.RegisterService(&CedarPerformanceMetrics_ServiceDesc, srv)
}

func _CedarPerformanceMetrics_CreateMetricSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarPerformanceMetricsServer).CreateMetricSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarPerformanceMetrics/CreateMetricSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarPerformanceMetricsServer).CreateMetricSeries(ctx, req.(*ResultData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarPerformanceMetrics_AttachArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarPerformanceMetricsServer).AttachArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarPerformanceMetrics/AttachArtifacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarPerformanceMetricsServer).AttachArtifacts(ctx, req.(*ArtifactData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarPerformanceMetrics_AttachRollups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollupData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarPerformanceMetricsServer).AttachRollups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarPerformanceMetrics/AttachRollups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarPerformanceMetricsServer).AttachRollups(ctx, req.(*RollupData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarPerformanceMetrics_SendMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CedarPerformanceMetricsServer).SendMetrics(&cedarPerformanceMetricsSendMetricsServer{stream})
}

type CedarPerformanceMetrics_SendMetricsServer interface {
	SendAndClose(*SendResponse) error
	Recv() (*MetricsEvent, error)
	grpc.ServerStream
}

type cedarPerformanceMetricsSendMetricsServer struct {
	grpc.ServerStream
}

func (x *cedarPerformanceMetricsSendMetricsServer) SendAndClose(m *SendResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cedarPerformanceMetricsSendMetricsServer) Recv() (*MetricsEvent, error) {
	m := new(MetricsEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CedarPerformanceMetrics_CloseMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsSeriesEnd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarPerformanceMetricsServer).CloseMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarPerformanceMetrics/CloseMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarPerformanceMetricsServer).CloseMetrics(ctx, req.(*MetricsSeriesEnd))
	}
	return interceptor(ctx, in, info, handler)
}

// CedarPerformanceMetrics_ServiceDesc is the grpc.ServiceDesc for CedarPerformanceMetrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CedarPerformanceMetrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cedar.CedarPerformanceMetrics",
	HandlerType: (*CedarPerformanceMetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMetricSeries",
			Handler:    _CedarPerformanceMetrics_CreateMetricSeries_Handler,
		},
		{
			MethodName: "AttachArtifacts",
			Handler:    _CedarPerformanceMetrics_AttachArtifacts_Handler,
		},
		{
			MethodName: "AttachRollups",
			Handler:    _CedarPerformanceMetrics_AttachRollups_Handler,
		},
		{
			MethodName: "CloseMetrics",
			Handler:    _CedarPerformanceMetrics_CloseMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMetrics",
			Handler:       _CedarPerformanceMetrics_SendMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "perf.proto",
}
