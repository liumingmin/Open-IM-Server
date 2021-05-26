// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relay/relay.proto

package pbRelay

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MsgToUserReq struct {
	SendID               string   `protobuf:"bytes,1,opt,name=SendID,proto3" json:"SendID,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	RecvSeq              int64    `protobuf:"varint,6,opt,name=RecvSeq,proto3" json:"RecvSeq,omitempty"`
	SendTime             int64    `protobuf:"varint,7,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	MsgFrom              int32    `protobuf:"varint,8,opt,name=MsgFrom,proto3" json:"MsgFrom,omitempty"`
	ContentType          int32    `protobuf:"varint,9,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	SessionType          int32    `protobuf:"varint,10,opt,name=SessionType,proto3" json:"SessionType,omitempty"`
	OperationID          string   `protobuf:"bytes,11,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	ServerMsgID          string   `protobuf:"bytes,12,opt,name=ServerMsgID,proto3" json:"ServerMsgID,omitempty"`
	PlatformID           int32    `protobuf:"varint,13,opt,name=PlatformID,proto3" json:"PlatformID,omitempty"`
	IsEmphasize          bool     `protobuf:"varint,14,opt,name=IsEmphasize,proto3" json:"IsEmphasize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgToUserReq) Reset()         { *m = MsgToUserReq{} }
func (m *MsgToUserReq) String() string { return proto.CompactTextString(m) }
func (*MsgToUserReq) ProtoMessage()    {}
func (*MsgToUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6da3b5c0d1535b3, []int{0}
}

func (m *MsgToUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgToUserReq.Unmarshal(m, b)
}
func (m *MsgToUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgToUserReq.Marshal(b, m, deterministic)
}
func (m *MsgToUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgToUserReq.Merge(m, src)
}
func (m *MsgToUserReq) XXX_Size() int {
	return xxx_messageInfo_MsgToUserReq.Size(m)
}
func (m *MsgToUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgToUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgToUserReq proto.InternalMessageInfo

func (m *MsgToUserReq) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *MsgToUserReq) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *MsgToUserReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgToUserReq) GetRecvSeq() int64 {
	if m != nil {
		return m.RecvSeq
	}
	return 0
}

func (m *MsgToUserReq) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *MsgToUserReq) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *MsgToUserReq) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *MsgToUserReq) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *MsgToUserReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *MsgToUserReq) GetServerMsgID() string {
	if m != nil {
		return m.ServerMsgID
	}
	return ""
}

func (m *MsgToUserReq) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *MsgToUserReq) GetIsEmphasize() bool {
	if m != nil {
		return m.IsEmphasize
	}
	return false
}

type MsgToUserResp struct {
	Resp                 []*SingleMsgToUser `protobuf:"bytes,1,rep,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MsgToUserResp) Reset()         { *m = MsgToUserResp{} }
func (m *MsgToUserResp) String() string { return proto.CompactTextString(m) }
func (*MsgToUserResp) ProtoMessage()    {}
func (*MsgToUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6da3b5c0d1535b3, []int{1}
}

func (m *MsgToUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgToUserResp.Unmarshal(m, b)
}
func (m *MsgToUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgToUserResp.Marshal(b, m, deterministic)
}
func (m *MsgToUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgToUserResp.Merge(m, src)
}
func (m *MsgToUserResp) XXX_Size() int {
	return xxx_messageInfo_MsgToUserResp.Size(m)
}
func (m *MsgToUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgToUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgToUserResp proto.InternalMessageInfo

func (m *MsgToUserResp) GetResp() []*SingleMsgToUser {
	if m != nil {
		return m.Resp
	}
	return nil
}

//message SendMsgByWSReq{
//  string SendID = 1;
//  string RecvID = 2;
//  string Content = 3;
//  int64 SendTime = 4;
//  int64  MsgFrom = 5;
//  int64  ContentType = 6;
//  int64  SessionType = 7;
//  string OperationID = 8;
//  int64  PlatformID = 9;
//}
type SingleMsgToUser struct {
	ResultCode           int64    `protobuf:"varint,1,opt,name=ResultCode,proto3" json:"ResultCode,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	RecvPlatFormID       int32    `protobuf:"varint,3,opt,name=RecvPlatFormID,proto3" json:"RecvPlatFormID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleMsgToUser) Reset()         { *m = SingleMsgToUser{} }
func (m *SingleMsgToUser) String() string { return proto.CompactTextString(m) }
func (*SingleMsgToUser) ProtoMessage()    {}
func (*SingleMsgToUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6da3b5c0d1535b3, []int{2}
}

func (m *SingleMsgToUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleMsgToUser.Unmarshal(m, b)
}
func (m *SingleMsgToUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleMsgToUser.Marshal(b, m, deterministic)
}
func (m *SingleMsgToUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleMsgToUser.Merge(m, src)
}
func (m *SingleMsgToUser) XXX_Size() int {
	return xxx_messageInfo_SingleMsgToUser.Size(m)
}
func (m *SingleMsgToUser) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleMsgToUser.DiscardUnknown(m)
}

var xxx_messageInfo_SingleMsgToUser proto.InternalMessageInfo

func (m *SingleMsgToUser) GetResultCode() int64 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *SingleMsgToUser) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *SingleMsgToUser) GetRecvPlatFormID() int32 {
	if m != nil {
		return m.RecvPlatFormID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgToUserReq)(nil), "relay.MsgToUserReq")
	proto.RegisterType((*MsgToUserResp)(nil), "relay.MsgToUserResp")
	proto.RegisterType((*SingleMsgToUser)(nil), "relay.SingleMsgToUser")
}

