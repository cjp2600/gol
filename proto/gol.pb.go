// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gol.proto

package gol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Methods int32

const (
	Methods_post  Methods = 0
	Methods_get   Methods = 1
	Methods_put   Methods = 2
	Methods_patch Methods = 3
)

var Methods_name = map[int32]string{
	0: "post",
	1: "get",
	2: "put",
	3: "patch",
}

var Methods_value = map[string]int32{
	"post":  0,
	"get":   1,
	"put":   2,
	"patch": 3,
}

func (x Methods) String() string {
	return proto.EnumName(Methods_name, int32(x))
}

func (Methods) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{0}
}

type SequenceType int32

const (
	SequenceType_parallel SequenceType = 0
	SequenceType_sync     SequenceType = 1
)

var SequenceType_name = map[int32]string{
	0: "parallel",
	1: "sync",
}

var SequenceType_value = map[string]int32{
	"parallel": 0,
	"sync":     1,
}

func (x SequenceType) String() string {
	return proto.EnumName(SequenceType_name, int32(x))
}

func (SequenceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{1}
}

type ExecuteResponse struct {
	Jobs                 map[string]string `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExecuteResponse) Reset()         { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()    {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{0}
}

func (m *ExecuteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteResponse.Unmarshal(m, b)
}
func (m *ExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteResponse.Merge(m, src)
}
func (m *ExecuteResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteResponse.Size(m)
}
func (m *ExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteResponse proto.InternalMessageInfo

func (m *ExecuteResponse) GetJobs() map[string]string {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type ExecuteRequest struct {
	Sequence             []*Sequence `protobuf:"bytes,1,rep,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{1}
}

func (m *ExecuteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteRequest.Unmarshal(m, b)
}
func (m *ExecuteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteRequest.Merge(m, src)
}
func (m *ExecuteRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteRequest.Size(m)
}
func (m *ExecuteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteRequest proto.InternalMessageInfo

func (m *ExecuteRequest) GetSequence() []*Sequence {
	if m != nil {
		return m.Sequence
	}
	return nil
}

type Sequence struct {
	Type                 SequenceType `protobuf:"varint,1,opt,name=type,proto3,enum=gol.SequenceType" json:"type,omitempty"`
	Jobs                 []*Job       `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Sequence) Reset()         { *m = Sequence{} }
func (m *Sequence) String() string { return proto.CompactTextString(m) }
func (*Sequence) ProtoMessage()    {}
func (*Sequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{2}
}

func (m *Sequence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sequence.Unmarshal(m, b)
}
func (m *Sequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sequence.Marshal(b, m, deterministic)
}
func (m *Sequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequence.Merge(m, src)
}
func (m *Sequence) XXX_Size() int {
	return xxx_messageInfo_Sequence.Size(m)
}
func (m *Sequence) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequence.DiscardUnknown(m)
}

var xxx_messageInfo_Sequence proto.InternalMessageInfo

func (m *Sequence) GetType() SequenceType {
	if m != nil {
		return m.Type
	}
	return SequenceType_parallel
}

func (m *Sequence) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type Job struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string              `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Method               Methods             `protobuf:"varint,3,opt,name=method,proto3,enum=gol.Methods" json:"method,omitempty"`
	Body                 map[string]*any.Any `protobuf:"bytes,4,rep,name=body,proto3" json:"body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Header               map[string]*any.Any `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Var                  []*Var              `protobuf:"bytes,6,rep,name=var,proto3" json:"var,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{3}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Job) GetMethod() Methods {
	if m != nil {
		return m.Method
	}
	return Methods_post
}

func (m *Job) GetBody() map[string]*any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Job) GetHeader() map[string]*any.Any {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Job) GetVar() []*Var {
	if m != nil {
		return m.Var
	}
	return nil
}

type Var struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	JPath                string   `protobuf:"bytes,3,opt,name=jPath,proto3" json:"jPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Var) Reset()         { *m = Var{} }
func (m *Var) String() string { return proto.CompactTextString(m) }
func (*Var) ProtoMessage()    {}
func (*Var) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7052d12ff7cf86b, []int{4}
}

func (m *Var) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Var.Unmarshal(m, b)
}
func (m *Var) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Var.Marshal(b, m, deterministic)
}
func (m *Var) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Var.Merge(m, src)
}
func (m *Var) XXX_Size() int {
	return xxx_messageInfo_Var.Size(m)
}
func (m *Var) XXX_DiscardUnknown() {
	xxx_messageInfo_Var.DiscardUnknown(m)
}

var xxx_messageInfo_Var proto.InternalMessageInfo

func (m *Var) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Var) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Var) GetJPath() string {
	if m != nil {
		return m.JPath
	}
	return ""
}

