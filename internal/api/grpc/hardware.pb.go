// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: hardware.proto

package grpc

import (
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type EVCC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Status string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Serial *Serial `protobuf:"bytes,3,opt,name=Serial,proto3,oneof" json:"Serial,omitempty"`
	ModBus *Modbus `protobuf:"bytes,4,opt,name=ModBus,proto3,oneof" json:"ModBus,omitempty"`
}

func (x *EVCC) Reset() {
	*x = EVCC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVCC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVCC) ProtoMessage() {}

func (x *EVCC) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVCC.ProtoReflect.Descriptor instead.
func (*EVCC) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{0}
}

func (x *EVCC) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EVCC) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EVCC) GetSerial() *Serial {
	if x != nil {
		return x.Serial
	}
	return nil
}

func (x *EVCC) GetModBus() *Modbus {
	if x != nil {
		return x.ModBus
	}
	return nil
}

type PowerMeter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Enabled bool    `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Modbus  *Modbus `protobuf:"bytes,3,opt,name=modbus,proto3,oneof" json:"modbus,omitempty"`
	Spi     *Spi    `protobuf:"bytes,4,opt,name=spi,proto3,oneof" json:"spi,omitempty"`
}

func (x *PowerMeter) Reset() {
	*x = PowerMeter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerMeter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerMeter) ProtoMessage() {}

func (x *PowerMeter) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerMeter.ProtoReflect.Descriptor instead.
func (*PowerMeter) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{1}
}

func (x *PowerMeter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PowerMeter) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PowerMeter) GetModbus() *Modbus {
	if x != nil {
		return x.Modbus
	}
	return nil
}

func (x *PowerMeter) GetSpi() *Spi {
	if x != nil {
		return x.Spi
	}
	return nil
}

type Display struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Enabled  bool    `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Language *string `protobuf:"bytes,3,opt,name=language,proto3,oneof" json:"language,omitempty"`
	I2C      *I2C    `protobuf:"bytes,4,opt,name=i2c,proto3,oneof" json:"i2c,omitempty"`
}

func (x *Display) Reset() {
	*x = Display{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Display) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Display) ProtoMessage() {}

func (x *Display) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Display.ProtoReflect.Descriptor instead.
func (*Display) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{2}
}

func (x *Display) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Display) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Display) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *Display) GetI2C() *I2C {
	if x != nil {
		return x.I2C
	}
	return nil
}

type TagReader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Enabled       bool    `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DeviceAddress *string `protobuf:"bytes,3,opt,name=deviceAddress,proto3,oneof" json:"deviceAddress,omitempty"`
	ResetPin      *string `protobuf:"bytes,4,opt,name=resetPin,proto3,oneof" json:"resetPin,omitempty"`
}

func (x *TagReader) Reset() {
	*x = TagReader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagReader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagReader) ProtoMessage() {}

func (x *TagReader) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagReader.ProtoReflect.Descriptor instead.
func (*TagReader) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{3}
}

func (x *TagReader) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TagReader) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TagReader) GetDeviceAddress() string {
	if x != nil && x.DeviceAddress != nil {
		return *x.DeviceAddress
	}
	return ""
}

func (x *TagReader) GetResetPin() string {
	if x != nil && x.ResetPin != nil {
		return *x.ResetPin
	}
	return ""
}

type Indicator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Enabled          bool    `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IndicateCardRead *bool   `protobuf:"varint,3,opt,name=indicateCardRead,proto3,oneof" json:"indicateCardRead,omitempty"`
	DataPin          *string `protobuf:"bytes,4,opt,name=dataPin,proto3,oneof" json:"dataPin,omitempty"`
	Invert           *bool   `protobuf:"varint,5,opt,name=invert,proto3,oneof" json:"invert,omitempty"`
}

func (x *Indicator) Reset() {
	*x = Indicator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Indicator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Indicator) ProtoMessage() {}

