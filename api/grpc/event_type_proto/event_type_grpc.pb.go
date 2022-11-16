// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: api/grpc/event_type_proto/event_type.proto

package event_type_proto

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

// EventTypeServiceClient is the client API for EventTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventTypeServiceClient interface {
	Create(ctx context.Context, in *CreateEventTypeRequest, opts ...grpc.CallOption) (*EventType, error)
	Get(ctx context.Context, in *GetEventTypeRequest, opts ...grpc.CallOption) (*EventType, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventTypes, error)
	Update(ctx context.Context, in *UpdateEventTypeRequest, opts ...grpc.CallOption) (*EventType, error)
	Delete(ctx context.Context, in *GetEventTypeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eventTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventTypeServiceClient(cc grpc.ClientConnInterface) EventTypeServiceClient {
	return &eventTypeServiceClient{cc}
}

func (c *eventTypeServiceClient) Create(ctx context.Context, in *CreateEventTypeRequest, opts ...grpc.CallOption) (*EventType, error) {
	out := new(EventType)
	err := c.cc.Invoke(ctx, "/event_type_proto.EventTypeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTypeServiceClient) Get(ctx context.Context, in *GetEventTypeRequest, opts ...grpc.CallOption) (*EventType, error) {
	out := new(EventType)
	err := c.cc.Invoke(ctx, "/event_type_proto.EventTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTypeServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventTypes, error) {
	out := new(EventTypes)
	err := c.cc.Invoke(ctx, "/event_type_proto.EventTypeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTypeServiceClient) Update(ctx context.Context, in *UpdateEventTypeRequest, opts ...grpc.CallOption) (*EventType, error) {
	out := new(EventType)
	err := c.cc.Invoke(ctx, "/event_type_proto.EventTypeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTypeServiceClient) Delete(ctx context.Context, in *GetEventTypeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/event_type_proto.EventTypeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventTypeServiceServer is the server API for EventTypeService service.
// All implementations must embed UnimplementedEventTypeServiceServer
// for forward compatibility
type EventTypeServiceServer interface {
	Create(context.Context, *CreateEventTypeRequest) (*EventType, error)
	Get(context.Context, *GetEventTypeRequest) (*EventType, error)
	List(context.Context, *emptypb.Empty) (*EventTypes, error)
	Update(context.Context, *UpdateEventTypeRequest) (*EventType, error)
	Delete(context.Context, *GetEventTypeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedEventTypeServiceServer()
}

// UnimplementedEventTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventTypeServiceServer struct {
}

func (UnimplementedEventTypeServiceServer) Create(context.Context, *CreateEventTypeRequest) (*EventType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEventTypeServiceServer) Get(context.Context, *GetEventTypeRequest) (*EventType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEventTypeServiceServer) List(context.Context, *emptypb.Empty) (*EventTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEventTypeServiceServer) Update(context.Context, *UpdateEventTypeRequest) (*EventType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEventTypeServiceServer) Delete(context.Context, *GetEventTypeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEventTypeServiceServer) mustEmbedUnimplementedEventTypeServiceServer() {}

// UnsafeEventTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventTypeServiceServer will
// result in compilation errors.
type UnsafeEventTypeServiceServer interface {
	mustEmbedUnimplementedEventTypeServiceServer()
}

func RegisterEventTypeServiceServer(s grpc.ServiceRegistrar, srv EventTypeServiceServer) {
	s.RegisterService(&EventTypeService_ServiceDesc, srv)
}

func _EventTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_type_proto.EventTypeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).Create(ctx, req.(*CreateEventTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_type_proto.EventTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).Get(ctx, req.(*GetEventTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_type_proto.EventTypeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_type_proto.EventTypeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).Update(ctx, req.(*UpdateEventTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_type_proto.EventTypeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).Delete(ctx, req.(*GetEventTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventTypeService_ServiceDesc is the grpc.ServiceDesc for EventTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event_type_proto.EventTypeService",
	HandlerType: (*EventTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EventTypeService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EventTypeService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EventTypeService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EventTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EventTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/event_type_proto/event_type.proto",
}