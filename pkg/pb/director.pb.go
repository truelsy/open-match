// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/director.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SynRegisterArena struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	MaxUserCount         int32    `protobuf:"varint,3,opt,name=max_user_count,json=maxUserCount,proto3" json:"max_user_count,omitempty"`
	SendTime             int64    `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynRegisterArena) Reset()         { *m = SynRegisterArena{} }
func (m *SynRegisterArena) String() string { return proto.CompactTextString(m) }
func (*SynRegisterArena) ProtoMessage()    {}
func (*SynRegisterArena) Descriptor() ([]byte, []int) {
	return fileDescriptor_419999d44e42e2ad, []int{0}
}

func (m *SynRegisterArena) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynRegisterArena.Unmarshal(m, b)
}
func (m *SynRegisterArena) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynRegisterArena.Marshal(b, m, deterministic)
}
func (m *SynRegisterArena) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynRegisterArena.Merge(m, src)
}
func (m *SynRegisterArena) XXX_Size() int {
	return xxx_messageInfo_SynRegisterArena.Size(m)
}
func (m *SynRegisterArena) XXX_DiscardUnknown() {
	xxx_messageInfo_SynRegisterArena.DiscardUnknown(m)
}

var xxx_messageInfo_SynRegisterArena proto.InternalMessageInfo

func (m *SynRegisterArena) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *SynRegisterArena) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SynRegisterArena) GetMaxUserCount() int32 {
	if m != nil {
		return m.MaxUserCount
	}
	return 0
}

func (m *SynRegisterArena) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

type SynHealthCheckArena struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	UserCount            int32    `protobuf:"varint,3,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	SendTime             int64    `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynHealthCheckArena) Reset()         { *m = SynHealthCheckArena{} }
func (m *SynHealthCheckArena) String() string { return proto.CompactTextString(m) }
func (*SynHealthCheckArena) ProtoMessage()    {}
func (*SynHealthCheckArena) Descriptor() ([]byte, []int) {
	return fileDescriptor_419999d44e42e2ad, []int{1}
}

func (m *SynHealthCheckArena) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynHealthCheckArena.Unmarshal(m, b)
}
func (m *SynHealthCheckArena) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynHealthCheckArena.Marshal(b, m, deterministic)
}
func (m *SynHealthCheckArena) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynHealthCheckArena.Merge(m, src)
}
func (m *SynHealthCheckArena) XXX_Size() int {
	return xxx_messageInfo_SynHealthCheckArena.Size(m)
}
func (m *SynHealthCheckArena) XXX_DiscardUnknown() {
	xxx_messageInfo_SynHealthCheckArena.DiscardUnknown(m)
}

var xxx_messageInfo_SynHealthCheckArena proto.InternalMessageInfo

func (m *SynHealthCheckArena) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *SynHealthCheckArena) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SynHealthCheckArena) GetUserCount() int32 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *SynHealthCheckArena) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

type SynTerminateArena struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynTerminateArena) Reset()         { *m = SynTerminateArena{} }
func (m *SynTerminateArena) String() string { return proto.CompactTextString(m) }
func (*SynTerminateArena) ProtoMessage()    {}
func (*SynTerminateArena) Descriptor() ([]byte, []int) {
	return fileDescriptor_419999d44e42e2ad, []int{2}
}

func (m *SynTerminateArena) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynTerminateArena.Unmarshal(m, b)
}
func (m *SynTerminateArena) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynTerminateArena.Marshal(b, m, deterministic)
}
func (m *SynTerminateArena) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynTerminateArena.Merge(m, src)
}
func (m *SynTerminateArena) XXX_Size() int {
	return xxx_messageInfo_SynTerminateArena.Size(m)
}
func (m *SynTerminateArena) XXX_DiscardUnknown() {
	xxx_messageInfo_SynTerminateArena.DiscardUnknown(m)
}

var xxx_messageInfo_SynTerminateArena proto.InternalMessageInfo

