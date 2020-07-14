// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: valprotocol.proto

package valprotocol

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PreconditionBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeforeHash  *common.HashBuf `protobuf:"bytes,1,opt,name=beforeHash,proto3" json:"beforeHash,omitempty"`
	BeforeInbox *common.HashBuf `protobuf:"bytes,2,opt,name=beforeInbox,proto3" json:"beforeInbox,omitempty"`
}

func (x *PreconditionBuf) Reset() {
	*x = PreconditionBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionBuf) ProtoMessage() {}

func (x *PreconditionBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionBuf.ProtoReflect.Descriptor instead.
func (*PreconditionBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{0}
}

func (x *PreconditionBuf) GetBeforeHash() *common.HashBuf {
	if x != nil {
		return x.BeforeHash
	}
	return nil
}

func (x *PreconditionBuf) GetBeforeInbox() *common.HashBuf {
	if x != nil {
		return x.BeforeInbox
	}
	return nil
}

type ExecutionAssertionStubBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AfterHash        *common.HashBuf `protobuf:"bytes,1,opt,name=afterHash,proto3" json:"afterHash,omitempty"`
	DidInboxInsn     bool            `protobuf:"varint,2,opt,name=didInboxInsn,proto3" json:"didInboxInsn,omitempty"`
	NumGas           uint64          `protobuf:"varint,3,opt,name=numGas,proto3" json:"numGas,omitempty"`
	FirstMessageHash *common.HashBuf `protobuf:"bytes,4,opt,name=firstMessageHash,proto3" json:"firstMessageHash,omitempty"`
	LastMessageHash  *common.HashBuf `protobuf:"bytes,5,opt,name=lastMessageHash,proto3" json:"lastMessageHash,omitempty"`
	FirstLogHash     *common.HashBuf `protobuf:"bytes,6,opt,name=firstLogHash,proto3" json:"firstLogHash,omitempty"`
	LastLogHash      *common.HashBuf `protobuf:"bytes,7,opt,name=lastLogHash,proto3" json:"lastLogHash,omitempty"`
}

func (x *ExecutionAssertionStubBuf) Reset() {
	*x = ExecutionAssertionStubBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionAssertionStubBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionAssertionStubBuf) ProtoMessage() {}

func (x *ExecutionAssertionStubBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionAssertionStubBuf.ProtoReflect.Descriptor instead.
func (*ExecutionAssertionStubBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionAssertionStubBuf) GetAfterHash() *common.HashBuf {
	if x != nil {
		return x.AfterHash
	}
	return nil
}

func (x *ExecutionAssertionStubBuf) GetDidInboxInsn() bool {
	if x != nil {
		return x.DidInboxInsn
	}
	return false
}

func (x *ExecutionAssertionStubBuf) GetNumGas() uint64 {
	if x != nil {
		return x.NumGas
	}
	return 0
}

func (x *ExecutionAssertionStubBuf) GetFirstMessageHash() *common.HashBuf {
	if x != nil {
		return x.FirstMessageHash
	}
	return nil
}

func (x *ExecutionAssertionStubBuf) GetLastMessageHash() *common.HashBuf {
	if x != nil {
		return x.LastMessageHash
	}
	return nil
}

func (x *ExecutionAssertionStubBuf) GetFirstLogHash() *common.HashBuf {
	if x != nil {
		return x.FirstLogHash
	}
	return nil
}

func (x *ExecutionAssertionStubBuf) GetLastLogHash() *common.HashBuf {
	if x != nil {
		return x.LastLogHash
	}
	return nil
}

type ChainParamsBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StakeRequirement        *common.BigIntegerBuf `protobuf:"bytes,1,opt,name=stakeRequirement,proto3" json:"stakeRequirement,omitempty"`
	GracePeriod             *common.TimeTicksBuf  `protobuf:"bytes,2,opt,name=gracePeriod,proto3" json:"gracePeriod,omitempty"`
	MaxExecutionSteps       uint64                `protobuf:"varint,3,opt,name=maxExecutionSteps,proto3" json:"maxExecutionSteps,omitempty"`
	ArbGasSpeedLimitPerTick uint64                `protobuf:"varint,4,opt,name=ArbGasSpeedLimitPerTick,proto3" json:"ArbGasSpeedLimitPerTick,omitempty"`
}

func (x *ChainParamsBuf) Reset() {
	*x = ChainParamsBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainParamsBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainParamsBuf) ProtoMessage() {}