func init() {
	proto.RegisterEnum("gol.Methods", Methods_name, Methods_value)
	proto.RegisterEnum("gol.SequenceType", SequenceType_name, SequenceType_value)
	proto.RegisterType((*ExecuteResponse)(nil), "gol.ExecuteResponse")
	proto.RegisterMapType((map[string]string)(nil), "gol.ExecuteResponse.JobsEntry")
	proto.RegisterType((*ExecuteRequest)(nil), "gol.ExecuteRequest")
	proto.RegisterType((*Sequence)(nil), "gol.Sequence")
	proto.RegisterType((*Job)(nil), "gol.Job")
	proto.RegisterMapType((map[string]*any.Any)(nil), "gol.Job.BodyEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "gol.Job.HeaderEntry")
	proto.RegisterType((*Var)(nil), "gol.Var")
}

func init() {
	proto.RegisterFile("proto/gol.proto", fileDescriptor_d7052d12ff7cf86b)
}

var fileDescriptor_d7052d12ff7cf86b = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0xad, 0x2f, 0xb9, 0x7d, 0xcd, 0x9f, 0xf8, 0x1f, 0xb2, 0x30, 0x16, 0x42, 0x95, 0x05, 0x55,
	0x89, 0xc0, 0x86, 0xb0, 0x00, 0x95, 0x15, 0xa0, 0x0a, 0x14, 0xa9, 0x6a, 0x65, 0x50, 0x91, 0xd8,
	0xa0, 0xb1, 0x3d, 0xd8, 0x2e, 0x8e, 0x67, 0xf0, 0x8c, 0x03, 0xb3, 0x61, 0xc1, 0x2b, 0xf0, 0x68,
	0x3c, 0x00, 0x1b, 0x1e, 0x04, 0x79, 0x3c, 0x89, 0x52, 0xd4, 0x1d, 0xbb, 0x33, 0xe7, 0x3b, 0x3e,
	0xdf, 0xd5, 0x30, 0x65, 0x35, 0x15, 0x34, 0xcc, 0x68, 0x19, 0x28, 0x84, 0xac, 0x8c, 0x96, 0xde,
	0xfb, 0xac, 0x10, 0x79, 0x13, 0x07, 0x09, 0x5d, 0x85, 0x59, 0xcd, 0x92, 0x07, 0x24, 0xa1, 0x5c,
	0x72, 0x41, 0xf4, 0x33, 0xc3, 0x82, 0x7c, 0xc1, 0x32, 0x14, 0x79, 0x51, 0xa7, 0x1f, 0x18, 0xae,
	0x85, 0x0c, 0x33, 0x4a, 0xb3, 0x92, 0x60, 0x56, 0x70, 0x0d, 0x43, 0xcc, 0x8a, 0x10, 0x57, 0x15,
	0x15, 0x58, 0x14, 0xb4, 0xe2, 0x5d, 0x02, 0xef, 0xa6, 0x8e, 0xaa, 0x57, 0xdc, 0x7c, 0x0c, 0x71,
	0x25, 0xbb, 0x90, 0xff, 0x0d, 0xa6, 0x27, 0x5f, 0x49, 0xd2, 0x08, 0x12, 0x11, 0xce, 0x68, 0xc5,
	0x09, 0x5a, 0x80, 0x7d, 0x49, 0x63, 0xee, 0x1a, 0x07, 0xd6, 0xd1, 0xfe, 0xe2, 0x76, 0xd0, 0x16,
	0xfa, 0x97, 0x26, 0x58, 0xd2, 0x98, 0x9f, 0x54, 0xa2, 0x96, 0x91, 0xd2, 0x7a, 0x4f, 0x60, 0xb4,
	0xa5, 0x90, 0x03, 0xd6, 0x27, 0x22, 0x5d, 0xe3, 0xc0, 0x38, 0x1a, 0x45, 0x2d, 0x44, 0x33, 0xe8,
	0xad, 0x71, 0xd9, 0x10, 0xd7, 0x54, 0x5c, 0xf7, 0x38, 0x36, 0x9f, 0x1a, 0xfe, 0x33, 0x98, 0x6c,
	0xbd, 0x3f, 0x37, 0x84, 0x0b, 0x74, 0x0f, 0x86, 0xbc, 0x85, 0x55, 0x42, 0x74, 0x09, 0xff, 0xa9,
	0x12, 0xde, 0x68, 0x32, 0xda, 0x86, 0xfd, 0x33, 0x18, 0x6e, 0x58, 0x74, 0x17, 0x6c, 0x21, 0x19,
	0x51, 0x59, 0x27, 0x8b, 0xff, 0xaf, 0x7c, 0xf2, 0x56, 0x32, 0x12, 0xa9, 0x30, 0xba, 0xa5, 0x9b,
	0x33, 0x95, 0xf3, 0x50, 0xc9, 0x96, 0x34, 0xee, 0xda, 0xf0, 0x7f, 0x99, 0x60, 0x2d, 0x69, 0x8c,
	0x26, 0x60, 0x16, 0xa9, 0x6e, 0xc0, 0x2c, 0xd2, 0xb6, 0xa3, 0xa6, 0x2e, 0x75, 0xf5, 0x2d, 0x44,
	0x77, 0xa0, 0xbf, 0x22, 0x22, 0xa7, 0xa9, 0x6b, 0xa9, 0x84, 0x63, 0xe5, 0x74, 0xaa, 0x28, 0x1e,
	0xe9, 0x18, 0x3a, 0x04, 0x3b, 0xa6, 0xa9, 0x74, 0x6d, 0x95, 0x0d, 0x6d, 0xb2, 0x05, 0x2f, 0x68,
	0x2a, 0xf5, 0xf8, 0xda, 0x38, 0xba, 0x0f, 0xfd, 0x9c, 0xe0, 0x94, 0xd4, 0x6e, 0x4f, 0x29, 0x67,
	0x5b, 0xe5, 0x6b, 0x45, 0x77, 0x5a, 0xad, 0x41, 0x1e, 0x58, 0x6b, 0x5c, 0xbb, 0xfd, 0x9d, 0x16,
	0x2e, 0x70, 0x1d, 0xb5, 0xa4, 0x77, 0x0a, 0xa3, 0xad, 0xf9, 0x35, 0x8b, 0x98, 0xef, 0x2e, 0xa2,
	0xcb, 0xd3, 0x5e, 0x46, 0xb0, 0xb9, 0x8c, 0xe0, 0x79, 0x25, 0x77, 0xd6, 0xe3, 0x9d, 0xc1, 0xfe,
	0x4e, 0x05, 0xff, 0x6e, 0xe8, 0xbf, 0x04, 0xeb, 0x02, 0xd7, 0x08, 0x81, 0x5d, 0xe1, 0x15, 0xd1,
	0x4e, 0x0a, 0xb7, 0x9c, 0xda, 0x60, 0x37, 0xe5, 0x6e, 0x5d, 0x33, 0xe8, 0x5d, 0x9e, 0x63, 0x91,
	0xab, 0x29, 0x8f, 0xa2, 0xee, 0x31, 0x7f, 0x08, 0x03, 0x3d, 0x69, 0x34, 0x04, 0x9b, 0x51, 0x2e,
	0x9c, 0x3d, 0x34, 0x00, 0x2b, 0x23, 0xc2, 0x31, 0x5a, 0xc0, 0x1a, 0xe1, 0x98, 0x68, 0x04, 0x3d,
	0x86, 0x45, 0x92, 0x3b, 0xd6, 0xfc, 0x10, 0xc6, 0xbb, 0xc7, 0x80, 0xc6, 0x30, 0x64, 0xb8, 0xc6,
	0x65, 0x49, 0x4a, 0x67, 0xaf, 0x35, 0xe1, 0xb2, 0x4a, 0x1c, 0x63, 0xf1, 0x0e, 0xac, 0x57, 0xb4,
	0x44, 0xe7, 0x30, 0xd0, 0x57, 0x89, 0x6e, 0x5c, 0xbd, 0x7f, 0x75, 0xa3, 0xde, 0xec, 0xba, 0x9f,
	0xc2, 0xf7, 0xbe, 0xff, 0xfc, 0xfd, 0xc3, 0x9c, 0xf9, 0x53, 0xf5, 0x1b, 0xae, 0x1f, 0x85, 0xa4,
	0x13, 0x1c, 0x1b, 0xf3, 0xb8, 0xaf, 0x06, 0xf2, 0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0xf3, 0xb5, 0x99, 0xfd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GolClient is the client API for Gol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GolClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type golClient struct {
	cc grpc.ClientConnInterface
}

func NewGolClient(cc grpc.ClientConnInterface) GolClient {
	return &golClient{cc}
}

func (c *golClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/gol.Gol/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GolServer is the server API for Gol service.
type GolServer interface {
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

// UnimplementedGolServer can be embedded to have forward compatible implementations.
type UnimplementedGolServer struct {
}

func (*UnimplementedGolServer) Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterGolServer(s *grpc.Server, srv GolServer) {
	s.RegisterService(&_Gol_serviceDesc, srv)
}

func _Gol_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gol.Gol/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gol.Gol",
	HandlerType: (*GolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Gol_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gol.proto",
}