func init() { proto.RegisterFile("relay/relay.proto", fileDescriptor_b6da3b5c0d1535b3) }

var fileDescriptor_b6da3b5c0d1535b3 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcb, 0x6a, 0xe3, 0x30,
	0x18, 0x85, 0x71, 0x3c, 0xb9, 0x29, 0x37, 0x46, 0x33, 0x04, 0x4d, 0x16, 0x83, 0xc9, 0xa2, 0x98,
	0x2e, 0x52, 0x48, 0xa1, 0x9b, 0xec, 0x1a, 0x37, 0xe0, 0x85, 0x49, 0x91, 0xd3, 0x4d, 0x77, 0x4e,
	0xf2, 0xd7, 0x35, 0xd8, 0x92, 0x22, 0xb9, 0x81, 0xf4, 0xb1, 0xfb, 0x04, 0x45, 0x72, 0xd2, 0x88,
	0x94, 0x6e, 0x8c, 0xcf, 0xf7, 0x1f, 0x1f, 0x1f, 0x5d, 0xd0, 0x6f, 0x09, 0x79, 0x72, 0xb8, 0x31,
	0xcf, 0x89, 0x90, 0xbc, 0xe4, 0xb8, 0x6e, 0xc4, 0xf8, 0xa3, 0x86, 0xba, 0x91, 0x4a, 0x57, 0xfc,
	0x49, 0x81, 0xa4, 0xb0, 0xc3, 0x43, 0xd4, 0x88, 0x81, 0x6d, 0xc3, 0x80, 0x38, 0x9e, 0xe3, 0xb7,
	0xe9, 0x51, 0x69, 0x4e, 0x61, 0xb3, 0x0f, 0x03, 0x52, 0xab, 0x78, 0xa5, 0x30, 0x41, 0xcd, 0x39,
	0x67, 0x25, 0xb0, 0x92, 0xd4, 0xcd, 0xe0, 0x24, 0xf5, 0x44, 0x7b, 0x62, 0xd8, 0x91, 0x86, 0xe7,
	0xf8, 0x2e, 0x3d, 0x49, 0x3c, 0x42, 0x2d, 0x9d, 0xba, 0xca, 0x0a, 0x20, 0x4d, 0x33, 0xfa, 0xd2,
	0xfa, 0xab, 0x48, 0xa5, 0x0b, 0xc9, 0x0b, 0xd2, 0xf2, 0x1c, 0xbf, 0x4e, 0x4f, 0x12, 0x7b, 0xa8,
	0x73, 0x8c, 0x5e, 0x1d, 0x04, 0x90, 0xb6, 0x99, 0xda, 0x48, 0x3b, 0x62, 0x50, 0x2a, 0xe3, 0xcc,
	0x38, 0x50, 0xe5, 0xb0, 0x90, 0x76, 0x2c, 0x05, 0xc8, 0xa4, 0xcc, 0x38, 0x0b, 0x03, 0xd2, 0x31,
	0x8d, 0x6d, 0x54, 0x65, 0xc8, 0x3d, 0xc8, 0x48, 0xa5, 0x61, 0x40, 0xba, 0x95, 0xc3, 0x42, 0xf8,
	0x3f, 0x42, 0x8f, 0x79, 0x52, 0xbe, 0x70, 0x59, 0x84, 0x01, 0xe9, 0x99, 0x9f, 0x58, 0x44, 0x27,
	0x84, 0xea, 0xa1, 0x10, 0xaf, 0x89, 0xca, 0xde, 0x81, 0xf4, 0x3d, 0xc7, 0x6f, 0x51, 0x1b, 0x8d,
	0x67, 0xa8, 0x67, 0xed, 0xb9, 0x12, 0xf8, 0x1a, 0xfd, 0x92, 0xa0, 0x04, 0x71, 0x3c, 0xd7, 0xef,
	0x4c, 0x87, 0x93, 0xea, 0xa0, 0xe2, 0x8c, 0xa5, 0x39, 0x9c, 0x9d, 0xc6, 0x33, 0xde, 0xa1, 0xc1,
	0xc5, 0x40, 0x37, 0xa2, 0xa0, 0xde, 0xf2, 0x72, 0xce, 0xb7, 0x60, 0xce, 0xcd, 0xa5, 0x16, 0xf9,
	0xf1, 0xec, 0xae, 0x50, 0x5f, 0xbf, 0xe9, 0xee, 0x8b, 0x6a, 0x35, 0xae, 0x59, 0xcd, 0x05, 0x9d,
	0xc6, 0xe8, 0xdf, 0x92, 0xe5, 0x19, 0x83, 0x08, 0x94, 0x4a, 0x52, 0xa0, 0xba, 0x9e, 0xde, 0x92,
	0x6c, 0x03, 0xf8, 0x0e, 0xb5, 0xcf, 0x4d, 0xfe, 0x1c, 0xab, 0xdb, 0x57, 0x6a, 0xf4, 0xf7, 0x3b,
	0x54, 0xe2, 0x7e, 0xf0, 0xdc, 0x33, 0x78, 0x26, 0xd6, 0x26, 0x6f, 0xdd, 0x30, 0x17, 0xf3, 0xf6,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x57, 0x9a, 0x71, 0x03, 0xad, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnlineMessageRelayServiceClient is the client API for OnlineMessageRelayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnlineMessageRelayServiceClient interface {
	MsgToUser(ctx context.Context, in *MsgToUserReq, opts ...grpc.CallOption) (*MsgToUserResp, error)
}