func (x *ChainParamsBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainParamsBuf.ProtoReflect.Descriptor instead.
func (*ChainParamsBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{2}
}

func (x *ChainParamsBuf) GetStakeRequirement() *common.BigIntegerBuf {
	if x != nil {
		return x.StakeRequirement
	}
	return nil
}

func (x *ChainParamsBuf) GetGracePeriod() *common.TimeTicksBuf {
	if x != nil {
		return x.GracePeriod
	}
	return nil
}

func (x *ChainParamsBuf) GetMaxExecutionSteps() uint64 {
	if x != nil {
		return x.MaxExecutionSteps
	}
	return 0
}

func (x *ChainParamsBuf) GetArbGasSpeedLimitPerTick() uint64 {
	if x != nil {
		return x.ArbGasSpeedLimitPerTick
	}
	return 0
}

type VMProtoDataBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineHash *common.HashBuf       `protobuf:"bytes,1,opt,name=machineHash,proto3" json:"machineHash,omitempty"`
	InboxTop    *common.HashBuf       `protobuf:"bytes,2,opt,name=inboxTop,proto3" json:"inboxTop,omitempty"`
	InboxCount  *common.BigIntegerBuf `protobuf:"bytes,3,opt,name=inboxCount,proto3" json:"inboxCount,omitempty"`
}

func (x *VMProtoDataBuf) Reset() {
	*x = VMProtoDataBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMProtoDataBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMProtoDataBuf) ProtoMessage() {}

func (x *VMProtoDataBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMProtoDataBuf.ProtoReflect.Descriptor instead.
func (*VMProtoDataBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{3}
}

func (x *VMProtoDataBuf) GetMachineHash() *common.HashBuf {
	if x != nil {
		return x.MachineHash
	}
	return nil
}

func (x *VMProtoDataBuf) GetInboxTop() *common.HashBuf {
	if x != nil {
		return x.InboxTop
	}
	return nil
}

func (x *VMProtoDataBuf) GetInboxCount() *common.BigIntegerBuf {
	if x != nil {
		return x.InboxCount
	}
	return nil
}

type AssertionParamsBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumSteps             uint64                `protobuf:"varint,1,opt,name=numSteps,proto3" json:"numSteps,omitempty"`
	ImportedMessageCount *common.BigIntegerBuf `protobuf:"bytes,2,opt,name=importedMessageCount,proto3" json:"importedMessageCount,omitempty"`
}

func (x *AssertionParamsBuf) Reset() {
	*x = AssertionParamsBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssertionParamsBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssertionParamsBuf) ProtoMessage() {}

func (x *AssertionParamsBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssertionParamsBuf.ProtoReflect.Descriptor instead.
func (*AssertionParamsBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{4}
}

func (x *AssertionParamsBuf) GetNumSteps() uint64 {
	if x != nil {
		return x.NumSteps
	}
	return 0
}

func (x *AssertionParamsBuf) GetImportedMessageCount() *common.BigIntegerBuf {
	if x != nil {
		return x.ImportedMessageCount
	}
	return nil
}

type AssertionClaimBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AfterInboxTop         *common.HashBuf            `protobuf:"bytes,1,opt,name=afterInboxTop,proto3" json:"afterInboxTop,omitempty"`
	ImportedMessagesSlice *common.HashBuf            `protobuf:"bytes,2,opt,name=importedMessagesSlice,proto3" json:"importedMessagesSlice,omitempty"`
	AssertionStub         *ExecutionAssertionStubBuf `protobuf:"bytes,3,opt,name=assertionStub,proto3" json:"assertionStub,omitempty"`
}

func (x *AssertionClaimBuf) Reset() {
	*x = AssertionClaimBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssertionClaimBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssertionClaimBuf) ProtoMessage() {}

func (x *AssertionClaimBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssertionClaimBuf.ProtoReflect.Descriptor instead.
func (*AssertionClaimBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{5}
}

func (x *AssertionClaimBuf) GetAfterInboxTop() *common.HashBuf {
	if x != nil {
		return x.AfterInboxTop
	}
	return nil
}

func (x *AssertionClaimBuf) GetImportedMessagesSlice() *common.HashBuf {
	if x != nil {
		return x.ImportedMessagesSlice
	}
	return nil
}

func (x *AssertionClaimBuf) GetAssertionStub() *ExecutionAssertionStubBuf {
	if x != nil {
		return x.AssertionStub
	}
	return nil
}

type DisputableNodeBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssertionParams *AssertionParamsBuf   `protobuf:"bytes,1,opt,name=assertionParams,proto3" json:"assertionParams,omitempty"`
	AssertionClaim  *AssertionClaimBuf    `protobuf:"bytes,2,opt,name=assertionClaim,proto3" json:"assertionClaim,omitempty"`
	MaxInboxTop     *common.HashBuf       `protobuf:"bytes,3,opt,name=maxInboxTop,proto3" json:"maxInboxTop,omitempty"`
	MaxInboxCount   *common.BigIntegerBuf `protobuf:"bytes,4,opt,name=maxInboxCount,proto3" json:"maxInboxCount,omitempty"`
}

