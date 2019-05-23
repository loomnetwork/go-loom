// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/user_deployer_whitelist/user_deployer_whitelist.proto

package user_deployer_whitelist

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import deployer_whitelist "github.com/loomnetwork/go-loom/builtin/types/deployer_whitelist"
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

type TierID int32

const (
	TierID_DEFAULT TierID = 0
)

var TierID_name = map[int32]string{
	0: "DEFAULT",
}
var TierID_value = map[string]int32{
	"DEFAULT": 0,
}

func (x TierID) String() string {
	return proto.EnumName(TierID_name, int32(x))
}
func (TierID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{0}
}

type Tier struct {
	TierID               TierID   `protobuf:"varint,1,opt,name=id,proto3,enum=user_deployer_whitelist.TierID" json:"id,omitempty"`
	Fee                  uint64   `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tier) Reset()         { *m = Tier{} }
func (m *Tier) String() string { return proto.CompactTextString(m) }
func (*Tier) ProtoMessage()    {}
func (*Tier) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{0}
}
func (m *Tier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tier.Unmarshal(m, b)
}
func (m *Tier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tier.Marshal(b, m, deterministic)
}
func (dst *Tier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tier.Merge(dst, src)
}
func (m *Tier) XXX_Size() int {
	return xxx_messageInfo_Tier.Size(m)
}
func (m *Tier) XXX_DiscardUnknown() {
	xxx_messageInfo_Tier.DiscardUnknown(m)
}

var xxx_messageInfo_Tier proto.InternalMessageInfo

func (m *Tier) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

func (m *Tier) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Tier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TierInfo struct {
	Tiers                []*Tier  `protobuf:"bytes,1,rep,name=tiers" json:"tiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TierInfo) Reset()         { *m = TierInfo{} }
func (m *TierInfo) String() string { return proto.CompactTextString(m) }
func (*TierInfo) ProtoMessage()    {}
func (*TierInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{1}
}
func (m *TierInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TierInfo.Unmarshal(m, b)
}
func (m *TierInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TierInfo.Marshal(b, m, deterministic)
}
func (dst *TierInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TierInfo.Merge(dst, src)
}
func (m *TierInfo) XXX_Size() int {
	return xxx_messageInfo_TierInfo.Size(m)
}
func (m *TierInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TierInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TierInfo proto.InternalMessageInfo

func (m *TierInfo) GetTiers() []*Tier {
	if m != nil {
		return m.Tiers
	}
	return nil
}

type InitRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	TierInfo             *TierInfo      `protobuf:"bytes,2,opt,name=tier_info,json=tierInfo" json:"tier_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{2}
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

func (m *InitRequest) GetTierInfo() *TierInfo {
	if m != nil {
		return m.TierInfo
	}
	return nil
}

type DeployerContract struct {
	ContractAddress      *types.Address `protobuf:"bytes,1,opt,name=contractAddress" json:"contractAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeployerContract) Reset()         { *m = DeployerContract{} }
func (m *DeployerContract) String() string { return proto.CompactTextString(m) }
func (*DeployerContract) ProtoMessage()    {}
func (*DeployerContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{3}
}
func (m *DeployerContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployerContract.Unmarshal(m, b)
}
func (m *DeployerContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployerContract.Marshal(b, m, deterministic)
}
func (dst *DeployerContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployerContract.Merge(dst, src)
}
func (m *DeployerContract) XXX_Size() int {
	return xxx_messageInfo_DeployerContract.Size(m)
}
func (m *DeployerContract) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployerContract.DiscardUnknown(m)
}

var xxx_messageInfo_DeployerContract proto.InternalMessageInfo

func (m *DeployerContract) GetContractAddress() *types.Address {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

type WhitelistUserDeployerRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployerAddr" json:"deployerAddr,omitempty"`
	TierID               TierID         `protobuf:"varint,2,opt,name=tier_id,json=tierId,proto3,enum=user_deployer_whitelist.TierID" json:"tier_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WhitelistUserDeployerRequest) Reset()         { *m = WhitelistUserDeployerRequest{} }
func (m *WhitelistUserDeployerRequest) String() string { return proto.CompactTextString(m) }
func (*WhitelistUserDeployerRequest) ProtoMessage()    {}
func (*WhitelistUserDeployerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{4}
}
func (m *WhitelistUserDeployerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhitelistUserDeployerRequest.Unmarshal(m, b)
}
func (m *WhitelistUserDeployerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhitelistUserDeployerRequest.Marshal(b, m, deterministic)
}
func (dst *WhitelistUserDeployerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistUserDeployerRequest.Merge(dst, src)
}
func (m *WhitelistUserDeployerRequest) XXX_Size() int {
	return xxx_messageInfo_WhitelistUserDeployerRequest.Size(m)
}
func (m *WhitelistUserDeployerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistUserDeployerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistUserDeployerRequest proto.InternalMessageInfo

func (m *WhitelistUserDeployerRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

func (m *WhitelistUserDeployerRequest) GetTierID() TierID {
	if m != nil {
		return m.TierID
	}
	return TierID_DEFAULT
}

type UserState struct {
	UserAddr             *types.Address   `protobuf:"bytes,1,opt,name=userAddr" json:"userAddr,omitempty"`
	Deployers            []*types.Address `protobuf:"bytes,2,rep,name=deployers" json:"deployers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserState) Reset()         { *m = UserState{} }
func (m *UserState) String() string { return proto.CompactTextString(m) }
func (*UserState) ProtoMessage()    {}
func (*UserState) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{5}
}
func (m *UserState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserState.Unmarshal(m, b)
}
func (m *UserState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserState.Marshal(b, m, deterministic)
}
func (dst *UserState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserState.Merge(dst, src)
}
func (m *UserState) XXX_Size() int {
	return xxx_messageInfo_UserState.Size(m)
}
func (m *UserState) XXX_DiscardUnknown() {
	xxx_messageInfo_UserState.DiscardUnknown(m)
}

var xxx_messageInfo_UserState proto.InternalMessageInfo

func (m *UserState) GetUserAddr() *types.Address {
	if m != nil {
		return m.UserAddr
	}
	return nil
}

func (m *UserState) GetDeployers() []*types.Address {
	if m != nil {
		return m.Deployers
	}
	return nil
}

type UserDeployerState struct {
	Address              *types.Address      `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Contracts            []*DeployerContract `protobuf:"bytes,2,rep,name=contracts" json:"contracts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserDeployerState) Reset()         { *m = UserDeployerState{} }
func (m *UserDeployerState) String() string { return proto.CompactTextString(m) }
func (*UserDeployerState) ProtoMessage()    {}
func (*UserDeployerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{6}
}
func (m *UserDeployerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeployerState.Unmarshal(m, b)
}
func (m *UserDeployerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeployerState.Marshal(b, m, deterministic)
}
func (dst *UserDeployerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeployerState.Merge(dst, src)
}
func (m *UserDeployerState) XXX_Size() int {
	return xxx_messageInfo_UserDeployerState.Size(m)
}
func (m *UserDeployerState) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeployerState.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeployerState proto.InternalMessageInfo

func (m *UserDeployerState) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserDeployerState) GetContracts() []*DeployerContract {
	if m != nil {
		return m.Contracts
	}
	return nil
}

type GetUserDeployersRequest struct {
	UserAddr             *types.Address `protobuf:"bytes,1,opt,name=userAddr" json:"userAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUserDeployersRequest) Reset()         { *m = GetUserDeployersRequest{} }
func (m *GetUserDeployersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserDeployersRequest) ProtoMessage()    {}
func (*GetUserDeployersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{7}
}
func (m *GetUserDeployersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDeployersRequest.Unmarshal(m, b)
}
func (m *GetUserDeployersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDeployersRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserDeployersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDeployersRequest.Merge(dst, src)
}
func (m *GetUserDeployersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserDeployersRequest.Size(m)
}
func (m *GetUserDeployersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDeployersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDeployersRequest proto.InternalMessageInfo

func (m *GetUserDeployersRequest) GetUserAddr() *types.Address {
	if m != nil {
		return m.UserAddr
	}
	return nil
}

type GetUserDeployersResponse struct {
	Deployers            []*deployer_whitelist.Deployer `protobuf:"bytes,1,rep,name=deployers" json:"deployers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetUserDeployersResponse) Reset()         { *m = GetUserDeployersResponse{} }
func (m *GetUserDeployersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserDeployersResponse) ProtoMessage()    {}
func (*GetUserDeployersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{8}
}
func (m *GetUserDeployersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDeployersResponse.Unmarshal(m, b)
}
func (m *GetUserDeployersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDeployersResponse.Marshal(b, m, deterministic)
}
func (dst *GetUserDeployersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDeployersResponse.Merge(dst, src)
}
func (m *GetUserDeployersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserDeployersResponse.Size(m)
}
func (m *GetUserDeployersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDeployersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDeployersResponse proto.InternalMessageInfo

func (m *GetUserDeployersResponse) GetDeployers() []*deployer_whitelist.Deployer {
	if m != nil {
		return m.Deployers
	}
	return nil
}

type GetDeployedContractsRequest struct {
	DeployerAddr         *types.Address `protobuf:"bytes,1,opt,name=deployerAddr" json:"deployerAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDeployedContractsRequest) Reset()         { *m = GetDeployedContractsRequest{} }
func (m *GetDeployedContractsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeployedContractsRequest) ProtoMessage()    {}
func (*GetDeployedContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{9}
}
func (m *GetDeployedContractsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeployedContractsRequest.Unmarshal(m, b)
}
func (m *GetDeployedContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeployedContractsRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeployedContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeployedContractsRequest.Merge(dst, src)
}
func (m *GetDeployedContractsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeployedContractsRequest.Size(m)
}
func (m *GetDeployedContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeployedContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeployedContractsRequest proto.InternalMessageInfo

func (m *GetDeployedContractsRequest) GetDeployerAddr() *types.Address {
	if m != nil {
		return m.DeployerAddr
	}
	return nil
}

type GetDeployedContractsResponse struct {
	ContractAddresses    []*DeployerContract `protobuf:"bytes,1,rep,name=contractAddresses" json:"contractAddresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetDeployedContractsResponse) Reset()         { *m = GetDeployedContractsResponse{} }
func (m *GetDeployedContractsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeployedContractsResponse) ProtoMessage()    {}
func (*GetDeployedContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd, []int{10}
}
func (m *GetDeployedContractsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeployedContractsResponse.Unmarshal(m, b)
}
func (m *GetDeployedContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeployedContractsResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeployedContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeployedContractsResponse.Merge(dst, src)
}
func (m *GetDeployedContractsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeployedContractsResponse.Size(m)
}
func (m *GetDeployedContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeployedContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeployedContractsResponse proto.InternalMessageInfo

func (m *GetDeployedContractsResponse) GetContractAddresses() []*DeployerContract {
	if m != nil {
		return m.ContractAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Tier)(nil), "user_deployer_whitelist.Tier")
	proto.RegisterType((*TierInfo)(nil), "user_deployer_whitelist.TierInfo")
	proto.RegisterType((*InitRequest)(nil), "user_deployer_whitelist.InitRequest")
	proto.RegisterType((*DeployerContract)(nil), "user_deployer_whitelist.DeployerContract")
	proto.RegisterType((*WhitelistUserDeployerRequest)(nil), "user_deployer_whitelist.WhitelistUserDeployerRequest")
	proto.RegisterType((*UserState)(nil), "user_deployer_whitelist.UserState")
	proto.RegisterType((*UserDeployerState)(nil), "user_deployer_whitelist.UserDeployerState")
	proto.RegisterType((*GetUserDeployersRequest)(nil), "user_deployer_whitelist.GetUserDeployersRequest")
	proto.RegisterType((*GetUserDeployersResponse)(nil), "user_deployer_whitelist.GetUserDeployersResponse")
	proto.RegisterType((*GetDeployedContractsRequest)(nil), "user_deployer_whitelist.GetDeployedContractsRequest")
	proto.RegisterType((*GetDeployedContractsResponse)(nil), "user_deployer_whitelist.GetDeployedContractsResponse")
	proto.RegisterEnum("user_deployer_whitelist.TierID", TierID_name, TierID_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/user_deployer_whitelist/user_deployer_whitelist.proto", fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd)
}

var fileDescriptor_user_deployer_whitelist_cdf46aa52180c2bd = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0x76, 0x81, 0xf2, 0x73, 0x30, 0x4a, 0x27, 0x31, 0xdd, 0x54, 0xb4, 0x38, 0x31, 0x06, 0x8d,
	0x65, 0x0d, 0xbd, 0xd3, 0xc4, 0xa6, 0x8a, 0x25, 0x8d, 0x5e, 0xad, 0xad, 0xd5, 0x2b, 0x02, 0xec,
	0x59, 0x3a, 0x11, 0x66, 0x70, 0x66, 0x36, 0xa4, 0x77, 0x3e, 0x83, 0x0f, 0xe8, 0x85, 0x4f, 0x62,
	0x76, 0x76, 0x47, 0x28, 0xb0, 0x28, 0xde, 0x6c, 0xce, 0xcc, 0xf9, 0xf9, 0x7e, 0xe6, 0x64, 0x61,
	0x30, 0x62, 0xfa, 0x2a, 0x1a, 0xb4, 0x86, 0x62, 0xe2, 0x8d, 0x85, 0x98, 0x70, 0xd4, 0x33, 0x21,
	0xbf, 0x7a, 0x23, 0x71, 0x18, 0x1f, 0xbd, 0x41, 0xc4, 0xc6, 0x9a, 0x71, 0x4f, 0x5f, 0x4f, 0x51,
	0x79, 0x91, 0x42, 0xd9, 0x0b, 0x70, 0x3a, 0x16, 0xd7, 0x28, 0x7b, 0xb3, 0x2b, 0xa6, 0x71, 0xcc,
	0x94, 0xce, 0xba, 0x6f, 0x4d, 0xa5, 0xd0, 0x82, 0xec, 0x65, 0xa4, 0xf7, 0x0f, 0x17, 0xc0, 0x47,
	0x62, 0x24, 0x3c, 0x53, 0x3f, 0x88, 0x42, 0x73, 0x32, 0x07, 0x13, 0x25, 0x73, 0xf6, 0x5f, 0xfc,
	0x85, 0x6b, 0xc2, 0xd1, 0x7c, 0xd3, 0x8e, 0xcf, 0x5b, 0xa9, 0x5b, 0x23, 0x2c, 0x4b, 0x13, 0x65,
	0x50, 0x38, 0x67, 0x28, 0xc9, 0x2b, 0xc8, 0xb1, 0xc0, 0x75, 0x1a, 0x4e, 0xf3, 0x4e, 0xfb, 0xa0,
	0x95, 0xe5, 0x43, 0x5c, 0x7a, 0xd6, 0x79, 0x03, 0xbf, 0x7e, 0x1e, 0x14, 0x93, 0xd8, 0xcf, 0xb1,
	0x80, 0xd4, 0x20, 0x1f, 0x22, 0xba, 0xb9, 0x86, 0xd3, 0x2c, 0xf8, 0x71, 0x48, 0x08, 0x14, 0x78,
	0x7f, 0x82, 0x6e, 0xbe, 0xe1, 0x34, 0x2b, 0xbe, 0x89, 0xe9, 0x31, 0x94, 0x4d, 0x0f, 0x0f, 0x05,
	0x39, 0x82, 0x1d, 0xcd, 0x50, 0x2a, 0xd7, 0x69, 0xe4, 0x9b, 0xd5, 0xf6, 0x83, 0x8d, 0x88, 0x7e,
	0x52, 0x4b, 0x27, 0x50, 0x3d, 0xe3, 0x4c, 0xfb, 0xf8, 0x2d, 0x42, 0xa5, 0xc9, 0x43, 0xd8, 0x11,
	0x33, 0x8e, 0xd2, 0xb0, 0xae, 0xb6, 0xcb, 0xad, 0x93, 0x20, 0x90, 0xa8, 0x94, 0x9f, 0x5c, 0x93,
	0xd7, 0x50, 0x89, 0xfb, 0x7a, 0x8c, 0x87, 0xc2, 0x70, 0xab, 0xb6, 0x1f, 0x6d, 0x56, 0xc6, 0x43,
	0xe1, 0x97, 0x75, 0x1a, 0xd1, 0x53, 0xa8, 0x75, 0xd2, 0xc2, 0xb7, 0x82, 0x6b, 0xd9, 0x1f, 0x6a,
	0xd2, 0x86, 0xbb, 0xc3, 0x34, 0x4e, 0xd1, 0x56, 0xd0, 0x97, 0x0b, 0xe8, 0x0f, 0x07, 0xea, 0x97,
	0x16, 0xe8, 0x42, 0xa1, 0xb4, 0x53, 0xad, 0x90, 0xe7, 0x70, 0xdb, 0x32, 0x8a, 0x7b, 0x56, 0x26,
	0xde, 0xc8, 0x92, 0x0e, 0x94, 0x12, 0x59, 0x81, 0x11, 0xb5, 0xe5, 0x73, 0x15, 0x8d, 0xbc, 0x80,
	0x7e, 0x81, 0x4a, 0x4c, 0xe5, 0xa3, 0xee, 0x6b, 0x24, 0x8f, 0xa1, 0x1c, 0x8f, 0x58, 0x0b, 0xfe,
	0x27, 0x43, 0x9e, 0x40, 0xc5, 0x62, 0x28, 0x37, 0x67, 0xde, 0x6d, 0x5e, 0x36, 0x4f, 0xd1, 0xef,
	0x0e, 0xec, 0x2e, 0xca, 0x4c, 0x30, 0x28, 0x94, 0xfa, 0x19, 0x8e, 0xd9, 0x04, 0xe9, 0x42, 0xc5,
	0x9a, 0x67, 0x11, 0x9e, 0x66, 0x8a, 0x5b, 0x7e, 0x1b, 0x7f, 0xde, 0x4b, 0x8f, 0x61, 0xaf, 0x8b,
	0x37, 0xbc, 0x56, 0xd6, 0xec, 0x7f, 0xd2, 0x4a, 0x3f, 0x81, 0xbb, 0x3a, 0x40, 0x4d, 0x05, 0x57,
	0x48, 0x5e, 0x2e, 0xfa, 0x90, 0xec, 0x6f, 0xbd, 0xb5, 0x81, 0xe0, 0xa2, 0x37, 0xef, 0xe1, 0x7e,
	0x17, 0x75, 0x9a, 0x09, 0x2c, 0x75, 0xf5, 0x5f, 0x9b, 0x40, 0x67, 0x50, 0x5f, 0x3f, 0x2c, 0x25,
	0x7a, 0x09, 0xbb, 0x4b, 0xbb, 0x88, 0x96, 0xf0, 0x16, 0xb6, 0xae, 0xce, 0x78, 0x76, 0x0f, 0xd2,
	0x75, 0x22, 0x55, 0x28, 0x75, 0xde, 0x9d, 0x9e, 0x5c, 0x7c, 0x38, 0xaf, 0xdd, 0x1a, 0x14, 0xcd,
	0x2f, 0xe5, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x35, 0x59, 0x05, 0x8c, 0x05, 0x00,
	0x00,
}
