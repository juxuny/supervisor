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

type HealthCheckType int32

const (
	HealthCheckType_TypeDefault HealthCheckType = 0
	HealthCheckType_TypeTcp     HealthCheckType = 1
)

var HealthCheckType_name = map[int32]string{
	0: "TypeDefault",
	1: "TypeTcp",
}
var HealthCheckType_value = map[string]int32{
	"TypeDefault": 0,
	"TypeTcp":     1,
}

func (x HealthCheckType) String() string {
	return proto.EnumName(HealthCheckType_name, int32(x))
}
func (HealthCheckType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{0}
}

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
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{0}
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
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{1}
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
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{2}
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
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{3}
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
	Status               *Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UpdateResp) Reset()         { *m = UpdateResp{} }
func (m *UpdateResp) String() string { return proto.CompactTextString(m) }
func (*UpdateResp) ProtoMessage()    {}
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{4}
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

func (m *UpdateResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type CheckReq struct {
	Type                 HealthCheckType `protobuf:"varint,1,opt,name=type,enum=proxy.HealthCheckType" json:"type,omitempty"`
	Host                 string          `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Path                 string          `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Port                 uint32          `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" bson:"-"`
	XXX_unrecognized     []byte          `json:"-" bson:"-"`
	XXX_sizecache        int32           `json:"-" bson:"-"`
}

func (m *CheckReq) Reset()         { *m = CheckReq{} }
func (m *CheckReq) String() string { return proto.CompactTextString(m) }
func (*CheckReq) ProtoMessage()    {}
func (*CheckReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{5}
}
func (m *CheckReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckReq.Unmarshal(m, b)
}
func (m *CheckReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckReq.Marshal(b, m, deterministic)
}
func (dst *CheckReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckReq.Merge(dst, src)
}
func (m *CheckReq) XXX_Size() int {
	return xxx_messageInfo_CheckReq.Size(m)
}
func (m *CheckReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckReq proto.InternalMessageInfo

func (m *CheckReq) GetType() HealthCheckType {
	if m != nil {
		return m.Type
	}
	return HealthCheckType_TypeDefault
}

func (m *CheckReq) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *CheckReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CheckReq) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type CheckResp struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *CheckResp) Reset()         { *m = CheckResp{} }
func (m *CheckResp) String() string { return proto.CompactTextString(m) }
func (*CheckResp) ProtoMessage()    {}
func (*CheckResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_0a18fc2b4b1d035c, []int{6}
}
func (m *CheckResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResp.Unmarshal(m, b)
}
func (m *CheckResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResp.Marshal(b, m, deterministic)
}
func (dst *CheckResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResp.Merge(dst, src)
}
func (m *CheckResp) XXX_Size() int {
	return xxx_messageInfo_CheckResp.Size(m)
}
func (m *CheckResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResp proto.InternalMessageInfo

func (m *CheckResp) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CheckResp) GetMsg() string {
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
	proto.RegisterType((*CheckReq)(nil), "proxy.CheckReq")
	proto.RegisterType((*CheckResp)(nil), "proxy.CheckResp")
	proto.RegisterEnum("proxy.HealthCheckType", HealthCheckType_name, HealthCheckType_value)
}

func init() { proto.RegisterFile("proxy/proxy.proto", fileDescriptor_proxy_0a18fc2b4b1d035c) }