func (m *SynTerminateArena) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *SynTerminateArena) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type SynAvailableArena struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynAvailableArena) Reset()         { *m = SynAvailableArena{} }
func (m *SynAvailableArena) String() string { return proto.CompactTextString(m) }
func (*SynAvailableArena) ProtoMessage()    {}
func (*SynAvailableArena) Descriptor() ([]byte, []int) {
	return fileDescriptor_419999d44e42e2ad, []int{3}
}

func (m *SynAvailableArena) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynAvailableArena.Unmarshal(m, b)
}
func (m *SynAvailableArena) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynAvailableArena.Marshal(b, m, deterministic)
}
func (m *SynAvailableArena) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynAvailableArena.Merge(m, src)
}
func (m *SynAvailableArena) XXX_Size() int {
	return xxx_messageInfo_SynAvailableArena.Size(m)
}
func (m *SynAvailableArena) XXX_DiscardUnknown() {
	xxx_messageInfo_SynAvailableArena.DiscardUnknown(m)
}

var xxx_messageInfo_SynAvailableArena proto.InternalMessageInfo

func (m *SynAvailableArena) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type AckAvailableArena struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckAvailableArena) Reset()         { *m = AckAvailableArena{} }
func (m *AckAvailableArena) String() string { return proto.CompactTextString(m) }
func (*AckAvailableArena) ProtoMessage()    {}
func (*AckAvailableArena) Descriptor() ([]byte, []int) {
	return fileDescriptor_419999d44e42e2ad, []int{4}
}

func (m *AckAvailableArena) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckAvailableArena.Unmarshal(m, b)
}
func (m *AckAvailableArena) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckAvailableArena.Marshal(b, m, deterministic)
}
func (m *AckAvailableArena) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckAvailableArena.Merge(m, src)
}
func (m *AckAvailableArena) XXX_Size() int {
	return xxx_messageInfo_AckAvailableArena.Size(m)
}
func (m *AckAvailableArena) XXX_DiscardUnknown() {
	xxx_messageInfo_AckAvailableArena.DiscardUnknown(m)
}

var xxx_messageInfo_AckAvailableArena proto.InternalMessageInfo

func (m *AckAvailableArena) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*SynRegisterArena)(nil), "openmatch.SynRegisterArena")
	proto.RegisterType((*SynHealthCheckArena)(nil), "openmatch.SynHealthCheckArena")
	proto.RegisterType((*SynTerminateArena)(nil), "openmatch.SynTerminateArena")
	proto.RegisterType((*SynAvailableArena)(nil), "openmatch.SynAvailableArena")
	proto.RegisterType((*AckAvailableArena)(nil), "openmatch.AckAvailableArena")
}

func init() { proto.RegisterFile("api/director.proto", fileDescriptor_419999d44e42e2ad) }

