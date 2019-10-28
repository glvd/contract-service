// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type WorkMode int32

const (
	WorkMode_LocalMode    WorkMode = 0
	WorkMode_RemoteMode   WorkMode = 1
	WorkMode_DownloadMode WorkMode = 2
)

var WorkMode_name = map[int32]string{
	0: "LocalMode",
	1: "RemoteMode",
	2: "DownloadMode",
}

var WorkMode_value = map[string]int32{
	"LocalMode":    0,
	"RemoteMode":   1,
	"DownloadMode": 2,
}

func (x WorkMode) String() string {
	return proto.EnumName(WorkMode_name, int32(x))
}

func (WorkMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{0}
}

type MessageType int32

const (
	MessageType_Add      MessageType = 0
	MessageType_AsyncAdd MessageType = 1
	MessageType_Status   MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "Add",
	1: "AsyncAdd",
	2: "Status",
}

var MessageType_value = map[string]int32{
	"Add":      0,
	"AsyncAdd": 1,
	"Status":   2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{1}
}

type WorkRequest struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=api.pb.MessageType" json:"msg,omitempty"`
	WorkMode             WorkMode    `protobuf:"varint,2,opt,name=work_mode,json=workMode,proto3,enum=api.pb.WorkMode" json:"work_mode,omitempty"`
	VideoPath            []string    `protobuf:"bytes,3,rep,name=video_path,json=videoPath,proto3" json:"video_path,omitempty"`
	PosterPath           string      `protobuf:"bytes,4,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`
	ThumbPath            string      `protobuf:"bytes,5,opt,name=thumb_path,json=thumbPath,proto3" json:"thumb_path,omitempty"`
	SamplePath           []string    `protobuf:"bytes,6,rep,name=sample_path,json=samplePath,proto3" json:"sample_path,omitempty"`
	VideoInfo            string      `protobuf:"bytes,7,opt,name=video_info,json=videoInfo,proto3" json:"video_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WorkRequest) Reset()         { *m = WorkRequest{} }
func (m *WorkRequest) String() string { return proto.CompactTextString(m) }
func (*WorkRequest) ProtoMessage()    {}
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{0}
}

func (m *WorkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkRequest.Unmarshal(m, b)
}
func (m *WorkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkRequest.Marshal(b, m, deterministic)
}
func (m *WorkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkRequest.Merge(m, src)
}
func (m *WorkRequest) XXX_Size() int {
	return xxx_messageInfo_WorkRequest.Size(m)
}
func (m *WorkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkRequest proto.InternalMessageInfo

func (m *WorkRequest) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

func (m *WorkRequest) GetWorkMode() WorkMode {
	if m != nil {
		return m.WorkMode
	}
	return WorkMode_LocalMode
}

func (m *WorkRequest) GetVideoPath() []string {
	if m != nil {
		return m.VideoPath
	}
	return nil
}

func (m *WorkRequest) GetPosterPath() string {
	if m != nil {
		return m.PosterPath
	}
	return ""
}

func (m *WorkRequest) GetThumbPath() string {
	if m != nil {
		return m.ThumbPath
	}
	return ""
}

func (m *WorkRequest) GetSamplePath() []string {
	if m != nil {
		return m.SamplePath
	}
	return nil
}

func (m *WorkRequest) GetVideoInfo() string {
	if m != nil {
		return m.VideoInfo
	}
	return ""
}

type WorkReply struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=api.pb.MessageType" json:"msg,omitempty"`
	WorkId               string      `protobuf:"bytes,2,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
	Error                string      `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WorkReply) Reset()         { *m = WorkReply{} }
func (m *WorkReply) String() string { return proto.CompactTextString(m) }
func (*WorkReply) ProtoMessage()    {}
func (*WorkReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{1}
}

func (m *WorkReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkReply.Unmarshal(m, b)
}
func (m *WorkReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkReply.Marshal(b, m, deterministic)
}
func (m *WorkReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkReply.Merge(m, src)
}
func (m *WorkReply) XXX_Size() int {
	return xxx_messageInfo_WorkReply.Size(m)
}
func (m *WorkReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkReply.DiscardUnknown(m)
}

var xxx_messageInfo_WorkReply proto.InternalMessageInfo

func (m *WorkReply) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

