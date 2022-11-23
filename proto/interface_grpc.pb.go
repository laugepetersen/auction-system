// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/interface.proto

package proto

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

// AuctionServiceClient is the client API for AuctionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionServiceClient interface {
	HeartBeat(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error)
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Ack, error)
	AskResult(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ack, error)
	AskForLeadership(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*ElectorAnswer, error)
	GiveLeadership(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error)
}

type auctionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionServiceClient(cc grpc.ClientConnInterface) AuctionServiceClient {
	return &auctionServiceClient{cc}
}

func (c *auctionServiceClient) HeartBeat(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error) {
	out := new(Ping)
	err := c.cc.Invoke(ctx, "/proto.AuctionService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/proto.AuctionService/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) AskResult(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/proto.AuctionService/AskResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) AskForLeadership(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*ElectorAnswer, error) {
	out := new(ElectorAnswer)
	err := c.cc.Invoke(ctx, "/proto.AuctionService/AskForLeadership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) GiveLeadership(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error) {
	out := new(Ping)
	err := c.cc.Invoke(ctx, "/proto.AuctionService/GiveLeadership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServiceServer is the server API for AuctionService service.
// All implementations must embed UnimplementedAuctionServiceServer
// for forward compatibility
type AuctionServiceServer interface {
	HeartBeat(context.Context, *Ping) (*Ping, error)
	Bid(context.Context, *BidMessage) (*Ack, error)
	AskResult(context.Context, *Ping) (*Ack, error)
	AskForLeadership(context.Context, *Ping) (*ElectorAnswer, error)
	GiveLeadership(context.Context, *Ping) (*Ping, error)
	mustEmbedUnimplementedAuctionServiceServer()
}

// UnimplementedAuctionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServiceServer struct {
}

func (UnimplementedAuctionServiceServer) HeartBeat(context.Context, *Ping) (*Ping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedAuctionServiceServer) Bid(context.Context, *BidMessage) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServiceServer) AskResult(context.Context, *Ping) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskResult not implemented")
}
func (UnimplementedAuctionServiceServer) AskForLeadership(context.Context, *Ping) (*ElectorAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskForLeadership not implemented")
}
func (UnimplementedAuctionServiceServer) GiveLeadership(context.Context, *Ping) (*Ping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveLeadership not implemented")
}
func (UnimplementedAuctionServiceServer) mustEmbedUnimplementedAuctionServiceServer() {}

// UnsafeAuctionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServiceServer will
// result in compilation errors.
type UnsafeAuctionServiceServer interface {
	mustEmbedUnimplementedAuctionServiceServer()
}

func RegisterAuctionServiceServer(s grpc.ServiceRegistrar, srv AuctionServiceServer) {
	s.RegisterService(&AuctionService_ServiceDesc, srv)
}

func _AuctionService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuctionService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).HeartBeat(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuctionService/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_AskResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).AskResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuctionService/AskResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).AskResult(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_AskForLeadership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).AskForLeadership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuctionService/AskForLeadership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).AskForLeadership(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_GiveLeadership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).GiveLeadership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuctionService/GiveLeadership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).GiveLeadership(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionService_ServiceDesc is the grpc.ServiceDesc for AuctionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuctionService",
	HandlerType: (*AuctionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _AuctionService_HeartBeat_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _AuctionService_Bid_Handler,
		},
		{
			MethodName: "AskResult",
			Handler:    _AuctionService_AskResult_Handler,
		},
		{
			MethodName: "AskForLeadership",
			Handler:    _AuctionService_AskForLeadership_Handler,
		},
		{
			MethodName: "GiveLeadership",
			Handler:    _AuctionService_GiveLeadership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/interface.proto",
}
