// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: charge-point.proto

package grpc

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ConnectionSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChargePointId         *string         `protobuf:"bytes,1,opt,name=chargePointId,proto3,oneof" json:"chargePointId,omitempty"`
	Url                   *string         `protobuf:"bytes,2,opt,name=url,proto3,oneof" json:"url,omitempty"`
	ProtocolVersion       ProtocolVersion `protobuf:"varint,3,opt,name=ProtocolVersion,proto3,enum=api.ProtocolVersion" json:"ProtocolVersion,omitempty"`
	BasicAuthUsername     *string         `protobuf:"bytes,4,opt,name=BasicAuthUsername,proto3,oneof" json:"BasicAuthUsername,omitempty"`
	BasicAuthPassword     *string         `protobuf:"bytes,5,opt,name=BasicAuthPassword,proto3,oneof" json:"BasicAuthPassword,omitempty"`
	TLS_CACertificate     *string         `protobuf:"bytes,6,opt,name=TLS_CACertificate,json=TLSCACertificate,proto3,oneof" json:"TLS_CACertificate,omitempty"`
	TLS_ClientCertificate *string         `protobuf:"bytes,7,opt,name=TLS_ClientCertificate,json=TLSClientCertificate,proto3,oneof" json:"TLS_ClientCertificate,omitempty"`
	TLS_ClientPrivateKey  *string         `protobuf:"bytes,8,opt,name=TLS_ClientPrivateKey,json=TLSClientPrivateKey,proto3,oneof" json:"TLS_ClientPrivateKey,omitempty"`
}

func (x *ConnectionSettings) Reset() {
	*x = ConnectionSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charge_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionSettings) ProtoMessage() {}

func (x *ConnectionSettings) ProtoReflect() protoreflect.Message {
	mi := &file_charge_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionSettings.ProtoReflect.Descriptor instead.
func (*ConnectionSettings) Descriptor() ([]byte, []int) {
	return file_charge_point_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionSettings) GetChargePointId() string {
	if x != nil && x.ChargePointId != nil {
		return *x.ChargePointId
	}
	return ""
}

func (x *ConnectionSettings) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *ConnectionSettings) GetProtocolVersion() ProtocolVersion {
	if x != nil {
		return x.ProtocolVersion
	}
	return ProtocolVersion_OCPP16
}

func (x *ConnectionSettings) GetBasicAuthUsername() string {
	if x != nil && x.BasicAuthUsername != nil {
		return *x.BasicAuthUsername
	}
	return ""
}

func (x *ConnectionSettings) GetBasicAuthPassword() string {
	if x != nil && x.BasicAuthPassword != nil {
		return *x.BasicAuthPassword
	}
	return ""
}

func (x *ConnectionSettings) GetTLS_CACertificate() string {
	if x != nil && x.TLS_CACertificate != nil {
		return *x.TLS_CACertificate
	}
	return ""
}

func (x *ConnectionSettings) GetTLS_ClientCertificate() string {
	if x != nil && x.TLS_ClientCertificate != nil {
		return *x.TLS_ClientCertificate
	}
	return ""
}

func (x *ConnectionSettings) GetTLS_ClientPrivateKey() string {
	if x != nil && x.TLS_ClientPrivateKey != nil {
		return *x.TLS_ClientPrivateKey
	}
	return ""
}

type ChargePointInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         *string       `protobuf:"bytes,1,opt,name=Type,proto3,oneof" json:"Type,omitempty"`
	MaxPower     *float32      `protobuf:"fixed32,2,opt,name=MaxPower,proto3,oneof" json:"MaxPower,omitempty"`
	OcppInfo     *OcppInfo     `protobuf:"bytes,4,opt,name=OcppInfo,proto3" json:"OcppInfo,omitempty"`
	FreeCharging *FreeCharging `protobuf:"bytes,5,opt,name=FreeCharging,proto3" json:"FreeCharging,omitempty"`
}

func (x *ChargePointInfo) Reset() {
	*x = ChargePointInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charge_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargePointInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargePointInfo) ProtoMessage() {}

func (x *ChargePointInfo) ProtoReflect() protoreflect.Message {
	mi := &file_charge_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargePointInfo.ProtoReflect.Descriptor instead.
func (*ChargePointInfo) Descriptor() ([]byte, []int) {
	return file_charge_point_proto_rawDescGZIP(), []int{1}
}

func (x *ChargePointInfo) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *ChargePointInfo) GetMaxPower() float32 {
	if x != nil && x.MaxPower != nil {
		return *x.MaxPower
	}
	return 0
}

