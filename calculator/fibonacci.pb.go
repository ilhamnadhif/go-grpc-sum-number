// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: fibonacci.proto

package calculator

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

type FibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fibonacci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{0}
}

func (x *FibonacciRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fibonacci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{1}
}

func (x *FibonacciResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_fibonacci_proto protoreflect.FileDescriptor

var file_fibonacci_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x28, 0x0a,
	0x10, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x46, 0x69, 0x62, 0x6f, 0x6e,
	0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x75,
	0x6d, 0x2d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fibonacci_proto_rawDescOnce sync.Once
	file_fibonacci_proto_rawDescData = file_fibonacci_proto_rawDesc
)

func file_fibonacci_proto_rawDescGZIP() []byte {
	file_fibonacci_proto_rawDescOnce.Do(func() {
		file_fibonacci_proto_rawDescData = protoimpl.X.CompressGZIP(file_fibonacci_proto_rawDescData)
	})
	return file_fibonacci_proto_rawDescData
}

var file_fibonacci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fibonacci_proto_goTypes = []interface{}{
	(*FibonacciRequest)(nil),  // 0: calculator.FibonacciRequest
	(*FibonacciResponse)(nil), // 1: calculator.FibonacciResponse
}
var file_fibonacci_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fibonacci_proto_init() }
func file_fibonacci_proto_init() {
	if File_fibonacci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fibonacci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciRequest); i {
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
		file_fibonacci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciResponse); i {
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
			RawDescriptor: file_fibonacci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fibonacci_proto_goTypes,
		DependencyIndexes: file_fibonacci_proto_depIdxs,
		MessageInfos:      file_fibonacci_proto_msgTypes,
	}.Build()
	File_fibonacci_proto = out.File
	file_fibonacci_proto_rawDesc = nil
	file_fibonacci_proto_goTypes = nil
	file_fibonacci_proto_depIdxs = nil
}