var fileDescriptor_419999d44e42e2ad = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x4f, 0xdb, 0x48,
	0x14, 0xc7, 0x65, 0x87, 0x05, 0x32, 0xda, 0x45, 0x30, 0xbb, 0x1b, 0x45, 0x09, 0xac, 0x8c, 0xd9,
	0x03, 0xca, 0x6e, 0x3c, 0x90, 0xe5, 0x94, 0x55, 0x25, 0x52, 0x88, 0x54, 0x24, 0xda, 0x4a, 0x09,
	0xed, 0xa1, 0x17, 0x34, 0x19, 0x3f, 0xec, 0x29, 0xf1, 0x8c, 0x35, 0x33, 0x06, 0x72, 0x6d, 0x0f,
	0xbd, 0xb7, 0xb7, 0x7e, 0x84, 0x1e, 0xfb, 0x05, 0xfa, 0x21, 0x7a, 0xea, 0xbd, 0xdf, 0xa3, 0x95,
	0xed, 0x04, 0xe2, 0x04, 0xaa, 0x8a, 0x93, 0xf5, 0xe6, 0xbd, 0xf7, 0xff, 0xbd, 0x99, 0xf7, 0x97,
	0x11, 0xa6, 0x31, 0x27, 0x3e, 0x57, 0xc0, 0x8c, 0x54, 0x5e, 0xac, 0xa4, 0x91, 0xb8, 0x2c, 0x63,
	0x10, 0x11, 0x35, 0x2c, 0xac, 0xad, 0x07, 0x52, 0x06, 0x43, 0x20, 0x69, 0x15, 0x15, 0x42, 0x1a,
	0x6a, 0xb8, 0x14, 0x3a, 0x2f, 0xac, 0xfd, 0x9b, 0x7d, 0x58, 0x33, 0x00, 0xd1, 0xd4, 0x97, 0x34,
	0x08, 0x40, 0x11, 0x19, 0x67, 0x15, 0xb7, 0x54, 0xd7, 0xc7, 0x5a, 0x59, 0x34, 0x48, 0xce, 0x08,
	0x44, 0xb1, 0x19, 0xe5, 0x49, 0xf7, 0x8d, 0x85, 0x56, 0xfb, 0x23, 0xd1, 0x83, 0x80, 0x6b, 0x03,
	0xaa, 0xa3, 0x40, 0x50, 0x5c, 0x41, 0x8b, 0x0a, 0x02, 0x2e, 0x45, 0xd5, 0x72, 0xac, 0xed, 0x72,
	0x6f, 0x1c, 0xe1, 0x2a, 0x5a, 0xa2, 0xbe, 0xaf, 0x40, 0xeb, 0xaa, 0x9d, 0x25, 0x26, 0x21, 0xfe,
	0x1b, 0xad, 0x44, 0xf4, 0xea, 0x34, 0xd1, 0xa0, 0x4e, 0x99, 0x4c, 0x84, 0xa9, 0x96, 0x1c, 0x6b,
	0xfb, 0x97, 0xde, 0xaf, 0x11, 0xbd, 0x7a, 0xa6, 0x41, 0x1d, 0xa4, 0x67, 0xb8, 0x8e, 0xca, 0x1a,
	0x84, 0x7f, 0x6a, 0x78, 0x04, 0xd5, 0x05, 0xc7, 0xda, 0x2e, 0xf5, 0x96, 0xd3, 0x83, 0x13, 0x1e,
	0x81, 0xfb, 0xda, 0x42, 0xbf, 0xf7, 0x47, 0xe2, 0x11, 0xd0, 0xa1, 0x09, 0x0f, 0x42, 0x60, 0xe7,
	0xf7, 0x1d, 0x66, 0x03, 0xa1, 0xb9, 0x41, 0xca, 0xc9, 0xcf, 0x4d, 0xd1, 0x45, 0x6b, 0xfd, 0x91,
	0x38, 0x01, 0x15, 0x71, 0x41, 0x0d, 0xdc, 0x73, 0x04, 0xf7, 0x9f, 0x4c, 0xa6, 0x73, 0x41, 0xf9,
	0x90, 0x0e, 0x86, 0x3f, 0x96, 0x71, 0x9b, 0x68, 0xad, 0xc3, 0xce, 0x67, 0x8a, 0xa7, 0xb4, 0xad,
	0x82, 0x76, 0xeb, 0x53, 0x09, 0x2d, 0x1f, 0x8e, 0x9d, 0x83, 0xcf, 0xd0, 0x6f, 0xc5, 0xdd, 0xd5,
	0xbd, 0x6b, 0x17, 0x79, 0xb3, 0x8b, 0xad, 0x55, 0xbc, 0xdc, 0x0b, 0xde, 0xc4, 0x0b, 0x5e, 0x37,
	0xf5, 0x82, 0xeb, 0xbc, 0xfa, 0xfc, 0xf5, 0x9d, 0x5d, 0x73, 0xff, 0x24, 0x17, 0xbb, 0xd7, 0xae,
	0x24, 0x6a, 0xdc, 0xdb, 0xb6, 0x1a, 0x58, 0xa2, 0xd5, 0xb9, 0xcd, 0xfc, 0x55, 0x44, 0xcd, 0xe6,
	0xef, 0xa4, 0x6d, 0x65, 0xb4, 0x0d, 0xb7, 0x5a, 0xa0, 0x85, 0x59, 0x3b, 0x4b, 0xdb, 0x53, 0x20,
	0x47, 0x2b, 0x33, 0x5b, 0x58, 0x2f, 0xe2, 0x8a, 0xd9, 0x3b, 0x61, 0x9b, 0x19, 0xac, 0xee, 0x56,
	0x0a, 0x30, 0x33, 0x69, 0xce, 0xef, 0xb6, 0x32, 0xf3, 0xf8, 0x33, 0xa8, 0x62, 0xb6, 0x36, 0x9d,
	0x9d, 0x5b, 0xdc, 0x1d, 0x40, 0x3a, 0x29, 0x6a, 0x5b, 0x8d, 0x87, 0xdf, 0xec, 0xb7, 0x9d, 0x2f,
	0x36, 0xfe, 0x68, 0xdd, 0x2c, 0xd2, 0x3d, 0x42, 0xe8, 0x69, 0x0c, 0xc2, 0x79, 0x9c, 0xca, 0xe2,
	0x4a, 0x68, 0x4c, 0xac, 0xdb, 0x84, 0xa4, 0xa4, 0x66, 0x8e, 0xf2, 0xe1, 0xa2, 0xb6, 0x75, 0x13,
	0x37, 0x7d, 0xae, 0x59, 0xa2, 0xf5, 0x7e, 0x7e, 0xef, 0x40, 0xc9, 0x24, 0xd6, 0x1e, 0x93, 0x51,
	0xe3, 0x39, 0xc2, 0x9d, 0x98, 0xb2, 0x10, 0x9c, 0x96, 0xb7, 0xe3, 0x1c, 0x73, 0x06, 0x42, 0x03,
	0xde, 0x9f, 0x48, 0x06, 0xdc, 0x84, 0xc9, 0x20, 0xad, 0x24, 0x79, 0xeb, 0x99, 0x54, 0x01, 0x8d,
	0x40, 0x4f, 0xc1, 0xc8, 0x60, 0x28, 0x07, 0x24, 0xa2, 0xe9, 0xfa, 0xc9, 0xf1, 0xd1, 0x41, 0xf7,
	0x49, 0xbf, 0xdb, 0x2a, 0xed, 0x7a, 0x3b, 0x0d, 0xdb, 0xb2, 0x5b, 0xab, 0x34, 0x8e, 0x87, 0x9c,
	0x65, 0x7f, 0x19, 0xf2, 0x52, 0x4b, 0xd1, 0x9e, 0x3b, 0xe9, 0xfd, 0x8f, 0x4a, 0x7b, 0x3b, 0x7b,
	0x78, 0x0f, 0x35, 0x7a, 0x60, 0x12, 0x25, 0xc0, 0x77, 0x2e, 0x43, 0x10, 0x8e, 0x09, 0xc1, 0x51,
	0xa0, 0x65, 0xa2, 0x18, 0x38, 0xbe, 0x04, 0xed, 0x08, 0x69, 0x1c, 0xb8, 0xe2, 0xda, 0x78, 0x78,
	0x11, 0x2d, 0xbc, 0xb7, 0xad, 0x25, 0xf5, 0x00, 0x55, 0x6f, 0x1e, 0xc3, 0x39, 0x94, 0x2c, 0x89,
	0x40, 0xe4, 0x7f, 0x35, 0xbc, 0x79, 0xfb, 0xd3, 0x10, 0xcd, 0x0d, 0x10, 0x5f, 0x32, 0x4d, 0x5e,
	0xfc, 0x21, 0x74, 0x73, 0xea, 0x2e, 0xf1, 0x79, 0x40, 0xe2, 0xc1, 0x07, 0xbb, 0x9c, 0x6a, 0x66,
	0x92, 0x83, 0xc5, 0xcc, 0x25, 0xff, 0x7d, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x30, 0x65, 0x3c, 0xe8,
	0x87, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DirectorClient is the client API for Director service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DirectorClient interface {
	RegisterArena(ctx context.Context, in *SynRegisterArena, opts ...grpc.CallOption) (*empty.Empty, error)
	HealthCheckArena(ctx context.Context, in *SynHealthCheckArena, opts ...grpc.CallOption) (*empty.Empty, error)
	TerminateArena(ctx context.Context, in *SynTerminateArena, opts ...grpc.CallOption) (*empty.Empty, error)
	AvailableArena(ctx context.Context, in *SynAvailableArena, opts ...grpc.CallOption) (*AckAvailableArena, error)
}

