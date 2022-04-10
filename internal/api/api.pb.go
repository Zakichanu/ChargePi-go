// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: internal/models/api/api.proto

package api

import (
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

type GetConnectorStatusResponseConnectorType int32

const (
	GetConnectorStatusResponse_TYPE1   GetConnectorStatusResponseConnectorType = 0
	GetConnectorStatusResponse_TYPE2   GetConnectorStatusResponseConnectorType = 1
	GetConnectorStatusResponse_SCHUKO  GetConnectorStatusResponseConnectorType = 2
	GetConnectorStatusResponse_CHADEMO GetConnectorStatusResponseConnectorType = 3
)

// Enum value maps for GetConnectorStatusResponseConnectorType.
var (
	GetConnectorStatusResponseConnectorType_name = map[int32]string{
		0: "TYPE1",
		1: "TYPE2",
		2: "SCHUKO",
		3: "CHADEMO",
	}
	GetConnectorStatusResponseConnectorType_value = map[string]int32{
		"TYPE1":   0,
		"TYPE2":   1,
		"SCHUKO":  2,
		"CHADEMO": 3,
	}
)

func (x GetConnectorStatusResponseConnectorType) Enum() *GetConnectorStatusResponseConnectorType {
	p := new(GetConnectorStatusResponseConnectorType)
	*p = x
	return p
}

func (x GetConnectorStatusResponseConnectorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetConnectorStatusResponseConnectorType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_models_api_api_proto_enumTypes[0].Descriptor()
}

func (GetConnectorStatusResponseConnectorType) Type() protoreflect.EnumType {
	return &file_internal_models_api_api_proto_enumTypes[0]
}

func (x GetConnectorStatusResponseConnectorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetConnectorStatusResponseConnectorType.Descriptor instead.
func (GetConnectorStatusResponseConnectorType) EnumDescriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{1, 0}
}

type GetConnectorStatusResponseConnectorStatus int32

const (
	GetConnectorStatusResponse_Available     GetConnectorStatusResponseConnectorStatus = 0
	GetConnectorStatusResponse_Preparing     GetConnectorStatusResponseConnectorStatus = 1
	GetConnectorStatusResponse_Charging      GetConnectorStatusResponseConnectorStatus = 2
	GetConnectorStatusResponse_Finishing     GetConnectorStatusResponseConnectorStatus = 3
	GetConnectorStatusResponse_Unavailable   GetConnectorStatusResponseConnectorStatus = 4
	GetConnectorStatusResponse_SuspendedEVSE GetConnectorStatusResponseConnectorStatus = 5
	GetConnectorStatusResponse_SuspendedEV   GetConnectorStatusResponseConnectorStatus = 6
	GetConnectorStatusResponse_Reserved      GetConnectorStatusResponseConnectorStatus = 7
	GetConnectorStatusResponse_Faulted       GetConnectorStatusResponseConnectorStatus = 8
)

// Enum value maps for GetConnectorStatusResponseConnectorStatus.
var (
	GetConnectorStatusResponseConnectorStatus_name = map[int32]string{
		0: "Available",
		1: "Preparing",
		2: "Charging",
		3: "Finishing",
		4: "Unavailable",
		5: "SuspendedEVSE",
		6: "SuspendedEV",
		7: "Reserved",
		8: "Faulted",
	}
	GetConnectorStatusResponseConnectorStatus_value = map[string]int32{
		"Available":     0,
		"Preparing":     1,
		"Charging":      2,
		"Finishing":     3,
		"Unavailable":   4,
		"SuspendedEVSE": 5,
		"SuspendedEV":   6,
		"Reserved":      7,
		"Faulted":       8,
	}
)

func (x GetConnectorStatusResponseConnectorStatus) Enum() *GetConnectorStatusResponseConnectorStatus {
	p := new(GetConnectorStatusResponseConnectorStatus)
	*p = x
	return p
}

func (x GetConnectorStatusResponseConnectorStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetConnectorStatusResponseConnectorStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_models_api_api_proto_enumTypes[1].Descriptor()
}

func (GetConnectorStatusResponseConnectorStatus) Type() protoreflect.EnumType {
	return &file_internal_models_api_api_proto_enumTypes[1]
}

func (x GetConnectorStatusResponseConnectorStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetConnectorStatusResponseConnectorStatus.Descriptor instead.
func (GetConnectorStatusResponseConnectorStatus) EnumDescriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{1, 1}
}

type GetConnectorStatusResponseErrorCode int32

const (
	GetConnectorStatusResponse_NoError              GetConnectorStatusResponseErrorCode = 0
	GetConnectorStatusResponse_OtherError           GetConnectorStatusResponseErrorCode = 1
	GetConnectorStatusResponse_ConnectorLockFailure GetConnectorStatusResponseErrorCode = 3
	GetConnectorStatusResponse_EVCommunicationError GetConnectorStatusResponseErrorCode = 4
	GetConnectorStatusResponse_GroundFailure        GetConnectorStatusResponseErrorCode = 5
	GetConnectorStatusResponse_HighTemperature      GetConnectorStatusResponseErrorCode = 6
	GetConnectorStatusResponse_InternalError        GetConnectorStatusResponseErrorCode = 7
	GetConnectorStatusResponse_LocalListConflict    GetConnectorStatusResponseErrorCode = 8
	GetConnectorStatusResponse_OverCurrentFailure   GetConnectorStatusResponseErrorCode = 9
	GetConnectorStatusResponse_OverVoltage          GetConnectorStatusResponseErrorCode = 10
	GetConnectorStatusResponse_PowerMeterFailure    GetConnectorStatusResponseErrorCode = 11
	GetConnectorStatusResponse_PowerSwitchFailure   GetConnectorStatusResponseErrorCode = 12
	GetConnectorStatusResponse_ReaderFailure        GetConnectorStatusResponseErrorCode = 13
	GetConnectorStatusResponse_ResetFailure         GetConnectorStatusResponseErrorCode = 14
	GetConnectorStatusResponse_UnderVoltage         GetConnectorStatusResponseErrorCode = 15
	GetConnectorStatusResponse_WeakSignal           GetConnectorStatusResponseErrorCode = 16
)

// Enum value maps for GetConnectorStatusResponseErrorCode.
var (
	GetConnectorStatusResponseErrorCode_name = map[int32]string{
		0:  "NoError",
		1:  "OtherError",
		3:  "ConnectorLockFailure",
		4:  "EVCommunicationError",
		5:  "GroundFailure",
		6:  "HighTemperature",
		7:  "InternalError",
		8:  "LocalListConflict",
		9:  "OverCurrentFailure",
		10: "OverVoltage",
		11: "PowerMeterFailure",
		12: "PowerSwitchFailure",
		13: "ReaderFailure",
		14: "ResetFailure",
		15: "UnderVoltage",
		16: "WeakSignal",
	}
	GetConnectorStatusResponseErrorCode_value = map[string]int32{
		"NoError":              0,
		"OtherError":           1,
		"ConnectorLockFailure": 3,
		"EVCommunicationError": 4,
		"GroundFailure":        5,
		"HighTemperature":      6,
		"InternalError":        7,
		"LocalListConflict":    8,
		"OverCurrentFailure":   9,
		"OverVoltage":          10,
		"PowerMeterFailure":    11,
		"PowerSwitchFailure":   12,
		"ReaderFailure":        13,
		"ResetFailure":         14,
		"UnderVoltage":         15,
		"WeakSignal":           16,
	}
)

func (x GetConnectorStatusResponseErrorCode) Enum() *GetConnectorStatusResponseErrorCode {
	p := new(GetConnectorStatusResponseErrorCode)
	*p = x
	return p
}

func (x GetConnectorStatusResponseErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetConnectorStatusResponseErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_models_api_api_proto_enumTypes[2].Descriptor()
}

func (GetConnectorStatusResponseErrorCode) Type() protoreflect.EnumType {
	return &file_internal_models_api_api_proto_enumTypes[2]
}

func (x GetConnectorStatusResponseErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetConnectorStatusResponseErrorCode.Descriptor instead.
func (GetConnectorStatusResponseErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{1, 2}
}

type GetConnectorStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectorId int32 `protobuf:"varint,1,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
	EvseId      int32 `protobuf:"varint,2,opt,name=evseId,proto3" json:"evseId,omitempty"`
}

func (x *GetConnectorStatusRequest) Reset() {
	*x = GetConnectorStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectorStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectorStatusRequest) ProtoMessage() {}

func (x *GetConnectorStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectorStatusRequest.ProtoReflect.Descriptor instead.
func (*GetConnectorStatusRequest) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetConnectorStatusRequest) GetConnectorId() int32 {
	if x != nil {
		return x.ConnectorId
	}
	return 0
}

func (x *GetConnectorStatusRequest) GetEvseId() int32 {
	if x != nil {
		return x.EvseId
	}
	return 0
}

type GetConnectorStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId  string  `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	TimeElapsed    int32   `protobuf:"varint,2,opt,name=timeElapsed,proto3" json:"timeElapsed,omitempty"`
	EnergyConsumed float32 `protobuf:"fixed32,3,opt,name=energyConsumed,proto3" json:"energyConsumed,omitempty"`
	CurrentPower   float32 `protobuf:"fixed32,4,opt,name=currentPower,proto3" json:"currentPower,omitempty"`
}

func (x *GetConnectorStatusResponse) Reset() {
	*x = GetConnectorStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectorStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectorStatusResponse) ProtoMessage() {}

func (x *GetConnectorStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectorStatusResponse.ProtoReflect.Descriptor instead.
func (*GetConnectorStatusResponse) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetConnectorStatusResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *GetConnectorStatusResponse) GetTimeElapsed() int32 {
	if x != nil {
		return x.TimeElapsed
	}
	return 0
}

func (x *GetConnectorStatusResponse) GetEnergyConsumed() float32 {
	if x != nil {
		return x.EnergyConsumed
	}
	return 0
}

func (x *GetConnectorStatusResponse) GetCurrentPower() float32 {
	if x != nil {
		return x.CurrentPower
	}
	return 0
}

type StartTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId       string `protobuf:"bytes,1,opt,name=tagId,proto3" json:"tagId,omitempty"`
	ConnectorId int32  `protobuf:"varint,2,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
}

func (x *StartTransactionRequest) Reset() {
	*x = StartTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTransactionRequest) ProtoMessage() {}

func (x *StartTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTransactionRequest.ProtoReflect.Descriptor instead.
func (*StartTransactionRequest) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *StartTransactionRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *StartTransactionRequest) GetConnectorId() int32 {
	if x != nil {
		return x.ConnectorId
	}
	return 0
}

type StartTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	ConnectorId  int32  `protobuf:"varint,3,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
}

