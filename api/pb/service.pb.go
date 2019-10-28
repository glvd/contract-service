// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// WorkMode ...
type WorkMode int32

// WorkMode_LocalMode ...
const (
	WorkMode_LocalMode    WorkMode = 0
	WorkMode_RemoteMode   WorkMode = 1
	WorkMode_DownloadMode WorkMode = 2
)

// WorkMode_name ...
var WorkMode_name = map[int32]string{
	0: "LocalMode",
	1: "RemoteMode",
	2: "DownloadMode",
}

// WorkMode_value ...
var WorkMode_value = map[string]int32{
	"LocalMode":    0,
	"RemoteMode":   1,
	"DownloadMode": 2,
}

// String ...
func (x WorkMode) String() string {
	return proto.EnumName(WorkMode_name, int32(x))
}

// EnumDescriptor ...
func (WorkMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{0}
}

// WorkStatus ...
type WorkStatus int32

// WorkStatus_WorkAbnormal ...
const (
	WorkStatus_WorkAbnormal WorkStatus = 0
	WorkStatus_WorkWaiting  WorkStatus = 1
	WorkStatus_WorkRunning  WorkStatus = 2
	WorkStatus_WorkStopped  WorkStatus = 3
	WorkStatus_WorkFinish   WorkStatus = 4
)

// WorkStatus_name ...
var WorkStatus_name = map[int32]string{
	0: "WorkAbnormal",
	1: "WorkWaiting",
	2: "WorkRunning",
	3: "WorkStopped",
	4: "WorkFinish",
}

// WorkStatus_value ...
var WorkStatus_value = map[string]int32{
	"WorkAbnormal": 0,
	"WorkWaiting":  1,
	"WorkRunning":  2,
	"WorkStopped":  3,
	"WorkFinish":   4,
}

// String ...
func (x WorkStatus) String() string {
	return proto.EnumName(WorkStatus_name, int32(x))
}

// EnumDescriptor ...
func (WorkStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{1}
}

// MessageType ...
type MessageType int32

// MessageType_Add ...
const (
	MessageType_Add      MessageType = 0
	MessageType_AsyncAdd MessageType = 1
	MessageType_Status   MessageType = 2
)

// MessageType_name ...
var MessageType_name = map[int32]string{
	0: "Add",
	1: "AsyncAdd",
	2: "Status",
}

// MessageType_value ...
var MessageType_value = map[string]int32{
	"Add":      0,
	"AsyncAdd": 1,
	"Status":   2,
}

// String ...
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

// EnumDescriptor ...
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{2}
}

// Work ...
type Work struct {
	VideoPath            []string `protobuf:"bytes,4,rep,name=video_path,json=videoPath,proto3" json:"video_path,omitempty"`
	PosterPath           string   `protobuf:"bytes,5,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`
	ThumbPath            string   `protobuf:"bytes,6,opt,name=thumb_path,json=thumbPath,proto3" json:"thumb_path,omitempty"`
	SamplePath           []string `protobuf:"bytes,7,rep,name=sample_path,json=samplePath,proto3" json:"sample_path,omitempty"`
	VideoInfo            string   `protobuf:"bytes,8,opt,name=video_info,json=videoInfo,proto3" json:"video_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *Work) Reset() { *m = Work{} }

// String ...
func (m *Work) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*Work) ProtoMessage() {}

// Descriptor ...
func (*Work) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{0}
}

// XXX_Unmarshal ...
func (m *Work) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Work.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *Work) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Work.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *Work) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Work.Merge(m, src)
}

// XXX_Size ...
func (m *Work) XXX_Size() int {
	return xxx_messageInfo_Work.Size(m)
}

// XXX_DiscardUnknown ...
func (m *Work) XXX_DiscardUnknown() {
	xxx_messageInfo_Work.DiscardUnknown(m)
}

var xxx_messageInfo_Work proto.InternalMessageInfo

// GetVideoPath ...
func (m *Work) GetVideoPath() []string {
	if m != nil {
		return m.VideoPath
	}
	return nil
}

// GetPosterPath ...
func (m *Work) GetPosterPath() string {
	if m != nil {
		return m.PosterPath
	}
	return ""
}

// GetThumbPath ...
func (m *Work) GetThumbPath() string {
	if m != nil {
		return m.ThumbPath
	}
	return ""
}

// GetSamplePath ...
func (m *Work) GetSamplePath() []string {
	if m != nil {
		return m.SamplePath
	}
	return nil
}

// GetVideoInfo ...
func (m *Work) GetVideoInfo() string {
	if m != nil {
		return m.VideoInfo
	}
	return ""
}

// WorkRequest ...
type WorkRequest struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=pb.MessageType" json:"msg,omitempty"`
	WorkMode             WorkMode    `protobuf:"varint,2,opt,name=work_mode,json=workMode,proto3,enum=pb.WorkMode" json:"work_mode,omitempty"`
	ID                   string      `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	Work                 *Work       `protobuf:"bytes,5,opt,name=work,proto3" json:"work,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

// Reset ...
func (m *WorkRequest) Reset() { *m = WorkRequest{} }

// String ...
func (m *WorkRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*WorkRequest) ProtoMessage() {}

// Descriptor ...
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{1}
}

