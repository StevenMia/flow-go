// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insecure/network.proto

package insecure

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Protocol int32

const (
	Protocol_UNKNOWN   Protocol = 0
	Protocol_UNICAST   Protocol = 1
	Protocol_MULTICAST Protocol = 2
	Protocol_PUBLISH   Protocol = 3
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "UNICAST",
	2: "MULTICAST",
	3: "PUBLISH",
}

var Protocol_value = map[string]int32{
	"UNKNOWN":   0,
	"UNICAST":   1,
	"MULTICAST": 2,
	"PUBLISH":   3,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2e799291353e003, []int{0}
}

// Message represents the messages exchanged between the CorruptNetwork (server) and Attacker (client).
// This is a wrapper for both egress and ingress messages.
type Message struct {
	Egress               *EgressMessage  `protobuf:"bytes,1,opt,name=Egress,proto3" json:"Egress,omitempty"`
	Ingress              *IngressMessage `protobuf:"bytes,2,opt,name=Ingress,proto3" json:"Ingress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2e799291353e003, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetEgress() *EgressMessage {
	if m != nil {
		return m.Egress
	}
	return nil
}

func (m *Message) GetIngress() *IngressMessage {
	if m != nil {
		return m.Ingress
	}
	return nil
}

// EgressMessage represents an outgoing message from a corrupt node to another (honest or corrupt) node.
// The exchanged message is between the CorruptConduitFactory and Attacker.
type EgressMessage struct {
	ChannelID string `protobuf:"bytes,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	// CorruptOriginID represents the corrupt node id where the outgoing message is coming from.
	CorruptOriginID      []byte   `protobuf:"bytes,2,opt,name=CorruptOriginID,proto3" json:"CorruptOriginID,omitempty"`
	TargetNum            uint32   `protobuf:"varint,3,opt,name=TargetNum,proto3" json:"TargetNum,omitempty"`
	TargetIDs            [][]byte `protobuf:"bytes,4,rep,name=TargetIDs,proto3" json:"TargetIDs,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Protocol             Protocol `protobuf:"varint,6,opt,name=protocol,proto3,enum=net.Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EgressMessage) Reset()         { *m = EgressMessage{} }
func (m *EgressMessage) String() string { return proto.CompactTextString(m) }
func (*EgressMessage) ProtoMessage()    {}
func (*EgressMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2e799291353e003, []int{1}
}

func (m *EgressMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EgressMessage.Unmarshal(m, b)
}
func (m *EgressMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EgressMessage.Marshal(b, m, deterministic)
}
func (m *EgressMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EgressMessage.Merge(m, src)
}
func (m *EgressMessage) XXX_Size() int {
	return xxx_messageInfo_EgressMessage.Size(m)
}
func (m *EgressMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EgressMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EgressMessage proto.InternalMessageInfo

func (m *EgressMessage) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *EgressMessage) GetCorruptOriginID() []byte {
	if m != nil {
		return m.CorruptOriginID
	}
	return nil
}

func (m *EgressMessage) GetTargetNum() uint32 {
	if m != nil {
		return m.TargetNum
	}
	return 0
}

func (m *EgressMessage) GetTargetIDs() [][]byte {
	if m != nil {
		return m.TargetIDs
	}
	return nil
}

func (m *EgressMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *EgressMessage) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_UNKNOWN
}