type onlineMessageRelayServiceClient struct {
	cc *grpc.ClientConn
}

func NewOnlineMessageRelayServiceClient(cc *grpc.ClientConn) OnlineMessageRelayServiceClient {
	return &onlineMessageRelayServiceClient{cc}
}

func (c *onlineMessageRelayServiceClient) MsgToUser(ctx context.Context, in *MsgToUserReq, opts ...grpc.CallOption) (*MsgToUserResp, error) {
	out := new(MsgToUserResp)
	err := c.cc.Invoke(ctx, "/relay.OnlineMessageRelayService/MsgToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlineMessageRelayServiceServer is the server API for OnlineMessageRelayService service.
type OnlineMessageRelayServiceServer interface {
	MsgToUser(context.Context, *MsgToUserReq) (*MsgToUserResp, error)
}

// UnimplementedOnlineMessageRelayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnlineMessageRelayServiceServer struct {
}

func (*UnimplementedOnlineMessageRelayServiceServer) MsgToUser(ctx context.Context, req *MsgToUserReq) (*MsgToUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgToUser not implemented")
}

func RegisterOnlineMessageRelayServiceServer(s *grpc.Server, srv OnlineMessageRelayServiceServer) {
	s.RegisterService(&_OnlineMessageRelayService_serviceDesc, srv)
}

func _OnlineMessageRelayService_MsgToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgToUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineMessageRelayServiceServer).MsgToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relay.OnlineMessageRelayService/MsgToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineMessageRelayServiceServer).MsgToUser(ctx, req.(*MsgToUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnlineMessageRelayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "relay.OnlineMessageRelayService",
	HandlerType: (*OnlineMessageRelayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgToUser",
			Handler:    _OnlineMessageRelayService_MsgToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relay/relay.proto",
}
