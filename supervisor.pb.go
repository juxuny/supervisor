// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor.proto

package supervisor // import "/supervisor"

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

type ApplyReq struct {
	Config               *DeployConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" bson:"-"`
	XXX_unrecognized     []byte        `json:"-" bson:"-"`
	XXX_sizecache        int32         `json:"-" bson:"-"`
}

func (m *ApplyReq) Reset()         { *m = ApplyReq{} }
func (m *ApplyReq) String() string { return proto.CompactTextString(m) }
func (*ApplyReq) ProtoMessage()    {}
func (*ApplyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_supervisor_be9e4e6670cde35a, []int{0}
}
func (m *ApplyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyReq.Unmarshal(m, b)
}
func (m *ApplyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyReq.Marshal(b, m, deterministic)
}
func (dst *ApplyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyReq.Merge(dst, src)
}
func (m *ApplyReq) XXX_Size() int {
	return xxx_messageInfo_ApplyReq.Size(m)
}
func (m *ApplyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyReq proto.InternalMessageInfo

func (m *ApplyReq) GetConfig() *DeployConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type ApplyResp struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *ApplyResp) Reset()         { *m = ApplyResp{} }
func (m *ApplyResp) String() string { return proto.CompactTextString(m) }
func (*ApplyResp) ProtoMessage()    {}
func (*ApplyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_supervisor_be9e4e6670cde35a, []int{1}
}
func (m *ApplyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyResp.Unmarshal(m, b)
}
func (m *ApplyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyResp.Marshal(b, m, deterministic)
}
func (dst *ApplyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyResp.Merge(dst, src)
}
func (m *ApplyResp) XXX_Size() int {
	return xxx_messageInfo_ApplyResp.Size(m)
}
func (m *ApplyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyResp.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyResp proto.InternalMessageInfo

func (m *ApplyResp) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ApplyResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DeployConfig struct {
	ServicePort          uint32      `protobuf:"varint,1,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image                string      `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Tag                  string      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Mounts               []*Mount    `protobuf:"bytes,5,rep,name=mounts" json:"mounts,omitempty"`
	EnvData              string      `protobuf:"bytes,6,opt,name=env_data,json=envData" json:"env_data,omitempty"`
	Envs                 []*KeyValue `protobuf:"bytes,7,rep,name=envs" json:"envs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-" bson:"-"`
	XXX_unrecognized     []byte      `json:"-" bson:"-"`
	XXX_sizecache        int32       `json:"-" bson:"-"`
}

func (m *DeployConfig) Reset()         { *m = DeployConfig{} }
func (m *DeployConfig) String() string { return proto.CompactTextString(m) }
func (*DeployConfig) ProtoMessage()    {}
func (*DeployConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_supervisor_be9e4e6670cde35a, []int{2}
}
func (m *DeployConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployConfig.Unmarshal(m, b)
}
func (m *DeployConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployConfig.Marshal(b, m, deterministic)
}
func (dst *DeployConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployConfig.Merge(dst, src)
}
func (m *DeployConfig) XXX_Size() int {
	return xxx_messageInfo_DeployConfig.Size(m)
}
func (m *DeployConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DeployConfig proto.InternalMessageInfo

func (m *DeployConfig) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *DeployConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeployConfig) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DeployConfig) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DeployConfig) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *DeployConfig) GetEnvData() string {
	if m != nil {
		return m.EnvData
	}
	return ""
}

func (m *DeployConfig) GetEnvs() []*KeyValue {
	if m != nil {
		return m.Envs
	}
	return nil
}

type Mount struct {
	HostPath             string   `protobuf:"bytes,1,opt,name=host_path,json=hostPath" json:"host_path,omitempty"`
	MountPath            string   `protobuf:"bytes,2,opt,name=mount_path,json=mountPath" json:"mount_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Mount) Reset()         { *m = Mount{} }
func (m *Mount) String() string { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()    {}
func (*Mount) Descriptor() ([]byte, []int) {
	return fileDescriptor_supervisor_be9e4e6670cde35a, []int{3}
}
func (m *Mount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mount.Unmarshal(m, b)
}
func (m *Mount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mount.Marshal(b, m, deterministic)
}
func (dst *Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mount.Merge(dst, src)
}
func (m *Mount) XXX_Size() int {
	return xxx_messageInfo_Mount.Size(m)
}
func (m *Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_Mount proto.InternalMessageInfo

func (m *Mount) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *Mount) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values               string   `protobuf:"bytes,2,opt,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_supervisor_be9e4e6670cde35a, []int{4}
}
func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (dst *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(dst, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplyReq)(nil), "ApplyReq")
	proto.RegisterType((*ApplyResp)(nil), "ApplyResp")
	proto.RegisterType((*DeployConfig)(nil), "DeployConfig")
	proto.RegisterType((*Mount)(nil), "Mount")
	proto.RegisterType((*KeyValue)(nil), "KeyValue")
}

