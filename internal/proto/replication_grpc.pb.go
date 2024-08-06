// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.2
// source: replication.proto

package proto

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
	Replication_EnableVolumeReplication_FullMethodName  = "/proto.Replication/EnableVolumeReplication"
	Replication_DisableVolumeReplication_FullMethodName = "/proto.Replication/DisableVolumeReplication"
	Replication_PromoteVolume_FullMethodName            = "/proto.Replication/PromoteVolume"
	Replication_DemoteVolume_FullMethodName             = "/proto.Replication/DemoteVolume"
	Replication_ResyncVolume_FullMethodName             = "/proto.Replication/ResyncVolume"
	Replication_GetVolumeReplicationInfo_FullMethodName = "/proto.Replication/GetVolumeReplicationInfo"
)

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Replication holds the RPC method for allowing the communication between
// the CSIAddons controller and the sidecar for replication operation.
type ReplicationClient interface {
	// EnableVolumeReplication RPC call to enable the volume replication.
	EnableVolumeReplication(ctx context.Context, in *EnableVolumeReplicationRequest, opts ...grpc.CallOption) (*EnableVolumeReplicationResponse, error)
	// DisableVolumeReplication RPC call to disable the volume replication.
	DisableVolumeReplication(ctx context.Context, in *DisableVolumeReplicationRequest, opts ...grpc.CallOption) (*DisableVolumeReplicationResponse, error)
	// PromoteVolume RPC call to promote the volume.
	PromoteVolume(ctx context.Context, in *PromoteVolumeRequest, opts ...grpc.CallOption) (*PromoteVolumeResponse, error)
	// DemoteVolume RPC call to demote the volume.
	DemoteVolume(ctx context.Context, in *DemoteVolumeRequest, opts ...grpc.CallOption) (*DemoteVolumeResponse, error)
	// ResyncVolume RPC call to resync the volume.
	ResyncVolume(ctx context.Context, in *ResyncVolumeRequest, opts ...grpc.CallOption) (*ResyncVolumeResponse, error)
	// GetVolumeReplicationInfo RPC call to get the volume replication info.
	GetVolumeReplicationInfo(ctx context.Context, in *GetVolumeReplicationInfoRequest, opts ...grpc.CallOption) (*GetVolumeReplicationInfoResponse, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) EnableVolumeReplication(ctx context.Context, in *EnableVolumeReplicationRequest, opts ...grpc.CallOption) (*EnableVolumeReplicationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnableVolumeReplicationResponse)
	err := c.cc.Invoke(ctx, Replication_EnableVolumeReplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) DisableVolumeReplication(ctx context.Context, in *DisableVolumeReplicationRequest, opts ...grpc.CallOption) (*DisableVolumeReplicationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisableVolumeReplicationResponse)
	err := c.cc.Invoke(ctx, Replication_DisableVolumeReplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) PromoteVolume(ctx context.Context, in *PromoteVolumeRequest, opts ...grpc.CallOption) (*PromoteVolumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromoteVolumeResponse)
	err := c.cc.Invoke(ctx, Replication_PromoteVolume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) DemoteVolume(ctx context.Context, in *DemoteVolumeRequest, opts ...grpc.CallOption) (*DemoteVolumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DemoteVolumeResponse)
	err := c.cc.Invoke(ctx, Replication_DemoteVolume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) ResyncVolume(ctx context.Context, in *ResyncVolumeRequest, opts ...grpc.CallOption) (*ResyncVolumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResyncVolumeResponse)
	err := c.cc.Invoke(ctx, Replication_ResyncVolume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationClient) GetVolumeReplicationInfo(ctx context.Context, in *GetVolumeReplicationInfoRequest, opts ...grpc.CallOption) (*GetVolumeReplicationInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVolumeReplicationInfoResponse)
	err := c.cc.Invoke(ctx, Replication_GetVolumeReplicationInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServer is the server API for Replication service.
// All implementations must embed UnimplementedReplicationServer
// for forward compatibility.
//
// Replication holds the RPC method for allowing the communication between
// the CSIAddons controller and the sidecar for replication operation.
type ReplicationServer interface {
	// EnableVolumeReplication RPC call to enable the volume replication.
	EnableVolumeReplication(context.Context, *EnableVolumeReplicationRequest) (*EnableVolumeReplicationResponse, error)
	// DisableVolumeReplication RPC call to disable the volume replication.
	DisableVolumeReplication(context.Context, *DisableVolumeReplicationRequest) (*DisableVolumeReplicationResponse, error)
	// PromoteVolume RPC call to promote the volume.
	PromoteVolume(context.Context, *PromoteVolumeRequest) (*PromoteVolumeResponse, error)
	// DemoteVolume RPC call to demote the volume.
	DemoteVolume(context.Context, *DemoteVolumeRequest) (*DemoteVolumeResponse, error)
	// ResyncVolume RPC call to resync the volume.
	ResyncVolume(context.Context, *ResyncVolumeRequest) (*ResyncVolumeResponse, error)
	// GetVolumeReplicationInfo RPC call to get the volume replication info.
	GetVolumeReplicationInfo(context.Context, *GetVolumeReplicationInfoRequest) (*GetVolumeReplicationInfoResponse, error)
	mustEmbedUnimplementedReplicationServer()
}

// UnimplementedReplicationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReplicationServer struct{}

func (UnimplementedReplicationServer) EnableVolumeReplication(context.Context, *EnableVolumeReplicationRequest) (*EnableVolumeReplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableVolumeReplication not implemented")
}
func (UnimplementedReplicationServer) DisableVolumeReplication(context.Context, *DisableVolumeReplicationRequest) (*DisableVolumeReplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableVolumeReplication not implemented")
}
func (UnimplementedReplicationServer) PromoteVolume(context.Context, *PromoteVolumeRequest) (*PromoteVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteVolume not implemented")
}
func (UnimplementedReplicationServer) DemoteVolume(context.Context, *DemoteVolumeRequest) (*DemoteVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoteVolume not implemented")
}
func (UnimplementedReplicationServer) ResyncVolume(context.Context, *ResyncVolumeRequest) (*ResyncVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResyncVolume not implemented")
}
func (UnimplementedReplicationServer) GetVolumeReplicationInfo(context.Context, *GetVolumeReplicationInfoRequest) (*GetVolumeReplicationInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolumeReplicationInfo not implemented")
}
func (UnimplementedReplicationServer) mustEmbedUnimplementedReplicationServer() {}
func (UnimplementedReplicationServer) testEmbeddedByValue()                     {}

// UnsafeReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServer will
// result in compilation errors.
type UnsafeReplicationServer interface {
	mustEmbedUnimplementedReplicationServer()
}

func RegisterReplicationServer(s grpc.ServiceRegistrar, srv ReplicationServer) {
	// If the following call pancis, it indicates UnimplementedReplicationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Replication_ServiceDesc, srv)
}

func _Replication_EnableVolumeReplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableVolumeReplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).EnableVolumeReplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_EnableVolumeReplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).EnableVolumeReplication(ctx, req.(*EnableVolumeReplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_DisableVolumeReplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableVolumeReplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).DisableVolumeReplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_DisableVolumeReplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).DisableVolumeReplication(ctx, req.(*DisableVolumeReplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_PromoteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).PromoteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_PromoteVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).PromoteVolume(ctx, req.(*PromoteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_DemoteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).DemoteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_DemoteVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).DemoteVolume(ctx, req.(*DemoteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_ResyncVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResyncVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).ResyncVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_ResyncVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).ResyncVolume(ctx, req.(*ResyncVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replication_GetVolumeReplicationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeReplicationInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).GetVolumeReplicationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Replication_GetVolumeReplicationInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).GetVolumeReplicationInfo(ctx, req.(*GetVolumeReplicationInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Replication_ServiceDesc is the grpc.ServiceDesc for Replication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnableVolumeReplication",
			Handler:    _Replication_EnableVolumeReplication_Handler,
		},
		{
			MethodName: "DisableVolumeReplication",
			Handler:    _Replication_DisableVolumeReplication_Handler,
		},
		{
			MethodName: "PromoteVolume",
			Handler:    _Replication_PromoteVolume_Handler,
		},
		{
			MethodName: "DemoteVolume",
			Handler:    _Replication_DemoteVolume_Handler,
		},
		{
			MethodName: "ResyncVolume",
			Handler:    _Replication_ResyncVolume_Handler,
		},
		{
			MethodName: "GetVolumeReplicationInfo",
			Handler:    _Replication_GetVolumeReplicationInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replication.proto",
}
