// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: protofile/person.proto

package protofile

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

type JsonStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonData []byte `protobuf:"bytes,1,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
}

func (x *JsonStruct) Reset() {
	*x = JsonStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofile_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonStruct) ProtoMessage() {}

func (x *JsonStruct) ProtoReflect() protoreflect.Message {
	mi := &file_protofile_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonStruct.ProtoReflect.Descriptor instead.
func (*JsonStruct) Descriptor() ([]byte, []int) {
	return file_protofile_person_proto_rawDescGZIP(), []int{0}
}

func (x *JsonStruct) GetJsonData() []byte {
	if x != nil {
		return x.JsonData
	}
	return nil
}

var File_protofile_person_proto protoreflect.FileDescriptor

var file_protofile_person_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x32, 0xb0, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x0b, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x0b, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2d,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x0b, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x1a, 0x0b, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x26, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x0b, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofile_person_proto_rawDescOnce sync.Once
	file_protofile_person_proto_rawDescData = file_protofile_person_proto_rawDesc
)

func file_protofile_person_proto_rawDescGZIP() []byte {
	file_protofile_person_proto_rawDescOnce.Do(func() {
		file_protofile_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofile_person_proto_rawDescData)
	})
	return file_protofile_person_proto_rawDescData
}

var file_protofile_person_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protofile_person_proto_goTypes = []interface{}{
	(*JsonStruct)(nil), // 0: JsonStruct
}
var file_protofile_person_proto_depIdxs = []int32{
	0, // 0: Person.CreateData:input_type -> JsonStruct
	0, // 1: Person.GetDataById:input_type -> JsonStruct
	0, // 2: Person.GetAllDataByField:input_type -> JsonStruct
	0, // 3: Person.DeleteData:input_type -> JsonStruct
	0, // 4: Person.CreateData:output_type -> JsonStruct
	0, // 5: Person.GetDataById:output_type -> JsonStruct
	0, // 6: Person.GetAllDataByField:output_type -> JsonStruct
	0, // 7: Person.DeleteData:output_type -> JsonStruct
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofile_person_proto_init() }
func file_protofile_person_proto_init() {
	if File_protofile_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofile_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonStruct); i {
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
			RawDescriptor: file_protofile_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofile_person_proto_goTypes,
		DependencyIndexes: file_protofile_person_proto_depIdxs,
		MessageInfos:      file_protofile_person_proto_msgTypes,
	}.Build()
	File_protofile_person_proto = out.File
	file_protofile_person_proto_rawDesc = nil
	file_protofile_person_proto_goTypes = nil
	file_protofile_person_proto_depIdxs = nil
}