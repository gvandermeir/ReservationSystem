// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: models.proto

package protos

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

const (
	ReservationService_SetAvailability_FullMethodName    = "/protos.ReservationService/SetAvailability"
	ReservationService_GetAppointments_FullMethodName    = "/protos.ReservationService/GetAppointments"
	ReservationService_ReserveAppointment_FullMethodName = "/protos.ReservationService/ReserveAppointment"
	ReservationService_ConfrimAppointment_FullMethodName = "/protos.ReservationService/ConfrimAppointment"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	SetAvailability(ctx context.Context, in *SetAvailabilityRequest, opts ...grpc.CallOption) (*SetAvailabilityResponse, error)
	GetAppointments(ctx context.Context, in *GetAppointmentsRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error)
	ReserveAppointment(ctx context.Context, in *ReserveAppointmentRequest, opts ...grpc.CallOption) (*ReserveAppointmentResponse, error)
	ConfrimAppointment(ctx context.Context, in *ConfrimAppointmentRequest, opts ...grpc.CallOption) (*ConfrimAppointmentResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) SetAvailability(ctx context.Context, in *SetAvailabilityRequest, opts ...grpc.CallOption) (*SetAvailabilityResponse, error) {
	out := new(SetAvailabilityResponse)
	err := c.cc.Invoke(ctx, ReservationService_SetAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAppointments(ctx context.Context, in *GetAppointmentsRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error) {
	out := new(GetAppointmentsResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAppointments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ReserveAppointment(ctx context.Context, in *ReserveAppointmentRequest, opts ...grpc.CallOption) (*ReserveAppointmentResponse, error) {
	out := new(ReserveAppointmentResponse)
	err := c.cc.Invoke(ctx, ReservationService_ReserveAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ConfrimAppointment(ctx context.Context, in *ConfrimAppointmentRequest, opts ...grpc.CallOption) (*ConfrimAppointmentResponse, error) {
	out := new(ConfrimAppointmentResponse)
	err := c.cc.Invoke(ctx, ReservationService_ConfrimAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	SetAvailability(context.Context, *SetAvailabilityRequest) (*SetAvailabilityResponse, error)
	GetAppointments(context.Context, *GetAppointmentsRequest) (*GetAppointmentsResponse, error)
	ReserveAppointment(context.Context, *ReserveAppointmentRequest) (*ReserveAppointmentResponse, error)
	ConfrimAppointment(context.Context, *ConfrimAppointmentRequest) (*ConfrimAppointmentResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) SetAvailability(context.Context, *SetAvailabilityRequest) (*SetAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAvailability not implemented")
}
func (UnimplementedReservationServiceServer) GetAppointments(context.Context, *GetAppointmentsRequest) (*GetAppointmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppointments not implemented")
}
func (UnimplementedReservationServiceServer) ReserveAppointment(context.Context, *ReserveAppointmentRequest) (*ReserveAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveAppointment not implemented")
}
func (UnimplementedReservationServiceServer) ConfrimAppointment(context.Context, *ConfrimAppointmentRequest) (*ConfrimAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfrimAppointment not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_SetAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).SetAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_SetAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).SetAvailability(ctx, req.(*SetAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAppointments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppointmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAppointments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAppointments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAppointments(ctx, req.(*GetAppointmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ReserveAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ReserveAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_ReserveAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ReserveAppointment(ctx, req.(*ReserveAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ConfrimAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfrimAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ConfrimAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_ConfrimAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ConfrimAppointment(ctx, req.(*ConfrimAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAvailability",
			Handler:    _ReservationService_SetAvailability_Handler,
		},
		{
			MethodName: "GetAppointments",
			Handler:    _ReservationService_GetAppointments_Handler,
		},
		{
			MethodName: "ReserveAppointment",
			Handler:    _ReservationService_ReserveAppointment_Handler,
		},
		{
			MethodName: "ConfrimAppointment",
			Handler:    _ReservationService_ConfrimAppointment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models.proto",
}
