// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/template/v1/template.proto

package templatev1

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Templates move settings that are shared by many experiments into a single
// YAML file.
type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the template.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The template value.
	Config *_struct.Struct `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// The id of the workspace associated with this model.
	WorkspaceId int32 `protobuf:"varint,5,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_template_v1_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_determined_template_v1_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_determined_template_v1_template_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template) GetConfig() *_struct.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Template) GetWorkspaceId() int32 {
	if x != nil {
		return x.WorkspaceId
	}
	return 0
}

var File_determined_template_v1_template_proto protoreflect.FileDescriptor

var file_determined_template_v1_template_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x64, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x08,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x3a, 0x24, 0x92, 0x41, 0x21, 0x0a, 0x1f, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xd2, 0x01, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2d,
	0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_template_v1_template_proto_rawDescOnce sync.Once
	file_determined_template_v1_template_proto_rawDescData = file_determined_template_v1_template_proto_rawDesc
)

func file_determined_template_v1_template_proto_rawDescGZIP() []byte {
	file_determined_template_v1_template_proto_rawDescOnce.Do(func() {
		file_determined_template_v1_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_template_v1_template_proto_rawDescData)
	})
	return file_determined_template_v1_template_proto_rawDescData
}

var file_determined_template_v1_template_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_determined_template_v1_template_proto_goTypes = []interface{}{
	(*Template)(nil),       // 0: determined.template.v1.Template
	(*_struct.Struct)(nil), // 1: google.protobuf.Struct
}
var file_determined_template_v1_template_proto_depIdxs = []int32{
	1, // 0: determined.template.v1.Template.config:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_determined_template_v1_template_proto_init() }
func file_determined_template_v1_template_proto_init() {
	if File_determined_template_v1_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_template_v1_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
			RawDescriptor: file_determined_template_v1_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_template_v1_template_proto_goTypes,
		DependencyIndexes: file_determined_template_v1_template_proto_depIdxs,
		MessageInfos:      file_determined_template_v1_template_proto_msgTypes,
	}.Build()
	File_determined_template_v1_template_proto = out.File
	file_determined_template_v1_template_proto_rawDesc = nil
	file_determined_template_v1_template_proto_goTypes = nil
	file_determined_template_v1_template_proto_depIdxs = nil
}
