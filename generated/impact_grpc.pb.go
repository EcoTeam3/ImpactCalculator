// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: impact.proto

package generated

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

// ImpactClient is the client API for Impact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImpactClient interface {
	CreateFootprint(ctx context.Context, in *CarbonFootprint, opts ...grpc.CallOption) (*Status, error)
	GetUserImpact(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*CarbonFootprints, error)
	GetGroupImpact(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*CarbonFootprint, error)
	GetLeaderBoardUsers(ctx context.Context, in *LeaderBoard, opts ...grpc.CallOption) (*LeaderBoardUsers, error)
	GetLeaderBoardGroups(ctx context.Context, in *LeaderBoard, opts ...grpc.CallOption) (*LeaderBoardGroups, error)
	CreateDonation(ctx context.Context, in *Donation, opts ...grpc.CallOption) (*Status, error)
	GetDonations(ctx context.Context, in *DonationCause, opts ...grpc.CallOption) (*Donations, error)
}

type impactClient struct {
	cc grpc.ClientConnInterface
}

func NewImpactClient(cc grpc.ClientConnInterface) ImpactClient {
	return &impactClient{cc}
}

func (c *impactClient) CreateFootprint(ctx context.Context, in *CarbonFootprint, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/impact.Impact/CreateFootprint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impactClient) GetUserImpact(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*CarbonFootprints, error) {
	out := new(CarbonFootprints)
	err := c.cc.Invoke(ctx, "/impact.Impact/GetUserImpact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impactClient) GetGroupImpact(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*CarbonFootprint, error) {
	out := new(CarbonFootprint)
	err := c.cc.Invoke(ctx, "/impact.Impact/GetGroupImpact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impactClient) GetLeaderBoardUsers(ctx context.Context, in *LeaderBoard, opts ...grpc.CallOption) (*LeaderBoardUsers, error) {
	out := new(LeaderBoardUsers)
	err := c.cc.Invoke(ctx, "/impact.Impact/GetLeaderBoardUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impactClient) GetLeaderBoardGroups(ctx context.Context, in *LeaderBoard, opts ...grpc.CallOption) (*LeaderBoardGroups, error) {
	out := new(LeaderBoardGroups)
	err := c.cc.Invoke(ctx, "/impact.Impact/GetLeaderBoardGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impactClient) CreateDonation(ctx context.Context, in *Donation, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/impact.Impact/CreateDonation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impactClient) GetDonations(ctx context.Context, in *DonationCause, opts ...grpc.CallOption) (*Donations, error) {
	out := new(Donations)
	err := c.cc.Invoke(ctx, "/impact.Impact/GetDonations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImpactServer is the server API for Impact service.
// All implementations must embed UnimplementedImpactServer
// for forward compatibility
type ImpactServer interface {
	CreateFootprint(context.Context, *CarbonFootprint) (*Status, error)
	GetUserImpact(context.Context, *UserId) (*CarbonFootprints, error)
	GetGroupImpact(context.Context, *GroupId) (*CarbonFootprint, error)
	GetLeaderBoardUsers(context.Context, *LeaderBoard) (*LeaderBoardUsers, error)
	GetLeaderBoardGroups(context.Context, *LeaderBoard) (*LeaderBoardGroups, error)
	CreateDonation(context.Context, *Donation) (*Status, error)
	GetDonations(context.Context, *DonationCause) (*Donations, error)
	mustEmbedUnimplementedImpactServer()
}

// UnimplementedImpactServer must be embedded to have forward compatible implementations.
type UnimplementedImpactServer struct {
}

func (UnimplementedImpactServer) CreateFootprint(context.Context, *CarbonFootprint) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFootprint not implemented")
}
func (UnimplementedImpactServer) GetUserImpact(context.Context, *UserId) (*CarbonFootprints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserImpact not implemented")
}
func (UnimplementedImpactServer) GetGroupImpact(context.Context, *GroupId) (*CarbonFootprint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupImpact not implemented")
}
func (UnimplementedImpactServer) GetLeaderBoardUsers(context.Context, *LeaderBoard) (*LeaderBoardUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderBoardUsers not implemented")
}
func (UnimplementedImpactServer) GetLeaderBoardGroups(context.Context, *LeaderBoard) (*LeaderBoardGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderBoardGroups not implemented")
}
func (UnimplementedImpactServer) CreateDonation(context.Context, *Donation) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDonation not implemented")
}
func (UnimplementedImpactServer) GetDonations(context.Context, *DonationCause) (*Donations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonations not implemented")
}
func (UnimplementedImpactServer) mustEmbedUnimplementedImpactServer() {}

// UnsafeImpactServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImpactServer will
// result in compilation errors.
type UnsafeImpactServer interface {
	mustEmbedUnimplementedImpactServer()
}

func RegisterImpactServer(s grpc.ServiceRegistrar, srv ImpactServer) {
	s.RegisterService(&Impact_ServiceDesc, srv)
}

func _Impact_CreateFootprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarbonFootprint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).CreateFootprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/CreateFootprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).CreateFootprint(ctx, req.(*CarbonFootprint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Impact_GetUserImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).GetUserImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/GetUserImpact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).GetUserImpact(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Impact_GetGroupImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).GetGroupImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/GetGroupImpact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).GetGroupImpact(ctx, req.(*GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Impact_GetLeaderBoardUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderBoard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).GetLeaderBoardUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/GetLeaderBoardUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).GetLeaderBoardUsers(ctx, req.(*LeaderBoard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Impact_GetLeaderBoardGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderBoard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).GetLeaderBoardGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/GetLeaderBoardGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).GetLeaderBoardGroups(ctx, req.(*LeaderBoard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Impact_CreateDonation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Donation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).CreateDonation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/CreateDonation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).CreateDonation(ctx, req.(*Donation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Impact_GetDonations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DonationCause)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpactServer).GetDonations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/impact.Impact/GetDonations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpactServer).GetDonations(ctx, req.(*DonationCause))
	}
	return interceptor(ctx, in, info, handler)
}

// Impact_ServiceDesc is the grpc.ServiceDesc for Impact service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Impact_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "impact.Impact",
	HandlerType: (*ImpactServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFootprint",
			Handler:    _Impact_CreateFootprint_Handler,
		},
		{
			MethodName: "GetUserImpact",
			Handler:    _Impact_GetUserImpact_Handler,
		},
		{
			MethodName: "GetGroupImpact",
			Handler:    _Impact_GetGroupImpact_Handler,
		},
		{
			MethodName: "GetLeaderBoardUsers",
			Handler:    _Impact_GetLeaderBoardUsers_Handler,
		},
		{
			MethodName: "GetLeaderBoardGroups",
			Handler:    _Impact_GetLeaderBoardGroups_Handler,
		},
		{
			MethodName: "CreateDonation",
			Handler:    _Impact_CreateDonation_Handler,
		},
		{
			MethodName: "GetDonations",
			Handler:    _Impact_GetDonations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "impact.proto",
}
