// Code generated by protoc-gen-go.
// source: squidbot.proto
// DO NOT EDIT!

package protocols

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FrontendRequest struct {
	// The version number of the frontend making the start request.
	FrontendVersion string `protobuf:"bytes,1,opt,name=frontend_version,json=frontendVersion" json:"frontend_version,omitempty"`
	// The name of the frontend making the start request.
	FrontendName string `protobuf:"bytes,2,opt,name=frontend_name,json=frontendName" json:"frontend_name,omitempty"`
}

func (m *FrontendRequest) Reset()                    { *m = FrontendRequest{} }
func (m *FrontendRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontendRequest) ProtoMessage()               {}
func (*FrontendRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *FrontendRequest) GetFrontendVersion() string {
	if m != nil {
		return m.FrontendVersion
	}
	return ""
}

func (m *FrontendRequest) GetFrontendName() string {
	if m != nil {
		return m.FrontendName
	}
	return ""
}

type FrontendResponse struct {
	// The version of the brain server the frontend is speaking to.
	BrainVersion string `protobuf:"bytes,1,opt,name=brain_version,json=brainVersion" json:"brain_version,omitempty"`
}

func (m *FrontendResponse) Reset()                    { *m = FrontendResponse{} }
func (m *FrontendResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontendResponse) ProtoMessage()               {}
func (*FrontendResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FrontendResponse) GetBrainVersion() string {
	if m != nil {
		return m.BrainVersion
	}
	return ""
}

// A message coming in from a frontend driver containing the sender's handle
type MessageRequest struct {
	// Used to help coordinate requests to replies.
	UniqueId int64 `protobuf:"varint,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	// Timestamp of the message as specified by the upstream service -- not the
	// frontend driver for squidbot.
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// The name of the sender who sent this message.
	SenderHandle string `protobuf:"bytes,3,opt,name=sender_handle,json=senderHandle" json:"sender_handle,omitempty"`
	// The body of the message.
	MessageBody string `protobuf:"bytes,4,opt,name=message_body,json=messageBody" json:"message_body,omitempty"`
	// The channel the message was sent in.
	ChannelName string `protobuf:"bytes,5,opt,name=channel_name,json=channelName" json:"channel_name,omitempty"`
	// Was this a 1:1 message?
	IsOneToOne bool `protobuf:"varint,6,opt,name=is_one_to_one,json=isOneToOne" json:"is_one_to_one,omitempty"`
}

func (m *MessageRequest) Reset()                    { *m = MessageRequest{} }
func (m *MessageRequest) String() string            { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()               {}
func (*MessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *MessageRequest) GetUniqueId() int64 {
	if m != nil {
		return m.UniqueId
	}
	return 0
}

func (m *MessageRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageRequest) GetSenderHandle() string {
	if m != nil {
		return m.SenderHandle
	}
	return ""
}

func (m *MessageRequest) GetMessageBody() string {
	if m != nil {
		return m.MessageBody
	}
	return ""
}

func (m *MessageRequest) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func (m *MessageRequest) GetIsOneToOne() bool {
	if m != nil {
		return m.IsOneToOne
	}
	return false
}

type MessageResponse struct {
	// The id of the message sent from a frontend to the brain. Helps to
	// coordinate the responses to the reuqests.
	UniqueId int64 `protobuf:"varint,1,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	// The message containing the response.
	MessageBody string `protobuf:"bytes,2,opt,name=message_body,json=messageBody" json:"message_body,omitempty"`
}