type directorClient struct {
	cc *grpc.ClientConn
}

func NewDirectorClient(cc *grpc.ClientConn) DirectorClient {
	return &directorClient{cc}
}

func (c *directorClient) RegisterArena(ctx context.Context, in *SynRegisterArena, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openmatch.Director/RegisterArena", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) HealthCheckArena(ctx context.Context, in *SynHealthCheckArena, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openmatch.Director/HealthCheckArena", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) TerminateArena(ctx context.Context, in *SynTerminateArena, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openmatch.Director/TerminateArena", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) AvailableArena(ctx context.Context, in *SynAvailableArena, opts ...grpc.CallOption) (*AckAvailableArena, error) {
	out := new(AckAvailableArena)
	err := c.cc.Invoke(ctx, "/openmatch.Director/AvailableArena", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectorServer is the server API for Director service.
type DirectorServer interface {
	RegisterArena(context.Context, *SynRegisterArena) (*empty.Empty, error)
	HealthCheckArena(context.Context, *SynHealthCheckArena) (*empty.Empty, error)
	TerminateArena(context.Context, *SynTerminateArena) (*empty.Empty, error)
	AvailableArena(context.Context, *SynAvailableArena) (*AckAvailableArena, error)
}

// UnimplementedDirectorServer can be embedded to have forward compatible implementations.
type UnimplementedDirectorServer struct {
}

func (*UnimplementedDirectorServer) RegisterArena(ctx context.Context, req *SynRegisterArena) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterArena not implemented")
}
func (*UnimplementedDirectorServer) HealthCheckArena(ctx context.Context, req *SynHealthCheckArena) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheckArena not implemented")
}
func (*UnimplementedDirectorServer) TerminateArena(ctx context.Context, req *SynTerminateArena) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateArena not implemented")
}
func (*UnimplementedDirectorServer) AvailableArena(ctx context.Context, req *SynAvailableArena) (*AckAvailableArena, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailableArena not implemented")
}

