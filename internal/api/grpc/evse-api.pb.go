// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: evse-api.proto

package grpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GetEvsesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EVSEs []*EVSE `protobuf:"bytes,1,rep,name=EVSEs,proto3" json:"EVSEs,omitempty"`
}

func (x *GetEvsesResponse) Reset() {
	*x = GetEvsesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvsesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvsesResponse) ProtoMessage() {}

func (x *GetEvsesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvsesResponse.ProtoReflect.Descriptor instead.
func (*GetEvsesResponse) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetEvsesResponse) GetEVSEs() []*EVSE {
	if x != nil {
		return x.EVSEs
	}
	return nil
}

type GetEvseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvseId int32 `protobuf:"varint,1,opt,name=evseId,proto3" json:"evseId,omitempty"`
}

func (x *GetEvseRequest) Reset() {
	*x = GetEvseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvseRequest) ProtoMessage() {}

func (x *GetEvseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvseRequest.ProtoReflect.Descriptor instead.
func (*GetEvseRequest) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetEvseRequest) GetEvseId() int32 {
	if x != nil {
		return x.EvseId
	}
	return 0
}

type GetEvseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EVSE *EVSE `protobuf:"bytes,1,opt,name=EVSE,proto3" json:"EVSE,omitempty"`
}

func (x *GetEvseResponse) Reset() {
	*x = GetEvseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvseResponse) ProtoMessage() {}

func (x *GetEvseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvseResponse.ProtoReflect.Descriptor instead.
func (*GetEvseResponse) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetEvseResponse) GetEVSE() *EVSE {
	if x != nil {
		return x.EVSE
	}
	return nil
}

type AddEvseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvseId     int32       `protobuf:"varint,1,opt,name=evseId,proto3" json:"evseId,omitempty"`
	EVCC       *EVCC       `protobuf:"bytes,2,opt,name=EVCC,proto3" json:"EVCC,omitempty"`
	PowerMeter *PowerMeter `protobuf:"bytes,3,opt,name=PowerMeter,proto3" json:"PowerMeter,omitempty"`
}

func (x *AddEvseRequest) Reset() {
	*x = AddEvseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEvseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEvseRequest) ProtoMessage() {}

func (x *AddEvseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEvseRequest.ProtoReflect.Descriptor instead.
func (*AddEvseRequest) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{3}
}

func (x *AddEvseRequest) GetEvseId() int32 {
	if x != nil {
		return x.EvseId
	}
	return 0
}

func (x *AddEvseRequest) GetEVCC() *EVCC {
	if x != nil {
		return x.EVCC
	}
	return nil
}

func (x *AddEvseRequest) GetPowerMeter() *PowerMeter {
	if x != nil {
		return x.PowerMeter
	}
	return nil
}

type AddEvseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	EVSE   *EVSE  `protobuf:"bytes,2,opt,name=EVSE,proto3" json:"EVSE,omitempty"`
}

func (x *AddEvseResponse) Reset() {
	*x = AddEvseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEvseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEvseResponse) ProtoMessage() {}

func (x *AddEvseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEvseResponse.ProtoReflect.Descriptor instead.
func (*AddEvseResponse) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{4}
}

func (x *AddEvseResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AddEvseResponse) GetEVSE() *EVSE {
	if x != nil {
		return x.EVSE
	}
	return nil
}

type SetEVCCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvseId int32 `protobuf:"varint,1,opt,name=evseId,proto3" json:"evseId,omitempty"`
	EVCC   *EVCC `protobuf:"bytes,2,opt,name=EVCC,proto3" json:"EVCC,omitempty"`
}

func (x *SetEVCCRequest) Reset() {
	*x = SetEVCCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEVCCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEVCCRequest) ProtoMessage() {}

func (x *SetEVCCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEVCCRequest.ProtoReflect.Descriptor instead.
func (*SetEVCCRequest) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{5}
}

func (x *SetEVCCRequest) GetEvseId() int32 {
	if x != nil {
		return x.EvseId
	}
	return 0
}

func (x *SetEVCCRequest) GetEVCC() *EVCC {
	if x != nil {
		return x.EVCC
	}
	return nil
}

type SetEvccResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SetEvccResponse) Reset() {
	*x = SetEvccResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEvccResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEvccResponse) ProtoMessage() {}

func (x *SetEvccResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEvccResponse.ProtoReflect.Descriptor instead.
func (*SetEvccResponse) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{6}
}

func (x *SetEvccResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type SetPowerMeterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvseId     int32       `protobuf:"varint,1,opt,name=evseId,proto3" json:"evseId,omitempty"`
	PowerMeter *PowerMeter `protobuf:"bytes,2,opt,name=PowerMeter,proto3" json:"PowerMeter,omitempty"`
}

