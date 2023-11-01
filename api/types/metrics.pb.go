//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/types/metrics.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ID        string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Data      *anypb.Any             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_types_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_types_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_types_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Metric) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Metric) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_github_com_containerd_containerd_api_types_metrics_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_types_metrics_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_types_metrics_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_types_metrics_proto_rawDescData = file_github_com_containerd_containerd_api_types_metrics_proto_rawDesc
)

func file_github_com_containerd_containerd_api_types_metrics_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_types_metrics_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_types_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_types_metrics_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_types_metrics_proto_rawDescData
}

var file_github_com_containerd_containerd_api_types_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_containerd_containerd_api_types_metrics_proto_goTypes = []interface{}{
	(*Metric)(nil),                // 0: containerd.types.Metric
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
}
var file_github_com_containerd_containerd_api_types_metrics_proto_depIdxs = []int32{
	1, // 0: containerd.types.Metric.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: containerd.types.Metric.data:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_types_metrics_proto_init() }
func file_github_com_containerd_containerd_api_types_metrics_proto_init() {
	if File_github_com_containerd_containerd_api_types_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_types_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_types_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_types_metrics_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_types_metrics_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_types_metrics_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_types_metrics_proto = out.File
	file_github_com_containerd_containerd_api_types_metrics_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_types_metrics_proto_goTypes = nil
	file_github_com_containerd_containerd_api_types_metrics_proto_depIdxs = nil
}
