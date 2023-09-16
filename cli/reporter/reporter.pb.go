// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: reporter.proto

package reporter

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pass        bool         `protobuf:"varint,1,opt,name=pass,proto3" json:"pass,omitempty"`
	Title       string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Error       *ResultError `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Diagnostics []byte       `protobuf:"bytes,5,opt,name=diagnostics,proto3" json:"diagnostics,omitempty"` // yaml encoded
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetPass() bool {
	if x != nil {
		return x.Pass
	}
	return false
}

func (x *Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Result) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Result) GetError() *ResultError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Result) GetDiagnostics() []byte {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type ResultError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	At       string `protobuf:"bytes,2,opt,name=at,proto3" json:"at,omitempty"`
	Got      string `protobuf:"bytes,3,opt,name=got,proto3" json:"got,omitempty"`
	Expected string `protobuf:"bytes,4,opt,name=expected,proto3" json:"expected,omitempty"`
}

func (x *ResultError) Reset() {
	*x = ResultError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultError) ProtoMessage() {}

func (x *ResultError) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultError.ProtoReflect.Descriptor instead.
func (*ResultError) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{1}
}

func (x *ResultError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultError) GetAt() string {
	if x != nil {
		return x.At
	}
	return ""
}

func (x *ResultError) GetGot() string {
	if x != nil {
		return x.Got
	}
	return ""
}

func (x *ResultError) GetExpected() string {
	if x != nil {
		return x.Expected
	}
	return ""
}

var File_reporter_proto protoreflect.FileDescriptor

var file_reporter_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x63, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0xad, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70,
	0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x22, 0x65, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x67, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x6f, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reporter_proto_rawDescOnce sync.Once
	file_reporter_proto_rawDescData = file_reporter_proto_rawDesc
)

func file_reporter_proto_rawDescGZIP() []byte {
	file_reporter_proto_rawDescOnce.Do(func() {
		file_reporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_reporter_proto_rawDescData)
	})
	return file_reporter_proto_rawDescData
}

var file_reporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reporter_proto_goTypes = []interface{}{
	(*Result)(nil),      // 0: cnspec.reporter.v1.Result
	(*ResultError)(nil), // 1: cnspec.reporter.v1.ResultError
}
var file_reporter_proto_depIdxs = []int32{
	1, // 0: cnspec.reporter.v1.Result.error:type_name -> cnspec.reporter.v1.ResultError
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reporter_proto_init() }
func file_reporter_proto_init() {
	if File_reporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_reporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultError); i {
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
			RawDescriptor: file_reporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reporter_proto_goTypes,
		DependencyIndexes: file_reporter_proto_depIdxs,
		MessageInfos:      file_reporter_proto_msgTypes,
	}.Build()
	File_reporter_proto = out.File
	file_reporter_proto_rawDesc = nil
	file_reporter_proto_goTypes = nil
	file_reporter_proto_depIdxs = nil
}
