// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/log/config.proto

package log

import (
	log "github.com/v2fly/v2ray-core/v4/common/log"
	_ "github.com/v2fly/v2ray-core/v4/common/protoext"
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

type LogType int32

const (
	LogType_None    LogType = 0
	LogType_Console LogType = 1
	LogType_File    LogType = 2
	LogType_Event   LogType = 3
)

// Enum value maps for LogType.
var (
	LogType_name = map[int32]string{
		0: "None",
		1: "Console",
		2: "File",
		3: "Event",
	}
	LogType_value = map[string]int32{
		"None":    0,
		"Console": 1,
		"File":    2,
		"Event":   3,
	}
)

func (x LogType) Enum() *LogType {
	p := new(LogType)
	*p = x
	return p
}

func (x LogType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_log_config_proto_enumTypes[0].Descriptor()
}

func (LogType) Type() protoreflect.EnumType {
	return &file_app_log_config_proto_enumTypes[0]
}

func (x LogType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogType.Descriptor instead.
func (LogType) EnumDescriptor() ([]byte, []int) {
	return file_app_log_config_proto_rawDescGZIP(), []int{0}
}

type LogSpecification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  LogType      `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.log.LogType" json:"type,omitempty"`
	Level log.Severity `protobuf:"varint,2,opt,name=level,proto3,enum=v2ray.core.common.log.Severity" json:"level,omitempty"`
	Path  string       `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *LogSpecification) Reset() {
	*x = LogSpecification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_log_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSpecification) ProtoMessage() {}

func (x *LogSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_app_log_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSpecification.ProtoReflect.Descriptor instead.
func (*LogSpecification) Descriptor() ([]byte, []int) {
	return file_app_log_config_proto_rawDescGZIP(), []int{0}
}

func (x *LogSpecification) GetType() LogType {
	if x != nil {
		return x.Type
	}
	return LogType_None
}

func (x *LogSpecification) GetLevel() log.Severity {
	if x != nil {
		return x.Level
	}
	return log.Severity(0)
}

func (x *LogSpecification) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *LogSpecification `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Access *LogSpecification `protobuf:"bytes,7,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_log_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_log_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_log_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetError() *LogSpecification {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Config) GetAccess() *LogSpecification {
	if x != nil {
		return x.Access
	}
	return nil
}

var File_app_log_config_proto protoreflect.FileDescriptor

var file_app_log_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78,
	0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0xb8, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x4c, 0x6f, 0x67, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x16, 0x82, 0xb5, 0x18, 0x09, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x82, 0xb5, 0x18, 0x05, 0x12, 0x03, 0x6c, 0x6f, 0x67,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x2a, 0x35,
	0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x10, 0x03, 0x42, 0x57, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x50,
	0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32,
	0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x34, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0xaa, 0x02, 0x12, 0x56, 0x32, 0x52, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x4c, 0x6f, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_log_config_proto_rawDescOnce sync.Once
	file_app_log_config_proto_rawDescData = file_app_log_config_proto_rawDesc
)

func file_app_log_config_proto_rawDescGZIP() []byte {
	file_app_log_config_proto_rawDescOnce.Do(func() {
		file_app_log_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_log_config_proto_rawDescData)
	})
	return file_app_log_config_proto_rawDescData
}

var file_app_log_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_log_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_log_config_proto_goTypes = []interface{}{
	(LogType)(0),             // 0: v2ray.core.app.log.LogType
	(*LogSpecification)(nil), // 1: v2ray.core.app.log.LogSpecification
	(*Config)(nil),           // 2: v2ray.core.app.log.Config
	(log.Severity)(0),        // 3: v2ray.core.common.log.Severity
}
var file_app_log_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.log.LogSpecification.type:type_name -> v2ray.core.app.log.LogType
	3, // 1: v2ray.core.app.log.LogSpecification.level:type_name -> v2ray.core.common.log.Severity
	1, // 2: v2ray.core.app.log.Config.error:type_name -> v2ray.core.app.log.LogSpecification
	1, // 3: v2ray.core.app.log.Config.access:type_name -> v2ray.core.app.log.LogSpecification
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_app_log_config_proto_init() }
func file_app_log_config_proto_init() {
	if File_app_log_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_log_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSpecification); i {
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
		file_app_log_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_app_log_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_log_config_proto_goTypes,
		DependencyIndexes: file_app_log_config_proto_depIdxs,
		EnumInfos:         file_app_log_config_proto_enumTypes,
		MessageInfos:      file_app_log_config_proto_msgTypes,
	}.Build()
	File_app_log_config_proto = out.File
	file_app_log_config_proto_rawDesc = nil
	file_app_log_config_proto_goTypes = nil
	file_app_log_config_proto_depIdxs = nil
}
