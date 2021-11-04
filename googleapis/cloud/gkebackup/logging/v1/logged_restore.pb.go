// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/gkebackup/logging/v1/logged_restore.proto

package logging

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible values for state of the Restore.
type LoggedRestore_State int32

const (
	// The Restore resource is in the process of being created.
	LoggedRestore_STATE_UNSPECIFIED LoggedRestore_State = 0
	// The Restore resource has been created and the associated RestoreJob
	// Kubernetes resource has been injected into target cluster.
	LoggedRestore_CREATING LoggedRestore_State = 1
	// The gkebackup agent in the cluster has begun executing the restore
	// operation.
	LoggedRestore_IN_PROGRESS LoggedRestore_State = 2
	// The restore operation has completed successfully. Restored workloads may
	// not yet be operational.
	LoggedRestore_SUCCEEDED LoggedRestore_State = 3
	// The restore operation has failed.
	LoggedRestore_FAILED LoggedRestore_State = 4
	// This Restore resource is in the process of being deleted.
	LoggedRestore_DELETING LoggedRestore_State = 5
)

// Enum value maps for LoggedRestore_State.
var (
	LoggedRestore_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "IN_PROGRESS",
		3: "SUCCEEDED",
		4: "FAILED",
		5: "DELETING",
	}
	LoggedRestore_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"IN_PROGRESS":       2,
		"SUCCEEDED":         3,
		"FAILED":            4,
		"DELETING":          5,
	}
)

func (x LoggedRestore_State) Enum() *LoggedRestore_State {
	p := new(LoggedRestore_State)
	*p = x
	return p
}

func (x LoggedRestore_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggedRestore_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_gkebackup_logging_v1_logged_restore_proto_enumTypes[0].Descriptor()
}

func (LoggedRestore_State) Type() protoreflect.EnumType {
	return &file_google_cloud_gkebackup_logging_v1_logged_restore_proto_enumTypes[0]
}

func (x LoggedRestore_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoggedRestore_State.Descriptor instead.
func (LoggedRestore_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescGZIP(), []int{0, 0}
}

// Restore as stored in Platform log. It's used to log the update details of a
// updateRestore request, so only mutable and non-output_only fields are
// included here..
type LoggedRestore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of the Backup resource this Restore resource used to restore
	// from. Format: projects/*/locations/*/backupPlans/*/backups/*.
	Backup string `protobuf:"bytes,1,opt,name=backup,proto3" json:"backup,omitempty"`
	// GCP Labels.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified descriptive string for this Restore.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The current state of the Restore.
	State LoggedRestore_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.gkebackup.logging.v1.LoggedRestore_State" json:"state,omitempty"`
	// Human-readable description of why the Restore is in its current state.
	StateReason string `protobuf:"bytes,5,opt,name=state_reason,json=stateReason,proto3" json:"state_reason,omitempty"`
}

func (x *LoggedRestore) Reset() {
	*x = LoggedRestore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_logging_v1_logged_restore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggedRestore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggedRestore) ProtoMessage() {}

func (x *LoggedRestore) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_logging_v1_logged_restore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggedRestore.ProtoReflect.Descriptor instead.
func (*LoggedRestore) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescGZIP(), []int{0}
}

func (x *LoggedRestore) GetBackup() string {
	if x != nil {
		return x.Backup
	}
	return ""
}

func (x *LoggedRestore) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LoggedRestore) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LoggedRestore) GetState() LoggedRestore_State {
	if x != nil {
		return x.State
	}
	return LoggedRestore_STATE_UNSPECIFIED
}

func (x *LoggedRestore) GetStateReason() string {
	if x != nil {
		return x.StateReason
	}
	return ""
}

var File_google_cloud_gkebackup_logging_v1_logged_restore_proto protoreflect.FileDescriptor

var file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xb3, 0x03, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x54, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x05, 0x42, 0xf3, 0x01, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x6b, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5c, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3a, 0x3a, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescOnce sync.Once
	file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescData = file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDesc
)

func file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescGZIP() []byte {
	file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescOnce.Do(func() {
		file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescData)
	})
	return file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDescData
}

var file_google_cloud_gkebackup_logging_v1_logged_restore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_gkebackup_logging_v1_logged_restore_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_gkebackup_logging_v1_logged_restore_proto_goTypes = []interface{}{
	(LoggedRestore_State)(0), // 0: google.cloud.gkebackup.logging.v1.LoggedRestore.State
	(*LoggedRestore)(nil),    // 1: google.cloud.gkebackup.logging.v1.LoggedRestore
	nil,                      // 2: google.cloud.gkebackup.logging.v1.LoggedRestore.LabelsEntry
}
var file_google_cloud_gkebackup_logging_v1_logged_restore_proto_depIdxs = []int32{
	2, // 0: google.cloud.gkebackup.logging.v1.LoggedRestore.labels:type_name -> google.cloud.gkebackup.logging.v1.LoggedRestore.LabelsEntry
	0, // 1: google.cloud.gkebackup.logging.v1.LoggedRestore.state:type_name -> google.cloud.gkebackup.logging.v1.LoggedRestore.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_gkebackup_logging_v1_logged_restore_proto_init() }
func file_google_cloud_gkebackup_logging_v1_logged_restore_proto_init() {
	if File_google_cloud_gkebackup_logging_v1_logged_restore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_gkebackup_logging_v1_logged_restore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggedRestore); i {
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
			RawDescriptor: file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_gkebackup_logging_v1_logged_restore_proto_goTypes,
		DependencyIndexes: file_google_cloud_gkebackup_logging_v1_logged_restore_proto_depIdxs,
		EnumInfos:         file_google_cloud_gkebackup_logging_v1_logged_restore_proto_enumTypes,
		MessageInfos:      file_google_cloud_gkebackup_logging_v1_logged_restore_proto_msgTypes,
	}.Build()
	File_google_cloud_gkebackup_logging_v1_logged_restore_proto = out.File
	file_google_cloud_gkebackup_logging_v1_logged_restore_proto_rawDesc = nil
	file_google_cloud_gkebackup_logging_v1_logged_restore_proto_goTypes = nil
	file_google_cloud_gkebackup_logging_v1_logged_restore_proto_depIdxs = nil
}