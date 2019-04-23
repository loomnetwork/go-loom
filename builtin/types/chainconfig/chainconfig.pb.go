// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/chainconfig/chainconfig.proto

package chainconfig

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

type Feature_FeatureStatus int32

const (
	Feature_PENDING  Feature_FeatureStatus = 0
	Feature_WAITING  Feature_FeatureStatus = 1
	Feature_ENABLED  Feature_FeatureStatus = 2
	Feature_DISABLED Feature_FeatureStatus = 3
)

var Feature_FeatureStatus_name = map[int32]string{
	0: "PENDING",
	1: "WAITING",
	2: "ENABLED",
	3: "DISABLED",
}
var Feature_FeatureStatus_value = map[string]int32{
	"PENDING":  0,
	"WAITING":  1,
	"ENABLED":  2,
	"DISABLED": 3,
}

func (x Feature_FeatureStatus) String() string {
	return proto.EnumName(Feature_FeatureStatus_name, int32(x))
}
func (Feature_FeatureStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{5, 0}
}

type InitRequest struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Params               *Params        `protobuf:"bytes,2,opt,name=params" json:"params,omitempty"`
	Features             []*Feature     `protobuf:"bytes,3,rep,name=features" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{0}
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

func (m *InitRequest) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *InitRequest) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type GetFeatureRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeatureRequest) Reset()         { *m = GetFeatureRequest{} }
func (m *GetFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeatureRequest) ProtoMessage()    {}
func (*GetFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{1}
}
func (m *GetFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureRequest.Unmarshal(m, b)
}
func (m *GetFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureRequest.Merge(dst, src)
}
func (m *GetFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeatureRequest.Size(m)
}
func (m *GetFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureRequest proto.InternalMessageInfo

func (m *GetFeatureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddFeatureRequest struct {
	Names                []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	BuildNumber          uint64   `protobuf:"varint,2,opt,name=build_number,json=buildNumber,proto3" json:"build_number,omitempty"`
	AutoEnabled          bool     `protobuf:"varint,3,opt,name=auto_enabled,json=autoEnabled,proto3" json:"auto_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeatureRequest) Reset()         { *m = AddFeatureRequest{} }
func (m *AddFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*AddFeatureRequest) ProtoMessage()    {}
func (*AddFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{2}
}
func (m *AddFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeatureRequest.Unmarshal(m, b)
}
func (m *AddFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *AddFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeatureRequest.Merge(dst, src)
}
func (m *AddFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_AddFeatureRequest.Size(m)
}
func (m *AddFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeatureRequest proto.InternalMessageInfo

func (m *AddFeatureRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *AddFeatureRequest) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *AddFeatureRequest) GetAutoEnabled() bool {
	if m != nil {
		return m.AutoEnabled
	}
	return false
}

type AddFeatureResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeatureResponse) Reset()         { *m = AddFeatureResponse{} }
func (m *AddFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*AddFeatureResponse) ProtoMessage()    {}
func (*AddFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{3}
}
func (m *AddFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeatureResponse.Unmarshal(m, b)
}
func (m *AddFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *AddFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeatureResponse.Merge(dst, src)
}
func (m *AddFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_AddFeatureResponse.Size(m)
}
func (m *AddFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeatureResponse proto.InternalMessageInfo

type GetFeatureResponse struct {
	Feature              *Feature `protobuf:"bytes,1,opt,name=feature" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeatureResponse) Reset()         { *m = GetFeatureResponse{} }
func (m *GetFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeatureResponse) ProtoMessage()    {}
func (*GetFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{4}
}
func (m *GetFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureResponse.Unmarshal(m, b)
}
func (m *GetFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *GetFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureResponse.Merge(dst, src)
}
func (m *GetFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeatureResponse.Size(m)
}
func (m *GetFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureResponse proto.InternalMessageInfo

func (m *GetFeatureResponse) GetFeature() *Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

type Feature struct {
	Name       string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status     Feature_FeatureStatus `protobuf:"varint,2,opt,name=status,proto3,enum=chainconfig.Feature_FeatureStatus" json:"status,omitempty"`
	Validators []*types.Address      `protobuf:"bytes,3,rep,name=validators" json:"validators,omitempty"`
	// Percentage of feature enabled validators exceeds vote_threshold at this block_height
	BlockHeight uint64 `protobuf:"varint,4,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Feature changes status from `pending` to `waiting` with at this percentage
	Percentage uint64 `protobuf:"varint,5,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Minimum build number that supports this feature
	BuildNumber uint64 `protobuf:"varint,6,opt,name=build_number,json=buildNumber,proto3" json:"build_number,omitempty"`
	// Pending feature is automatically enabled if it is supported in current build
	AutoEnabled          bool     `protobuf:"varint,7,opt,name=auto_enabled,json=autoEnabled,proto3" json:"auto_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{5}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetStatus() Feature_FeatureStatus {
	if m != nil {
		return m.Status
	}
	return Feature_PENDING
}

func (m *Feature) GetValidators() []*types.Address {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *Feature) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Feature) GetPercentage() uint64 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *Feature) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *Feature) GetAutoEnabled() bool {
	if m != nil {
		return m.AutoEnabled
	}
	return false
}

