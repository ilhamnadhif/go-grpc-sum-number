// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: calculator.proto

package calculator

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	DevidedBy(ctx context.Context, in *DevidedByRequest, opts ...grpc.CallOption) (*DevidedByResponse, error)
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciClient, error)
	AVG(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AVGClient, error)
	FilterSameNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FilterSameNumberClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) DevidedBy(ctx context.Context, in *DevidedByRequest, opts ...grpc.CallOption) (*DevidedByResponse, error) {
	out := new(DevidedByResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/DevidedBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/Fibonacci", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_FibonacciClient interface {
	Recv() (*FibonacciResponse, error)
	grpc.ClientStream
}

type calculatorServiceFibonacciClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFibonacciClient) Recv() (*FibonacciResponse, error) {
	m := new(FibonacciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) AVG(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AVGClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/AVG", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAVGClient{stream}
	return x, nil
}

type CalculatorService_AVGClient interface {
	Send(*AVGRequest) error
	CloseAndRecv() (*AVGResponse, error)
	grpc.ClientStream
}

type calculatorServiceAVGClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAVGClient) Send(m *AVGRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAVGClient) CloseAndRecv() (*AVGResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AVGResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FilterSameNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FilterSameNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], "/calculator.CalculatorService/FilterSameNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFilterSameNumberClient{stream}
	return x, nil
}

type CalculatorService_FilterSameNumberClient interface {
	Send(*FilterSumNumberRequest) error
	Recv() (*FilterSumNumberResponse, error)
	grpc.ClientStream
}

type calculatorServiceFilterSameNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFilterSameNumberClient) Send(m *FilterSumNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFilterSameNumberClient) Recv() (*FilterSumNumberResponse, error) {
	m := new(FilterSumNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	DevidedBy(context.Context, *DevidedByRequest) (*DevidedByResponse, error)
	Fibonacci(*FibonacciRequest, CalculatorService_FibonacciServer) error
	AVG(CalculatorService_AVGServer) error
	FilterSameNumber(CalculatorService_FilterSameNumberServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) DevidedBy(context.Context, *DevidedByRequest) (*DevidedByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DevidedBy not implemented")
}
func (UnimplementedCalculatorServiceServer) Fibonacci(*FibonacciRequest, CalculatorService_FibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (UnimplementedCalculatorServiceServer) AVG(CalculatorService_AVGServer) error {
	return status.Errorf(codes.Unimplemented, "method AVG not implemented")
}
func (UnimplementedCalculatorServiceServer) FilterSameNumber(CalculatorService_FilterSameNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method FilterSameNumber not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_DevidedBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevidedByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).DevidedBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/DevidedBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).DevidedBy(ctx, req.(*DevidedByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Fibonacci(m, &calculatorServiceFibonacciServer{stream})
}

type CalculatorService_FibonacciServer interface {
	Send(*FibonacciResponse) error
	grpc.ServerStream
}

type calculatorServiceFibonacciServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFibonacciServer) Send(m *FibonacciResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_AVG_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).AVG(&calculatorServiceAVGServer{stream})
}

type CalculatorService_AVGServer interface {
	SendAndClose(*AVGResponse) error
	Recv() (*AVGRequest, error)
	grpc.ServerStream
}

type calculatorServiceAVGServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAVGServer) SendAndClose(m *AVGResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAVGServer) Recv() (*AVGRequest, error) {
	m := new(AVGRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FilterSameNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FilterSameNumber(&calculatorServiceFilterSameNumberServer{stream})
}

type CalculatorService_FilterSameNumberServer interface {
	Send(*FilterSumNumberResponse) error
	Recv() (*FilterSumNumberRequest, error)
	grpc.ServerStream
}

type calculatorServiceFilterSameNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFilterSameNumberServer) Send(m *FilterSumNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFilterSameNumberServer) Recv() (*FilterSumNumberRequest, error) {
	m := new(FilterSumNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "DevidedBy",
			Handler:    _CalculatorService_DevidedBy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fibonacci",
			Handler:       _CalculatorService_Fibonacci_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AVG",
			Handler:       _CalculatorService_AVG_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FilterSameNumber",
			Handler:       _CalculatorService_FilterSameNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