func init() { proto.RegisterFile("supervisor.proto", fileDescriptor_supervisor_be9e4e6670cde35a) }

var fileDescriptor_supervisor_be9e4e6670cde35a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xd1, 0x4a, 0xc3, 0x40,
	0x10, 0x34, 0xb6, 0x49, 0x93, 0x4d, 0x0b, 0x65, 0x11, 0x89, 0x4a, 0x25, 0x06, 0x84, 0x3e, 0x45,
	0x5a, 0xfd, 0x01, 0x6d, 0xdf, 0x44, 0x28, 0x11, 0x7c, 0xf0, 0xa5, 0x9c, 0xed, 0x99, 0x16, 0x9b,
	0xdc, 0x99, 0xbb, 0x06, 0xf2, 0x85, 0xfe, 0x96, 0xdc, 0xe6, 0x5a, 0x7d, 0x9b, 0x9d, 0x61, 0x66,
	0x77, 0x19, 0x18, 0xaa, 0xbd, 0xe4, 0x55, 0xbd, 0x55, 0xa2, 0x4a, 0x65, 0x25, 0xb4, 0x48, 0x26,
	0xe0, 0x3f, 0x4a, 0xb9, 0x6b, 0x32, 0xfe, 0x8d, 0xb7, 0xe0, 0xad, 0x44, 0xf9, 0xb9, 0xcd, 0x23,
	0x27, 0x76, 0xc6, 0xe1, 0x74, 0x90, 0xce, 0xb9, 0xdc, 0x89, 0x66, 0x46, 0x64, 0x66, 0xc5, 0x64,
	0x02, 0x81, 0xb5, 0x28, 0x89, 0x08, 0xdd, 0x95, 0x58, 0x73, 0x72, 0x0c, 0x32, 0xc2, 0x38, 0x84,
	0x4e, 0xa1, 0xf2, 0xe8, 0x34, 0x76, 0xc6, 0x41, 0x66, 0x60, 0xf2, 0xe3, 0x40, 0xff, 0x7f, 0x16,
	0xde, 0x40, 0x5f, 0x99, 0x43, 0x56, 0x7c, 0x29, 0x45, 0xa5, 0xad, 0x3d, 0xb4, 0xdc, 0x42, 0x54,
	0xda, 0x24, 0x97, 0xac, 0xe0, 0x36, 0x86, 0x30, 0x9e, 0x81, 0xbb, 0x2d, 0x58, 0xce, 0xa3, 0x0e,
	0x91, 0xed, 0x60, 0xf6, 0x69, 0x96, 0x47, 0xdd, 0x76, 0x9f, 0x66, 0x39, 0x5e, 0x83, 0x57, 0x88,
	0x7d, 0xa9, 0x55, 0xe4, 0xc6, 0x9d, 0x71, 0x38, 0xf5, 0xd2, 0x17, 0x33, 0x66, 0x96, 0xc5, 0x0b,
	0xf0, 0x79, 0x59, 0x2f, 0xd7, 0x4c, 0xb3, 0xc8, 0x23, 0x5b, 0x8f, 0x97, 0xf5, 0x9c, 0x69, 0x86,
	0x23, 0xe8, 0xf2, 0xb2, 0x56, 0x51, 0x8f, 0x8c, 0x41, 0xfa, 0xcc, 0x9b, 0x37, 0xb6, 0xdb, 0xf3,
	0x8c, 0xe8, 0x64, 0x06, 0x2e, 0x45, 0xe1, 0x15, 0x04, 0x1b, 0xa1, 0xf4, 0x52, 0x32, 0xbd, 0xa1,
	0xf3, 0x83, 0xcc, 0x37, 0xc4, 0x82, 0xe9, 0x0d, 0x8e, 0x00, 0x68, 0x53, 0xab, 0xb6, 0x1f, 0x04,
	0xc4, 0x18, 0x39, 0x79, 0x00, 0xff, 0x10, 0x6b, 0x8e, 0xff, 0xe2, 0x8d, 0x4d, 0x30, 0x10, 0xcf,
	0xc1, 0xab, 0x8d, 0xa4, 0xac, 0xd1, 0x4e, 0xd3, 0x14, 0xe0, 0xf5, 0x58, 0x1f, 0xc6, 0xe0, 0x52,
	0x0b, 0x18, 0xa4, 0x87, 0x02, 0x2f, 0x21, 0x3d, 0x16, 0x93, 0x9c, 0x3c, 0x0d, 0xde, 0xc3, 0xbb,
	0xbf, 0xbe, 0x3f, 0x3c, 0x2a, 0xfc, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x84, 0xe6, 0xd4,
	0x04, 0x02, 0x00, 0x00,
}