func (x *DisputableNodeBuf) Reset() {
	*x = DisputableNodeBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valprotocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisputableNodeBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisputableNodeBuf) ProtoMessage() {}

func (x *DisputableNodeBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valprotocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisputableNodeBuf.ProtoReflect.Descriptor instead.
func (*DisputableNodeBuf) Descriptor() ([]byte, []int) {
	return file_valprotocol_proto_rawDescGZIP(), []int{6}
}

func (x *DisputableNodeBuf) GetAssertionParams() *AssertionParamsBuf {
	if x != nil {
		return x.AssertionParams
	}
	return nil
}

func (x *DisputableNodeBuf) GetAssertionClaim() *AssertionClaimBuf {
	if x != nil {
		return x.AssertionClaim
	}
	return nil
}

func (x *DisputableNodeBuf) GetMaxInboxTop() *common.HashBuf {
	if x != nil {
		return x.MaxInboxTop
	}
	return nil
}

func (x *DisputableNodeBuf) GetMaxInboxCount() *common.BigIntegerBuf {
	if x != nil {
		return x.MaxInboxCount
	}
	return nil
}

var File_valprotocol_proto protoreflect.FileDescriptor

var file_valprotocol_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x1a, 0x1c, 0x61, 0x72, 0x62, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75,
	0x0a, 0x0f, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x75,
	0x66, 0x12, 0x2f, 0x0a, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x31, 0x0a, 0x0b, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x62, 0x6f,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0b, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x22, 0xe6, 0x02, 0x0a, 0x19, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x75, 0x62,
	0x42, 0x75, 0x66, 0x12, 0x2d, 0x0a, 0x09, 0x61, 0x66, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x09, 0x61, 0x66, 0x74, 0x65, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x49, 0x6e,
	0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x64, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x49, 0x6e, 0x73, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x47, 0x61, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x47, 0x61, 0x73, 0x12, 0x3b,
	0x0a, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x39, 0x0a, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x33, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0c, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75,
	0x66, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x73, 0x68, 0x22, 0xf3,
	0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x75,
	0x66, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x42,
	0x75, 0x66, 0x52, 0x10, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x42, 0x75, 0x66, 0x52,
	0x0b, 0x67, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x11,
	0x6d, 0x61, 0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x41, 0x72,
	0x62, 0x47, 0x61, 0x73, 0x53, 0x70, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x41, 0x72, 0x62,
	0x47, 0x61, 0x73, 0x53, 0x70, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x56, 0x4d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x75, 0x66, 0x12, 0x31, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0b, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2b, 0x0a, 0x08, 0x69, 0x6e,
	0x62, 0x6f, 0x78, 0x54, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x08, 0x69,
	0x6e, 0x62, 0x6f, 0x78, 0x54, 0x6f, 0x70, 0x12, 0x35, 0x0a, 0x0a, 0x69, 0x6e, 0x62, 0x6f, 0x78,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x42,
	0x75, 0x66, 0x52, 0x0a, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7b,
	0x0a, 0x12, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x42, 0x75, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x65, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x49, 0x0a, 0x14, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x42, 0x75, 0x66, 0x52, 0x14, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x11,
	0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x42, 0x75,
	0x66, 0x12, 0x35, 0x0a, 0x0d, 0x61, 0x66, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x54,
	0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0d, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x54, 0x6f, 0x70, 0x12, 0x45, 0x0a, 0x15, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x15, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x75, 0x62,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73,
	0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x75, 0x62, 0x42, 0x75, 0x66, 0x52, 0x0d,
	0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x75, 0x62, 0x22, 0x96, 0x02,
	0x0a, 0x11, 0x44, 0x69, 0x73, 0x70, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x42, 0x75, 0x66, 0x12, 0x49, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76,
	0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x75, 0x66, 0x52, 0x0f, 0x61,
	0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x46,
	0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x42, 0x75, 0x66, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x31, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x54, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0b, 0x6d, 0x61,
	0x78, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x54, 0x6f, 0x70, 0x12, 0x3b, 0x0a, 0x0d, 0x6d, 0x61, 0x78,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x42, 0x75, 0x66, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x62, 0x6f,
	0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_valprotocol_proto_rawDescOnce sync.Once
	file_valprotocol_proto_rawDescData = file_valprotocol_proto_rawDesc
)