func (m *WorkReply) GetWorkId() string {
	if m != nil {
		return m.WorkId
	}
	return ""
}

func (m *WorkReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type NodeRequest struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=api.pb.MessageType" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NodeRequest) Reset()         { *m = NodeRequest{} }
func (m *NodeRequest) String() string { return proto.CompactTextString(m) }
func (*NodeRequest) ProtoMessage()    {}
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{2}
}

func (m *NodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRequest.Unmarshal(m, b)
}
func (m *NodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRequest.Marshal(b, m, deterministic)
}
func (m *NodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRequest.Merge(m, src)
}
func (m *NodeRequest) XXX_Size() int {
	return xxx_messageInfo_NodeRequest.Size(m)
}
func (m *NodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRequest proto.InternalMessageInfo

func (m *NodeRequest) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

type NodeReply struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=api.pb.MessageType" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NodeReply) Reset()         { *m = NodeReply{} }
func (m *NodeReply) String() string { return proto.CompactTextString(m) }
func (*NodeReply) ProtoMessage()    {}
func (*NodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{3}
}

func (m *NodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeReply.Unmarshal(m, b)
}
func (m *NodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeReply.Marshal(b, m, deterministic)
}
func (m *NodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeReply.Merge(m, src)
}
func (m *NodeReply) XXX_Size() int {
	return xxx_messageInfo_NodeReply.Size(m)
}
func (m *NodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_NodeReply proto.InternalMessageInfo

func (m *NodeReply) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

func init() {
	proto.RegisterEnum("api.pb.WorkMode", WorkMode_name, WorkMode_value)
	proto.RegisterEnum("api.pb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*WorkRequest)(nil), "api.pb.WorkRequest")
	proto.RegisterType((*WorkReply)(nil), "api.pb.WorkReply")
	proto.RegisterType((*NodeRequest)(nil), "api.pb.NodeRequest")
	proto.RegisterType((*NodeReply)(nil), "api.pb.NodeReply")
}

func init() { proto.RegisterFile("api/pb/service.proto", fileDescriptor_785183a1bcf95a49) }

var fileDescriptor_785183a1bcf95a49 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xf6, 0x8f, 0xd6, 0x8e, 0x9e, 0xbb, 0xe2, 0xa9, 0x85, 0x99, 0x41, 0x59, 0x30, 0x0c, 0x42,
	0x61, 0x6e, 0xc9, 0x76, 0xdb, 0xa9, 0xb0, 0x4b, 0x61, 0x1d, 0xc3, 0x19, 0xec, 0x18, 0xe4, 0x48,
	0x49, 0x4c, 0x6c, 0x4b, 0x93, 0x94, 0x04, 0x5f, 0xf7, 0x97, 0x0f, 0x49, 0xf1, 0x92, 0xdd, 0xb2,
	0x9b, 0xbf, 0xef, 0x7b, 0xdf, 0x7b, 0x7e, 0x9f, 0x1e, 0xdc, 0x12, 0x51, 0x3f, 0x88, 0xea, 0x41,
	0x31, 0xb9, 0xab, 0x17, 0xac, 0x10, 0x92, 0x6b, 0x8e, 0x23, 0x22, 0xea, 0x42, 0x54, 0xf9, 0xef,
	0x00, 0x92, 0x9f, 0x5c, 0x6e, 0x4a, 0xf6, 0x6b, 0xcb, 0x94, 0xc6, 0xef, 0x21, 0x6c, 0xd5, 0x2a,
	0xf3, 0xc7, 0xfe, 0xe4, 0x7a, 0x7a, 0x53, 0xb8, 0xaa, 0xe2, 0x85, 0x29, 0x45, 0x56, 0xec, 0x47,
	0x2f, 0x58, 0x69, 0x74, 0xfc, 0x01, 0xd0, 0x9e, 0xcb, 0xcd, 0xbc, 0xe5, 0x94, 0x65, 0x81, 0x2d,
	0x4e, 0x87, 0x62, 0xd3, 0xee, 0x85, 0x53, 0x56, 0x8e, 0xf6, 0x87, 0x2f, 0x7c, 0x07, 0xb0, 0xab,
	0x29, 0xe3, 0x73, 0x41, 0xf4, 0x3a, 0x0b, 0xc7, 0xe1, 0x04, 0x95, 0xc8, 0x32, 0xdf, 0x89, 0x5e,
	0xe3, 0x77, 0x90, 0x08, 0xae, 0x34, 0x93, 0x4e, 0xbf, 0x18, 0xfb, 0x13, 0x54, 0x82, 0xa3, 0x6c,
	0xc1, 0x1d, 0x80, 0x5e, 0x6f, 0xdb, 0xca, 0xe9, 0x97, 0x56, 0x47, 0x96, 0x19, 0xfc, 0x8a, 0xb4,
	0xa2, 0x61, 0x4e, 0x8f, 0x6c, 0x7f, 0x70, 0xd4, 0xe0, 0x77, 0xf3, 0xeb, 0x6e, 0xc9, 0xb3, 0xd8,
	0xf9, 0x2d, 0xf3, 0xdc, 0x2d, 0x79, 0x4e, 0x00, 0xb9, 0x0c, 0x44, 0xd3, 0x9f, 0x9b, 0xc0, 0x1b,
	0x88, 0x6d, 0x02, 0x35, 0xb5, 0xfb, 0xa3, 0x32, 0x32, 0xf0, 0x99, 0xe2, 0x5b, 0xb8, 0x64, 0x52,
	0x72, 0x99, 0x85, 0x96, 0x76, 0x20, 0xff, 0x04, 0xc9, 0x37, 0x93, 0xc9, 0x7f, 0xc5, 0x9c, 0x4f,
	0x01, 0x39, 0xd7, 0xf9, 0x3f, 0x76, 0xff, 0x19, 0x46, 0xc3, 0x0b, 0xe0, 0x57, 0x80, 0xbe, 0xf2,
	0x05, 0x69, 0x0c, 0x48, 0x3d, 0x7c, 0x0d, 0x50, 0xb2, 0x96, 0x6b, 0x66, 0xb1, 0x8f, 0x53, 0xb8,
	0xfa, 0xc2, 0xf7, 0x5d, 0xc3, 0x09, 0xb5, 0x4c, 0x70, 0xff, 0x08, 0xc9, 0x49, 0x43, 0x1c, 0x43,
	0xf8, 0x44, 0x69, 0xea, 0xe1, 0x2b, 0x18, 0x3d, 0xa9, 0xbe, 0x5b, 0x18, 0xe4, 0x63, 0x80, 0x68,
	0xa6, 0x89, 0xde, 0xaa, 0x34, 0x98, 0xb6, 0x10, 0xcf, 0xdc, 0x65, 0xe1, 0x47, 0xb8, 0x30, 0x93,
	0xf1, 0xcd, 0xe9, 0x25, 0x1c, 0x36, 0x7e, 0xfb, 0xfa, 0x5f, 0x52, 0x34, 0x7d, 0xee, 0x19, 0x87,
	0xd9, 0xef, 0xe8, 0x38, 0xc9, 0xe8, 0xe8, 0xf8, 0x1b, 0x41, 0xee, 0x55, 0x91, 0x3d, 0xdf, 0x8f,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xca, 0x0b, 0x60, 0x22, 0xd6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Work(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error)
	//    rpc AsyncAddWork (WorkRequest) returns (WorkReply) {
	//    }
	//    rpc WorkStatus (WorkRequest) returns (WorkReply) {
	//    }
	Node(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeReply, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Work(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error) {
	out := new(WorkReply)
	err := c.cc.Invoke(ctx, "/api.pb.Service/Work", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Node(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeReply, error) {
	out := new(NodeReply)
	err := c.cc.Invoke(ctx, "/api.pb.Service/Node", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Work(context.Context, *WorkRequest) (*WorkReply, error)
	//    rpc AsyncAddWork (WorkRequest) returns (WorkReply) {
	//    }
	//    rpc WorkStatus (WorkRequest) returns (WorkReply) {
	//    }
	Node(context.Context, *NodeRequest) (*NodeReply, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pb.Service/Work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Work(ctx, req.(*WorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Node_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Node(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pb.Service/Node",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Node(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.pb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Work",
			Handler:    _Service_Work_Handler,
		},
		{
			MethodName: "Node",
			Handler:    _Service_Node_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pb/service.proto",
}
