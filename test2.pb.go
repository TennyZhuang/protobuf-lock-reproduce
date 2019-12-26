// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test2.proto

package test

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BB struct {
	X                    uint64   `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint64   `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    uint64   `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BB) Reset()         { *m = BB{} }
func (m *BB) String() string { return proto.CompactTextString(m) }
func (*BB) ProtoMessage()    {}
func (*BB) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5535eff96cd18, []int{0}
}
func (m *BB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BB.Unmarshal(m, b)
}
func (m *BB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BB.Marshal(b, m, deterministic)
}
func (m *BB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BB.Merge(m, src)
}
func (m *BB) XXX_Size() int {
	return xxx_messageInfo_BB.Size(m)
}
func (m *BB) XXX_DiscardUnknown() {
	xxx_messageInfo_BB.DiscardUnknown(m)
}

var xxx_messageInfo_BB proto.InternalMessageInfo

func (m *BB) GetX() uint64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *BB) GetY() uint64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *BB) GetZ() uint64 {
	if m != nil {
		return m.Z
	}
	return 0
}

type AA struct {
	Bs                   []*BB    `protobuf:"bytes,1,rep,name=bs,proto3" json:"bs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AA) Reset()         { *m = AA{} }
func (m *AA) String() string { return proto.CompactTextString(m) }
func (*AA) ProtoMessage()    {}
func (*AA) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5535eff96cd18, []int{1}
}
func (m *AA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AA.Unmarshal(m, b)
}
func (m *AA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AA.Marshal(b, m, deterministic)
}
func (m *AA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AA.Merge(m, src)
}
func (m *AA) XXX_Size() int {
	return xxx_messageInfo_AA.Size(m)
}
func (m *AA) XXX_DiscardUnknown() {
	xxx_messageInfo_AA.DiscardUnknown(m)
}

var xxx_messageInfo_AA proto.InternalMessageInfo

func (m *AA) GetBs() []*BB {
	if m != nil {
		return m.Bs
	}
	return nil
}

func init() {
	proto.RegisterType((*BB)(nil), "test.BB")
	proto.RegisterType((*AA)(nil), "test.AA")
}

func init() { proto.RegisterFile("test2.proto", fileDescriptor_61e5535eff96cd18) }

var fileDescriptor_61e5535eff96cd18 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x49, 0x2d, 0x2e,
	0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x71, 0xa4, 0x74, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x92, 0x49,
	0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x34, 0x29, 0xe9, 0x71, 0x31, 0x39, 0x39, 0x09,
	0xf1, 0x70, 0x31, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x31, 0x56, 0x80, 0x78, 0x95,
	0x12, 0x4c, 0x10, 0x5e, 0x25, 0x88, 0x57, 0x25, 0xc1, 0x0c, 0xe1, 0x55, 0x29, 0xc9, 0x71, 0x31,
	0x39, 0x3a, 0x0a, 0x49, 0x70, 0x31, 0x25, 0x15, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x71,
	0xe8, 0x81, 0xec, 0xd5, 0x73, 0x72, 0x0a, 0x62, 0x4a, 0x2a, 0x4e, 0x62, 0x03, 0x1b, 0x6b, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x67, 0x35, 0xf0, 0x9a, 0x00, 0x00, 0x00,
}
