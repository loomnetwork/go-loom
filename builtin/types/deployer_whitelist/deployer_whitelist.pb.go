// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/deployer_whitelist/deployer_whitelist.proto

package deployer_whitelist

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Flags int32

const (
	Flags_NONE      Flags = 0
	Flags_GO        Flags = 1
	Flags_EVM       Flags = 2
	Flags_MIGRATION Flags = 4
)

var Flags_name = map[int32]string{
	0: "NONE",
	1: "GO",
	2: "EVM",
	4: "MIGRATION",
}
var Flags_value = map[string]int32{
	"NONE":      0,
	"GO":        1,
	"EVM":       2,
	"MIGRATION": 4,
}

func (x Flags) String() string {
	return proto.EnumName(Flags_name, int32(x))
}
func (Flags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{0}
}

type InitRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Deployers            []*Deployer    `protobuf:"bytes,2,rep,name=deployers" json:"deployers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{0}
}
func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (dst *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(dst, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *InitRequest) GetDeployers() []*Deployer {
	if m != nil {
		return m.Deployers
	}
	return nil
}

type Deployer struct {
	Address              *types.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Flags                uint32         `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Deployer) Reset()         { *m = Deployer{} }
func (m *Deployer) String() string { return proto.CompactTextString(m) }
func (*Deployer) ProtoMessage()    {}
func (*Deployer) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{1}
}
func (m *Deployer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployer.Unmarshal(m, b)
}
func (m *Deployer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployer.Marshal(b, m, deterministic)
}
func (dst *Deployer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployer.Merge(dst, src)
}
func (m *Deployer) XXX_Size() int {
	return xxx_messageInfo_Deployer.Size(m)
}
func (m *Deployer) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployer.DiscardUnknown(m)
}

var xxx_messageInfo_Deployer proto.InternalMessageInfo

func (m *Deployer) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Deployer) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type AddDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployerAddr" json:"deployerAddr,omitempty"`
	Flags                uint32         `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddDeployerRequest) Reset()         { *m = AddDeployerRequest{} }
func (m *AddDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*AddDeployerRequest) ProtoMessage()    {}
func (*AddDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{2}
}
func (m *AddDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeployerRequest.Unmarshal(m, b)
}
func (m *AddDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *AddDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeployerRequest.Merge(dst, src)
}
func (m *AddDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_AddDeployerRequest.Size(m)
}
func (m *AddDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeployerRequest proto.InternalMessageInfo

func (m *AddDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

func (m *AddDeployerRequest) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type AddUserDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployerAddr" json:"deployerAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddUserDeployerRequest) Reset()         { *m = AddUserDeployerRequest{} }
func (m *AddUserDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserDeployerRequest) ProtoMessage()    {}
func (*AddUserDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{3}
}
func (m *AddUserDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserDeployerRequest.Unmarshal(m, b)
}
func (m *AddUserDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *AddUserDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserDeployerRequest.Merge(dst, src)
}
func (m *AddUserDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserDeployerRequest.Size(m)
}
func (m *AddUserDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserDeployerRequest proto.InternalMessageInfo

func (m *AddUserDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

type AddDeployerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDeployerResponse) Reset()         { *m = AddDeployerResponse{} }
func (m *AddDeployerResponse) String() string { return proto.CompactTextString(m) }
func (*AddDeployerResponse) ProtoMessage()    {}
func (*AddDeployerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{4}
}
func (m *AddDeployerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeployerResponse.Unmarshal(m, b)
}
func (m *AddDeployerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeployerResponse.Marshal(b, m, deterministic)
}
func (dst *AddDeployerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeployerResponse.Merge(dst, src)
}
func (m *AddDeployerResponse) XXX_Size() int {
	return xxx_messageInfo_AddDeployerResponse.Size(m)
}
func (m *AddDeployerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeployerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeployerResponse proto.InternalMessageInfo

type GetDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployerAddr" json:"deployerAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDeployerRequest) Reset()         { *m = GetDeployerRequest{} }
func (m *GetDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeployerRequest) ProtoMessage()    {}
func (*GetDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{5}
}
func (m *GetDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeployerRequest.Unmarshal(m, b)
}
func (m *GetDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeployerRequest.Merge(dst, src)
}
func (m *GetDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeployerRequest.Size(m)
}
func (m *GetDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeployerRequest proto.InternalMessageInfo

func (m *GetDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

type GetDeployerResponse struct {
	Deployer             *Deployer `protobuf:"bytes,1,opt,name=deployer" json:"deployer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetDeployerResponse) Reset()         { *m = GetDeployerResponse{} }
