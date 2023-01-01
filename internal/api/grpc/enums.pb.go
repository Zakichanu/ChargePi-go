// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: enums.proto

package grpc

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

type ProtocolVersion int32

const (
	ProtocolVersion_OCPP16  ProtocolVersion = 0
	ProtocolVersion_OCPP201 ProtocolVersion = 1
)

// Enum value maps for ProtocolVersion.
var (
	ProtocolVersion_name = map[int32]string{
		0: "OCPP16",
		1: "OCPP201",
	}
	ProtocolVersion_value = map[string]int32{
		"OCPP16":  0,
		"OCPP201": 1,
	}
)

func (x ProtocolVersion) Enum() *ProtocolVersion {
	p := new(ProtocolVersion)
	*p = x
	return p
}

func (x ProtocolVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (ProtocolVersion) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x ProtocolVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtocolVersion.Descriptor instead.
func (ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

type ConnectorStatus int32

const (
	ConnectorStatus_Available     ConnectorStatus = 0
	ConnectorStatus_Preparing     ConnectorStatus = 1
	ConnectorStatus_Charging      ConnectorStatus = 2
	ConnectorStatus_Finishing     ConnectorStatus = 3
	ConnectorStatus_Unavailable   ConnectorStatus = 4
	ConnectorStatus_SuspendedEVSE ConnectorStatus = 5
	ConnectorStatus_SuspendedEV   ConnectorStatus = 6
	ConnectorStatus_Reserved      ConnectorStatus = 7
	ConnectorStatus_Faulted       ConnectorStatus = 8
)

// Enum value maps for ConnectorStatus.
var (
	ConnectorStatus_name = map[int32]string{
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
	ConnectorStatus_value = map[string]int32{
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

func (x ConnectorStatus) Enum() *ConnectorStatus {
	p := new(ConnectorStatus)
	*p = x
	return p
}

func (x ConnectorStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectorStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[1].Descriptor()
}

func (ConnectorStatus) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[1]
}

func (x ConnectorStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectorStatus.Descriptor instead.
func (ConnectorStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

type ErrorCode int32

const (
	ErrorCode_NoError              ErrorCode = 0
	ErrorCode_OtherError           ErrorCode = 1
	ErrorCode_ConnectorLockFailure ErrorCode = 3
	ErrorCode_EVCommunicationError ErrorCode = 4
	ErrorCode_GroundFailure        ErrorCode = 5
	ErrorCode_HighTemperature      ErrorCode = 6
	ErrorCode_InternalError        ErrorCode = 7
	ErrorCode_LocalListConflict    ErrorCode = 8
	ErrorCode_OverCurrentFailure   ErrorCode = 9
	ErrorCode_OverVoltage          ErrorCode = 10
	ErrorCode_PowerMeterFailure    ErrorCode = 11
	ErrorCode_PowerSwitchFailure   ErrorCode = 12
	ErrorCode_ReaderFailure        ErrorCode = 13
	ErrorCode_ResetFailure         ErrorCode = 14
	ErrorCode_UnderVoltage         ErrorCode = 15
	ErrorCode_WeakSignal           ErrorCode = 16
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
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
	ErrorCode_value = map[string]int32{
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

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[2].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[2]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{2}
}

type ConnectorType int32

const (
	ConnectorType_TYPE1   ConnectorType = 0
	ConnectorType_TYPE2   ConnectorType = 1
	ConnectorType_SCHUKO  ConnectorType = 2
	ConnectorType_CHADEMO ConnectorType = 3
)

// Enum value maps for ConnectorType.
var (
	ConnectorType_name = map[int32]string{
		0: "TYPE1",
		1: "TYPE2",
		2: "SCHUKO",
		3: "CHADEMO",
	}
	ConnectorType_value = map[string]int32{
		"TYPE1":   0,
		"TYPE2":   1,
		"SCHUKO":  2,
		"CHADEMO": 3,
	}
)

func (x ConnectorType) Enum() *ConnectorType {
	p := new(ConnectorType)
	*p = x
	return p
}

func (x ConnectorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectorType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[3].Descriptor()
}

func (ConnectorType) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[3]
}

func (x ConnectorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectorType.Descriptor instead.
func (ConnectorType) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{3}
}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x2a, 0x2a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x43, 0x50, 0x50, 0x31, 0x36, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x43, 0x50, 0x50, 0x32, 0x30, 0x31, 0x10, 0x01, 0x2a, 0x9c,
	0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x04, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x56, 0x53, 0x45, 0x10,
	0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x56,
	0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x10, 0x07,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x10, 0x08, 0x2a, 0xcd, 0x02,
	0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e,
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
	0x0a, 0x57, 0x65, 0x61, 0x6b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x10, 0x10, 0x2a, 0x3e, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x54, 0x59, 0x50, 0x45, 0x31, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x59, 0x50,
	0x45, 0x32, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x55, 0x4b, 0x4f, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x44, 0x45, 0x4d, 0x4f, 0x10, 0x03, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x42, 0x6c, 0x61,
	0x7a, 0x33, 0x6b, 0x78, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_enums_proto_goTypes = []interface{}{
	(ProtocolVersion)(0), // 0: api.ProtocolVersion
	(ConnectorStatus)(0), // 1: api.ConnectorStatus
	(ErrorCode)(0),       // 2: api.ErrorCode
	(ConnectorType)(0),   // 3: api.ConnectorType
}
var file_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
