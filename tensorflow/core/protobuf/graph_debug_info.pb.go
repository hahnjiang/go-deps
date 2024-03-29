// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/graph_debug_info.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GraphDebugInfo struct {
	// This stores all the source code file names and can be indexed by the
	// `file_index`.
	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// This maps a node name to a stack trace in the source code.
	Traces               map[string]*GraphDebugInfo_StackTrace `protobuf:"bytes,2,rep,name=traces,proto3" json:"traces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *GraphDebugInfo) Reset()         { *m = GraphDebugInfo{} }
func (m *GraphDebugInfo) String() string { return proto.CompactTextString(m) }
func (*GraphDebugInfo) ProtoMessage()    {}
func (*GraphDebugInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d49d5c184d173e1, []int{0}
}

func (m *GraphDebugInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDebugInfo.Unmarshal(m, b)
}
func (m *GraphDebugInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDebugInfo.Marshal(b, m, deterministic)
}
func (m *GraphDebugInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDebugInfo.Merge(m, src)
}
func (m *GraphDebugInfo) XXX_Size() int {
	return xxx_messageInfo_GraphDebugInfo.Size(m)
}
func (m *GraphDebugInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDebugInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDebugInfo proto.InternalMessageInfo

func (m *GraphDebugInfo) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *GraphDebugInfo) GetTraces() map[string]*GraphDebugInfo_StackTrace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// This represents a file/line location in the source code.
type GraphDebugInfo_FileLineCol struct {
	// File name index, which can be used to retrieve the file name string from
	// `files`. The value should be between 0 and (len(files)-1)
	FileIndex int32 `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
	// Line number in the file.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Col number in the file line.
	Col int32 `protobuf:"varint,3,opt,name=col,proto3" json:"col,omitempty"`
	// Name of function contains the file line.
	Func string `protobuf:"bytes,4,opt,name=func,proto3" json:"func,omitempty"`
	// Source code contained in this file line.
	Code                 string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphDebugInfo_FileLineCol) Reset()         { *m = GraphDebugInfo_FileLineCol{} }
func (m *GraphDebugInfo_FileLineCol) String() string { return proto.CompactTextString(m) }
func (*GraphDebugInfo_FileLineCol) ProtoMessage()    {}
func (*GraphDebugInfo_FileLineCol) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d49d5c184d173e1, []int{0, 0}
}

func (m *GraphDebugInfo_FileLineCol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDebugInfo_FileLineCol.Unmarshal(m, b)
}
func (m *GraphDebugInfo_FileLineCol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDebugInfo_FileLineCol.Marshal(b, m, deterministic)
}
func (m *GraphDebugInfo_FileLineCol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDebugInfo_FileLineCol.Merge(m, src)
}
func (m *GraphDebugInfo_FileLineCol) XXX_Size() int {
	return xxx_messageInfo_GraphDebugInfo_FileLineCol.Size(m)
}
func (m *GraphDebugInfo_FileLineCol) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDebugInfo_FileLineCol.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDebugInfo_FileLineCol proto.InternalMessageInfo

func (m *GraphDebugInfo_FileLineCol) GetFileIndex() int32 {
	if m != nil {
		return m.FileIndex
	}
	return 0
}

func (m *GraphDebugInfo_FileLineCol) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *GraphDebugInfo_FileLineCol) GetCol() int32 {
	if m != nil {
		return m.Col
	}
	return 0
}