func (x *SetPowerMeterRequest) Reset() {
	*x = SetPowerMeterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPowerMeterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerMeterRequest) ProtoMessage() {}

func (x *SetPowerMeterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerMeterRequest.ProtoReflect.Descriptor instead.
func (*SetPowerMeterRequest) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{7}
}

func (x *SetPowerMeterRequest) GetEvseId() int32 {
	if x != nil {
		return x.EvseId
	}
	return 0
}

func (x *SetPowerMeterRequest) GetPowerMeter() *PowerMeter {
	if x != nil {
		return x.PowerMeter
	}
	return nil
}

type SetPowerMeterDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SetPowerMeterDetails) Reset() {
	*x = SetPowerMeterDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPowerMeterDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerMeterDetails) ProtoMessage() {}

func (x *SetPowerMeterDetails) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerMeterDetails.ProtoReflect.Descriptor instead.
func (*SetPowerMeterDetails) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{8}
}

func (x *SetPowerMeterDetails) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetUsageForEVSERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvseId int32 `protobuf:"varint,1,opt,name=evseId,proto3" json:"evseId,omitempty"`
}

func (x *GetUsageForEVSERequest) Reset() {
	*x = GetUsageForEVSERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsageForEVSERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsageForEVSERequest) ProtoMessage() {}

func (x *GetUsageForEVSERequest) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsageForEVSERequest.ProtoReflect.Descriptor instead.
func (*GetUsageForEVSERequest) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetUsageForEVSERequest) GetEvseId() int32 {
	if x != nil {
		return x.EvseId
	}
	return 0
}

type GetUsageForEVSEResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*Sample `protobuf:"bytes,1,rep,name=Samples,proto3" json:"Samples,omitempty"`
}

func (x *GetUsageForEVSEResponse) Reset() {
	*x = GetUsageForEVSEResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evse_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsageForEVSEResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsageForEVSEResponse) ProtoMessage() {}

func (x *GetUsageForEVSEResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evse_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsageForEVSEResponse.ProtoReflect.Descriptor instead.
func (*GetUsageForEVSEResponse) Descriptor() ([]byte, []int) {
	return file_evse_api_proto_rawDescGZIP(), []int{10}
}

func (x *GetUsageForEVSEResponse) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

var File_evse_api_proto protoreflect.FileDescriptor

var file_evse_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x76, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x45, 0x56, 0x53, 0x45, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x56, 0x53,
	0x45, 0x52, 0x05, 0x45, 0x56, 0x53, 0x45, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76,
	0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x73, 0x65,
	0x49, 0x64, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x45, 0x56, 0x53, 0x45, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x56, 0x53, 0x45, 0x52, 0x04,
	0x45, 0x56, 0x53, 0x45, 0x22, 0x78, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x04, 0x45, 0x56, 0x43, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x56, 0x43, 0x43, 0x52, 0x04, 0x45, 0x56, 0x43, 0x43, 0x12, 0x2f, 0x0a,
	0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x52, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x48,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x76, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x45, 0x56, 0x53,
	0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x56,
	0x53, 0x45, 0x52, 0x04, 0x45, 0x56, 0x53, 0x45, 0x22, 0x47, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x45,
	0x56, 0x43, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76,
	0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x45, 0x56, 0x43, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x56, 0x43, 0x43, 0x52, 0x04, 0x45, 0x56, 0x43,
	0x43, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x45, 0x76, 0x63, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x14, 0x53, 0x65, 0x74,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x45, 0x56, 0x53, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x45, 0x56, 0x53, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x07, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x32, 0x86, 0x03, 0x0a, 0x04, 0x45, 0x76,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x45, 0x56, 0x53, 0x45, 0x12, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x56, 0x43, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x45, 0x56, 0x53, 0x45, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45, 0x56,
	0x53, 0x45, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x45, 0x56, 0x43, 0x43, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x74, 0x45, 0x56, 0x43, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x76, 0x63, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x45,
	0x56, 0x53, 0x45, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x46, 0x6f, 0x72, 0x45, 0x56, 0x53, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x46,
	0x6f, 0x72, 0x45, 0x56, 0x53, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x42, 0x6c, 0x61, 0x7a, 0x33, 0x6b, 0x78, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x50, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evse_api_proto_rawDescOnce sync.Once
	file_evse_api_proto_rawDescData = file_evse_api_proto_rawDesc
)

func file_evse_api_proto_rawDescGZIP() []byte {
	file_evse_api_proto_rawDescOnce.Do(func() {
		file_evse_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_evse_api_proto_rawDescData)
	})
	return file_evse_api_proto_rawDescData
}

