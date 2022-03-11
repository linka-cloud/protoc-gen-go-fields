// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tests/pb/test.proto

package test

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

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AString      *string     `protobuf:"bytes,1,opt,name=a_string,json=aString,proto3,oneof" json:"a_string,omitempty"`
	AnInt        int64       `protobuf:"varint,2,opt,name=an_int,json=anInt,proto3" json:"an_int,omitempty"`
	SubMessage   *Sub        `protobuf:"bytes,3,opt,name=sub_message,json=subMessage,proto3" json:"sub_message,omitempty"`
	InnerMessage *Test_Inner `protobuf:"bytes,4,opt,name=inner_message,json=innerMessage,proto3" json:"inner_message,omitempty"`
	// Types that are assignable to Oneof:
	//	*Test_SubOneof
	//	*Test_InnerOneof
	Oneof isTest_Oneof `protobuf_oneof:"oneof"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetAString() string {
	if x != nil && x.AString != nil {
		return *x.AString
	}
	return ""
}

func (x *Test) GetAnInt() int64 {
	if x != nil {
		return x.AnInt
	}
	return 0
}

func (x *Test) GetSubMessage() *Sub {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

func (x *Test) GetInnerMessage() *Test_Inner {
	if x != nil {
		return x.InnerMessage
	}
	return nil
}

func (m *Test) GetOneof() isTest_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *Test) GetSubOneof() *Sub {
	if x, ok := x.GetOneof().(*Test_SubOneof); ok {
		return x.SubOneof
	}
	return nil
}

func (x *Test) GetInnerOneof() *Test_Inner {
	if x, ok := x.GetOneof().(*Test_InnerOneof); ok {
		return x.InnerOneof
	}
	return nil
}

type isTest_Oneof interface {
	isTest_Oneof()
}

type Test_SubOneof struct {
	SubOneof *Sub `protobuf:"bytes,5,opt,name=sub_oneof,json=subOneof,proto3,oneof"`
}

type Test_InnerOneof struct {
	InnerOneof *Test_Inner `protobuf:"bytes,6,opt,name=inner_oneof,json=innerOneof,proto3,oneof"`
}

func (*Test_SubOneof) isTest_Oneof() {}

func (*Test_InnerOneof) isTest_Oneof() {}

type Sub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldOne string `protobuf:"bytes,1,opt,name=field_one,json=fieldOne,proto3" json:"field_one,omitempty"`
	FieldTwo int32  `protobuf:"varint,2,opt,name=field_two,json=fieldTwo,proto3" json:"field_two,omitempty"`
}

func (x *Sub) Reset() {
	*x = Sub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub) ProtoMessage() {}

func (x *Sub) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub.ProtoReflect.Descriptor instead.
func (*Sub) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{1}
}

func (x *Sub) GetFieldOne() string {
	if x != nil {
		return x.FieldOne
	}
	return ""
}

func (x *Sub) GetFieldTwo() int32 {
	if x != nil {
		return x.FieldTwo
	}
	return 0
}

type TestOptional struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AString      *string             `protobuf:"bytes,1,opt,name=a_string,json=aString,proto3,oneof" json:"a_string,omitempty"`
	AnInt        *int64              `protobuf:"varint,2,opt,name=an_int,json=anInt,proto3,oneof" json:"an_int,omitempty"`
	SubMessage   *Sub                `protobuf:"bytes,3,opt,name=sub_message,json=subMessage,proto3" json:"sub_message,omitempty"`
	InnerMessage *TestOptional_Inner `protobuf:"bytes,4,opt,name=inner_message,json=innerMessage,proto3" json:"inner_message,omitempty"`
	// Types that are assignable to Oneof:
	//	*TestOptional_SubOneof
	//	*TestOptional_InnerOneof
	Oneof isTestOptional_Oneof `protobuf_oneof:"oneof"`
}

func (x *TestOptional) Reset() {
	*x = TestOptional{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestOptional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOptional) ProtoMessage() {}

func (x *TestOptional) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestOptional.ProtoReflect.Descriptor instead.
func (*TestOptional) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestOptional) GetAString() string {
	if x != nil && x.AString != nil {
		return *x.AString
	}
	return ""
}

func (x *TestOptional) GetAnInt() int64 {
	if x != nil && x.AnInt != nil {
		return *x.AnInt
	}
	return 0
}

func (x *TestOptional) GetSubMessage() *Sub {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

func (x *TestOptional) GetInnerMessage() *TestOptional_Inner {
	if x != nil {
		return x.InnerMessage
	}
	return nil
}

func (m *TestOptional) GetOneof() isTestOptional_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *TestOptional) GetSubOneof() *Sub {
	if x, ok := x.GetOneof().(*TestOptional_SubOneof); ok {
		return x.SubOneof
	}
	return nil
}

func (x *TestOptional) GetInnerOneof() *TestOptional_Inner {
	if x, ok := x.GetOneof().(*TestOptional_InnerOneof); ok {
		return x.InnerOneof
	}
	return nil
}

type isTestOptional_Oneof interface {
	isTestOptional_Oneof()
}

type TestOptional_SubOneof struct {
	SubOneof *Sub `protobuf:"bytes,5,opt,name=sub_oneof,json=subOneof,proto3,oneof"`
}

type TestOptional_InnerOneof struct {
	InnerOneof *TestOptional_Inner `protobuf:"bytes,6,opt,name=inner_oneof,json=innerOneof,proto3,oneof"`
}

func (*TestOptional_SubOneof) isTestOptional_Oneof() {}

func (*TestOptional_InnerOneof) isTestOptional_Oneof() {}

type SubOptional struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldOne string `protobuf:"bytes,1,opt,name=field_one,json=fieldOne,proto3" json:"field_one,omitempty"`
	FieldTwo int32  `protobuf:"varint,2,opt,name=field_two,json=fieldTwo,proto3" json:"field_two,omitempty"`
}

func (x *SubOptional) Reset() {
	*x = SubOptional{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubOptional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubOptional) ProtoMessage() {}

func (x *SubOptional) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubOptional.ProtoReflect.Descriptor instead.
func (*SubOptional) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{3}
}

func (x *SubOptional) GetFieldOne() string {
	if x != nil {
		return x.FieldOne
	}
	return ""
}

func (x *SubOptional) GetFieldTwo() int32 {
	if x != nil {
		return x.FieldTwo
	}
	return 0
}

type Test_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One string `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two string `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
}

func (x *Test_Inner) Reset() {
	*x = Test_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_Inner) ProtoMessage() {}

func (x *Test_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_Inner.ProtoReflect.Descriptor instead.
func (*Test_Inner) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Test_Inner) GetOne() string {
	if x != nil {
		return x.One
	}
	return ""
}

func (x *Test_Inner) GetTwo() string {
	if x != nil {
		return x.Two
	}
	return ""
}

type TestOptional_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One *string `protobuf:"bytes,1,opt,name=one,proto3,oneof" json:"one,omitempty"`
	Two *string `protobuf:"bytes,2,opt,name=two,proto3,oneof" json:"two,omitempty"`
}

func (x *TestOptional_Inner) Reset() {
	*x = TestOptional_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestOptional_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOptional_Inner) ProtoMessage() {}

func (x *TestOptional_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestOptional_Inner.ProtoReflect.Descriptor instead.
func (*TestOptional_Inner) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TestOptional_Inner) GetOne() string {
	if x != nil && x.One != nil {
		return *x.One
	}
	return ""
}

func (x *TestOptional_Inner) GetTwo() string {
	if x != nil && x.Two != nil {
		return *x.Two
	}
	return ""
}

var File_tests_pb_test_proto protoreflect.FileDescriptor

var file_tests_pb_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0xea, 0x02, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x08, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x07, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x61, 0x6e, 0x49, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x2e,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0c,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x73, 0x75, 0x62, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x75, 0x62, 0x48, 0x00, 0x52, 0x08, 0x73, 0x75, 0x62, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x12, 0x3d, 0x0a, 0x0b, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x1a,
	0x2b, 0x0a, 0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x77,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x42, 0x07, 0x0a, 0x05,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x03, 0x53, 0x75, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x77, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x08, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x05, 0x61, 0x6e, 0x49, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x34, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x0c, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x32, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x48, 0x00, 0x52, 0x08, 0x73, 0x75, 0x62, 0x4f, 0x6e,
	0x65, 0x6f, 0x66, 0x12, 0x45, 0x0a, 0x0b, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0a,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x1a, 0x45, 0x0a, 0x05, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x74, 0x77,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x88, 0x01,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6f, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x77,
	0x6f, 0x42, 0x07, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6e, 0x5f, 0x69,
	0x6e, 0x74, 0x22, 0x47, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x77, 0x6f, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_pb_test_proto_rawDescOnce sync.Once
	file_tests_pb_test_proto_rawDescData = file_tests_pb_test_proto_rawDesc
)

func file_tests_pb_test_proto_rawDescGZIP() []byte {
	file_tests_pb_test_proto_rawDescOnce.Do(func() {
		file_tests_pb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_pb_test_proto_rawDescData)
	})
	return file_tests_pb_test_proto_rawDescData
}

var file_tests_pb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tests_pb_test_proto_goTypes = []interface{}{
	(*Test)(nil),               // 0: go.fields.test.Test
	(*Sub)(nil),                // 1: go.fields.test.Sub
	(*TestOptional)(nil),       // 2: go.fields.test.TestOptional
	(*SubOptional)(nil),        // 3: go.fields.test.SubOptional
	(*Test_Inner)(nil),         // 4: go.fields.test.Test.Inner
	(*TestOptional_Inner)(nil), // 5: go.fields.test.TestOptional.Inner
}
var file_tests_pb_test_proto_depIdxs = []int32{
	1, // 0: go.fields.test.Test.sub_message:type_name -> go.fields.test.Sub
	4, // 1: go.fields.test.Test.inner_message:type_name -> go.fields.test.Test.Inner
	1, // 2: go.fields.test.Test.sub_oneof:type_name -> go.fields.test.Sub
	4, // 3: go.fields.test.Test.inner_oneof:type_name -> go.fields.test.Test.Inner
	1, // 4: go.fields.test.TestOptional.sub_message:type_name -> go.fields.test.Sub
	5, // 5: go.fields.test.TestOptional.inner_message:type_name -> go.fields.test.TestOptional.Inner
	1, // 6: go.fields.test.TestOptional.sub_oneof:type_name -> go.fields.test.Sub
	5, // 7: go.fields.test.TestOptional.inner_oneof:type_name -> go.fields.test.TestOptional.Inner
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tests_pb_test_proto_init() }
func file_tests_pb_test_proto_init() {
	if File_tests_pb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_pb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_tests_pb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub); i {
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
		file_tests_pb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestOptional); i {
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
		file_tests_pb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubOptional); i {
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
		file_tests_pb_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test_Inner); i {
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
		file_tests_pb_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestOptional_Inner); i {
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
	file_tests_pb_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Test_SubOneof)(nil),
		(*Test_InnerOneof)(nil),
	}
	file_tests_pb_test_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TestOptional_SubOneof)(nil),
		(*TestOptional_InnerOneof)(nil),
	}
	file_tests_pb_test_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_pb_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_pb_test_proto_goTypes,
		DependencyIndexes: file_tests_pb_test_proto_depIdxs,
		MessageInfos:      file_tests_pb_test_proto_msgTypes,
	}.Build()
	File_tests_pb_test_proto = out.File
	file_tests_pb_test_proto_rawDesc = nil
	file_tests_pb_test_proto_goTypes = nil
	file_tests_pb_test_proto_depIdxs = nil
}
