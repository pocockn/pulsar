//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: kv.proto

package kv

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

type Event_EventType int32

const (
	Event_PUT    Event_EventType = 0
	Event_DELETE Event_EventType = 1
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0: "PUT",
		1: "DELETE",
	}
	Event_EventType_value = map[string]int32{
		"PUT":    0,
		"DELETE": 1,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_kv_proto_enumTypes[0].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_kv_proto_enumTypes[0]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{1, 0}
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the key in bytes. An empty key is not allowed.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the key. A deletion resets
	// the version to zero and any modification of the key
	// increases its version.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// value is the value held by the key, in bytes.
	Value       []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	IsNumber    bool   `protobuf:"varint,6,opt,name=is_number,json=isNumber,proto3" json:"is_number,omitempty"`
	NumberValue int64  `protobuf:"varint,7,opt,name=number_value,json=numberValue,proto3" json:"number_value,omitempty"`
	// time in milliseconds when the record expires (0 for none)
	ExpireTime int64 `protobuf:"varint,8,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValue) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValue) GetCreateRevision() int64 {
	if x != nil {
		return x.CreateRevision
	}
	return 0
}

func (x *KeyValue) GetModRevision() int64 {
	if x != nil {
		return x.ModRevision
	}
	return 0
}

func (x *KeyValue) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyValue) GetIsNumber() bool {
	if x != nil {
		return x.IsNumber
	}
	return false
}

func (x *KeyValue) GetNumberValue() int64 {
	if x != nil {
		return x.NumberValue
	}
	return 0
}

func (x *KeyValue) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the kind of event. If type is a PUT, it indicates
	// new data has been stored to the key. If type is a DELETE,
	// it indicates the key was deleted.
	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=bookkeeper.proto.kv.Event_EventType" json:"type,omitempty"`
	// kv holds the KeyValue for the event.
	// A PUT event contains current kv pair.
	// A PUT event with kv.Version=1 indicates the creation of a key.
	// A DELETE/EXPIRE event contains the deleted key with
	// its modification revision set to the revision of deletion.
	Kv *KeyValue `protobuf:"bytes,2,opt,name=kv,proto3" json:"kv,omitempty"`
	// prev_kv holds the key-value pair before the event happens.
	PrevKv *KeyValue `protobuf:"bytes,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetType() Event_EventType {
	if x != nil {
		return x.Type
	}
	return Event_PUT
}

func (x *Event) GetKv() *KeyValue {
	if x != nil {
		return x.Kv
	}
	return nil
}

func (x *Event) GetPrevKv() *KeyValue {
	if x != nil {
		return x.PrevKv
	}
	return nil
}

var File_kv_proto protoreflect.FileDescriptor

var file_kv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x6f, 0x6f, 0x6b,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x76, 0x22,
	0xf9, 0x01, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d,
	0x6f, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x76, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2d, 0x0a, 0x02, 0x6b, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b,
	0x76, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x6b, 0x76, 0x12, 0x36,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6b, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6b, 0x76, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x70, 0x72, 0x65, 0x76, 0x4b, 0x76, 0x22, 0x20, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x42, 0x11, 0x50, 0x01, 0x5a, 0x0d, 0x62, 0x6f,
	0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x6b, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kv_proto_rawDescOnce sync.Once
	file_kv_proto_rawDescData = file_kv_proto_rawDesc
)

func file_kv_proto_rawDescGZIP() []byte {
	file_kv_proto_rawDescOnce.Do(func() {
		file_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_proto_rawDescData)
	})
	return file_kv_proto_rawDescData
}

var file_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kv_proto_goTypes = []interface{}{
	(Event_EventType)(0), // 0: bookkeeper.proto.kv.Event.EventType
	(*KeyValue)(nil),     // 1: bookkeeper.proto.kv.KeyValue
	(*Event)(nil),        // 2: bookkeeper.proto.kv.Event
}
var file_kv_proto_depIdxs = []int32{
	0, // 0: bookkeeper.proto.kv.Event.type:type_name -> bookkeeper.proto.kv.Event.EventType
	1, // 1: bookkeeper.proto.kv.Event.kv:type_name -> bookkeeper.proto.kv.KeyValue
	1, // 2: bookkeeper.proto.kv.Event.prev_kv:type_name -> bookkeeper.proto.kv.KeyValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kv_proto_init() }
func file_kv_proto_init() {
	if File_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_kv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kv_proto_goTypes,
		DependencyIndexes: file_kv_proto_depIdxs,
		EnumInfos:         file_kv_proto_enumTypes,
		MessageInfos:      file_kv_proto_msgTypes,
	}.Build()
	File_kv_proto = out.File
	file_kv_proto_rawDesc = nil
	file_kv_proto_goTypes = nil
	file_kv_proto_depIdxs = nil
}