// XXX_Unmarshal ...
func (m *WorkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *WorkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *WorkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkRequest.Merge(m, src)
}

// XXX_Size ...
func (m *WorkRequest) XXX_Size() int {
	return xxx_messageInfo_WorkRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *WorkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkRequest proto.InternalMessageInfo

// GetMsg ...
func (m *WorkRequest) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

// GetWorkMode ...
func (m *WorkRequest) GetWorkMode() WorkMode {
	if m != nil {
		return m.WorkMode
	}
	return WorkMode_LocalMode
}

// GetID ...
func (m *WorkRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// GetWork ...
func (m *WorkRequest) GetWork() *Work {
	if m != nil {
		return m.Work
	}
	return nil
}

// WorkReply ...
type WorkReply struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=pb.MessageType" json:"msg,omitempty"`
	Status               WorkStatus  `protobuf:"varint,2,opt,name=status,proto3,enum=pb.WorkStatus" json:"status,omitempty"`
	Work                 *Work       `protobuf:"bytes,3,opt,name=work,proto3" json:"work,omitempty"`
	WorkId               string      `protobuf:"bytes,4,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
	Error                string      `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

// Reset ...
func (m *WorkReply) Reset() { *m = WorkReply{} }

// String ...
func (m *WorkReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*WorkReply) ProtoMessage() {}

// Descriptor ...
func (*WorkReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{2}
}

// XXX_Unmarshal ...
func (m *WorkReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkReply.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *WorkReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkReply.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *WorkReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkReply.Merge(m, src)
}

// XXX_Size ...
func (m *WorkReply) XXX_Size() int {
	return xxx_messageInfo_WorkReply.Size(m)
}

// XXX_DiscardUnknown ...
func (m *WorkReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkReply.DiscardUnknown(m)
}

var xxx_messageInfo_WorkReply proto.InternalMessageInfo

// GetMsg ...
func (m *WorkReply) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

// GetStatus ...
func (m *WorkReply) GetStatus() WorkStatus {
	if m != nil {
		return m.Status
	}
	return WorkStatus_WorkAbnormal
}

// GetWork ...
func (m *WorkReply) GetWork() *Work {
	if m != nil {
		return m.Work
	}
	return nil
}

// GetWorkId ...
func (m *WorkReply) GetWorkId() string {
	if m != nil {
		return m.WorkId
	}
	return ""
}

// GetError ...
func (m *WorkReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// NodeRequest ...
type NodeRequest struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=pb.MessageType" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

// Reset ...
func (m *NodeRequest) Reset() { *m = NodeRequest{} }

// String ...
func (m *NodeRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*NodeRequest) ProtoMessage() {}

// Descriptor ...
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{3}
}

// XXX_Unmarshal ...
func (m *NodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *NodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *NodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRequest.Merge(m, src)
}

// XXX_Size ...
func (m *NodeRequest) XXX_Size() int {
	return xxx_messageInfo_NodeRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *NodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRequest proto.InternalMessageInfo

// GetMsg ...
func (m *NodeRequest) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

// NodeReply ...
type NodeReply struct {
	Msg                  MessageType `protobuf:"varint,1,opt,name=msg,proto3,enum=pb.MessageType" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

// Reset ...
func (m *NodeReply) Reset() { *m = NodeReply{} }

// String ...
func (m *NodeReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*NodeReply) ProtoMessage() {}

// Descriptor ...
func (*NodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_785183a1bcf95a49, []int{4}
}

// XXX_Unmarshal ...
func (m *NodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeReply.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *NodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeReply.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *NodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeReply.Merge(m, src)
}

// XXX_Size ...
func (m *NodeReply) XXX_Size() int {
	return xxx_messageInfo_NodeReply.Size(m)
}

// XXX_DiscardUnknown ...
func (m *NodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_NodeReply proto.InternalMessageInfo

// GetMsg ...
func (m *NodeReply) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_Add
}

func init() {
	proto.RegisterEnum("pb.WorkMode", WorkMode_name, WorkMode_value)
	proto.RegisterEnum("pb.WorkStatus", WorkStatus_name, WorkStatus_value)
	proto.RegisterEnum("pb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Work)(nil), "pb.Work")
	proto.RegisterType((*WorkRequest)(nil), "pb.WorkRequest")
	proto.RegisterType((*WorkReply)(nil), "pb.WorkReply")
	proto.RegisterType((*NodeRequest)(nil), "pb.NodeRequest")
	proto.RegisterType((*NodeReply)(nil), "pb.NodeReply")
}

func init() { proto.RegisterFile("api/pb/service.proto", fileDescriptor_785183a1bcf95a49) }

var fileDescriptor_785183a1bcf95a49 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xfd, 0x6f, 0x1d, 0x7b, 0x9c, 0x4d, 0x8c, 0x58, 0xa8, 0x29, 0x5d, 0x9a, 0xfa, 0x50,
	0xd2, 0x1c, 0xb2, 0x4b, 0x7a, 0xec, 0x29, 0x10, 0x0a, 0x81, 0x6e, 0x29, 0x4e, 0x61, 0x4f, 0x65,
	0x91, 0x63, 0x6d, 0x62, 0xd6, 0xb6, 0x54, 0x49, 0xd9, 0x90, 0x47, 0xe8, 0x73, 0x94, 0xbe, 0x67,
	0x91, 0x64, 0x67, 0x73, 0xd8, 0x43, 0x6e, 0x9a, 0xdf, 0x37, 0xa3, 0x99, 0xf9, 0x84, 0xe0, 0x0a,
	0xb3, 0xf2, 0x86, 0xe5, 0x37, 0x82, 0xf0, 0xe7, 0x72, 0x4d, 0xa6, 0x8c, 0x53, 0x49, 0x91, 0xc3,
	0xf2, 0xf4, 0x9f, 0x0d, 0xde, 0x3d, 0xe5, 0x4f, 0xe8, 0x1a, 0xe0, 0xb9, 0x2c, 0x08, 0x7d, 0x60,
	0x58, 0x6e, 0x13, 0x6f, 0xe4, 0x8e, 0xc3, 0x2c, 0xd4, 0xe4, 0x07, 0x96, 0x5b, 0xf4, 0x1e, 0x22,
	0x46, 0x85, 0x24, 0xdc, 0xe8, 0x17, 0x23, 0x7b, 0x1c, 0x66, 0x60, 0x90, 0x4e, 0xb8, 0x06, 0x90,
	0xdb, 0x5d, 0x9d, 0x1b, 0xdd, 0xd7, 0x7a, 0xa8, 0x49, 0x57, 0x2f, 0x70, 0xcd, 0x2a, 0x62, 0xf4,
	0x9e, 0xbe, 0x1f, 0x0c, 0xea, 0xea, 0x4d, 0xff, 0xb2, 0x79, 0xa4, 0x49, 0x60, 0xea, 0x35, 0x59,
	0x36, 0x8f, 0x34, 0xfd, 0x63, 0x43, 0xa4, 0xe6, 0xcc, 0xc8, 0xef, 0x1d, 0x11, 0x12, 0x7d, 0x00,
	0xb7, 0x16, 0x9b, 0xc4, 0x1e, 0xd9, 0xe3, 0xc1, 0x6c, 0x38, 0x65, 0xf9, 0xf4, 0x8e, 0x08, 0x81,
	0x37, 0xe4, 0xe7, 0x81, 0x91, 0x4c, 0x69, 0xe8, 0x13, 0x84, 0x7b, 0xca, 0x9f, 0x1e, 0x6a, 0x5a,
	0x90, 0xc4, 0xd1, 0x89, 0x7d, 0x95, 0xa8, 0xae, 0xb9, 0xa3, 0x05, 0xc9, 0x82, 0x7d, 0x7b, 0x42,
	0x03, 0x70, 0x96, 0x8b, 0xc4, 0xd5, 0x4d, 0x9d, 0xe5, 0x02, 0xbd, 0x03, 0x4f, 0x69, 0x7a, 0xcd,
	0x68, 0x16, 0x74, 0x55, 0x99, 0xa6, 0xe9, 0x5f, 0x1b, 0x42, 0x33, 0x0b, 0xab, 0x0e, 0xe7, 0x4c,
	0xf2, 0x11, 0x7c, 0x21, 0xb1, 0xdc, 0x89, 0x76, 0x8c, 0x41, 0x77, 0xe1, 0x4a, 0xd3, 0xac, 0x55,
	0x8f, 0x6d, 0xdd, 0xd7, 0xda, 0xa2, 0x37, 0xd0, 0xd3, 0xfb, 0x94, 0x45, 0xe2, 0xe9, 0x49, 0x7d,
	0x15, 0x2e, 0x0b, 0x74, 0x05, 0x17, 0x84, 0x73, 0xca, 0xdb, 0x57, 0x31, 0x41, 0x7a, 0x0b, 0xd1,
	0x77, 0xb5, 0xe5, 0xd9, 0x86, 0xa5, 0x53, 0x08, 0x4d, 0xc5, 0x79, 0x6b, 0x4d, 0xbe, 0x40, 0xd0,
	0x79, 0x89, 0x2e, 0x21, 0xfc, 0x46, 0xd7, 0xb8, 0x52, 0x41, 0x6c, 0xa1, 0x01, 0x40, 0x46, 0x6a,
	0x2a, 0x89, 0x8e, 0x6d, 0x14, 0x43, 0x7f, 0x41, 0xf7, 0x4d, 0x45, 0x71, 0xa1, 0x89, 0x33, 0xc1,
	0x00, 0x2f, 0x0e, 0x28, 0x5d, 0x45, 0xf3, 0xbc, 0xa1, 0xbc, 0xc6, 0x55, 0x6c, 0xa1, 0xa1, 0x79,
	0xef, 0x7b, 0x5c, 0xca, 0xb2, 0xd9, 0xc4, 0x76, 0x07, 0xb2, 0x5d, 0xd3, 0x28, 0xe0, 0x74, 0x60,
	0x25, 0x29, 0x63, 0xa4, 0x88, 0x5d, 0xd5, 0x54, 0x81, 0xaf, 0x65, 0x53, 0x8a, 0x6d, 0xec, 0x4d,
	0x6e, 0x21, 0x3a, 0x99, 0x19, 0xf5, 0xc0, 0x9d, 0x17, 0x45, 0x6c, 0xa1, 0x3e, 0x04, 0x73, 0x71,
	0x68, 0xd6, 0x2a, 0xb2, 0x11, 0x80, 0x6f, 0x86, 0x88, 0x9d, 0xd9, 0x2f, 0xe8, 0xad, 0xcc, 0x17,
	0x41, 0xe3, 0xf6, 0x5f, 0x0c, 0x8f, 0xaf, 0x60, 0x8c, 0x7c, 0x7b, 0xf9, 0x02, 0x58, 0x75, 0x48,
	0x2d, 0x95, 0xa9, 0x6c, 0x33, 0x99, 0x27, 0x96, 0x9b, 0xcc, 0xa3, 0xa3, 0xa9, 0x95, 0xfb, 0xfa,
	0xdf, 0x7d, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x72, 0x07, 0xea, 0x8e, 0x8f, 0x03, 0x00, 0x00,
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

// NewServiceClient ...
func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

// Work ...
func (c *serviceClient) Work(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error) {
	out := new(WorkReply)
	err := c.cc.Invoke(ctx, "/pb.Service/Work", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Node ...
func (c *serviceClient) Node(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeReply, error) {
	out := new(NodeReply)
	err := c.cc.Invoke(ctx, "/pb.Service/Node", in, out, opts...)
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

// RegisterServiceServer ...
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
		FullMethod: "/pb.Service/Work",
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
		FullMethod: "/pb.Service/Node",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Node(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service",
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