func (m *GetDeployerResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeployerResponse) ProtoMessage()    {}
func (*GetDeployerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{6}
}
func (m *GetDeployerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeployerResponse.Unmarshal(m, b)
}
func (m *GetDeployerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeployerResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeployerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeployerResponse.Merge(dst, src)
}
func (m *GetDeployerResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeployerResponse.Size(m)
}
func (m *GetDeployerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeployerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeployerResponse proto.InternalMessageInfo

func (m *GetDeployerResponse) GetDeployer() *Deployer {
	if m != nil {
		return m.Deployer
	}
	return nil
}

type RemoveDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployerAddr" json:"deployerAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RemoveDeployerRequest) Reset()         { *m = RemoveDeployerRequest{} }
func (m *RemoveDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveDeployerRequest) ProtoMessage()    {}
func (*RemoveDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{7}
}
func (m *RemoveDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDeployerRequest.Unmarshal(m, b)
}
func (m *RemoveDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDeployerRequest.Merge(dst, src)
}
func (m *RemoveDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveDeployerRequest.Size(m)
}
func (m *RemoveDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDeployerRequest proto.InternalMessageInfo

func (m *RemoveDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

type RemoveDeployerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDeployerResponse) Reset()         { *m = RemoveDeployerResponse{} }
func (m *RemoveDeployerResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveDeployerResponse) ProtoMessage()    {}
func (*RemoveDeployerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{8}
}
func (m *RemoveDeployerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDeployerResponse.Unmarshal(m, b)
}
func (m *RemoveDeployerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDeployerResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveDeployerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDeployerResponse.Merge(dst, src)
}
func (m *RemoveDeployerResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveDeployerResponse.Size(m)
}
func (m *RemoveDeployerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDeployerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDeployerResponse proto.InternalMessageInfo

type ListDeployersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeployersRequest) Reset()         { *m = ListDeployersRequest{} }
func (m *ListDeployersRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeployersRequest) ProtoMessage()    {}
func (*ListDeployersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{9}
}
func (m *ListDeployersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeployersRequest.Unmarshal(m, b)
}
func (m *ListDeployersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeployersRequest.Marshal(b, m, deterministic)
}
func (dst *ListDeployersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeployersRequest.Merge(dst, src)
}
func (m *ListDeployersRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeployersRequest.Size(m)
}
func (m *ListDeployersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeployersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeployersRequest proto.InternalMessageInfo

type ListDeployersResponse struct {
	Deployers            []*Deployer `protobuf:"bytes,1,rep,name=deployers" json:"deployers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListDeployersResponse) Reset()         { *m = ListDeployersResponse{} }
func (m *ListDeployersResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeployersResponse) ProtoMessage()    {}
func (*ListDeployersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_88c0038a1e37c175, []int{10}
}
func (m *ListDeployersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeployersResponse.Unmarshal(m, b)
}
func (m *ListDeployersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeployersResponse.Marshal(b, m, deterministic)
}
func (dst *ListDeployersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeployersResponse.Merge(dst, src)
}
func (m *ListDeployersResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeployersResponse.Size(m)
}
func (m *ListDeployersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeployersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeployersResponse proto.InternalMessageInfo

func (m *ListDeployersResponse) GetDeployers() []*Deployer {
	if m != nil {
		return m.Deployers
	}
	return nil
}

func init() {
	proto.RegisterType((*InitRequest)(nil), "deployer_whitelist.InitRequest")
	proto.RegisterType((*Deployer)(nil), "deployer_whitelist.Deployer")
	proto.RegisterType((*AddDeployerRequest)(nil), "deployer_whitelist.AddDeployerRequest")
	proto.RegisterType((*AddUserDeployerRequest)(nil), "deployer_whitelist.AddUserDeployerRequest")
	proto.RegisterType((*AddDeployerResponse)(nil), "deployer_whitelist.AddDeployerResponse")
	proto.RegisterType((*GetDeployerRequest)(nil), "deployer_whitelist.GetDeployerRequest")
	proto.RegisterType((*GetDeployerResponse)(nil), "deployer_whitelist.GetDeployerResponse")
	proto.RegisterType((*RemoveDeployerRequest)(nil), "deployer_whitelist.RemoveDeployerRequest")
	proto.RegisterType((*RemoveDeployerResponse)(nil), "deployer_whitelist.RemoveDeployerResponse")
	proto.RegisterType((*ListDeployersRequest)(nil), "deployer_whitelist.ListDeployersRequest")
	proto.RegisterType((*ListDeployersResponse)(nil), "deployer_whitelist.ListDeployersResponse")
	proto.RegisterEnum("deployer_whitelist.Flags", Flags_name, Flags_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/deployer_whitelist/deployer_whitelist.proto", fileDescriptor_deployer_whitelist_88c0038a1e37c175)
}

var fileDescriptor_deployer_whitelist_88c0038a1e37c175 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdb, 0x6b, 0xf2, 0x40,
	0x10, 0xc5, 0xbf, 0x78, 0x8d, 0xe3, 0x27, 0x84, 0xf5, 0x82, 0x94, 0x52, 0x64, 0x9f, 0xa4, 0xb4,
	0x49, 0x6b, 0x5f, 0x4a, 0xdf, 0x52, 0xbc, 0x20, 0x54, 0x03, 0xe9, 0x05, 0xdf, 0x8a, 0x76, 0xb7,
	0xba, 0x34, 0x66, 0xd3, 0xec, 0x5a, 0xf1, 0xbf, 0x2f, 0x31, 0x89, 0x25, 0x6a, 0x5b, 0x8a, 0x2f,
	0x81, 0x3d, 0x33, 0xfc, 0xce, 0x19, 0x66, 0x02, 0xa3, 0x29, 0x93, 0xb3, 0xc5, 0x44, 0x7f, 0xe1,
	0x73, 0xc3, 0xe1, 0x7c, 0xee, 0x52, 0xb9, 0xe4, 0xfe, 0x9b, 0x31, 0xe5, 0xe7, 0xc1, 0xd3, 0x98,
	0x2c, 0x98, 0x23, 0x99, 0x6b, 0xc8, 0x95, 0x47, 0x85, 0x41, 0xa8, 0xe7, 0xf0, 0x15, 0xf5, 0x9f,
	0x97, 0x33, 0x26, 0xa9, 0xc3, 0x84, 0xdc, 0x23, 0xe9, 0x9e, 0xcf, 0x25, 0x47, 0x68, 0xb7, 0x72,
	0x74, 0xf1, 0x8b, 0x5b, 0xe8, 0xb2, 0xfe, 0x86, 0x14, 0xcc, 0xa0, 0xd8, 0x77, 0x99, 0xb4, 0xe9,
	0xfb, 0x82, 0x0a, 0x89, 0x4e, 0x20, 0xcb, 0x97, 0x2e, 0xf5, 0xeb, 0x4a, 0x43, 0x69, 0x16, 0x5b,
	0xaa, 0x6e, 0x12, 0xe2, 0x53, 0x21, 0xec, 0x50, 0x46, 0x37, 0x50, 0x88, 0x6d, 0x45, 0x3d, 0xd5,
	0x48, 0x37, 0x8b, 0xad, 0x63, 0x7d, 0x4f, 0xc4, 0x76, 0x24, 0xd9, 0x5f, 0xed, 0xb8, 0x0d, 0x6a,
	0x2c, 0x23, 0x0c, 0xf9, 0x71, 0x48, 0xde, 0x71, 0x8a, 0x0b, 0xa8, 0x02, 0xd9, 0x57, 0x67, 0x3c,
	0x0d, 0x7c, 0x94, 0x66, 0xc9, 0x0e, 0x1f, 0x78, 0x04, 0xc8, 0x24, 0x64, 0xc3, 0x8f, 0x72, 0x9f,
	0xc1, 0xff, 0xd8, 0x28, 0xe0, 0xec, 0x40, 0x13, 0xd5, 0x6f, 0xc8, 0x5d, 0xa8, 0x99, 0x84, 0x3c,
	0x0a, 0xea, 0x1f, 0x44, 0xc7, 0x55, 0x28, 0x27, 0x12, 0x0a, 0x8f, 0xbb, 0x82, 0xe2, 0x5b, 0x40,
	0x3d, 0x2a, 0x0f, 0x43, 0x5b, 0x50, 0x4e, 0x30, 0x42, 0x34, 0xba, 0x06, 0x35, 0x6e, 0x8b, 0x00,
	0x3f, 0x2f, 0x65, 0xd3, 0x8d, 0x3b, 0x50, 0xb5, 0xe9, 0x9c, 0x7f, 0xd0, 0xc3, 0x72, 0xd5, 0xa1,
	0xb6, 0x8d, 0x89, 0xa6, 0xae, 0x41, 0xe5, 0x8e, 0x89, 0x4d, 0x64, 0x11, 0xf1, 0xf1, 0x3d, 0x54,
	0xb7, 0xf4, 0x68, 0x96, 0xc4, 0x85, 0x29, 0x7f, 0xba, 0xb0, 0xd3, 0x4b, 0xc8, 0x76, 0x83, 0x55,
	0x22, 0x15, 0x32, 0x43, 0x6b, 0xd8, 0xd1, 0xfe, 0xa1, 0x1c, 0xa4, 0x7a, 0x96, 0xa6, 0xa0, 0x3c,
	0xa4, 0x3b, 0x4f, 0x03, 0x2d, 0x85, 0x4a, 0x50, 0x18, 0xf4, 0x7b, 0xb6, 0xf9, 0xd0, 0xb7, 0x86,
	0x5a, 0x66, 0x92, 0x5b, 0xff, 0x06, 0x57, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xc5, 0x5f,
	0x09, 0xa8, 0x03, 0x00, 0x00,
}
