// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/omp_template_api/omp_template_api.proto

package omp_template_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Foo uint64 `protobuf:"varint,2,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_omp_template_api_omp_template_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_api_omp_template_api_omp_template_api_proto_msgTypes[0]
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
	return file_api_omp_template_api_omp_template_api_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Template) GetFoo() uint64 {
	if x != nil {
		return x.Foo
	}
	return 0
}

type DescribeTemplateV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeTemplateV1Request) Reset() {
	*x = DescribeTemplateV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_omp_template_api_omp_template_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTemplateV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTemplateV1Request) ProtoMessage() {}

func (x *DescribeTemplateV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_omp_template_api_omp_template_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTemplateV1Request.ProtoReflect.Descriptor instead.
func (*DescribeTemplateV1Request) Descriptor() ([]byte, []int) {
	return file_api_omp_template_api_omp_template_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeTemplateV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeTemplateV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Template `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DescribeTemplateV1Response) Reset() {
	*x = DescribeTemplateV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_omp_template_api_omp_template_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTemplateV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTemplateV1Response) ProtoMessage() {}

func (x *DescribeTemplateV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_omp_template_api_omp_template_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTemplateV1Response.ProtoReflect.Descriptor instead.
func (*DescribeTemplateV1Response) Descriptor() ([]byte, []int) {
	return file_api_omp_template_api_omp_template_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeTemplateV1Response) GetValue() *Template {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_api_omp_template_api_omp_template_api_proto protoreflect.FileDescriptor

var file_api_omp_template_api_omp_template_api_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6f,
	0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x6f, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2c, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x22, 0x34,
	0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x6f, 0x6d, 0x70, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb7,
	0x01, 0x0a, 0x15, 0x4f, 0x6d, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x31, 0x12,
	0x35, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x6f, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e,
	0x6f, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x6f, 0x6d,
	0x70, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6f, 0x6d, 0x70, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x3b, 0x6f, 0x6d, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_omp_template_api_omp_template_api_proto_rawDescOnce sync.Once
	file_api_omp_template_api_omp_template_api_proto_rawDescData = file_api_omp_template_api_omp_template_api_proto_rawDesc
)

func file_api_omp_template_api_omp_template_api_proto_rawDescGZIP() []byte {
	file_api_omp_template_api_omp_template_api_proto_rawDescOnce.Do(func() {
		file_api_omp_template_api_omp_template_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_omp_template_api_omp_template_api_proto_rawDescData)
	})
	return file_api_omp_template_api_omp_template_api_proto_rawDescData
}

var file_api_omp_template_api_omp_template_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_omp_template_api_omp_template_api_proto_goTypes = []interface{}{
	(*Template)(nil),                   // 0: ozonmp.omp_template_api.v1.Template
	(*DescribeTemplateV1Request)(nil),  // 1: ozonmp.omp_template_api.v1.DescribeTemplateV1Request
	(*DescribeTemplateV1Response)(nil), // 2: ozonmp.omp_template_api.v1.DescribeTemplateV1Response
}
var file_api_omp_template_api_omp_template_api_proto_depIdxs = []int32{
	0, // 0: ozonmp.omp_template_api.v1.DescribeTemplateV1Response.value:type_name -> ozonmp.omp_template_api.v1.Template
	1, // 1: ozonmp.omp_template_api.v1.OmpTemplateApiService.DescribeTemplateV1:input_type -> ozonmp.omp_template_api.v1.DescribeTemplateV1Request
	2, // 2: ozonmp.omp_template_api.v1.OmpTemplateApiService.DescribeTemplateV1:output_type -> ozonmp.omp_template_api.v1.DescribeTemplateV1Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_omp_template_api_omp_template_api_proto_init() }
func file_api_omp_template_api_omp_template_api_proto_init() {
	if File_api_omp_template_api_omp_template_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_omp_template_api_omp_template_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_omp_template_api_omp_template_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTemplateV1Request); i {
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
		file_api_omp_template_api_omp_template_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTemplateV1Response); i {
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
			RawDescriptor: file_api_omp_template_api_omp_template_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_omp_template_api_omp_template_api_proto_goTypes,
		DependencyIndexes: file_api_omp_template_api_omp_template_api_proto_depIdxs,
		MessageInfos:      file_api_omp_template_api_omp_template_api_proto_msgTypes,
	}.Build()
	File_api_omp_template_api_omp_template_api_proto = out.File
	file_api_omp_template_api_omp_template_api_proto_rawDesc = nil
	file_api_omp_template_api_omp_template_api_proto_goTypes = nil
	file_api_omp_template_api_omp_template_api_proto_depIdxs = nil
}