func RegisterDirectorServer(s *grpc.Server, srv DirectorServer) {
	s.RegisterService(&_Director_serviceDesc, srv)
}

func _Director_RegisterArena_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynRegisterArena)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).RegisterArena(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.Director/RegisterArena",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).RegisterArena(ctx, req.(*SynRegisterArena))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_HealthCheckArena_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynHealthCheckArena)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).HealthCheckArena(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.Director/HealthCheckArena",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).HealthCheckArena(ctx, req.(*SynHealthCheckArena))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_TerminateArena_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynTerminateArena)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).TerminateArena(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.Director/TerminateArena",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).TerminateArena(ctx, req.(*SynTerminateArena))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_AvailableArena_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynAvailableArena)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).AvailableArena(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.Director/AvailableArena",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).AvailableArena(ctx, req.(*SynAvailableArena))
	}
	return interceptor(ctx, in, info, handler)
}

var _Director_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.Director",
	HandlerType: (*DirectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterArena",
			Handler:    _Director_RegisterArena_Handler,
		},
		{
			MethodName: "HealthCheckArena",
			Handler:    _Director_HealthCheckArena_Handler,
		},
		{
			MethodName: "TerminateArena",
			Handler:    _Director_TerminateArena_Handler,
		},
		{
			MethodName: "AvailableArena",
			Handler:    _Director_AvailableArena_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/director.proto",
}