// IngressMessage represents an incoming message from another node (honest or corrupt) to a corrupt node.
type IngressMessage struct {
	ChannelID string `protobuf:"bytes,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	// OriginID represents the node id where the incoming message is coming from - that node could be corrupt or honest.
	OriginID             []byte   `protobuf:"bytes,2,opt,name=OriginID,proto3" json:"OriginID,omitempty"`
	CorruptTargetID      []byte   `protobuf:"bytes,3,opt,name=CorruptTargetID,proto3" json:"CorruptTargetID,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngressMessage) Reset()         { *m = IngressMessage{} }
func (m *IngressMessage) String() string { return proto.CompactTextString(m) }
func (*IngressMessage) ProtoMessage()    {}
func (*IngressMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2e799291353e003, []int{2}
}

func (m *IngressMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressMessage.Unmarshal(m, b)
}
func (m *IngressMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressMessage.Marshal(b, m, deterministic)
}
func (m *IngressMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressMessage.Merge(m, src)
}
func (m *IngressMessage) XXX_Size() int {
	return xxx_messageInfo_IngressMessage.Size(m)
}
func (m *IngressMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IngressMessage proto.InternalMessageInfo

func (m *IngressMessage) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *IngressMessage) GetOriginID() []byte {
	if m != nil {
		return m.OriginID
	}
	return nil
}

func (m *IngressMessage) GetCorruptTargetID() []byte {
	if m != nil {
		return m.CorruptTargetID
	}
	return nil
}

func (m *IngressMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("net.Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*Message)(nil), "net.Message")
	proto.RegisterType((*EgressMessage)(nil), "net.EgressMessage")
	proto.RegisterType((*IngressMessage)(nil), "net.IngressMessage")
}

func init() { proto.RegisterFile("insecure/network.proto", fileDescriptor_b2e799291353e003) }

var fileDescriptor_b2e799291353e003 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0xa3, 0xba, 0x8b, 0x9d, 0xd3, 0x38, 0x0b, 0x1a, 0x04, 0xe3, 0xed, 0xc2, 0xf8, 0xca,
	0x2b, 0xcc, 0x19, 0xd9, 0xe5, 0x6e, 0xd6, 0x24, 0x85, 0x99, 0xb5, 0xae, 0x71, 0x13, 0x06, 0xbb,
	0x73, 0x5d, 0xcd, 0x0b, 0x75, 0xa5, 0x20, 0xc9, 0x8c, 0xbe, 0xc4, 0x60, 0x6f, 0xb7, 0xc7, 0x19,
	0x96, 0xad, 0x3a, 0x0e, 0x8c, 0x5d, 0x9e, 0xff, 0x7c, 0x3a, 0x3a, 0xff, 0xe1, 0x87, 0xd9, 0x8e,
	0x0a, 0x92, 0x57, 0x9c, 0xcc, 0x29, 0x91, 0x3f, 0x19, 0x7f, 0x08, 0xf7, 0x9c, 0x49, 0x86, 0x0d,
	0x4a, 0xa4, 0xfb, 0xba, 0x60, 0xac, 0x28, 0xc9, 0x5c, 0x49, 0x77, 0xd5, 0xf7, 0x39, 0x79, 0xdc,
	0xcb, 0xa7, 0x86, 0xf0, 0xef, 0xc1, 0xbc, 0x26, 0x42, 0x64, 0x05, 0xc1, 0xe7, 0x30, 0xbc, 0x2c,
	0x38, 0x11, 0xc2, 0x41, 0x1e, 0x0a, 0xce, 0x16, 0x38, 0xa4, 0x44, 0x86, 0x8d, 0xd4, 0x32, 0x69,
	0x4b, 0xe0, 0x77, 0x60, 0x46, 0xb4, 0x81, 0x4f, 0x14, 0xfc, 0x4a, 0xc1, 0xad, 0xa6, 0x69, 0xcd,
	0xf8, 0x7f, 0x10, 0xd8, 0xbd, 0x41, 0xf8, 0x0d, 0x8c, 0x56, 0x3f, 0x32, 0x4a, 0x49, 0x19, 0xad,
	0xd5, 0x7f, 0xa3, 0xb4, 0x13, 0x70, 0x00, 0x2f, 0x57, 0x8c, 0xf3, 0x6a, 0x2f, 0x6f, 0xf8, 0xae,
	0xd8, 0xd1, 0x68, 0xad, 0xbe, 0x19, 0xa7, 0xc7, 0x72, 0x3d, 0x67, 0x93, 0xf1, 0x82, 0xc8, 0xb8,
	0x7a, 0x74, 0x0c, 0x0f, 0x05, 0x76, 0xda, 0x09, 0x5d, 0x37, 0x5a, 0x0b, 0xe7, 0xd4, 0x33, 0x82,
	0x71, 0xda, 0x09, 0xd8, 0x01, 0x33, 0xc9, 0x9e, 0x4a, 0x96, 0xdd, 0x3b, 0x2f, 0xd4, 0x74, 0x5d,
	0xe2, 0xb7, 0x60, 0xa9, 0xf3, 0xe4, 0xac, 0x74, 0x86, 0x1e, 0x0a, 0x26, 0x0b, 0x5b, 0xf9, 0x4b,
	0x5a, 0x31, 0x7d, 0x6e, 0xfb, 0xbf, 0x10, 0x4c, 0xfa, 0xb6, 0xff, 0xe3, 0xcd, 0x05, 0xeb, 0xc8,
	0xd4, 0x73, 0x7d, 0xe0, 0x5b, 0x6f, 0xa9, 0x3c, 0x75, 0xbe, 0xb5, 0x7c, 0xb8, 0xfb, 0x69, 0x6f,
	0xf7, 0xf3, 0x4f, 0x60, 0xe9, 0x35, 0xf1, 0x19, 0x98, 0xdb, 0xf8, 0x4b, 0x7c, 0xf3, 0x35, 0x9e,
	0x0e, 0x9a, 0x22, 0x5a, 0x5d, 0xdc, 0x6e, 0xa6, 0x08, 0xdb, 0x30, 0xba, 0xde, 0x5e, 0x6d, 0x9a,
	0xf2, 0xa4, 0xee, 0x25, 0xdb, 0xe5, 0x55, 0x74, 0xfb, 0x79, 0x6a, 0x2c, 0x7e, 0x23, 0x98, 0xb4,
	0xff, 0xc5, 0x4d, 0x9c, 0xf0, 0xc7, 0x7a, 0x31, 0x4a, 0x49, 0x2e, 0x2f, 0xa4, 0xcc, 0xf2, 0x07,
	0xc2, 0xf1, 0x2c, 0x6c, 0x72, 0x15, 0xea, 0x5c, 0x85, 0x97, 0x75, 0xae, 0xdc, 0xb1, 0xba, 0x54,
	0x7b, 0x0b, 0x7f, 0xf0, 0x1e, 0xe1, 0x25, 0xcc, 0x12, 0xce, 0x72, 0x22, 0x84, 0x7e, 0xac, 0x2f,
	0xd5, 0x63, 0xdd, 0x7f, 0x4c, 0xf4, 0x07, 0x01, 0x5a, 0xc2, 0x37, 0x4b, 0x67, 0xfc, 0x6e, 0xa8,
	0xfa, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xef, 0x6f, 0xa2, 0x8c, 0xf6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CorruptNetworkClient is the client API for CorruptNetwork service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CorruptNetworkClient interface {
	// ConnectAttacker registers an attacker to the corrupt network.
	ConnectAttacker(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CorruptNetwork_ConnectAttackerClient, error)
	// ProcessAttackerMessage is the central place for the corrupt network to process messages from an attacker.
	ProcessAttackerMessage(ctx context.Context, opts ...grpc.CallOption) (CorruptNetwork_ProcessAttackerMessageClient, error)
}

type corruptNetworkClient struct {
	cc *grpc.ClientConn
}

func NewCorruptNetworkClient(cc *grpc.ClientConn) CorruptNetworkClient {
	return &corruptNetworkClient{cc}
}

func (c *corruptNetworkClient) ConnectAttacker(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CorruptNetwork_ConnectAttackerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CorruptNetwork_serviceDesc.Streams[0], "/net.CorruptNetwork/ConnectAttacker", opts...)
	if err != nil {
		return nil, err
	}
	x := &corruptNetworkConnectAttackerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CorruptNetwork_ConnectAttackerClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type corruptNetworkConnectAttackerClient struct {
	grpc.ClientStream
}

func (x *corruptNetworkConnectAttackerClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *corruptNetworkClient) ProcessAttackerMessage(ctx context.Context, opts ...grpc.CallOption) (CorruptNetwork_ProcessAttackerMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CorruptNetwork_serviceDesc.Streams[1], "/net.CorruptNetwork/ProcessAttackerMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &corruptNetworkProcessAttackerMessageClient{stream}
	return x, nil
}

type CorruptNetwork_ProcessAttackerMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type corruptNetworkProcessAttackerMessageClient struct {
	grpc.ClientStream
}

func (x *corruptNetworkProcessAttackerMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *corruptNetworkProcessAttackerMessageClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CorruptNetworkServer is the server API for CorruptNetwork service.
type CorruptNetworkServer interface {
	// ConnectAttacker registers an attacker to the corrupt network.
	ConnectAttacker(*emptypb.Empty, CorruptNetwork_ConnectAttackerServer) error
	// ProcessAttackerMessage is the central place for the corrupt network to process messages from an attacker.
	ProcessAttackerMessage(CorruptNetwork_ProcessAttackerMessageServer) error
}

// UnimplementedCorruptNetworkServer can be embedded to have forward compatible implementations.
type UnimplementedCorruptNetworkServer struct {
}

func (*UnimplementedCorruptNetworkServer) ConnectAttacker(req *emptypb.Empty, srv CorruptNetwork_ConnectAttackerServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectAttacker not implemented")
}
func (*UnimplementedCorruptNetworkServer) ProcessAttackerMessage(srv CorruptNetwork_ProcessAttackerMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessAttackerMessage not implemented")
}

func RegisterCorruptNetworkServer(s *grpc.Server, srv CorruptNetworkServer) {
	s.RegisterService(&_CorruptNetwork_serviceDesc, srv)
}

func _CorruptNetwork_ConnectAttacker_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CorruptNetworkServer).ConnectAttacker(m, &corruptNetworkConnectAttackerServer{stream})
}

type CorruptNetwork_ConnectAttackerServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type corruptNetworkConnectAttackerServer struct {
	grpc.ServerStream
}

func (x *corruptNetworkConnectAttackerServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _CorruptNetwork_ProcessAttackerMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CorruptNetworkServer).ProcessAttackerMessage(&corruptNetworkProcessAttackerMessageServer{stream})
}

type CorruptNetwork_ProcessAttackerMessageServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type corruptNetworkProcessAttackerMessageServer struct {
	grpc.ServerStream
}

func (x *corruptNetworkProcessAttackerMessageServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *corruptNetworkProcessAttackerMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CorruptNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "net.CorruptNetwork",
	HandlerType: (*CorruptNetworkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectAttacker",
			Handler:       _CorruptNetwork_ConnectAttacker_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProcessAttackerMessage",
			Handler:       _CorruptNetwork_ProcessAttackerMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "insecure/network.proto",
}