func (x *Indicator) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Indicator.ProtoReflect.Descriptor instead.
func (*Indicator) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{4}
}

func (x *Indicator) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Indicator) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Indicator) GetIndicateCardRead() bool {
	if x != nil && x.IndicateCardRead != nil {
		return *x.IndicateCardRead
	}
	return false
}

func (x *Indicator) GetDataPin() string {
	if x != nil && x.DataPin != nil {
		return *x.DataPin
	}
	return ""
}

func (x *Indicator) GetInvert() bool {
	if x != nil && x.Invert != nil {
		return *x.Invert
	}
	return false
}

type Spi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChipSelect int32 `protobuf:"varint,1,opt,name=chipSelect,proto3" json:"chipSelect,omitempty"`
	Bus        int32 `protobuf:"varint,2,opt,name=bus,proto3" json:"bus,omitempty"`
}

func (x *Spi) Reset() {
	*x = Spi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spi) ProtoMessage() {}

func (x *Spi) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spi.ProtoReflect.Descriptor instead.
func (*Spi) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{5}
}

func (x *Spi) GetChipSelect() int32 {
	if x != nil {
		return x.ChipSelect
	}
	return 0
}

func (x *Spi) GetBus() int32 {
	if x != nil {
		return x.Bus
	}
	return 0
}

type Modbus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol      string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	DeviceAddress string `protobuf:"bytes,2,opt,name=deviceAddress,proto3" json:"deviceAddress,omitempty"`
}

func (x *Modbus) Reset() {
	*x = Modbus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Modbus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Modbus) ProtoMessage() {}

func (x *Modbus) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Modbus.ProtoReflect.Descriptor instead.
func (*Modbus) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{6}
}

func (x *Modbus) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Modbus) GetDeviceAddress() string {
	if x != nil {
		return x.DeviceAddress
	}
	return ""
}

type I2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Bus     int32  `protobuf:"varint,2,opt,name=bus,proto3" json:"bus,omitempty"`
}

func (x *I2C) Reset() {
	*x = I2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I2C) ProtoMessage() {}

func (x *I2C) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I2C.ProtoReflect.Descriptor instead.
func (*I2C) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{7}
}

func (x *I2C) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *I2C) GetBus() int32 {
	if x != nil {
		return x.Bus
	}
	return 0
}

type Serial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceAddress string `protobuf:"bytes,1,opt,name=deviceAddress,proto3" json:"deviceAddress,omitempty"`
	BaudRate      uint32 `protobuf:"varint,2,opt,name=baudRate,proto3" json:"baudRate,omitempty"`
	Parity        uint32 `protobuf:"varint,3,opt,name=parity,proto3" json:"parity,omitempty"`
	DataBits      uint32 `protobuf:"varint,4,opt,name=dataBits,proto3" json:"dataBits,omitempty"`
	StopBits      uint32 `protobuf:"varint,5,opt,name=stopBits,proto3" json:"stopBits,omitempty"`
}

func (x *Serial) Reset() {
	*x = Serial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hardware_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Serial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Serial) ProtoMessage() {}

func (x *Serial) ProtoReflect() protoreflect.Message {
	mi := &file_hardware_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Serial.ProtoReflect.Descriptor instead.
func (*Serial) Descriptor() ([]byte, []int) {
	return file_hardware_proto_rawDescGZIP(), []int{8}
}

func (x *Serial) GetDeviceAddress() string {
	if x != nil {
		return x.DeviceAddress
	}
	return ""
}

func (x *Serial) GetBaudRate() uint32 {
	if x != nil {
		return x.BaudRate
	}
	return 0
}

func (x *Serial) GetParity() uint32 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *Serial) GetDataBits() uint32 {
	if x != nil {
		return x.DataBits
	}
	return 0
}

func (x *Serial) GetStopBits() uint32 {
	if x != nil {
		return x.StopBits
	}
	return 0
}

var File_hardware_proto protoreflect.FileDescriptor