func (m *MessageResponse) Reset()                    { *m = MessageResponse{} }
func (m *MessageResponse) String() string            { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()               {}
func (*MessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *MessageResponse) GetUniqueId() int64 {
	if m != nil {
		return m.UniqueId
	}
	return 0
}

func (m *MessageResponse) GetMessageBody() string {
	if m != nil {
		return m.MessageBody
	}
	return ""
}

func init() {
	proto.RegisterType((*FrontendRequest)(nil), "protocols.FrontendRequest")
	proto.RegisterType((*FrontendResponse)(nil), "protocols.FrontendResponse")
	proto.RegisterType((*MessageRequest)(nil), "protocols.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "protocols.MessageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Brain service

type BrainClient interface {
	// First call that a frontend must make to the brain to ensure that it
	// supports the correct version of the brain and also ensure the brain is up.
	FrontendStarted(ctx context.Context, in *FrontendRequest, opts ...grpc.CallOption) (*FrontendResponse, error)
	// Posts a message from a frontend
	OnNewMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type brainClient struct {
	cc *grpc.ClientConn
}

func NewBrainClient(cc *grpc.ClientConn) BrainClient {
	return &brainClient{cc}
}

func (c *brainClient) FrontendStarted(ctx context.Context, in *FrontendRequest, opts ...grpc.CallOption) (*FrontendResponse, error) {
	out := new(FrontendResponse)
	err := grpc.Invoke(ctx, "/protocols.Brain/FrontendStarted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brainClient) OnNewMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := grpc.Invoke(ctx, "/protocols.Brain/OnNewMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Brain service

type BrainServer interface {
	// First call that a frontend must make to the brain to ensure that it
	// supports the correct version of the brain and also ensure the brain is up.
	FrontendStarted(context.Context, *FrontendRequest) (*FrontendResponse, error)
	// Posts a message from a frontend
	OnNewMessage(context.Context, *MessageRequest) (*MessageResponse, error)
}

func RegisterBrainServer(s *grpc.Server, srv BrainServer) {
	s.RegisterService(&_Brain_serviceDesc, srv)
}

func _Brain_FrontendStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FrontendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrainServer).FrontendStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.Brain/FrontendStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrainServer).FrontendStarted(ctx, req.(*FrontendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brain_OnNewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrainServer).OnNewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.Brain/OnNewMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrainServer).OnNewMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Brain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocols.Brain",
	HandlerType: (*BrainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FrontendStarted",
			Handler:    _Brain_FrontendStarted_Handler,
		},
		{
			MethodName: "OnNewMessage",
			Handler:    _Brain_OnNewMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squidbot.proto",
}

func init() { proto.RegisterFile("squidbot.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xcb, 0xae, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0x4b, 0xab, 0xc6, 0xa4, 0x17, 0x79, 0x15, 0x52, 0x16, 0x10, 0x84, 0x04, 0x9b,
	0x2c, 0x60, 0x01, 0xeb, 0x2c, 0xb8, 0x48, 0xd0, 0x96, 0x14, 0xb1, 0x8d, 0x9c, 0x7a, 0x28, 0x96,
	0x1a, 0xbb, 0xb5, 0x9d, 0x22, 0x5e, 0x86, 0xf7, 0xe2, 0x6d, 0xf0, 0x25, 0x69, 0x69, 0x4f, 0x75,
	0x36, 0x71, 0xf4, 0xcd, 0x3f, 0x33, 0xff, 0x78, 0x8c, 0x26, 0xea, 0xd0, 0x30, 0x5a, 0x09, 0x9d,
	0xed, 0xa5, 0xd0, 0x02, 0x87, 0xee, 0xd8, 0x88, 0x9d, 0x4a, 0x09, 0x9a, 0xbe, 0x97, 0x82, 0x6b,
	0xe0, 0xb4, 0x80, 0x43, 0x03, 0x4a, 0xe3, 0x57, 0x68, 0xf6, 0xa3, 0x45, 0xe5, 0x11, 0xa4, 0x62,
	0x82, 0xc7, 0xc1, 0xd3, 0xe0, 0x65, 0x58, 0x4c, 0x3b, 0xfe, 0xdd, 0x63, 0xfc, 0x1c, 0x8d, 0x4f,
	0x52, 0x4e, 0x6a, 0x88, 0x7b, 0x4e, 0x17, 0x75, 0x70, 0x61, 0x58, 0xfa, 0x16, 0xcd, 0xce, 0x2d,
	0xd4, 0x5e, 0x70, 0x05, 0x36, 0xb1, 0x92, 0x84, 0xf1, 0xab, 0x06, 0x91, 0x83, 0x6d, 0xf5, 0xf4,
	0x6f, 0x80, 0x26, 0x5f, 0x40, 0x29, 0xb2, 0x85, 0xce, 0xdb, 0x1c, 0x85, 0x0d, 0x67, 0xe6, 0xbf,
	0x64, 0xd4, 0xe5, 0xf4, 0x8b, 0x91, 0x07, 0x9f, 0x28, 0x7e, 0x82, 0x42, 0xcd, 0x6a, 0x23, 0x23,
	0xf5, 0xde, 0x39, 0xe9, 0x17, 0x67, 0x60, 0x5b, 0x2a, 0x63, 0x01, 0x64, 0xf9, 0x93, 0x70, 0xba,
	0x83, 0xb8, 0xef, 0x5b, 0x7a, 0xf8, 0xd1, 0x31, 0xfc, 0x0c, 0x45, 0xb5, 0xef, 0x58, 0x56, 0x82,
	0xfe, 0x8e, 0x1f, 0x3a, 0xcd, 0xa3, 0x96, 0xe5, 0x06, 0x59, 0xc9, 0xc6, 0x54, 0xe0, 0xb0, 0xf3,
	0x23, 0x0f, 0xbc, 0xa4, 0x65, 0x76, 0x62, 0x23, 0x19, 0x33, 0x55, 0x0a, 0x0e, 0xa5, 0x16, 0xf6,
	0x88, 0x87, 0x46, 0x33, 0x2a, 0x10, 0x53, 0x4b, 0x0e, 0xdf, 0x84, 0xf9, 0xa4, 0x5f, 0xd1, 0xf4,
	0x34, 0x5a, 0x7b, 0x27, 0xf7, 0xce, 0x76, 0x6d, 0xac, 0x77, 0xc7, 0xd8, 0xeb, 0x3f, 0x01, 0x1a,
	0xe4, 0xf6, 0xfe, 0xf0, 0xe7, 0xf3, 0x52, 0xd7, 0x9a, 0x48, 0x0d, 0x14, 0x27, 0xd9, 0x69, 0xe7,
	0xd9, 0xd5, 0xc2, 0x93, 0xf9, 0xcd, 0x98, 0x77, 0x95, 0x3e, 0xc0, 0x1f, 0x50, 0xb4, 0xe4, 0x0b,
	0xf8, 0xd5, 0xfa, 0xc5, 0x8f, 0xff, 0x93, 0x5f, 0xae, 0x27, 0x49, 0x6e, 0x85, 0xba, 0x42, 0xf9,
	0x3b, 0xf4, 0x42, 0xc8, 0x6d, 0x46, 0xe8, 0x11, 0xb8, 0x6e, 0x24, 0x98, 0x57, 0x49, 0x68, 0x76,
	0xf9, 0x36, 0x6d, 0x6a, 0x3e, 0x5e, 0xb7, 0x6c, 0x65, 0xd1, 0x2a, 0xa8, 0x86, 0x2e, 0xf6, 0xe6,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x6c, 0x26, 0x40, 0xc9, 0x02, 0x00, 0x00,
}
