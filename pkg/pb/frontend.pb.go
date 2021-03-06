// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/frontend.proto

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

type CreateTicketRequest struct {
	// A Ticket object with SearchFields defined.
	Ticket               *Ticket  `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTicketRequest) Reset()         { *m = CreateTicketRequest{} }
func (m *CreateTicketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTicketRequest) ProtoMessage()    {}
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{0}
}

func (m *CreateTicketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTicketRequest.Unmarshal(m, b)
}
func (m *CreateTicketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTicketRequest.Marshal(b, m, deterministic)
}
func (m *CreateTicketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTicketRequest.Merge(m, src)
}
func (m *CreateTicketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTicketRequest.Size(m)
}
func (m *CreateTicketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTicketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTicketRequest proto.InternalMessageInfo

func (m *CreateTicketRequest) GetTicket() *Ticket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type DeleteTicketRequest struct {
	// A TicketId of a generated Ticket to be deleted.
	TicketId             string   `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTicketRequest) Reset()         { *m = DeleteTicketRequest{} }
func (m *DeleteTicketRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTicketRequest) ProtoMessage()    {}
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{1}
}

func (m *DeleteTicketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTicketRequest.Unmarshal(m, b)
}
func (m *DeleteTicketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTicketRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTicketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTicketRequest.Merge(m, src)
}
func (m *DeleteTicketRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTicketRequest.Size(m)
}
func (m *DeleteTicketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTicketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTicketRequest proto.InternalMessageInfo

func (m *DeleteTicketRequest) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

type GetTicketRequest struct {
	// A TicketId of a generated Ticket.
	TicketId             string   `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTicketRequest) Reset()         { *m = GetTicketRequest{} }
func (m *GetTicketRequest) String() string { return proto.CompactTextString(m) }
func (*GetTicketRequest) ProtoMessage()    {}
func (*GetTicketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{2}
}

func (m *GetTicketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTicketRequest.Unmarshal(m, b)
}
func (m *GetTicketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTicketRequest.Marshal(b, m, deterministic)
}
func (m *GetTicketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTicketRequest.Merge(m, src)
}
func (m *GetTicketRequest) XXX_Size() int {
	return xxx_messageInfo_GetTicketRequest.Size(m)
}
func (m *GetTicketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTicketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTicketRequest proto.InternalMessageInfo

func (m *GetTicketRequest) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

type WatchAssignmentsRequest struct {
	// A TicketId of a generated Ticket to get updates on.
	TicketId             string   `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchAssignmentsRequest) Reset()         { *m = WatchAssignmentsRequest{} }
func (m *WatchAssignmentsRequest) String() string { return proto.CompactTextString(m) }
func (*WatchAssignmentsRequest) ProtoMessage()    {}
func (*WatchAssignmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{3}
}

func (m *WatchAssignmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchAssignmentsRequest.Unmarshal(m, b)
}
func (m *WatchAssignmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchAssignmentsRequest.Marshal(b, m, deterministic)
}
func (m *WatchAssignmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchAssignmentsRequest.Merge(m, src)
}
func (m *WatchAssignmentsRequest) XXX_Size() int {
	return xxx_messageInfo_WatchAssignmentsRequest.Size(m)
}
func (m *WatchAssignmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchAssignmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchAssignmentsRequest proto.InternalMessageInfo

func (m *WatchAssignmentsRequest) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

type WatchAssignmentsResponse struct {
	// An updated Assignment of the requested Ticket.
	Assignment           *Assignment `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WatchAssignmentsResponse) Reset()         { *m = WatchAssignmentsResponse{} }
func (m *WatchAssignmentsResponse) String() string { return proto.CompactTextString(m) }
func (*WatchAssignmentsResponse) ProtoMessage()    {}
func (*WatchAssignmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06c902cf58d2ae57, []int{4}
}

func (m *WatchAssignmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchAssignmentsResponse.Unmarshal(m, b)
}
func (m *WatchAssignmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchAssignmentsResponse.Marshal(b, m, deterministic)
}
func (m *WatchAssignmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchAssignmentsResponse.Merge(m, src)
}
func (m *WatchAssignmentsResponse) XXX_Size() int {
	return xxx_messageInfo_WatchAssignmentsResponse.Size(m)
}
func (m *WatchAssignmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchAssignmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchAssignmentsResponse proto.InternalMessageInfo

func (m *WatchAssignmentsResponse) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTicketRequest)(nil), "openmatch.CreateTicketRequest")
	proto.RegisterType((*DeleteTicketRequest)(nil), "openmatch.DeleteTicketRequest")
	proto.RegisterType((*GetTicketRequest)(nil), "openmatch.GetTicketRequest")
	proto.RegisterType((*WatchAssignmentsRequest)(nil), "openmatch.WatchAssignmentsRequest")
	proto.RegisterType((*WatchAssignmentsResponse)(nil), "openmatch.WatchAssignmentsResponse")
}

func init() { proto.RegisterFile("api/frontend.proto", fileDescriptor_06c902cf58d2ae57) }

var fileDescriptor_06c902cf58d2ae57 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x4e, 0x13, 0x41,
	0x14, 0xcd, 0x2e, 0x04, 0xe9, 0x48, 0x22, 0x0e, 0x8a, 0xa4, 0x35, 0x66, 0x5d, 0x12, 0x95, 0xc6,
	0xee, 0x94, 0x02, 0x3e, 0x40, 0x4c, 0x40, 0x40, 0x43, 0x82, 0x1a, 0x8b, 0xd1, 0xc4, 0x17, 0xb3,
	0xdd, 0xbd, 0x6c, 0x47, 0xba, 0x33, 0xe3, 0xde, 0x59, 0xd0, 0x18, 0x13, 0xe3, 0x27, 0xe8, 0x1b,
	0x9f, 0xe0, 0xa3, 0xbf, 0xe2, 0x93, 0xef, 0xfe, 0x87, 0x66, 0x67, 0xb7, 0xed, 0x0a, 0x95, 0xe0,
	0x53, 0xb3, 0x73, 0xcf, 0x3d, 0xe7, 0xde, 0x73, 0x6e, 0x4a, 0xa8, 0xaf, 0x38, 0xdb, 0x4f, 0xa4,
	0xd0, 0x20, 0x42, 0x4f, 0x25, 0x52, 0x4b, 0x5a, 0x91, 0x0a, 0x44, 0xec, 0xeb, 0xa0, 0x5b, 0x35,
	0xe5, 0x18, 0x10, 0xfd, 0x08, 0x30, 0x2f, 0x57, 0xaf, 0x47, 0x52, 0x46, 0x3d, 0x60, 0x59, 0xc9,
	0x17, 0x42, 0x6a, 0x5f, 0x73, 0x29, 0xfa, 0xd5, 0xbb, 0xe6, 0x27, 0x68, 0x44, 0x20, 0x1a, 0x78,
	0xe4, 0x47, 0x11, 0x24, 0x4c, 0x2a, 0x83, 0x18, 0x81, 0xae, 0x15, 0x5c, 0xe6, 0xab, 0x93, 0xee,
	0x33, 0x88, 0x95, 0x7e, 0x9f, 0x17, 0xdd, 0x75, 0x32, 0xb3, 0x99, 0x80, 0xaf, 0xe1, 0x39, 0x0f,
	0x0e, 0x40, 0xb7, 0xe1, 0x6d, 0x0a, 0xa8, 0xe9, 0x02, 0x99, 0xd0, 0xe6, 0x61, 0xce, 0x72, 0xac,
	0x3b, 0x17, 0x5b, 0x97, 0xbd, 0xc1, 0xbc, 0x5e, 0x81, 0x2c, 0x00, 0x6e, 0x8b, 0xcc, 0x6c, 0x41,
	0x0f, 0x4e, 0x32, 0xd4, 0x48, 0x25, 0x07, 0xbc, 0xe6, 0xa1, 0x21, 0xa9, 0xb4, 0x27, 0xf3, 0x87,
	0x9d, 0xd0, 0x65, 0x64, 0xfa, 0x11, 0xe8, 0xff, 0x68, 0xb8, 0x47, 0xae, 0xbd, 0xcc, 0xc4, 0x37,
	0x10, 0x79, 0x24, 0x62, 0x10, 0x1a, 0xcf, 0xd5, 0xf7, 0x8c, 0xcc, 0x9d, 0xee, 0x43, 0x25, 0x05,
	0x02, 0x5d, 0x21, 0xc4, 0x1f, 0x3c, 0x17, 0x7b, 0x5e, 0x2d, 0xed, 0x39, 0xec, 0x69, 0x97, 0x80,
	0xad, 0x4f, 0xe3, 0xe4, 0xd2, 0xc3, 0x22, 0xcc, 0x3d, 0x48, 0x0e, 0x79, 0x00, 0x94, 0x93, 0xa9,
	0xb2, 0x8b, 0xf4, 0x46, 0x89, 0x66, 0x84, 0xbd, 0xd5, 0xd3, 0x76, 0xba, 0xb7, 0x3e, 0xff, 0xf8,
	0xf5, 0xd5, 0x76, 0xdc, 0x1a, 0x3b, 0x5c, 0x1c, 0x1c, 0x0b, 0xe6, 0xfc, 0x2c, 0xdf, 0x07, 0x57,
	0xad, 0x3a, 0x3d, 0x22, 0x53, 0x65, 0xbb, 0xff, 0x92, 0x1a, 0x91, 0x43, 0x75, 0xd6, 0xcb, 0xe3,
	0xf7, 0xfa, 0xf1, 0x7b, 0xdb, 0x59, 0xfc, 0x2e, 0x33, 0x7a, 0x0b, 0xf5, 0xdb, 0x67, 0xe8, 0xb1,
	0x0f, 0x03, 0x67, 0x3f, 0xd2, 0x1e, 0xa9, 0x0c, 0x32, 0xa3, 0xb5, 0x92, 0xea, 0xc9, 0x24, 0x47,
	0x6d, 0x57, 0xa8, 0xd1, 0x73, 0xab, 0x1d, 0x5b, 0x64, 0xfa, 0x64, 0x72, 0xd4, 0x2d, 0x11, 0xff,
	0xe3, 0x1c, 0xaa, 0xf3, 0x67, 0x62, 0xf2, 0xe8, 0xdd, 0x35, 0x33, 0xce, 0x0a, 0x5d, 0x3a, 0xe7,
	0x38, 0x6c, 0x98, 0x3f, 0x36, 0xad, 0x07, 0xbf, 0xed, 0x2f, 0x1b, 0x3f, 0x6d, 0xfa, 0xdd, 0x22,
	0x93, 0xfd, 0x4b, 0x70, 0x77, 0x08, 0x79, 0xaa, 0x40, 0x38, 0x8f, 0x33, 0x4d, 0x3a, 0xdb, 0xd5,
	0x5a, 0xe1, 0x2a, 0x63, 0xd9, 0x30, 0x8d, 0x7c, 0x9a, 0x10, 0x0e, 0xab, 0xf3, 0xc3, 0xef, 0x46,
	0xc8, 0x31, 0x48, 0x11, 0xd7, 0xf3, 0x7c, 0xa2, 0x44, 0xa6, 0x0a, 0xbd, 0x40, 0xc6, 0xf5, 0x17,
	0x84, 0x6e, 0x28, 0x3f, 0xe8, 0x82, 0xd3, 0xf2, 0x9a, 0xce, 0x2e, 0x0f, 0x20, 0x3b, 0xd7, 0xf5,
	0x3e, 0x65, 0xc4, 0x75, 0x37, 0xed, 0x64, 0x48, 0x96, 0xb7, 0xee, 0xcb, 0x24, 0xf2, 0x63, 0xc0,
	0x92, 0x18, 0xeb, 0xf4, 0x64, 0x87, 0xc5, 0x3e, 0x6a, 0x48, 0xd8, 0xee, 0xce, 0xe6, 0xf6, 0x93,
	0xbd, 0xed, 0xd6, 0xd8, 0xa2, 0xd7, 0xac, 0xdb, 0x96, 0xdd, 0x9a, 0xf6, 0x95, 0xea, 0xf1, 0xc0,
	0xfc, 0x4b, 0xb0, 0x37, 0x28, 0xc5, 0xea, 0xa9, 0x97, 0xf6, 0x1a, 0x19, 0x5b, 0x6e, 0x2e, 0xd3,
	0x65, 0x52, 0x6f, 0x83, 0x4e, 0x13, 0x01, 0xa1, 0x73, 0xd4, 0x05, 0xe1, 0xe8, 0x2e, 0x38, 0x09,
	0xa0, 0x4c, 0x93, 0x00, 0x9c, 0x50, 0x02, 0x3a, 0x42, 0x6a, 0x07, 0xde, 0x71, 0xd4, 0x1e, 0x9d,
	0x20, 0xe3, 0xc7, 0xb6, 0x75, 0x21, 0xb9, 0x4f, 0xe6, 0x86, 0x66, 0x38, 0x5b, 0x32, 0x48, 0x33,
	0xeb, 0x0c, 0x3b, 0xbd, 0x39, 0xda, 0x1a, 0x86, 0x5c, 0x03, 0x0b, 0x65, 0x80, 0xec, 0xd5, 0x15,
	0x81, 0x8d, 0xd2, 0x2e, 0xea, 0x20, 0x62, 0xaa, 0xf3, 0xcd, 0xae, 0x64, 0x9c, 0x86, 0xb2, 0x33,
	0x61, 0xae, 0x79, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x7a, 0xba, 0x80, 0x5b, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendServiceClient is the client API for FrontendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendServiceClient interface {
	// CreateTicket assigns an unique TicketId to the input Ticket and record it in state storage.
	// A ticket is considered as ready for matchmaking once it is created.
	//   - If a TicketId exists in a Ticket request, an auto-generated TicketId will override this field.
	//   - If SearchFields exist in a Ticket, CreateTicket will also index these fields such that one can query the ticket with query.QueryTickets function.
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	// DeleteTicket immediately stops Open Match from using the Ticket for matchmaking and removes the Ticket from state storage.
	// The client should delete the Ticket when finished matchmaking with it.
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetTicket get the Ticket associated with the specified TicketId.
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	// WatchAssignments stream back Assignment of the specified TicketId if it is updated.
	//   - If the Assignment is not updated, GetAssignment will retry using the configured backoff strategy.
	WatchAssignments(ctx context.Context, in *WatchAssignmentsRequest, opts ...grpc.CallOption) (FrontendService_WatchAssignmentsClient, error)
}

type frontendServiceClient struct {
	cc *grpc.ClientConn
}

func NewFrontendServiceClient(cc *grpc.ClientConn) FrontendServiceClient {
	return &frontendServiceClient{cc}
}

func (c *frontendServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/openmatch.FrontendService/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openmatch.FrontendService/DeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/openmatch.FrontendService/GetTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendServiceClient) WatchAssignments(ctx context.Context, in *WatchAssignmentsRequest, opts ...grpc.CallOption) (FrontendService_WatchAssignmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FrontendService_serviceDesc.Streams[0], "/openmatch.FrontendService/WatchAssignments", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendServiceWatchAssignmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FrontendService_WatchAssignmentsClient interface {
	Recv() (*WatchAssignmentsResponse, error)
	grpc.ClientStream
}

type frontendServiceWatchAssignmentsClient struct {
	grpc.ClientStream
}

func (x *frontendServiceWatchAssignmentsClient) Recv() (*WatchAssignmentsResponse, error) {
	m := new(WatchAssignmentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontendServiceServer is the server API for FrontendService service.
type FrontendServiceServer interface {
	// CreateTicket assigns an unique TicketId to the input Ticket and record it in state storage.
	// A ticket is considered as ready for matchmaking once it is created.
	//   - If a TicketId exists in a Ticket request, an auto-generated TicketId will override this field.
	//   - If SearchFields exist in a Ticket, CreateTicket will also index these fields such that one can query the ticket with query.QueryTickets function.
	CreateTicket(context.Context, *CreateTicketRequest) (*Ticket, error)
	// DeleteTicket immediately stops Open Match from using the Ticket for matchmaking and removes the Ticket from state storage.
	// The client should delete the Ticket when finished matchmaking with it.
	DeleteTicket(context.Context, *DeleteTicketRequest) (*empty.Empty, error)
	// GetTicket get the Ticket associated with the specified TicketId.
	GetTicket(context.Context, *GetTicketRequest) (*Ticket, error)
	// WatchAssignments stream back Assignment of the specified TicketId if it is updated.
	//   - If the Assignment is not updated, GetAssignment will retry using the configured backoff strategy.
	WatchAssignments(*WatchAssignmentsRequest, FrontendService_WatchAssignmentsServer) error
}

// UnimplementedFrontendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendServiceServer struct {
}

func (*UnimplementedFrontendServiceServer) CreateTicket(ctx context.Context, req *CreateTicketRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (*UnimplementedFrontendServiceServer) DeleteTicket(ctx context.Context, req *DeleteTicketRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (*UnimplementedFrontendServiceServer) GetTicket(ctx context.Context, req *GetTicketRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (*UnimplementedFrontendServiceServer) WatchAssignments(req *WatchAssignmentsRequest, srv FrontendService_WatchAssignmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAssignments not implemented")
}

func RegisterFrontendServiceServer(s *grpc.Server, srv FrontendServiceServer) {
	s.RegisterService(&_FrontendService_serviceDesc, srv)
}

func _FrontendService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.FrontendService/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontendService_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServiceServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.FrontendService/DeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServiceServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontendService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.FrontendService/GetTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServiceServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontendService_WatchAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchAssignmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FrontendServiceServer).WatchAssignments(m, &frontendServiceWatchAssignmentsServer{stream})
}

type FrontendService_WatchAssignmentsServer interface {
	Send(*WatchAssignmentsResponse) error
	grpc.ServerStream
}

type frontendServiceWatchAssignmentsServer struct {
	grpc.ServerStream
}

func (x *frontendServiceWatchAssignmentsServer) Send(m *WatchAssignmentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _FrontendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.FrontendService",
	HandlerType: (*FrontendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _FrontendService_CreateTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _FrontendService_DeleteTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _FrontendService_GetTicket_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchAssignments",
			Handler:       _FrontendService_WatchAssignments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/frontend.proto",
}