func (x *StartTransactionResponse) Reset() {
	*x = StartTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTransactionResponse) ProtoMessage() {}

func (x *StartTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTransactionResponse.ProtoReflect.Descriptor instead.
func (*StartTransactionResponse) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *StartTransactionResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StartTransactionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *StartTransactionResponse) GetConnectorId() int32 {
	if x != nil {
		return x.ConnectorId
	}
	return 0
}

type StopTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId       string `protobuf:"bytes,1,opt,name=tagId,proto3" json:"tagId,omitempty"`
	ConnectorId int32  `protobuf:"varint,2,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
}

func (x *StopTransactionRequest) Reset() {
	*x = StopTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTransactionRequest) ProtoMessage() {}

func (x *StopTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTransactionRequest.ProtoReflect.Descriptor instead.
func (*StopTransactionRequest) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *StopTransactionRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *StopTransactionRequest) GetConnectorId() int32 {
	if x != nil {
		return x.ConnectorId
	}
	return 0
}

type StopTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *StopTransactionResponse) Reset() {
	*x = StopTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTransactionResponse) ProtoMessage() {}

func (x *StopTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTransactionResponse.ProtoReflect.Descriptor instead.
func (*StopTransactionResponse) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *StopTransactionResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StopTransactionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type HandleChargingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId string `protobuf:"bytes,1,opt,name=tagId,proto3" json:"tagId,omitempty"`
}

func (x *HandleChargingRequest) Reset() {
	*x = HandleChargingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleChargingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleChargingRequest) ProtoMessage() {}

func (x *HandleChargingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleChargingRequest.ProtoReflect.Descriptor instead.
func (*HandleChargingRequest) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *HandleChargingRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type HandleChargingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	ConnectorId  int32  `protobuf:"varint,3,opt,name=connectorId,proto3" json:"connectorId,omitempty"`
}

func (x *HandleChargingResponse) Reset() {
	*x = HandleChargingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_models_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleChargingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleChargingResponse) ProtoMessage() {}

func (x *HandleChargingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_models_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleChargingResponse.ProtoReflect.Descriptor instead.
func (*HandleChargingResponse) Descriptor() ([]byte, []int) {
	return file_internal_models_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *HandleChargingResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HandleChargingResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *HandleChargingResponse) GetConnectorId() int32 {
	if x != nil {
		return x.ConnectorId
	}
	return 0
}

var File_internal_models_api_api_proto protoreflect.FileDescriptor

var file_internal_models_api_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x55, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x73, 0x65, 0x49, 0x64, 0x22, 0xdf, 0x05, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61, 0x70, 0x73,
	0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x3e,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x54, 0x59, 0x50, 0x45, 0x31, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x59,
	0x50, 0x45, 0x32, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x55, 0x4b, 0x4f, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x44, 0x45, 0x4d, 0x4f, 0x10, 0x03, 0x22, 0x9c,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x04, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x56, 0x53, 0x45, 0x10,
	0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x56,
	0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x10, 0x07,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x10, 0x08, 0x22, 0xcd, 0x02,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e,
	0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d,
	0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x05, 0x12,
	0x13, 0x0a, 0x0f, 0x48, 0x69, 0x67, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0x08, 0x12, 0x16,
	0x0a, 0x12, 0x4f, 0x76, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x76, 0x65, 0x72, 0x56, 0x6f,
	0x6c, 0x74, 0x61, 0x67, 0x65, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x0b, 0x12, 0x16,
	0x0a, 0x12, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x6e, 0x64, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x10, 0x0f, 0x12, 0x0e, 0x0a,
	0x0a, 0x57, 0x65, 0x61, 0x6b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x10, 0x10, 0x22, 0x51, 0x0a,
	0x17, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x78, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x53, 0x74,
	0x6f, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x17,
	0x53, 0x74, 0x6f, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67,
	0x49, 0x64, 0x22, 0x76, 0x0a, 0x16, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xda, 0x02, 0x0a, 0x0b, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x53, 0x74,
	0x6f, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x42, 0x6c, 0x61, 0x7a, 0x33, 0x6b, 0x78, 0x2f, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_models_api_api_proto_rawDescOnce sync.Once
	file_internal_models_api_api_proto_rawDescData = file_internal_models_api_api_proto_rawDesc
)

func file_internal_models_api_api_proto_rawDescGZIP() []byte {
	file_internal_models_api_api_proto_rawDescOnce.Do(func() {
		file_internal_models_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_models_api_api_proto_rawDescData)
	})
	return file_internal_models_api_api_proto_rawDescData
}

var file_internal_models_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_internal_models_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_models_api_api_proto_goTypes = []interface{}{
	(GetConnectorStatusResponseConnectorType)(0),   // 0: api.GetConnectorStatusResponse.connectorType
	(GetConnectorStatusResponseConnectorStatus)(0), // 1: api.GetConnectorStatusResponse.connectorStatus
	(GetConnectorStatusResponseErrorCode)(0),       // 2: api.GetConnectorStatusResponse.errorCode
	(*GetConnectorStatusRequest)(nil),              // 3: api.GetConnectorStatusRequest
	(*GetConnectorStatusResponse)(nil),             // 4: api.GetConnectorStatusResponse
	(*StartTransactionRequest)(nil),                // 5: api.StartTransactionRequest
	(*StartTransactionResponse)(nil),               // 6: api.StartTransactionResponse
	(*StopTransactionRequest)(nil),                 // 7: api.StopTransactionRequest
	(*StopTransactionResponse)(nil),                // 8: api.StopTransactionResponse
	(*HandleChargingRequest)(nil),                  // 9: api.HandleChargingRequest
	(*HandleChargingResponse)(nil),                 // 10: api.HandleChargingResponse
}
var file_internal_models_api_api_proto_depIdxs = []int32{
	3,  // 0: api.ChargePoint.GetConnectorStatus:input_type -> api.GetConnectorStatusRequest
	5,  // 1: api.ChargePoint.StartTransaction:input_type -> api.StartTransactionRequest
	7,  // 2: api.ChargePoint.StopTransaction:input_type -> api.StopTransactionRequest
	9,  // 3: api.ChargePoint.HandleCharging:input_type -> api.HandleChargingRequest
	4,  // 4: api.ChargePoint.GetConnectorStatus:output_type -> api.GetConnectorStatusResponse
	6,  // 5: api.ChargePoint.StartTransaction:output_type -> api.StartTransactionResponse
	8,  // 6: api.ChargePoint.StopTransaction:output_type -> api.StopTransactionResponse
	10, // 7: api.ChargePoint.HandleCharging:output_type -> api.HandleChargingResponse
	4,  // [4:8] is the sub-list for method output_type
	0,  // [0:4] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_internal_models_api_api_proto_init() }
func file_internal_models_api_api_proto_init() {
	if File_internal_models_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_models_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectorStatusRequest); i {
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
		file_internal_models_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectorStatusResponse); i {
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
		file_internal_models_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTransactionRequest); i {
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
		file_internal_models_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTransactionResponse); i {
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
		file_internal_models_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTransactionRequest); i {
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
		file_internal_models_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTransactionResponse); i {
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
		file_internal_models_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleChargingRequest); i {
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
		file_internal_models_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleChargingResponse); i {
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
			RawDescriptor: file_internal_models_api_api_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_models_api_api_proto_goTypes,
		DependencyIndexes: file_internal_models_api_api_proto_depIdxs,
		EnumInfos:         file_internal_models_api_api_proto_enumTypes,
		MessageInfos:      file_internal_models_api_api_proto_msgTypes,
	}.Build()
	File_internal_models_api_api_proto = out.File
	file_internal_models_api_api_proto_rawDesc = nil
	file_internal_models_api_api_proto_goTypes = nil
	file_internal_models_api_api_proto_depIdxs = nil
}