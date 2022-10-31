// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: connector.proto

package grpc

import (
	_ "github.com/golang/protobuf/ptypes/empty"
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

type EVSE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32           `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	EVCC       *EVCC           `protobuf:"bytes,2,opt,name=EVCC,proto3" json:"EVCC,omitempty"`
	PowerMeter *PowerMeter     `protobuf:"bytes,3,opt,name=PowerMeter,proto3" json:"PowerMeter,omitempty"`
	Status     ConnectorStatus `protobuf:"varint,4,opt,name=status,proto3,enum=api.ConnectorStatus" json:"status,omitempty"`
	Session    *Session        `protobuf:"bytes,5,opt,name=Session,proto3" json:"Session,omitempty"`
}

func (x *EVSE) Reset() {
	*x = EVSE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVSE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVSE) ProtoMessage() {}

func (x *EVSE) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVSE.ProtoReflect.Descriptor instead.
func (*EVSE) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{0}
}

func (x *EVSE) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EVSE) GetEVCC() *EVCC {
	if x != nil {
		return x.EVCC
	}
	return nil
}

func (x *EVSE) GetPowerMeter() *PowerMeter {
	if x != nil {
		return x.PowerMeter
	}
	return nil
}

func (x *EVSE) GetStatus() ConnectorStatus {
	if x != nil {
		return x.Status
	}
	return ConnectorStatus_Available
}

func (x *EVSE) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string               `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	TagId         string               `protobuf:"bytes,2,opt,name=tagId,proto3" json:"tagId,omitempty"`
	Started       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=started,proto3" json:"started,omitempty"`
	Consumption   string               `protobuf:"bytes,4,opt,name=Consumption,proto3" json:"Consumption,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Session) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *Session) GetStarted() *timestamp.Timestamp {
	if x != nil {
		return x.Started
	}
	return nil
}

func (x *Session) GetConsumption() string {
	if x != nil {
		return x.Consumption
	}
	return ""
}

var File_connector_proto protoreflect.FileDescriptor

var file_connector_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x04, 0x45, 0x56, 0x53, 0x45, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x45, 0x56,
	0x43, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x56, 0x43, 0x43, 0x52, 0x04, 0x45, 0x56, 0x43, 0x43, 0x12, 0x2f, 0x0a, 0x0a, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0a,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x9d, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x42, 0x6c, 0x61, 0x7a, 0x33, 0x6b, 0x78, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connector_proto_rawDescOnce sync.Once
	file_connector_proto_rawDescData = file_connector_proto_rawDesc
)

func file_connector_proto_rawDescGZIP() []byte {
	file_connector_proto_rawDescOnce.Do(func() {
		file_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_connector_proto_rawDescData)
	})
	return file_connector_proto_rawDescData
}

var file_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_connector_proto_goTypes = []interface{}{
	(*EVSE)(nil),                // 0: api.EVSE
	(*Session)(nil),             // 1: api.Session
	(*EVCC)(nil),                // 2: api.EVCC
	(*PowerMeter)(nil),          // 3: api.PowerMeter
	(ConnectorStatus)(0),        // 4: api.ConnectorStatus
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_connector_proto_depIdxs = []int32{
	2, // 0: api.EVSE.EVCC:type_name -> api.EVCC
	3, // 1: api.EVSE.PowerMeter:type_name -> api.PowerMeter
	4, // 2: api.EVSE.status:type_name -> api.ConnectorStatus
	1, // 3: api.EVSE.Session:type_name -> api.Session
	5, // 4: api.Session.started:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_connector_proto_init() }
func file_connector_proto_init() {
	if File_connector_proto != nil {
		return
	}
	file_enums_proto_init()
	file_hardware_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_connector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVSE); i {
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
		file_connector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_connector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connector_proto_goTypes,
		DependencyIndexes: file_connector_proto_depIdxs,
		MessageInfos:      file_connector_proto_msgTypes,
	}.Build()
	File_connector_proto = out.File
	file_connector_proto_rawDesc = nil
	file_connector_proto_goTypes = nil
	file_connector_proto_depIdxs = nil
}