type Params struct {
	VoteThreshold         uint64   `protobuf:"varint,1,opt,name=vote_threshold,json=voteThreshold,proto3" json:"vote_threshold,omitempty"`
	NumBlockConfirmations uint64   `protobuf:"varint,2,opt,name=num_block_confirmations,json=numBlockConfirmations,proto3" json:"num_block_confirmations,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{6}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (dst *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(dst, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetVoteThreshold() uint64 {
	if m != nil {
		return m.VoteThreshold
	}
	return 0
}

func (m *Params) GetNumBlockConfirmations() uint64 {
	if m != nil {
		return m.NumBlockConfirmations
	}
	return 0
}

type SetParamsRequest struct {
	Params               *Params  `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetParamsRequest) Reset()         { *m = SetParamsRequest{} }
func (m *SetParamsRequest) String() string { return proto.CompactTextString(m) }
func (*SetParamsRequest) ProtoMessage()    {}
func (*SetParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{7}
}
func (m *SetParamsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetParamsRequest.Unmarshal(m, b)
}
func (m *SetParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetParamsRequest.Marshal(b, m, deterministic)
}
func (dst *SetParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetParamsRequest.Merge(dst, src)
}
func (m *SetParamsRequest) XXX_Size() int {
	return xxx_messageInfo_SetParamsRequest.Size(m)
}
func (m *SetParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetParamsRequest proto.InternalMessageInfo

func (m *SetParamsRequest) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type GetParamsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParamsRequest) Reset()         { *m = GetParamsRequest{} }
func (m *GetParamsRequest) String() string { return proto.CompactTextString(m) }
func (*GetParamsRequest) ProtoMessage()    {}
func (*GetParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{8}
}
func (m *GetParamsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParamsRequest.Unmarshal(m, b)
}
func (m *GetParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParamsRequest.Marshal(b, m, deterministic)
}
func (dst *GetParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParamsRequest.Merge(dst, src)
}
func (m *GetParamsRequest) XXX_Size() int {
	return xxx_messageInfo_GetParamsRequest.Size(m)
}
func (m *GetParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParamsRequest proto.InternalMessageInfo

type GetParamsResponse struct {
	Params               *Params  `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParamsResponse) Reset()         { *m = GetParamsResponse{} }
func (m *GetParamsResponse) String() string { return proto.CompactTextString(m) }
func (*GetParamsResponse) ProtoMessage()    {}
func (*GetParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{9}
}
func (m *GetParamsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParamsResponse.Unmarshal(m, b)
}
func (m *GetParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParamsResponse.Marshal(b, m, deterministic)
}
func (dst *GetParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParamsResponse.Merge(dst, src)
}
func (m *GetParamsResponse) XXX_Size() int {
	return xxx_messageInfo_GetParamsResponse.Size(m)
}
func (m *GetParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetParamsResponse proto.InternalMessageInfo

func (m *GetParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type ListFeaturesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFeaturesRequest) Reset()         { *m = ListFeaturesRequest{} }
func (m *ListFeaturesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFeaturesRequest) ProtoMessage()    {}
func (*ListFeaturesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{10}
}
func (m *ListFeaturesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFeaturesRequest.Unmarshal(m, b)
}
func (m *ListFeaturesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFeaturesRequest.Marshal(b, m, deterministic)
}
func (dst *ListFeaturesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFeaturesRequest.Merge(dst, src)
}
func (m *ListFeaturesRequest) XXX_Size() int {
	return xxx_messageInfo_ListFeaturesRequest.Size(m)
}
func (m *ListFeaturesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFeaturesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFeaturesRequest proto.InternalMessageInfo

type ListFeaturesResponse struct {
	Features             []*Feature `protobuf:"bytes,1,rep,name=features" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListFeaturesResponse) Reset()         { *m = ListFeaturesResponse{} }
func (m *ListFeaturesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFeaturesResponse) ProtoMessage()    {}
func (*ListFeaturesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{11}
}
func (m *ListFeaturesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFeaturesResponse.Unmarshal(m, b)
}
func (m *ListFeaturesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFeaturesResponse.Marshal(b, m, deterministic)
}
func (dst *ListFeaturesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFeaturesResponse.Merge(dst, src)
}
func (m *ListFeaturesResponse) XXX_Size() int {
	return xxx_messageInfo_ListFeaturesResponse.Size(m)
}
func (m *ListFeaturesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFeaturesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFeaturesResponse proto.InternalMessageInfo

func (m *ListFeaturesResponse) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type EnableFeatureRequest struct {
	Names                []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnableFeatureRequest) Reset()         { *m = EnableFeatureRequest{} }
func (m *EnableFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*EnableFeatureRequest) ProtoMessage()    {}
func (*EnableFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{12}
}
func (m *EnableFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnableFeatureRequest.Unmarshal(m, b)
}
func (m *EnableFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnableFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *EnableFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnableFeatureRequest.Merge(dst, src)
}
func (m *EnableFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_EnableFeatureRequest.Size(m)
}
func (m *EnableFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnableFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnableFeatureRequest proto.InternalMessageInfo

func (m *EnableFeatureRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type EnableFeatureResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnableFeatureResponse) Reset()         { *m = EnableFeatureResponse{} }
func (m *EnableFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*EnableFeatureResponse) ProtoMessage()    {}
func (*EnableFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{13}
}
func (m *EnableFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnableFeatureResponse.Unmarshal(m, b)
}
func (m *EnableFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnableFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *EnableFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnableFeatureResponse.Merge(dst, src)
}
func (m *EnableFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_EnableFeatureResponse.Size(m)
}
func (m *EnableFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnableFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnableFeatureResponse proto.InternalMessageInfo

type UpdateFeatureRequest struct {
	Names                []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	BuildNumber          uint64   `protobuf:"varint,2,opt,name=build_number,json=buildNumber,proto3" json:"build_number,omitempty"`
	AutoEnabled          bool     `protobuf:"varint,3,opt,name=auto_enabled,json=autoEnabled,proto3" json:"auto_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFeatureRequest) Reset()         { *m = UpdateFeatureRequest{} }
func (m *UpdateFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFeatureRequest) ProtoMessage()    {}
func (*UpdateFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{14}
}
func (m *UpdateFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeatureRequest.Unmarshal(m, b)
}
func (m *UpdateFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeatureRequest.Merge(dst, src)
}
func (m *UpdateFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFeatureRequest.Size(m)
}
func (m *UpdateFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeatureRequest proto.InternalMessageInfo

func (m *UpdateFeatureRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *UpdateFeatureRequest) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *UpdateFeatureRequest) GetAutoEnabled() bool {
	if m != nil {
		return m.AutoEnabled
	}
	return false
}

type UpdateFeatureResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFeatureResponse) Reset()         { *m = UpdateFeatureResponse{} }
func (m *UpdateFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateFeatureResponse) ProtoMessage()    {}
func (*UpdateFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e143e9cac04517af, []int{15}
}
func (m *UpdateFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeatureResponse.Unmarshal(m, b)
}
func (m *UpdateFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeatureResponse.Merge(dst, src)
}
func (m *UpdateFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateFeatureResponse.Size(m)
}
func (m *UpdateFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeatureResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InitRequest)(nil), "chainconfig.InitRequest")
	proto.RegisterType((*GetFeatureRequest)(nil), "chainconfig.GetFeatureRequest")
	proto.RegisterType((*AddFeatureRequest)(nil), "chainconfig.AddFeatureRequest")
	proto.RegisterType((*AddFeatureResponse)(nil), "chainconfig.AddFeatureResponse")
	proto.RegisterType((*GetFeatureResponse)(nil), "chainconfig.GetFeatureResponse")
	proto.RegisterType((*Feature)(nil), "chainconfig.Feature")
	proto.RegisterType((*Params)(nil), "chainconfig.Params")
	proto.RegisterType((*SetParamsRequest)(nil), "chainconfig.SetParamsRequest")
	proto.RegisterType((*GetParamsRequest)(nil), "chainconfig.GetParamsRequest")
	proto.RegisterType((*GetParamsResponse)(nil), "chainconfig.GetParamsResponse")
	proto.RegisterType((*ListFeaturesRequest)(nil), "chainconfig.ListFeaturesRequest")
	proto.RegisterType((*ListFeaturesResponse)(nil), "chainconfig.ListFeaturesResponse")
	proto.RegisterType((*EnableFeatureRequest)(nil), "chainconfig.EnableFeatureRequest")
	proto.RegisterType((*EnableFeatureResponse)(nil), "chainconfig.EnableFeatureResponse")
	proto.RegisterType((*UpdateFeatureRequest)(nil), "chainconfig.UpdateFeatureRequest")
	proto.RegisterType((*UpdateFeatureResponse)(nil), "chainconfig.UpdateFeatureResponse")
	proto.RegisterEnum("chainconfig.Feature_FeatureStatus", Feature_FeatureStatus_name, Feature_FeatureStatus_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/chainconfig/chainconfig.proto", fileDescriptor_chainconfig_e143e9cac04517af)
}

var fileDescriptor_chainconfig_e143e9cac04517af = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x51, 0x4f, 0xdb, 0x3c,
	0x14, 0xfd, 0x42, 0x4b, 0x5b, 0x6e, 0x00, 0x15, 0xd3, 0x8a, 0xe8, 0x7b, 0x40, 0x9d, 0xa5, 0x69,
	0x95, 0xb6, 0xb5, 0x88, 0x49, 0x7b, 0xd8, 0xcb, 0x56, 0xd6, 0x0e, 0x3a, 0xa1, 0x0a, 0x05, 0xa6,
	0x3d, 0x56, 0x6e, 0x63, 0x92, 0x88, 0xc4, 0xce, 0x6c, 0x07, 0xb4, 0xff, 0xb0, 0x5f, 0xb7, 0x5f,
	0x34, 0xc5, 0x76, 0x20, 0x14, 0x24, 0xd8, 0xcb, 0x5e, 0xda, 0xdc, 0x73, 0x4f, 0xae, 0xef, 0x3d,
	0xbe, 0x27, 0xf0, 0x35, 0x8c, 0x55, 0x94, 0x2f, 0x06, 0x4b, 0x9e, 0x0e, 0x13, 0xce, 0x53, 0x46,
	0xd5, 0x0d, 0x17, 0x57, 0xc3, 0x90, 0xbf, 0x2d, 0xc2, 0xe1, 0x22, 0x8f, 0x13, 0x15, 0xb3, 0xa1,
	0xfa, 0x99, 0x51, 0x39, 0x5c, 0x46, 0x24, 0x66, 0x4b, 0xce, 0x2e, 0xe3, 0xb0, 0xfa, 0x3c, 0xc8,
	0x04, 0x57, 0x1c, 0xb9, 0x15, 0xe8, 0xff, 0x83, 0x27, 0x0a, 0x9b, 0x82, 0xfa, 0xd7, 0xbc, 0x8e,
	0x7f, 0x39, 0xe0, 0x4e, 0x59, 0xac, 0x7c, 0xfa, 0x23, 0xa7, 0x52, 0xa1, 0x7d, 0x58, 0xe7, 0x37,
	0x8c, 0x0a, 0xcf, 0xe9, 0x39, 0x7d, 0xf7, 0xb0, 0x35, 0x18, 0x05, 0x81, 0xa0, 0x52, 0xfa, 0x06,
	0x46, 0xaf, 0xa1, 0x91, 0x11, 0x41, 0x52, 0xe9, 0xad, 0x69, 0xc2, 0xee, 0xa0, 0xda, 0xd2, 0x99,
	0x4e, 0xf9, 0x96, 0x82, 0x0e, 0xa0, 0x75, 0x49, 0x89, 0xca, 0x05, 0x95, 0x5e, 0xad, 0x57, 0xeb,
	0xbb, 0x87, 0x9d, 0x7b, 0xf4, 0x2f, 0x26, 0xe9, 0xdf, 0xb2, 0xf0, 0x2b, 0xd8, 0x39, 0xa6, 0xaa,
	0xc4, 0x6d, 0x4f, 0x08, 0xea, 0x8c, 0xa4, 0x54, 0xb7, 0xb4, 0xe1, 0xeb, 0x67, 0xcc, 0x61, 0x67,
	0x14, 0x04, 0x2b, 0xc4, 0x0e, 0xac, 0x17, 0x49, 0xe9, 0x39, 0xbd, 0x5a, 0x7f, 0xc3, 0x37, 0x01,
	0x7a, 0x01, 0x9b, 0x85, 0xa0, 0xc1, 0x9c, 0xe5, 0xe9, 0x82, 0x0a, 0xdd, 0x78, 0xdd, 0x77, 0x35,
	0x36, 0xd3, 0x50, 0x41, 0x21, 0xb9, 0xe2, 0x73, 0xca, 0xc8, 0x22, 0xa1, 0x81, 0x57, 0xeb, 0x39,
	0xfd, 0x96, 0xef, 0x16, 0xd8, 0xc4, 0x40, 0xb8, 0x03, 0xa8, 0x7a, 0xa0, 0xcc, 0x38, 0x93, 0x14,
	0x8f, 0x01, 0x55, 0xfb, 0x35, 0x28, 0x1a, 0x40, 0xd3, 0x4e, 0x64, 0x65, 0x7c, 0x7c, 0xec, 0x92,
	0x84, 0x7f, 0xaf, 0x41, 0xd3, 0x82, 0x8f, 0x0d, 0x8b, 0x3e, 0x40, 0x43, 0x2a, 0xa2, 0x72, 0x23,
	0xfa, 0xf6, 0x21, 0x7e, 0xac, 0x5c, 0xf9, 0x7f, 0xae, 0x99, 0xbe, 0x7d, 0x03, 0xf5, 0x01, 0xae,
	0x49, 0x12, 0x07, 0x44, 0x71, 0x51, 0xde, 0xc2, 0xdd, 0xad, 0x56, 0x72, 0x5a, 0xa7, 0x84, 0x2f,
	0xaf, 0xe6, 0x11, 0x8d, 0xc3, 0x48, 0x79, 0x75, 0xab, 0x53, 0x81, 0x9d, 0x68, 0x08, 0xed, 0x03,
	0x64, 0x54, 0x2c, 0x29, 0x53, 0x24, 0xa4, 0xde, 0xba, 0x26, 0x54, 0x90, 0x07, 0x52, 0x37, 0x9e,
	0x96, 0xba, 0xf9, 0x50, 0xea, 0x31, 0x6c, 0xdd, 0x9b, 0x05, 0xb9, 0xd0, 0x3c, 0x9b, 0xcc, 0xc6,
	0xd3, 0xd9, 0x71, 0xfb, 0xbf, 0x22, 0xf8, 0x3e, 0x9a, 0x5e, 0x14, 0x81, 0x53, 0x04, 0x93, 0xd9,
	0xe8, 0xe8, 0x74, 0x32, 0x6e, 0xaf, 0xa1, 0x4d, 0x68, 0x8d, 0xa7, 0xe7, 0x26, 0xaa, 0xe1, 0x10,
	0x1a, 0x66, 0x1d, 0xd1, 0x4b, 0xd8, 0xbe, 0xe6, 0x8a, 0xce, 0x55, 0x24, 0xa8, 0x8c, 0x78, 0x12,
	0x68, 0x71, 0xeb, 0xfe, 0x56, 0x81, 0x5e, 0x94, 0x20, 0x7a, 0x0f, 0x7b, 0x2c, 0x4f, 0xe7, 0x46,
	0x03, 0xad, 0xad, 0x48, 0x89, 0x8a, 0x39, 0x93, 0x76, 0x65, 0xba, 0x2c, 0x4f, 0x8f, 0x8a, 0xec,
	0xe7, 0x6a, 0x12, 0x7f, 0x84, 0xf6, 0x39, 0x55, 0x76, 0xf5, 0xed, 0x26, 0xde, 0xd9, 0xc4, 0x79,
	0xd2, 0x26, 0x18, 0x41, 0xfb, 0x78, 0xa5, 0x00, 0xfe, 0xa4, 0x8d, 0x50, 0x62, 0x76, 0xaf, 0xfe,
	0xaa, 0x6a, 0x17, 0x76, 0x4f, 0x63, 0x59, 0xee, 0xe6, 0x6d, 0xe1, 0x13, 0xe8, 0xdc, 0x87, 0x6d,
	0xed, 0xaa, 0x57, 0x9d, 0x67, 0x79, 0xf5, 0x0d, 0x74, 0xcc, 0x8d, 0x3d, 0xc7, 0x85, 0x78, 0x0f,
	0xba, 0x2b, 0x6c, 0x6b, 0x21, 0x01, 0x9d, 0x6f, 0x59, 0x40, 0x14, 0xfd, 0x87, 0x66, 0xde, 0x83,
	0xee, 0xca, 0x99, 0xa6, 0x99, 0x45, 0x43, 0x7f, 0x15, 0xdf, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xa5, 0x32, 0x63, 0xfc, 0xa2, 0x05, 0x00, 0x00,
}
