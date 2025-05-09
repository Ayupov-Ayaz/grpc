// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1/transactions.proto

package api

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
	TransactionService_Bet_FullMethodName = "/api.v1.TransactionService/Bet"
	TransactionService_Win_FullMethodName = "/api.v1.TransactionService/Win"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Bet(ctx context.Context, in *BetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	Win(ctx context.Context, in *WinRequest, opts ...grpc.CallOption) (*WinResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Bet(ctx context.Context, in *BetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, TransactionService_Bet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Win(ctx context.Context, in *WinRequest, opts ...grpc.CallOption) (*WinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WinResponse)
	err := c.cc.Invoke(ctx, TransactionService_Win_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility.
type TransactionServiceServer interface {
	Bet(context.Context, *BetRequest) (*BetResponse, error)
	Win(context.Context, *WinRequest) (*WinResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionServiceServer struct{}

func (UnimplementedTransactionServiceServer) Bet(context.Context, *BetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bet not implemented")
}
func (UnimplementedTransactionServiceServer) Win(context.Context, *WinRequest) (*WinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Win not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}
func (UnimplementedTransactionServiceServer) testEmbeddedByValue()                            {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_Bet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Bet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_Bet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Bet(ctx, req.(*BetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Win_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Win(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_Win_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Win(ctx, req.(*WinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bet",
			Handler:    _TransactionService_Bet_Handler,
		},
		{
			MethodName: "Win",
			Handler:    _TransactionService_Win_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/transactions.proto",
}
