// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: habitTracker.proto

package habit

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
	HabitTracker_CreateHabit_FullMethodName         = "/habbitTracker.HabitTracker/CreateHabit"
	HabitTracker_GetHabit_FullMethodName            = "/habbitTracker.HabitTracker/GetHabit"
	HabitTracker_UpdateHabit_FullMethodName         = "/habbitTracker.HabitTracker/UpdateHabit"
	HabitTracker_DeleteHabit_FullMethodName         = "/habbitTracker.HabitTracker/DeleteHabit"
	HabitTracker_GetUserHabits_FullMethodName       = "/habbitTracker.HabitTracker/GetUserHabits"
	HabitTracker_CreateHabitLog_FullMethodName      = "/habbitTracker.HabitTracker/CreateHabitLog"
	HabitTracker_GetHabitLogs_FullMethodName        = "/habbitTracker.HabitTracker/GetHabitLogs"
	HabitTracker_GetHabitSuggestions_FullMethodName = "/habbitTracker.HabitTracker/GetHabitSuggestions"
)

// HabitTrackerClient is the client API for HabitTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HabitTrackerClient interface {
	CreateHabit(ctx context.Context, in *Habit, opts ...grpc.CallOption) (*Status, error)
	GetHabit(ctx context.Context, in *HabitId, opts ...grpc.CallOption) (*Habit, error)
	UpdateHabit(ctx context.Context, in *Habit, opts ...grpc.CallOption) (*Status, error)
	DeleteHabit(ctx context.Context, in *HabitId, opts ...grpc.CallOption) (*Status, error)
	GetUserHabits(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserHabits, error)
	CreateHabitLog(ctx context.Context, in *HabitLog, opts ...grpc.CallOption) (*Status, error)
	GetHabitLogs(ctx context.Context, in *HabitId, opts ...grpc.CallOption) (*HabitLog, error)
	GetHabitSuggestions(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Habits, error)
}

type habitTrackerClient struct {
	cc grpc.ClientConnInterface
}

func NewHabitTrackerClient(cc grpc.ClientConnInterface) HabitTrackerClient {
	return &habitTrackerClient{cc}
}

func (c *habitTrackerClient) CreateHabit(ctx context.Context, in *Habit, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, HabitTracker_CreateHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) GetHabit(ctx context.Context, in *HabitId, opts ...grpc.CallOption) (*Habit, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Habit)
	err := c.cc.Invoke(ctx, HabitTracker_GetHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) UpdateHabit(ctx context.Context, in *Habit, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, HabitTracker_UpdateHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) DeleteHabit(ctx context.Context, in *HabitId, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, HabitTracker_DeleteHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) GetUserHabits(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserHabits, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserHabits)
	err := c.cc.Invoke(ctx, HabitTracker_GetUserHabits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) CreateHabitLog(ctx context.Context, in *HabitLog, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, HabitTracker_CreateHabitLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) GetHabitLogs(ctx context.Context, in *HabitId, opts ...grpc.CallOption) (*HabitLog, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HabitLog)
	err := c.cc.Invoke(ctx, HabitTracker_GetHabitLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerClient) GetHabitSuggestions(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Habits, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Habits)
	err := c.cc.Invoke(ctx, HabitTracker_GetHabitSuggestions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HabitTrackerServer is the server API for HabitTracker service.
// All implementations must embed UnimplementedHabitTrackerServer
// for forward compatibility
type HabitTrackerServer interface {
	CreateHabit(context.Context, *Habit) (*Status, error)
	GetHabit(context.Context, *HabitId) (*Habit, error)
	UpdateHabit(context.Context, *Habit) (*Status, error)
	DeleteHabit(context.Context, *HabitId) (*Status, error)
	GetUserHabits(context.Context, *UserId) (*UserHabits, error)
	CreateHabitLog(context.Context, *HabitLog) (*Status, error)
	GetHabitLogs(context.Context, *HabitId) (*HabitLog, error)
	GetHabitSuggestions(context.Context, *Req) (*Habits, error)
	mustEmbedUnimplementedHabitTrackerServer()
}

// UnimplementedHabitTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedHabitTrackerServer struct {
}