func file_valprotocol_proto_rawDescGZIP() []byte {
	file_valprotocol_proto_rawDescOnce.Do(func() {
		file_valprotocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_valprotocol_proto_rawDescData)
	})
	return file_valprotocol_proto_rawDescData
}

var file_valprotocol_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_valprotocol_proto_goTypes = []interface{}{
	(*PreconditionBuf)(nil),           // 0: valprotocol.PreconditionBuf
	(*ExecutionAssertionStubBuf)(nil), // 1: valprotocol.ExecutionAssertionStubBuf
	(*ChainParamsBuf)(nil),            // 2: valprotocol.ChainParamsBuf
	(*VMProtoDataBuf)(nil),            // 3: valprotocol.VMProtoDataBuf
	(*AssertionParamsBuf)(nil),        // 4: valprotocol.AssertionParamsBuf
	(*AssertionClaimBuf)(nil),         // 5: valprotocol.AssertionClaimBuf
	(*DisputableNodeBuf)(nil),         // 6: valprotocol.DisputableNodeBuf
	(*common.HashBuf)(nil),            // 7: common.HashBuf
	(*common.BigIntegerBuf)(nil),      // 8: common.BigIntegerBuf
	(*common.TimeTicksBuf)(nil),       // 9: common.TimeTicksBuf
}
var file_valprotocol_proto_depIdxs = []int32{
	7,  // 0: valprotocol.PreconditionBuf.beforeHash:type_name -> common.HashBuf
	7,  // 1: valprotocol.PreconditionBuf.beforeInbox:type_name -> common.HashBuf
	7,  // 2: valprotocol.ExecutionAssertionStubBuf.afterHash:type_name -> common.HashBuf
	7,  // 3: valprotocol.ExecutionAssertionStubBuf.firstMessageHash:type_name -> common.HashBuf
	7,  // 4: valprotocol.ExecutionAssertionStubBuf.lastMessageHash:type_name -> common.HashBuf
	7,  // 5: valprotocol.ExecutionAssertionStubBuf.firstLogHash:type_name -> common.HashBuf
	7,  // 6: valprotocol.ExecutionAssertionStubBuf.lastLogHash:type_name -> common.HashBuf
	8,  // 7: valprotocol.ChainParamsBuf.stakeRequirement:type_name -> common.BigIntegerBuf
	9,  // 8: valprotocol.ChainParamsBuf.gracePeriod:type_name -> common.TimeTicksBuf
	7,  // 9: valprotocol.VMProtoDataBuf.machineHash:type_name -> common.HashBuf
	7,  // 10: valprotocol.VMProtoDataBuf.inboxTop:type_name -> common.HashBuf
	8,  // 11: valprotocol.VMProtoDataBuf.inboxCount:type_name -> common.BigIntegerBuf
	8,  // 12: valprotocol.AssertionParamsBuf.importedMessageCount:type_name -> common.BigIntegerBuf
	7,  // 13: valprotocol.AssertionClaimBuf.afterInboxTop:type_name -> common.HashBuf
	7,  // 14: valprotocol.AssertionClaimBuf.importedMessagesSlice:type_name -> common.HashBuf
	1,  // 15: valprotocol.AssertionClaimBuf.assertionStub:type_name -> valprotocol.ExecutionAssertionStubBuf
	4,  // 16: valprotocol.DisputableNodeBuf.assertionParams:type_name -> valprotocol.AssertionParamsBuf
	5,  // 17: valprotocol.DisputableNodeBuf.assertionClaim:type_name -> valprotocol.AssertionClaimBuf
	7,  // 18: valprotocol.DisputableNodeBuf.maxInboxTop:type_name -> common.HashBuf
	8,  // 19: valprotocol.DisputableNodeBuf.maxInboxCount:type_name -> common.BigIntegerBuf
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_valprotocol_proto_init() }
func file_valprotocol_proto_init() {
	if File_valprotocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_valprotocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_valprotocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionAssertionStubBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_valprotocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainParamsBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_valprotocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMProtoDataBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_valprotocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssertionParamsBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_valprotocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssertionClaimBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_valprotocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisputableNodeBuf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_valprotocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_valprotocol_proto_goTypes,
		DependencyIndexes: file_valprotocol_proto_depIdxs,
		MessageInfos:      file_valprotocol_proto_msgTypes,
	}.Build()
	File_valprotocol_proto = out.File
	file_valprotocol_proto_rawDesc = nil
	file_valprotocol_proto_goTypes = nil
	file_valprotocol_proto_depIdxs = nil
}