var file_evse_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_evse_api_proto_goTypes = []interface{}{
	(*GetEvsesResponse)(nil),        // 0: api.GetEvsesResponse
	(*GetEvseRequest)(nil),          // 1: api.GetEvseRequest
	(*GetEvseResponse)(nil),         // 2: api.GetEvseResponse
	(*AddEvseRequest)(nil),          // 3: api.AddEvseRequest
	(*AddEvseResponse)(nil),         // 4: api.AddEvseResponse
	(*SetEVCCRequest)(nil),          // 5: api.SetEVCCRequest
	(*SetEvccResponse)(nil),         // 6: api.SetEvccResponse
	(*SetPowerMeterRequest)(nil),    // 7: api.SetPowerMeterRequest
	(*SetPowerMeterDetails)(nil),    // 8: api.SetPowerMeterDetails
	(*GetUsageForEVSERequest)(nil),  // 9: api.GetUsageForEVSERequest
	(*GetUsageForEVSEResponse)(nil), // 10: api.GetUsageForEVSEResponse
	(*EVSE)(nil),                    // 11: api.EVSE
	(*EVCC)(nil),                    // 12: api.EVCC
	(*PowerMeter)(nil),              // 13: api.PowerMeter
	(*Sample)(nil),                  // 14: api.Sample
	(*empty.Empty)(nil),             // 15: google.protobuf.Empty
}
var file_evse_api_proto_depIdxs = []int32{
	11, // 0: api.GetEvsesResponse.EVSEs:type_name -> api.EVSE
	11, // 1: api.GetEvseResponse.EVSE:type_name -> api.EVSE
	12, // 2: api.AddEvseRequest.EVCC:type_name -> api.EVCC
	13, // 3: api.AddEvseRequest.PowerMeter:type_name -> api.PowerMeter
	11, // 4: api.AddEvseResponse.EVSE:type_name -> api.EVSE
	12, // 5: api.SetEVCCRequest.EVCC:type_name -> api.EVCC
	13, // 6: api.SetPowerMeterRequest.PowerMeter:type_name -> api.PowerMeter
	14, // 7: api.GetUsageForEVSEResponse.Samples:type_name -> api.Sample
	5,  // 8: api.Evse.AddEVSE:input_type -> api.SetEVCCRequest
	15, // 9: api.Evse.GetEVSEs:input_type -> google.protobuf.Empty
	1,  // 10: api.Evse.GetEVSE:input_type -> api.GetEvseRequest
	5,  // 11: api.Evse.SetEVCC:input_type -> api.SetEVCCRequest
	7,  // 12: api.Evse.SetPowerMeter:input_type -> api.SetPowerMeterRequest
	9,  // 13: api.Evse.GetUsageForEVSE:input_type -> api.GetUsageForEVSERequest
	4,  // 14: api.Evse.AddEVSE:output_type -> api.AddEvseResponse
	0,  // 15: api.Evse.GetEVSEs:output_type -> api.GetEvsesResponse
	2,  // 16: api.Evse.GetEVSE:output_type -> api.GetEvseResponse
	6,  // 17: api.Evse.SetEVCC:output_type -> api.SetEvccResponse
	8,  // 18: api.Evse.SetPowerMeter:output_type -> api.SetPowerMeterDetails
	10, // 19: api.Evse.GetUsageForEVSE:output_type -> api.GetUsageForEVSEResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_evse_api_proto_init() }
func file_evse_api_proto_init() {
	if File_evse_api_proto != nil {
		return
	}
	file_connector_proto_init()
	file_hardware_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_evse_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvsesResponse); i {
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
		file_evse_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvseRequest); i {
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
		file_evse_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvseResponse); i {
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
		file_evse_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEvseRequest); i {
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
		file_evse_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEvseResponse); i {
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
		file_evse_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEVCCRequest); i {
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
		file_evse_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEvccResponse); i {
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
		file_evse_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPowerMeterRequest); i {
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
		file_evse_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPowerMeterDetails); i {
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
		file_evse_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsageForEVSERequest); i {
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
		file_evse_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsageForEVSEResponse); i {
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
			RawDescriptor: file_evse_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evse_api_proto_goTypes,
		DependencyIndexes: file_evse_api_proto_depIdxs,
		MessageInfos:      file_evse_api_proto_msgTypes,
	}.Build()
	File_evse_api_proto = out.File
	file_evse_api_proto_rawDesc = nil
	file_evse_api_proto_goTypes = nil
	file_evse_api_proto_depIdxs = nil
}