func (UnimplementedHabitTrackerServer) CreateHabit(context.Context, *Habit) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHabit not implemented")
}
func (UnimplementedHabitTrackerServer) GetHabit(context.Context, *HabitId) (*Habit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHabit not implemented")
}
func (UnimplementedHabitTrackerServer) UpdateHabit(context.Context, *Habit) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHabit not implemented")
}
func (UnimplementedHabitTrackerServer) DeleteHabit(context.Context, *HabitId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHabit not implemented")
}
func (UnimplementedHabitTrackerServer) GetUserHabits(context.Context, *UserId) (*UserHabits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserHabits not implemented")
}
func (UnimplementedHabitTrackerServer) CreateHabitLog(context.Context, *HabitLog) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHabitLog not implemented")
}
func (UnimplementedHabitTrackerServer) GetHabitLogs(context.Context, *HabitId) (*HabitLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHabitLogs not implemented")
}
func (UnimplementedHabitTrackerServer) GetHabitSuggestions(context.Context, *Req) (*Habits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHabitSuggestions not implemented")
}
func (UnimplementedHabitTrackerServer) mustEmbedUnimplementedHabitTrackerServer() {}

// UnsafeHabitTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HabitTrackerServer will
// result in compilation errors.
type UnsafeHabitTrackerServer interface {
	mustEmbedUnimplementedHabitTrackerServer()
}

func RegisterHabitTrackerServer(s grpc.ServiceRegistrar, srv HabitTrackerServer) {
	s.RegisterService(&HabitTracker_ServiceDesc, srv)
}

func _HabitTracker_CreateHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Habit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).CreateHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_CreateHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).CreateHabit(ctx, req.(*Habit))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_GetHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HabitId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).GetHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_GetHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).GetHabit(ctx, req.(*HabitId))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_UpdateHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Habit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).UpdateHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_UpdateHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).UpdateHabit(ctx, req.(*Habit))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_DeleteHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HabitId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).DeleteHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_DeleteHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).DeleteHabit(ctx, req.(*HabitId))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_GetUserHabits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).GetUserHabits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_GetUserHabits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).GetUserHabits(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_CreateHabitLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HabitLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).CreateHabitLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_CreateHabitLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).CreateHabitLog(ctx, req.(*HabitLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_GetHabitLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HabitId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).GetHabitLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_GetHabitLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).GetHabitLogs(ctx, req.(*HabitId))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTracker_GetHabitSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServer).GetHabitSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTracker_GetHabitSuggestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServer).GetHabitSuggestions(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// HabitTracker_ServiceDesc is the grpc.ServiceDesc for HabitTracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HabitTracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "habbitTracker.HabitTracker",
	HandlerType: (*HabitTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHabit",
			Handler:    _HabitTracker_CreateHabit_Handler,
		},
		{
			MethodName: "GetHabit",
			Handler:    _HabitTracker_GetHabit_Handler,
		},
		{
			MethodName: "UpdateHabit",
			Handler:    _HabitTracker_UpdateHabit_Handler,
		},
		{
			MethodName: "DeleteHabit",
			Handler:    _HabitTracker_DeleteHabit_Handler,
		},
		{
			MethodName: "GetUserHabits",
			Handler:    _HabitTracker_GetUserHabits_Handler,
		},
		{
			MethodName: "CreateHabitLog",
			Handler:    _HabitTracker_CreateHabitLog_Handler,
		},
		{
			MethodName: "GetHabitLogs",
			Handler:    _HabitTracker_GetHabitLogs_Handler,
		},
		{
			MethodName: "GetHabitSuggestions",
			Handler:    _HabitTracker_GetHabitSuggestions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "habitTracker.proto",
}