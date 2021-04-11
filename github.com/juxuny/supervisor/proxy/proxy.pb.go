// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy/proxy.proto

package proxy // import "github.com/juxuny/supervisor/proxy"

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

type Status struct {
	ListenPort           uint32   `protobuf:"varint,1,opt,name=listen_port,json=listenPort" json:"listen_port,omitempty"`
	Remote               string   `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_6f39bd17cd9cf329, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetListenPort() uint32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *Status) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

type StatusReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *StatusReq) Reset()         { *m = StatusReq{} }
func (m *StatusReq) String() string { return proto.CompactTextString(m) }
func (*StatusReq) ProtoMessage()    {}
func (*StatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_6f39bd17cd9cf329, []int{1}
}
func (m *StatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReq.Unmarshal(m, b)
}
func (m *StatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReq.Marshal(b, m, deterministic)
}
func (dst *StatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReq.Merge(dst, src)
}
func (m *StatusReq) XXX_Size() int {
	return xxx_messageInfo_StatusReq.Size(m)
}
func (m *StatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReq proto.InternalMessageInfo

type StatusResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *StatusResp) Reset()         { *m = StatusResp{} }
func (m *StatusResp) String() string { return proto.CompactTextString(m) }
func (*StatusResp) ProtoMessage()    {}
func (*StatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_6f39bd17cd9cf329, []int{2}
}
func (m *StatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResp.Unmarshal(m, b)
}
func (m *StatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResp.Marshal(b, m, deterministic)
}
func (dst *StatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResp.Merge(dst, src)
}
func (m *StatusResp) XXX_Size() int {
	return xxx_messageInfo_StatusResp.Size(m)
}
func (m *StatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResp proto.InternalMessageInfo

func (m *StatusResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdateReq struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_6f39bd17cd9cf329, []int{3}
}
func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (dst *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(dst, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdateResp struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UpdateResp) Reset()         { *m = UpdateResp{} }
func (m *UpdateResp) String() string { return proto.CompactTextString(m) }
func (*UpdateResp) ProtoMessage()    {}
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_6f39bd17cd9cf329, []int{4}
}
func (m *UpdateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResp.Unmarshal(m, b)
}
func (m *UpdateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResp.Marshal(b, m, deterministic)
}
func (dst *UpdateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResp.Merge(dst, src)
}
func (m *UpdateResp) XXX_Size() int {
	return xxx_messageInfo_UpdateResp.Size(m)
}
func (m *UpdateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResp proto.InternalMessageInfo

func (m *UpdateResp) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "proxy.Status")
	proto.RegisterType((*StatusReq)(nil), "proxy.StatusReq")
	proto.RegisterType((*StatusResp)(nil), "proxy.StatusResp")
	proto.RegisterType((*UpdateReq)(nil), "proxy.UpdateReq")
	proto.RegisterType((*UpdateResp)(nil), "proxy.UpdateResp")
}

func init() { proto.RegisterFile("proxy/proxy.proto", fileDescriptor_proxy_6f39bd17cd9cf329) }

var fileDescriptor_proxy_6f39bd17cd9cf329 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0xff, 0x04, 0x32, 0xa1, 0xd0, 0xee, 0x41, 0x8a, 0x17, 0xcb, 0x82, 0x10, 0x3c,
	0x24, 0x90, 0x1e, 0x3d, 0xe9, 0x27, 0x90, 0x88, 0x17, 0x2f, 0xd2, 0x3f, 0x43, 0x8d, 0x98, 0xce,
	0x66, 0x67, 0x56, 0xda, 0x6f, 0x2f, 0xdd, 0xdd, 0x46, 0xf4, 0xe4, 0x65, 0x98, 0x37, 0xc9, 0x7b,
	0xfb, 0x1b, 0x06, 0xa6, 0xc6, 0xd2, 0xfe, 0x50, 0xf9, 0x5a, 0x1a, 0x4b, 0x42, 0xea, 0xd2, 0x0b,
	0xfd, 0x00, 0xe9, 0xb3, 0x2c, 0xc5, 0xb1, 0xba, 0x81, 0xfc, 0xb3, 0x65, 0xc1, 0xdd, 0x9b, 0x21,
	0x2b, 0xb3, 0x64, 0x9e, 0x14, 0xe3, 0x06, 0xc2, 0xe8, 0x89, 0xac, 0xa8, 0x2b, 0x48, 0x2d, 0x76,
	0x24, 0x38, 0x3b, 0x9b, 0x27, 0x45, 0xd6, 0x44, 0xa5, 0x73, 0xc8, 0x42, 0x44, 0x83, 0xbd, 0x5e,
	0x00, 0x9c, 0x04, 0x1b, 0x75, 0x0b, 0x29, 0x7b, 0xe5, 0xe3, 0xf2, 0x7a, 0x5c, 0x06, 0x84, 0xf8,
	0x4b, 0xfc, 0xa8, 0x6b, 0xc8, 0x5e, 0xcc, 0x66, 0x29, 0xd8, 0x60, 0xff, 0x7f, 0x0f, 0x9c, 0x3c,
	0x6c, 0x94, 0x82, 0x8b, 0x35, 0x6d, 0x30, 0x52, 0xfb, 0x5e, 0x4d, 0xe0, 0xbc, 0xe3, 0x6d, 0x84,
	0x3d, 0xb6, 0x75, 0x0b, 0x61, 0x6b, 0x55, 0x0d, 0x5b, 0x4f, 0x7e, 0xa7, 0x63, 0x7f, 0x3d, 0xfd,
	0x33, 0x61, 0xa3, 0x47, 0x47, 0x43, 0x78, 0x6d, 0x30, 0x0c, 0xc0, 0x83, 0xe1, 0x07, 0x47, 0x8f,
	0x1e, 0xef, 0x5e, 0x8b, 0x6d, 0x2b, 0xef, 0x6e, 0x55, 0xae, 0xa9, 0xab, 0x3e, 0xdc, 0xde, 0xed,
	0x0e, 0x15, 0x3b, 0x83, 0xf6, 0xab, 0x65, 0xb2, 0xe1, 0x14, 0xf7, 0xbe, 0xae, 0x52, 0x7f, 0x91,
	0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xaa, 0xef, 0x5e, 0xa6, 0x01, 0x00, 0x00,
}
