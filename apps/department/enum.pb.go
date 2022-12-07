// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: apps/department/pb/enum.proto

package department

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

type ApplicationFormStatus int32

const (
	ApplicationFormStatus_NULL ApplicationFormStatus = 0
	// Pending todo
	ApplicationFormStatus_PENDDING ApplicationFormStatus = 1
	// Passed todo
	ApplicationFormStatus_PASSED ApplicationFormStatus = 2
	// Deny todo
	ApplicationFormStatus_DENY ApplicationFormStatus = 3
)

// Enum value maps for ApplicationFormStatus.
var (
	ApplicationFormStatus_name = map[int32]string{
		0: "NULL",
		1: "PENDDING",
		2: "PASSED",
		3: "DENY",
	}
	ApplicationFormStatus_value = map[string]int32{
		"NULL":     0,
		"PENDDING": 1,
		"PASSED":   2,
		"DENY":     3,
	}
)

func (x ApplicationFormStatus) Enum() *ApplicationFormStatus {
	p := new(ApplicationFormStatus)
	*p = x
	return p
}

func (x ApplicationFormStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationFormStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_department_pb_enum_proto_enumTypes[0].Descriptor()
}

func (ApplicationFormStatus) Type() protoreflect.EnumType {
	return &file_apps_department_pb_enum_proto_enumTypes[0]
}

func (x ApplicationFormStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationFormStatus.Descriptor instead.
func (ApplicationFormStatus) EnumDescriptor() ([]byte, []int) {
	return file_apps_department_pb_enum_proto_rawDescGZIP(), []int{0}
}

var File_apps_department_pb_enum_proto protoreflect.FileDescriptor

var file_apps_department_pb_enum_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x6b, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0x72, 0x61, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x45, 0x0a, 0x15,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x45, 0x4e, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e,
	0x59, 0x10, 0x03, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0x72, 0x61, 0x6d, 0x6f, 0x2f, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_department_pb_enum_proto_rawDescOnce sync.Once
	file_apps_department_pb_enum_proto_rawDescData = file_apps_department_pb_enum_proto_rawDesc
)

func file_apps_department_pb_enum_proto_rawDescGZIP() []byte {
	file_apps_department_pb_enum_proto_rawDescOnce.Do(func() {
		file_apps_department_pb_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_department_pb_enum_proto_rawDescData)
	})
	return file_apps_department_pb_enum_proto_rawDescData
}

var file_apps_department_pb_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_department_pb_enum_proto_goTypes = []interface{}{
	(ApplicationFormStatus)(0), // 0: kalandra.keyauth.department.ApplicationFormStatus
}
var file_apps_department_pb_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_department_pb_enum_proto_init() }
func file_apps_department_pb_enum_proto_init() {
	if File_apps_department_pb_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_department_pb_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_department_pb_enum_proto_goTypes,
		DependencyIndexes: file_apps_department_pb_enum_proto_depIdxs,
		EnumInfos:         file_apps_department_pb_enum_proto_enumTypes,
	}.Build()
	File_apps_department_pb_enum_proto = out.File
	file_apps_department_pb_enum_proto_rawDesc = nil
	file_apps_department_pb_enum_proto_goTypes = nil
	file_apps_department_pb_enum_proto_depIdxs = nil
}