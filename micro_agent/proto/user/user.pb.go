// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user_agent is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	Error
	ReqMsg
	ResMsg
*/
package user_agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ReqMsg struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *ReqMsg) Reset()                    { *m = ReqMsg{} }
func (m *ReqMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqMsg) ProtoMessage()               {}
func (*ReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReqMsg) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type ResMsg struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Info  string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *ResMsg) Reset()                    { *m = ResMsg{} }
func (m *ResMsg) String() string            { return proto.CompactTextString(m) }
func (*ResMsg) ProtoMessage()               {}
func (*ResMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResMsg) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ResMsg) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "user_agent.Error")
	proto.RegisterType((*ReqMsg)(nil), "user_agent.ReqMsg")
	proto.RegisterType((*ResMsg)(nil), "user_agent.ResMsg")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0xaf, 0x82, 0x30,
	0x14, 0x85, 0x1f, 0x4f, 0x20, 0x72, 0x59, 0xf4, 0x4e, 0x44, 0x17, 0xd2, 0xc4, 0xc8, 0x22, 0x03,
	0x0e, 0x8e, 0x4e, 0x0c, 0x0e, 0x3a, 0x34, 0x71, 0x36, 0x88, 0x85, 0x38, 0xd0, 0x62, 0x8b, 0xff,
	0xdf, 0xdc, 0x8b, 0x89, 0xc6, 0xed, 0xcb, 0xe9, 0xe9, 0xf9, 0x5a, 0x80, 0xa7, 0x53, 0x36, 0xef,
	0xad, 0x19, 0x0c, 0x32, 0x5f, 0xaa, 0x56, 0xe9, 0x41, 0x6c, 0x20, 0x28, 0xad, 0x35, 0x16, 0x11,
	0xfc, 0xda, 0xdc, 0x54, 0xe2, 0xa5, 0x5e, 0x16, 0x48, 0x66, 0x9c, 0xc1, 0xa4, 0x73, 0x6d, 0xf2,
	0x9f, 0x7a, 0x59, 0x24, 0x09, 0xc5, 0x0a, 0x42, 0xa9, 0x1e, 0x47, 0xd7, 0xe2, 0x12, 0x22, 0x9e,
	0xd1, 0x55, 0x37, 0x5e, 0x8a, 0xe4, 0x94, 0x82, 0x53, 0xd5, 0x29, 0x51, 0x52, 0xcd, 0x51, 0x6d,
	0x0d, 0x81, 0xa2, 0x7d, 0xae, 0xc4, 0xc5, 0x3c, 0xff, 0xb8, 0x73, 0x16, 0xcb, 0xf1, 0x9c, 0xfc,
	0x77, 0xdd, 0x98, 0xb7, 0x8c, 0xb9, 0xd8, 0x83, 0x7f, 0x76, 0xca, 0xe2, 0x0e, 0x62, 0xd9, 0xd7,
	0x84, 0x07, 0xdd, 0x18, 0xc4, 0xef, 0x91, 0xf1, 0x39, 0x8b, 0x9f, 0x8c, 0xdc, 0xe2, 0xef, 0x1a,
	0xf2, 0x87, 0xb7, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xd4, 0x66, 0x28, 0xfe, 0x00, 0x00,
	0x00,
}
