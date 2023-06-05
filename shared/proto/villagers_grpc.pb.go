// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: villagers.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VillagersService_FindAll_FullMethodName = "/proto.VillagersService/FindAll"
)

// VillagersServiceClient is the client API for VillagersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VillagersServiceClient interface {
	FindAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FindAllResponse, error)
}

type villagersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVillagersServiceClient(cc grpc.ClientConnInterface) VillagersServiceClient {
	return &villagersServiceClient{cc}
}

func (c *villagersServiceClient) FindAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, VillagersService_FindAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VillagersServiceServer is the server API for VillagersService service.
// All implementations must embed UnimplementedVillagersServiceServer
// for forward compatibility
type VillagersServiceServer interface {
	FindAll(context.Context, *emptypb.Empty) (*FindAllResponse, error)
	mustEmbedUnimplementedVillagersServiceServer()
}

// UnimplementedVillagersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVillagersServiceServer struct {
}

func (UnimplementedVillagersServiceServer) FindAll(context.Context, *emptypb.Empty) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedVillagersServiceServer) mustEmbedUnimplementedVillagersServiceServer() {}

// UnsafeVillagersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VillagersServiceServer will
// result in compilation errors.
type UnsafeVillagersServiceServer interface {
	mustEmbedUnimplementedVillagersServiceServer()
}

func RegisterVillagersServiceServer(s grpc.ServiceRegistrar, srv VillagersServiceServer) {
	s.RegisterService(&VillagersService_ServiceDesc, srv)
}

func _VillagersService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VillagersServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VillagersService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VillagersServiceServer).FindAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VillagersService_ServiceDesc is the grpc.ServiceDesc for VillagersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VillagersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VillagersService",
	HandlerType: (*VillagersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _VillagersService_FindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "villagers.proto",
}
