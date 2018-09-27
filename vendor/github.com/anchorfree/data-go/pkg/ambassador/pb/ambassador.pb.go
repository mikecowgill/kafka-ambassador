// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ambassador.proto

package pb

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProdRq struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	StreamOffset         uint64   `protobuf:"varint,3,opt,name=streamOffset" json:"streamOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRq) Reset()         { *m = ProdRq{} }
func (m *ProdRq) String() string { return proto.CompactTextString(m) }
func (*ProdRq) ProtoMessage()    {}
func (*ProdRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ambassador_dc02ba1e67403bba, []int{0}
}
func (m *ProdRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRq.Unmarshal(m, b)
}
func (m *ProdRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRq.Marshal(b, m, deterministic)
}
func (dst *ProdRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRq.Merge(dst, src)
}
func (m *ProdRq) XXX_Size() int {
	return xxx_messageInfo_ProdRq.Size(m)
}
func (m *ProdRq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRq.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRq proto.InternalMessageInfo

func (m *ProdRq) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ProdRq) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ProdRq) GetStreamOffset() uint64 {
	if m != nil {
		return m.StreamOffset
	}
	return 0
}

type ProdRs struct {
	StreamOffset         uint64   `protobuf:"varint,3,opt,name=streamOffset" json:"streamOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRs) Reset()         { *m = ProdRs{} }
func (m *ProdRs) String() string { return proto.CompactTextString(m) }
func (*ProdRs) ProtoMessage()    {}
func (*ProdRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ambassador_dc02ba1e67403bba, []int{1}
}
func (m *ProdRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRs.Unmarshal(m, b)
}
func (m *ProdRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRs.Marshal(b, m, deterministic)
}
func (dst *ProdRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRs.Merge(dst, src)
}
func (m *ProdRs) XXX_Size() int {
	return xxx_messageInfo_ProdRs.Size(m)
}
func (m *ProdRs) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRs.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRs proto.InternalMessageInfo

func (m *ProdRs) GetStreamOffset() uint64 {
	if m != nil {
		return m.StreamOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*ProdRq)(nil), "ProdRq")
	proto.RegisterType((*ProdRs)(nil), "ProdRs")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KafkaAmbassadorClient is the client API for KafkaAmbassador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KafkaAmbassadorClient interface {
	Produce(ctx context.Context, opts ...grpc.CallOption) (KafkaAmbassador_ProduceClient, error)
}

type kafkaAmbassadorClient struct {
	cc *grpc.ClientConn
}

func NewKafkaAmbassadorClient(cc *grpc.ClientConn) KafkaAmbassadorClient {
	return &kafkaAmbassadorClient{cc}
}

func (c *kafkaAmbassadorClient) Produce(ctx context.Context, opts ...grpc.CallOption) (KafkaAmbassador_ProduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KafkaAmbassador_serviceDesc.Streams[0], "/KafkaAmbassador/Produce", opts...)
	if err != nil {
		return nil, err
	}
	x := &kafkaAmbassadorProduceClient{stream}
	return x, nil
}

type KafkaAmbassador_ProduceClient interface {
	Send(*ProdRq) error
	Recv() (*ProdRs, error)
	grpc.ClientStream
}

type kafkaAmbassadorProduceClient struct {
	grpc.ClientStream
}

func (x *kafkaAmbassadorProduceClient) Send(m *ProdRq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kafkaAmbassadorProduceClient) Recv() (*ProdRs, error) {
	m := new(ProdRs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KafkaAmbassadorServer is the server API for KafkaAmbassador service.
type KafkaAmbassadorServer interface {
	Produce(KafkaAmbassador_ProduceServer) error
}

func RegisterKafkaAmbassadorServer(s *grpc.Server, srv KafkaAmbassadorServer) {
	s.RegisterService(&_KafkaAmbassador_serviceDesc, srv)
}

func _KafkaAmbassador_Produce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KafkaAmbassadorServer).Produce(&kafkaAmbassadorProduceServer{stream})
}

type KafkaAmbassador_ProduceServer interface {
	Send(*ProdRs) error
	Recv() (*ProdRq, error)
	grpc.ServerStream
}

type kafkaAmbassadorProduceServer struct {
	grpc.ServerStream
}

func (x *kafkaAmbassadorProduceServer) Send(m *ProdRs) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kafkaAmbassadorProduceServer) Recv() (*ProdRq, error) {
	m := new(ProdRq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KafkaAmbassador_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KafkaAmbassador",
	HandlerType: (*KafkaAmbassadorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Produce",
			Handler:       _KafkaAmbassador_Produce_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ambassador.proto",
}

func init() { proto.RegisterFile("ambassador.proto", fileDescriptor_ambassador_dc02ba1e67403bba) }

var fileDescriptor_ambassador_dc02ba1e67403bba = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcc, 0x4d, 0x4a,
	0x2c, 0x2e, 0x4e, 0x4c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe1, 0x62,
	0x0b, 0x28, 0xca, 0x4f, 0x09, 0x2a, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2f, 0xc8, 0x4c, 0x96,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x21, 0x25, 0x2e, 0x9e,
	0xe2, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0xff, 0xb4, 0xb4, 0xe2, 0xd4, 0x12, 0x09, 0x66, 0x05, 0x46,
	0x0d, 0x96, 0x20, 0x14, 0x31, 0x25, 0x1d, 0xa8, 0xe9, 0xc5, 0xc4, 0xa8, 0x36, 0x32, 0xe1, 0xe2,
	0xf7, 0x4e, 0x4c, 0xcb, 0x4e, 0x74, 0x84, 0x3b, 0x52, 0x48, 0x91, 0x8b, 0x1d, 0x64, 0x40, 0x69,
	0x72, 0xaa, 0x10, 0xbb, 0x1e, 0xc4, 0xa1, 0x52, 0x50, 0x46, 0xb1, 0x12, 0x83, 0x06, 0xa3, 0x01,
	0xa3, 0x13, 0x4b, 0x14, 0x53, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x3b, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb3, 0x33, 0xef, 0x49, 0xe2, 0x00, 0x00, 0x00,
}