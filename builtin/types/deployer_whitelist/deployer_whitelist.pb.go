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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{0}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{0}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{1}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{2}
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

type AddDeployerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDeployerResponse) Reset()         { *m = AddDeployerResponse{} }
func (m *AddDeployerResponse) String() string { return proto.CompactTextString(m) }
func (*AddDeployerResponse) ProtoMessage()    {}
func (*AddDeployerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{3}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{4}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{5}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{6}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{7}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{8}
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
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{9}
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

type SetOverrideRequest struct {
	Flags                uint32   `protobuf:"varint,1,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOverrideRequest) Reset()         { *m = SetOverrideRequest{} }
func (m *SetOverrideRequest) String() string { return proto.CompactTextString(m) }
func (*SetOverrideRequest) ProtoMessage()    {}
func (*SetOverrideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{10}
}
func (m *SetOverrideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOverrideRequest.Unmarshal(m, b)
}
func (m *SetOverrideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOverrideRequest.Marshal(b, m, deterministic)
}
func (dst *SetOverrideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOverrideRequest.Merge(dst, src)
}
func (m *SetOverrideRequest) XXX_Size() int {
	return xxx_messageInfo_SetOverrideRequest.Size(m)
}
func (m *SetOverrideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOverrideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetOverrideRequest proto.InternalMessageInfo

func (m *SetOverrideRequest) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type GetOverrideRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOverrideRequest) Reset()         { *m = GetOverrideRequest{} }
func (m *GetOverrideRequest) String() string { return proto.CompactTextString(m) }
func (*GetOverrideRequest) ProtoMessage()    {}
func (*GetOverrideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{11}
}
func (m *GetOverrideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOverrideRequest.Unmarshal(m, b)
}
func (m *GetOverrideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOverrideRequest.Marshal(b, m, deterministic)
}
func (dst *GetOverrideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOverrideRequest.Merge(dst, src)
}
func (m *GetOverrideRequest) XXX_Size() int {
	return xxx_messageInfo_GetOverrideRequest.Size(m)
}
func (m *GetOverrideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOverrideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOverrideRequest proto.InternalMessageInfo

type GetOverrideResponse struct {
	Override             *Override `protobuf:"bytes,1,opt,name=override" json:"override,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetOverrideResponse) Reset()         { *m = GetOverrideResponse{} }
func (m *GetOverrideResponse) String() string { return proto.CompactTextString(m) }
func (*GetOverrideResponse) ProtoMessage()    {}
func (*GetOverrideResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{12}
}
func (m *GetOverrideResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOverrideResponse.Unmarshal(m, b)
}
func (m *GetOverrideResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOverrideResponse.Marshal(b, m, deterministic)
}
func (dst *GetOverrideResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOverrideResponse.Merge(dst, src)
}
func (m *GetOverrideResponse) XXX_Size() int {
	return xxx_messageInfo_GetOverrideResponse.Size(m)
}
func (m *GetOverrideResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOverrideResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOverrideResponse proto.InternalMessageInfo

func (m *GetOverrideResponse) GetOverride() *Override {
	if m != nil {
		return m.Override
	}
	return nil
}

type Override struct {
	Flags                uint32   `protobuf:"varint,1,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Override) Reset()         { *m = Override{} }
func (m *Override) String() string { return proto.CompactTextString(m) }
func (*Override) ProtoMessage()    {}
func (*Override) Descriptor() ([]byte, []int) {
	return fileDescriptor_deployer_whitelist_a8204d55b887d487, []int{13}
}
func (m *Override) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Override.Unmarshal(m, b)
}
func (m *Override) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Override.Marshal(b, m, deterministic)
}
func (dst *Override) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Override.Merge(dst, src)
}
func (m *Override) XXX_Size() int {
	return xxx_messageInfo_Override.Size(m)
}
func (m *Override) XXX_DiscardUnknown() {
	xxx_messageInfo_Override.DiscardUnknown(m)
}

var xxx_messageInfo_Override proto.InternalMessageInfo

func (m *Override) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func init() {
	proto.RegisterType((*InitRequest)(nil), "deployer_whitelist.InitRequest")
	proto.RegisterType((*Deployer)(nil), "deployer_whitelist.Deployer")
	proto.RegisterType((*AddDeployerRequest)(nil), "deployer_whitelist.AddDeployerRequest")
	proto.RegisterType((*AddDeployerResponse)(nil), "deployer_whitelist.AddDeployerResponse")
	proto.RegisterType((*GetDeployerRequest)(nil), "deployer_whitelist.GetDeployerRequest")
	proto.RegisterType((*GetDeployerResponse)(nil), "deployer_whitelist.GetDeployerResponse")
	proto.RegisterType((*RemoveDeployerRequest)(nil), "deployer_whitelist.RemoveDeployerRequest")
	proto.RegisterType((*RemoveDeployerResponse)(nil), "deployer_whitelist.RemoveDeployerResponse")
	proto.RegisterType((*ListDeployersRequest)(nil), "deployer_whitelist.ListDeployersRequest")
	proto.RegisterType((*ListDeployersResponse)(nil), "deployer_whitelist.ListDeployersResponse")
	proto.RegisterType((*SetOverrideRequest)(nil), "deployer_whitelist.SetOverrideRequest")
	proto.RegisterType((*GetOverrideRequest)(nil), "deployer_whitelist.GetOverrideRequest")
	proto.RegisterType((*GetOverrideResponse)(nil), "deployer_whitelist.GetOverrideResponse")
	proto.RegisterType((*Override)(nil), "deployer_whitelist.Override")
	proto.RegisterEnum("deployer_whitelist.Flags", Flags_name, Flags_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/deployer_whitelist/deployer_whitelist.proto", fileDescriptor_deployer_whitelist_a8204d55b887d487)
}

var fileDescriptor_deployer_whitelist_a8204d55b887d487 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcb, 0xeb, 0xda, 0x40,
	0x10, 0xc7, 0x1b, 0x9f, 0x71, 0xac, 0x10, 0xd6, 0x07, 0x52, 0x4a, 0x91, 0x3d, 0x89, 0xb4, 0x49,
	0x6b, 0x2f, 0xa5, 0x37, 0x8b, 0x56, 0x84, 0x6a, 0x20, 0x96, 0xe2, 0xad, 0x68, 0x33, 0xd5, 0xa5,
	0x31, 0x9b, 0x66, 0x57, 0xc5, 0xff, 0xbe, 0xc4, 0x6c, 0x62, 0xa3, 0xd6, 0xf2, 0xc3, 0x4b, 0xc8,
	0x7e, 0x67, 0xf2, 0xf9, 0xce, 0x63, 0x03, 0x8b, 0x35, 0x93, 0x9b, 0xdd, 0xca, 0xfc, 0xc1, 0xb7,
	0x96, 0xc7, 0xf9, 0xd6, 0x47, 0x79, 0xe0, 0xe1, 0x2f, 0x6b, 0xcd, 0xdf, 0x44, 0x47, 0x6b, 0xb5,
	0x63, 0x9e, 0x64, 0xbe, 0x25, 0x8f, 0x01, 0x0a, 0xcb, 0xc5, 0xc0, 0xe3, 0x47, 0x0c, 0xbf, 0x1f,
	0x36, 0x4c, 0xa2, 0xc7, 0x84, 0xbc, 0x21, 0x99, 0x41, 0xc8, 0x25, 0x27, 0xe4, 0x3a, 0xf2, 0xe2,
	0xed, 0x7f, 0xdc, 0x62, 0x97, 0xd3, 0x33, 0xa6, 0x50, 0x06, 0xd5, 0x89, 0xcf, 0xa4, 0x83, 0xbf,
	0x77, 0x28, 0x24, 0x79, 0x05, 0x45, 0x7e, 0xf0, 0x31, 0x6c, 0x6b, 0x1d, 0xad, 0x5b, 0xed, 0xeb,
	0xe6, 0xc0, 0x75, 0x43, 0x14, 0xc2, 0x89, 0x65, 0xf2, 0x11, 0x2a, 0x89, 0xad, 0x68, 0xe7, 0x3a,
	0xf9, 0x6e, 0xb5, 0xff, 0xd2, 0xbc, 0x51, 0xe2, 0x50, 0x49, 0xce, 0x39, 0x9d, 0x0e, 0x41, 0x4f,
	0x64, 0x42, 0xa1, 0xbc, 0x8c, 0xc9, 0x57, 0x4e, 0x49, 0x80, 0x34, 0xa0, 0xf8, 0xd3, 0x5b, 0xae,
	0x23, 0x1f, 0xad, 0x5b, 0x73, 0xe2, 0x03, 0x5d, 0x00, 0x19, 0xb8, 0x6e, 0xca, 0x57, 0x75, 0xbf,
	0x86, 0xe7, 0x89, 0x51, 0xc4, 0xb9, 0x82, 0x66, 0xa2, 0xff, 0x20, 0x37, 0xa1, 0x9e, 0x21, 0x8b,
	0x80, 0xfb, 0x02, 0xe9, 0x27, 0x20, 0x63, 0x94, 0x0f, 0x19, 0x52, 0x1b, 0xea, 0x19, 0x46, 0x8c,
	0x26, 0x1f, 0x40, 0x4f, 0xd2, 0x14, 0xe0, 0xfe, 0x30, 0xd3, 0x6c, 0x3a, 0x82, 0xa6, 0x83, 0x5b,
	0xbe, 0xc7, 0xc7, 0xea, 0x6a, 0x43, 0xeb, 0x12, 0xa3, 0xba, 0x6e, 0x41, 0xe3, 0x0b, 0x13, 0x69,
	0xc9, 0x42, 0xf1, 0xe9, 0x1c, 0x9a, 0x17, 0xba, 0xea, 0x25, 0x73, 0x33, 0xb4, 0xa7, 0xdd, 0x8c,
	0x1e, 0x90, 0x39, 0x4a, 0x7b, 0x8f, 0x61, 0xc8, 0x5c, 0x4c, 0x5a, 0x49, 0xb7, 0xa4, 0xfd, 0xbd,
	0xa5, 0xc6, 0x69, 0x1d, 0x17, 0xb9, 0x6a, 0xc0, 0x67, 0xf5, 0x3c, 0x60, 0xae, 0xb4, 0x7b, 0x03,
	0x4e, 0xbf, 0x4b, 0xb3, 0x69, 0x07, 0xf4, 0x44, 0xbd, 0x5d, 0x48, 0xef, 0x1d, 0x14, 0x3f, 0x47,
	0x2f, 0x44, 0x87, 0xc2, 0xcc, 0x9e, 0x8d, 0x8c, 0x67, 0xa4, 0x04, 0xb9, 0xb1, 0x6d, 0x68, 0xa4,
	0x0c, 0xf9, 0xd1, 0xb7, 0xa9, 0x91, 0x23, 0x35, 0xa8, 0x4c, 0x27, 0x63, 0x67, 0xf0, 0x75, 0x62,
	0xcf, 0x8c, 0xc2, 0xaa, 0x74, 0xfa, 0xe7, 0xde, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x4e,
	0xe2, 0xfd, 0x15, 0x04, 0x00, 0x00,
}