func (m *GraphDebugInfo_FileLineCol) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *GraphDebugInfo_FileLineCol) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// This represents a stack trace which is a ordered list of `FileLineCol`.
type GraphDebugInfo_StackTrace struct {
	// Each line in the stack trace.
	FileLineCols         []*GraphDebugInfo_FileLineCol `protobuf:"bytes,1,rep,name=file_line_cols,json=fileLineCols,proto3" json:"file_line_cols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GraphDebugInfo_StackTrace) Reset()         { *m = GraphDebugInfo_StackTrace{} }
func (m *GraphDebugInfo_StackTrace) String() string { return proto.CompactTextString(m) }
func (*GraphDebugInfo_StackTrace) ProtoMessage()    {}
func (*GraphDebugInfo_StackTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d49d5c184d173e1, []int{0, 1}
}

func (m *GraphDebugInfo_StackTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphDebugInfo_StackTrace.Unmarshal(m, b)
}
func (m *GraphDebugInfo_StackTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphDebugInfo_StackTrace.Marshal(b, m, deterministic)
}
func (m *GraphDebugInfo_StackTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDebugInfo_StackTrace.Merge(m, src)
}
func (m *GraphDebugInfo_StackTrace) XXX_Size() int {
	return xxx_messageInfo_GraphDebugInfo_StackTrace.Size(m)
}
func (m *GraphDebugInfo_StackTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDebugInfo_StackTrace.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDebugInfo_StackTrace proto.InternalMessageInfo

func (m *GraphDebugInfo_StackTrace) GetFileLineCols() []*GraphDebugInfo_FileLineCol {
	if m != nil {
		return m.FileLineCols
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphDebugInfo)(nil), "tensorflow.GraphDebugInfo")
	proto.RegisterMapType((map[string]*GraphDebugInfo_StackTrace)(nil), "tensorflow.GraphDebugInfo.TracesEntry")
	proto.RegisterType((*GraphDebugInfo_FileLineCol)(nil), "tensorflow.GraphDebugInfo.FileLineCol")
	proto.RegisterType((*GraphDebugInfo_StackTrace)(nil), "tensorflow.GraphDebugInfo.StackTrace")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/graph_debug_info.proto", fileDescriptor_2d49d5c184d173e1)
}

var fileDescriptor_2d49d5c184d173e1 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0x25, 0xeb, 0x3a, 0xe8, 0xaf, 0x32, 0x24, 0xec, 0x10, 0x06, 0x42, 0x11, 0x94, 0x9e, 0x5a,
	0x98, 0x08, 0xa2, 0xe0, 0x61, 0xfe, 0x63, 0xb0, 0xc3, 0x88, 0x9e, 0xbc, 0xd4, 0x2e, 0x4b, 0x66,
	0x59, 0x4c, 0x46, 0xda, 0x3a, 0xe7, 0xa7, 0xf1, 0x63, 0x7a, 0x94, 0xa4, 0x42, 0xb7, 0xcb, 0x6e,
	0x2f, 0x2f, 0xef, 0xbd, 0xbc, 0x47, 0x20, 0xad, 0xb8, 0x2a, 0xb5, 0x11, 0x52, 0x6f, 0x52, 0xa6,
	0x0d, 0x4f, 0xd7, 0x46, 0x57, 0x7a, 0x5e, 0x8b, 0x74, 0x69, 0xf2, 0xf5, 0x7b, 0xb6, 0xe0, 0xf3,
	0x7a, 0x99, 0x15, 0x4a, 0xe8, 0xc4, 0xdd, 0x60, 0x68, 0x0d, 0xa7, 0x3f, 0x1e, 0xf4, 0x9f, 0xac,
	0xec, 0xde, 0xaa, 0x26, 0x4a, 0x68, 0x3c, 0x00, 0x5f, 0x14, 0x92, 0x97, 0x04, 0x45, 0x5e, 0x1c,
	0xd0, 0xe6, 0x80, 0x6f, 0xa1, 0x57, 0x99, 0x9c, 0xf1, 0x92, 0x74, 0x22, 0x2f, 0x0e, 0x47, 0xe7,
	0x49, 0x9b, 0x92, 0xec, 0x27, 0x24, 0x2f, 0x4e, 0xf8, 0xa0, 0x2a, 0xb3, 0xa5, 0xff, 0xae, 0xe1,
	0x37, 0x84, 0x8f, 0x85, 0xe4, 0xd3, 0x42, 0xf1, 0x3b, 0x2d, 0xf1, 0x09, 0x80, 0xcd, 0xcd, 0x0a,
	0xb5, 0xe0, 0x5f, 0x04, 0x45, 0x28, 0xf6, 0x69, 0x60, 0x99, 0x89, 0x25, 0x30, 0x86, 0xae, 0x2c,
	0x14, 0x27, 0x1d, 0x77, 0xe1, 0x30, 0x3e, 0x06, 0x8f, 0x69, 0x49, 0x3c, 0x47, 0x59, 0x68, 0x55,
	0xa2, 0x56, 0x8c, 0x74, 0x23, 0x14, 0x07, 0xd4, 0x61, 0xcb, 0x31, 0xbd, 0xe0, 0xc4, 0x6f, 0x38,
	0x8b, 0x87, 0xaf, 0x00, 0xcf, 0x55, 0xce, 0x56, 0xae, 0x17, 0x9e, 0x42, 0xdf, 0x3d, 0x6d, 0x43,
	0x33, 0xa6, 0x65, 0x33, 0xf4, 0xf0, 0xa2, 0x9d, 0xea, 0xf4, 0x48, 0xb4, 0x87, 0x72, 0xf8, 0x06,
	0xe1, 0xce, 0x5c, 0x5b, 0x72, 0xc5, 0xb7, 0x6e, 0x50, 0x40, 0x2d, 0xc4, 0x37, 0xe0, 0x7f, 0xe6,
	0xb2, 0x6e, 0xb6, 0x84, 0xa3, 0xb3, 0x03, 0xaf, 0xb4, 0x25, 0x69, 0xe3, 0xb9, 0xee, 0x5c, 0xa1,
	0xf1, 0x25, 0x10, 0x6d, 0x96, 0xbb, 0x36, 0x61, 0xf2, 0x0f, 0xbe, 0xd1, 0x66, 0x35, 0x1e, 0xec,
	0x27, 0xcc, 0xec, 0xff, 0x96, 0x33, 0xf4, 0x8b, 0xd0, 0xbc, 0xe7, 0x3e, 0xfb, 0xe2, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x38, 0x4d, 0xa8, 0x46, 0x1f, 0x02, 0x00, 0x00,
}