func (x *ChargePointInfo) GetOcppInfo() *OcppInfo {
	if x != nil {
		return x.OcppInfo
	}
	return nil
}

func (x *ChargePointInfo) GetFreeCharging() *FreeCharging {
	if x != nil {
		return x.FreeCharging
	}
	return nil
}

type OcppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vendor                  string  `protobuf:"bytes,1,opt,name=Vendor,proto3" json:"Vendor,omitempty"`
	Model                   string  `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`
	Iccid                   *string `protobuf:"bytes,3,opt,name=Iccid,proto3,oneof" json:"Iccid,omitempty"`
	Imsi                    *string `protobuf:"bytes,4,opt,name=Imsi,proto3,oneof" json:"Imsi,omitempty"`
	ChargeBoxSerialNumber   *string `protobuf:"bytes,5,opt,name=ChargeBoxSerialNumber,proto3,oneof" json:"ChargeBoxSerialNumber,omitempty"`
	ChargePointSerialNumber *string `protobuf:"bytes,6,opt,name=ChargePointSerialNumber,proto3,oneof" json:"ChargePointSerialNumber,omitempty"`
}

func (x *OcppInfo) Reset() {
	*x = OcppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charge_point_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcppInfo) ProtoMessage() {}

func (x *OcppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_charge_point_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcppInfo.ProtoReflect.Descriptor instead.
func (*OcppInfo) Descriptor() ([]byte, []int) {
	return file_charge_point_proto_rawDescGZIP(), []int{2}
}

func (x *OcppInfo) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *OcppInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *OcppInfo) GetIccid() string {
	if x != nil && x.Iccid != nil {
		return *x.Iccid
	}
	return ""
}

func (x *OcppInfo) GetImsi() string {
	if x != nil && x.Imsi != nil {
		return *x.Imsi
	}
	return ""
}

func (x *OcppInfo) GetChargeBoxSerialNumber() string {
	if x != nil && x.ChargeBoxSerialNumber != nil {
		return *x.ChargeBoxSerialNumber
	}
	return ""
}

func (x *OcppInfo) GetChargePointSerialNumber() string {
	if x != nil && x.ChargePointSerialNumber != nil {
		return *x.ChargePointSerialNumber
	}
	return ""
}

type FreeCharging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled         bool   `protobuf:"varint,1,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	MaxChargingTime *int32 `protobuf:"varint,2,opt,name=MaxChargingTime,proto3,oneof" json:"MaxChargingTime,omitempty"`
}

func (x *FreeCharging) Reset() {
	*x = FreeCharging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charge_point_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeCharging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeCharging) ProtoMessage() {}

func (x *FreeCharging) ProtoReflect() protoreflect.Message {
	mi := &file_charge_point_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeCharging.ProtoReflect.Descriptor instead.
func (*FreeCharging) Descriptor() ([]byte, []int) {
	return file_charge_point_proto_rawDescGZIP(), []int{3}
}

func (x *FreeCharging) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *FreeCharging) GetMaxChargingTime() int32 {
	if x != nil && x.MaxChargingTime != nil {
		return *x.MaxChargingTime
	}
	return 0
}

type AuthorizedCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId      string               `protobuf:"bytes,1,opt,name=TagId,proto3" json:"TagId,omitempty"`
	Status     string               `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	ExpiryDate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=ExpiryDate,proto3,oneof" json:"ExpiryDate,omitempty"`
}

func (x *AuthorizedCard) Reset() {
	*x = AuthorizedCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charge_point_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedCard) ProtoMessage() {}

func (x *AuthorizedCard) ProtoReflect() protoreflect.Message {
	mi := &file_charge_point_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizedCard.ProtoReflect.Descriptor instead.
func (*AuthorizedCard) Descriptor() ([]byte, []int) {
	return file_charge_point_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorizedCard) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *AuthorizedCard) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AuthorizedCard) GetExpiryDate() *timestamp.Timestamp {
	if x != nil {
		return x.ExpiryDate
	}
	return nil
}

type OcppVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string  `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Readonly bool    `protobuf:"varint,2,opt,name=Readonly,proto3" json:"Readonly,omitempty"`
	Value    *string `protobuf:"bytes,3,opt,name=Value,proto3,oneof" json:"Value,omitempty"`
}

func (x *OcppVariable) Reset() {
	*x = OcppVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charge_point_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcppVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcppVariable) ProtoMessage() {}

func (x *OcppVariable) ProtoReflect() protoreflect.Message {
	mi := &file_charge_point_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcppVariable.ProtoReflect.Descriptor instead.
func (*OcppVariable) Descriptor() ([]byte, []int) {
	return file_charge_point_proto_rawDescGZIP(), []int{5}
}

func (x *OcppVariable) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OcppVariable) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *OcppVariable) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_charge_point_proto protoreflect.FileDescriptor

var file_charge_point_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x29,
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x3e, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x31, 0x0a, 0x11, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x11, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x11, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x54, 0x4c, 0x53, 0x5f, 0x43, 0x41,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x10, 0x54, 0x4c, 0x53, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x15, 0x54, 0x4c, 0x53, 0x5f,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x14, 0x54, 0x4c, 0x53, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x36, 0x0a, 0x14, 0x54, 0x4c, 0x53, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x06, 0x52, 0x13, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x75, 0x72, 0x6c, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x54, 0x4c, 0x53, 0x5f, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x54, 0x4c, 0x53, 0x5f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x42, 0x17, 0x0a, 0x15, 0x5f, 0x54, 0x4c, 0x53, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x08, 0x4f, 0x63, 0x70, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4f, 0x63, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x4f, 0x63, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x35, 0x0a, 0x0c, 0x46, 0x72, 0x65, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46,
	0x72, 0x65, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x46, 0x72, 0x65,
	0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x22,
	0xaf, 0x02, 0x0a, 0x08, 0x4f, 0x63, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x63,
	0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x49, 0x63, 0x63,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x49, 0x6d, 0x73, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x49, 0x6d, 0x73, 0x69, 0x88, 0x01, 0x01, 0x12, 0x39,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x15, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x17, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x17, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x63, 0x63,
	0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x49, 0x6d, 0x73, 0x69, 0x42, 0x18, 0x0a, 0x16, 0x5f,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x6b, 0x0a, 0x0c, 0x46, 0x72, 0x65, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x0f, 0x4d,
	0x61, 0x78, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0f, 0x4d, 0x61, 0x78, 0x43, 0x68, 0x61, 0x72, 0x67,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x4d,
	0x61, 0x78, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8e,
	0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3f, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x00, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x61, 0x0a, 0x0c, 0x4f, 0x63, 0x70, 0x70, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x52, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x12, 0x19, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x42, 0x6c, 0x61, 0x7a, 0x33, 0x6b, 0x78, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x50, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_charge_point_proto_rawDescOnce sync.Once
	file_charge_point_proto_rawDescData = file_charge_point_proto_rawDesc
)

func file_charge_point_proto_rawDescGZIP() []byte {
	file_charge_point_proto_rawDescOnce.Do(func() {
		file_charge_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_charge_point_proto_rawDescData)
	})
	return file_charge_point_proto_rawDescData
}

var file_charge_point_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_charge_point_proto_goTypes = []interface{}{
	(*ConnectionSettings)(nil),  // 0: api.ConnectionSettings
	(*ChargePointInfo)(nil),     // 1: api.ChargePointInfo
	(*OcppInfo)(nil),            // 2: api.OcppInfo
	(*FreeCharging)(nil),        // 3: api.FreeCharging
	(*AuthorizedCard)(nil),      // 4: api.AuthorizedCard
	(*OcppVariable)(nil),        // 5: api.OcppVariable
	(ProtocolVersion)(0),        // 6: api.ProtocolVersion
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_charge_point_proto_depIdxs = []int32{
	6, // 0: api.ConnectionSettings.ProtocolVersion:type_name -> api.ProtocolVersion
	2, // 1: api.ChargePointInfo.OcppInfo:type_name -> api.OcppInfo
	3, // 2: api.ChargePointInfo.FreeCharging:type_name -> api.FreeCharging
	7, // 3: api.AuthorizedCard.ExpiryDate:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_charge_point_proto_init() }
func file_charge_point_proto_init() {
	if File_charge_point_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_charge_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionSettings); i {
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
		file_charge_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargePointInfo); i {
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
		file_charge_point_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcppInfo); i {
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
		file_charge_point_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeCharging); i {
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
		file_charge_point_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedCard); i {
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
		file_charge_point_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcppVariable); i {
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
	file_charge_point_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_charge_point_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_charge_point_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_charge_point_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_charge_point_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_charge_point_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_charge_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_charge_point_proto_goTypes,
		DependencyIndexes: file_charge_point_proto_depIdxs,
		MessageInfos:      file_charge_point_proto_msgTypes,
	}.Build()
	File_charge_point_proto = out.File
	file_charge_point_proto_rawDesc = nil
	file_charge_point_proto_goTypes = nil
	file_charge_point_proto_depIdxs = nil
}