var file_hardware_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x01, 0x0a, 0x04, 0x45, 0x56, 0x43, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x28, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x42, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x48, 0x01, 0x52, 0x06,
	0x4d, 0x6f, 0x64, 0x42, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4d, 0x6f, 0x64, 0x42, 0x75, 0x73, 0x22,
	0x98, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x06,
	0x6d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x6f, 0x64,
	0x62, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x70, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x69, 0x48, 0x01, 0x52,
	0x03, 0x73, 0x70, 0x69, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x64, 0x62,
	0x75, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x70, 0x69, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x03, 0x69, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x32, 0x63, 0x48, 0x01, 0x52, 0x03,
	0x69, 0x32, 0x63, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x69, 0x32, 0x63, 0x22, 0xa4, 0x01, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x69, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x69, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2f,
	0x0a, 0x10, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x10, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x50, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x50, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02,
	0x52, 0x06, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x69, 0x6e, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x22, 0x37, 0x0a, 0x03, 0x53, 0x70, 0x69, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x68, 0x69, 0x70, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x69, 0x70, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x62, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x75, 0x73,
	0x22, 0x4a, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x62, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x31, 0x0a, 0x03,
	0x49, 0x32, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x75, 0x73, 0x22,
	0x9a, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x75, 0x64, 0x52, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x62, 0x61, 0x75, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x42, 0x69, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x42, 0x69, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x74, 0x73, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x42, 0x6c, 0x61, 0x7a,
	0x33, 0x6b, 0x78, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hardware_proto_rawDescOnce sync.Once
	file_hardware_proto_rawDescData = file_hardware_proto_rawDesc
)

func file_hardware_proto_rawDescGZIP() []byte {
	file_hardware_proto_rawDescOnce.Do(func() {
		file_hardware_proto_rawDescData = protoimpl.X.CompressGZIP(file_hardware_proto_rawDescData)
	})
	return file_hardware_proto_rawDescData
}

var file_hardware_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hardware_proto_goTypes = []interface{}{
	(*EVCC)(nil),       // 0: api.EVCC
	(*PowerMeter)(nil), // 1: api.PowerMeter
	(*Display)(nil),    // 2: api.Display
	(*TagReader)(nil),  // 3: api.TagReader
	(*Indicator)(nil),  // 4: api.Indicator
	(*Spi)(nil),        // 5: api.Spi
	(*Modbus)(nil),     // 6: api.Modbus
	(*I2C)(nil),        // 7: api.I2c
	(*Serial)(nil),     // 8: api.Serial
}
var file_hardware_proto_depIdxs = []int32{
	8, // 0: api.EVCC.Serial:type_name -> api.Serial
	6, // 1: api.EVCC.ModBus:type_name -> api.Modbus
	6, // 2: api.PowerMeter.modbus:type_name -> api.Modbus
	5, // 3: api.PowerMeter.spi:type_name -> api.Spi
	7, // 4: api.Display.i2c:type_name -> api.I2c
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_hardware_proto_init() }
func file_hardware_proto_init() {
	if File_hardware_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hardware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVCC); i {
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
		file_hardware_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerMeter); i {
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
		file_hardware_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Display); i {
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
		file_hardware_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagReader); i {
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
		file_hardware_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Indicator); i {
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
		file_hardware_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spi); i {
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
		file_hardware_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Modbus); i {
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
		file_hardware_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I2C); i {
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
		file_hardware_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Serial); i {
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
	file_hardware_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_hardware_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_hardware_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_hardware_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_hardware_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hardware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hardware_proto_goTypes,
		DependencyIndexes: file_hardware_proto_depIdxs,
		MessageInfos:      file_hardware_proto_msgTypes,
	}.Build()
	File_hardware_proto = out.File
	file_hardware_proto_rawDesc = nil
	file_hardware_proto_goTypes = nil
	file_hardware_proto_depIdxs = nil
}