var fileDescriptor_proxy_0a18fc2b4b1d035c = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0x02, 0xd5, 0x4e, 0x83, 0x94, 0x3d, 0x10, 0xc2, 0x45, 0xb2, 0xd1, 0x84, 0x10, 0x43,
	0x63, 0xf9, 0x05, 0x7e, 0x1c, 0x3c, 0x9a, 0x8a, 0x17, 0x2f, 0xa6, 0x94, 0x95, 0xa2, 0xc0, 0x2e,
	0xbb, 0x5b, 0x43, 0x7f, 0x89, 0x7f, 0xd7, 0xec, 0x74, 0xa9, 0x81, 0x13, 0x97, 0x76, 0xe6, 0xed,
	0x7b, 0x33, 0xb3, 0x6f, 0x16, 0x3a, 0x42, 0xf2, 0x5d, 0x11, 0xe2, 0x77, 0x2c, 0x24, 0xd7, 0x9c,
	0x34, 0x31, 0xa1, 0xf7, 0xe0, 0xbe, 0xea, 0x44, 0xe7, 0x8a, 0x5c, 0x81, 0xbf, 0x5a, 0x2a, 0xcd,
	0x36, 0x1f, 0x82, 0x4b, 0xdd, 0x73, 0x06, 0xce, 0xb0, 0x15, 0x43, 0x09, 0xbd, 0x70, 0xa9, 0x49,
	0x17, 0x5c, 0xc9, 0xd6, 0x5c, 0xb3, 0xde, 0xd9, 0xc0, 0x19, 0x7a, 0xb1, 0xcd, 0xa8, 0x0f, 0x5e,
	0x59, 0x22, 0x66, 0x5b, 0x3a, 0x01, 0xd8, 0x27, 0x4a, 0x90, 0x1b, 0x70, 0x15, 0x66, 0x58, 0xce,
	0x8f, 0x5a, 0xe3, 0x72, 0x04, 0x4b, 0xb1, 0x87, 0x34, 0x02, 0xef, 0x4d, 0xcc, 0x13, 0xcd, 0x62,
	0xb6, 0x3d, 0x55, 0x33, 0x01, 0xd8, 0x6b, 0x4e, 0x6f, 0x24, 0xe1, 0xe2, 0x31, 0x63, 0xe9, 0xb7,
	0xe9, 0x33, 0x82, 0x86, 0x2e, 0x04, 0x43, 0xc1, 0x65, 0xd4, 0xb5, 0x82, 0x67, 0x96, 0xac, 0x74,
	0x86, 0xa4, 0x69, 0x21, 0x58, 0x8c, 0x1c, 0x42, 0xa0, 0x91, 0x71, 0xa5, 0xed, 0xc5, 0x31, 0x36,
	0x98, 0x48, 0x74, 0xd6, 0xab, 0x97, 0x98, 0x89, 0x11, 0x33, 0xe6, 0x35, 0xd0, 0x3c, 0x8c, 0xe9,
	0x1d, 0x78, 0xb6, 0xa7, 0x12, 0x86, 0x90, 0xf2, 0x39, 0xb3, 0xee, 0x62, 0x4c, 0x02, 0xa8, 0xaf,
	0xd5, 0xc2, 0xd6, 0x36, 0xe1, 0x28, 0x84, 0xf6, 0xd1, 0x1c, 0xa4, 0x0d, 0xbe, 0xf9, 0x3f, 0xb1,
	0xcf, 0x24, 0x5f, 0xe9, 0xa0, 0x46, 0x7c, 0x38, 0x37, 0xc0, 0x34, 0x15, 0x81, 0x13, 0xfd, 0x3a,
	0x50, 0xee, 0x93, 0x84, 0xd5, 0x3e, 0x83, 0x43, 0x0b, 0xd8, 0xb6, 0xdf, 0x39, 0x42, 0x94, 0xa0,
	0x35, 0x23, 0x28, 0x7d, 0xac, 0x04, 0xd5, 0x2a, 0x2a, 0xc1, 0xbf, 0xd1, 0xb4, 0x46, 0x6e, 0xa1,
	0x89, 0x63, 0x91, 0xb6, 0x3d, 0xdd, 0x3b, 0xda, 0x0f, 0x0e, 0x01, 0xc3, 0x7e, 0xb8, 0x7e, 0xa7,
	0x8b, 0xa5, 0xce, 0xf2, 0xd9, 0x38, 0xe5, 0xeb, 0xf0, 0x2b, 0xdf, 0xe5, 0x9b, 0x22, 0x54, 0xb9,
	0x60, 0xf2, 0x67, 0xa9, 0xb8, 0x2c, 0x9f, 0xe4, 0xcc, 0xc5, 0x37, 0x39, 0xf9, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0x34, 0x21, 0x50, 0xa8, 0x02, 0x00, 0x00,
}
