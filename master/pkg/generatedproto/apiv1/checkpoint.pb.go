// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/api/v1/checkpoint.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	checkpointv1 "github.com/determined-ai/determined/master/pkg/generatedproto/checkpointv1"
	trialv1 "github.com/determined-ai/determined/master/pkg/generatedproto/trialv1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Get the requested checkpoint.
type GetCheckpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uuid for the requested checkpoint.
	CheckpointUuid string `protobuf:"bytes,1,opt,name=checkpoint_uuid,json=checkpointUuid,proto3" json:"checkpoint_uuid,omitempty"`
}

func (x *GetCheckpointRequest) Reset() {
	*x = GetCheckpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCheckpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCheckpointRequest) ProtoMessage() {}

func (x *GetCheckpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCheckpointRequest.ProtoReflect.Descriptor instead.
func (*GetCheckpointRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{0}
}

func (x *GetCheckpointRequest) GetCheckpointUuid() string {
	if x != nil {
		return x.CheckpointUuid
	}
	return ""
}

// Response to GetCheckpointRequest.
type GetCheckpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requested checkpoint.
	Checkpoint *checkpointv1.Checkpoint `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (x *GetCheckpointResponse) Reset() {
	*x = GetCheckpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCheckpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCheckpointResponse) ProtoMessage() {}

func (x *GetCheckpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCheckpointResponse.ProtoReflect.Descriptor instead.
func (*GetCheckpointResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{1}
}

func (x *GetCheckpointResponse) GetCheckpoint() *checkpointv1.Checkpoint {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

// Request for updating a checkpoints metadata.
type PostCheckpointMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The desired checkpoint fields and values.
	Checkpoint *checkpointv1.Checkpoint `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (x *PostCheckpointMetadataRequest) Reset() {
	*x = PostCheckpointMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCheckpointMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCheckpointMetadataRequest) ProtoMessage() {}

func (x *PostCheckpointMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCheckpointMetadataRequest.ProtoReflect.Descriptor instead.
func (*PostCheckpointMetadataRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{2}
}

func (x *PostCheckpointMetadataRequest) GetCheckpoint() *checkpointv1.Checkpoint {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

// Response to PostCheckpointRequest.
type PostCheckpointMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The updated checkpoint.
	Checkpoint *checkpointv1.Checkpoint `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (x *PostCheckpointMetadataResponse) Reset() {
	*x = PostCheckpointMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCheckpointMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCheckpointMetadataResponse) ProtoMessage() {}

func (x *PostCheckpointMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCheckpointMetadataResponse.ProtoReflect.Descriptor instead.
func (*PostCheckpointMetadataResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{3}
}

func (x *PostCheckpointMetadataResponse) GetCheckpoint() *checkpointv1.Checkpoint {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

// Request to delete files matching globs in checkpoints.
type CheckpointsRemoveFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of checkpoint_uuids for the requested checkpoints.
	CheckpointUuids []string `protobuf:"bytes,1,rep,name=checkpoint_uuids,json=checkpointUuids,proto3" json:"checkpoint_uuids,omitempty"`
	// The list of checkpoint_globs for the requested checkpoints.
	// If a value is set to the empty string the checkpoint will only
	// have its metadata refreshed.
	CheckpointGlobs []string `protobuf:"bytes,2,rep,name=checkpoint_globs,json=checkpointGlobs,proto3" json:"checkpoint_globs,omitempty"`
}

func (x *CheckpointsRemoveFilesRequest) Reset() {
	*x = CheckpointsRemoveFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointsRemoveFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointsRemoveFilesRequest) ProtoMessage() {}

func (x *CheckpointsRemoveFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointsRemoveFilesRequest.ProtoReflect.Descriptor instead.
func (*CheckpointsRemoveFilesRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{4}
}

func (x *CheckpointsRemoveFilesRequest) GetCheckpointUuids() []string {
	if x != nil {
		return x.CheckpointUuids
	}
	return nil
}

func (x *CheckpointsRemoveFilesRequest) GetCheckpointGlobs() []string {
	if x != nil {
		return x.CheckpointGlobs
	}
	return nil
}

// Response to CheckpointRemoveFilesRequest.
type CheckpointsRemoveFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckpointsRemoveFilesResponse) Reset() {
	*x = CheckpointsRemoveFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointsRemoveFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointsRemoveFilesResponse) ProtoMessage() {}

func (x *CheckpointsRemoveFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointsRemoveFilesResponse.ProtoReflect.Descriptor instead.
func (*CheckpointsRemoveFilesResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{5}
}

// Request to patch database info about a checkpoint.
type PatchCheckpointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of checkpoints to patch.
	Checkpoints []*checkpointv1.PatchCheckpoint `protobuf:"bytes,1,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
}

func (x *PatchCheckpointsRequest) Reset() {
	*x = PatchCheckpointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchCheckpointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchCheckpointsRequest) ProtoMessage() {}

func (x *PatchCheckpointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchCheckpointsRequest.ProtoReflect.Descriptor instead.
func (*PatchCheckpointsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{6}
}

func (x *PatchCheckpointsRequest) GetCheckpoints() []*checkpointv1.PatchCheckpoint {
	if x != nil {
		return x.Checkpoints
	}
	return nil
}

// Intentionally don't send the updated response for performance reasons.
type PatchCheckpointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PatchCheckpointsResponse) Reset() {
	*x = PatchCheckpointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchCheckpointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchCheckpointsResponse) ProtoMessage() {}

func (x *PatchCheckpointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchCheckpointsResponse.ProtoReflect.Descriptor instead.
func (*PatchCheckpointsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{7}
}

// Request to Delete the list of checkpoints
type DeleteCheckpointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of checkpoint_uuids for the requested checkpoint.
	CheckpointUuids []string `protobuf:"bytes,1,rep,name=checkpoint_uuids,json=checkpointUuids,proto3" json:"checkpoint_uuids,omitempty"`
}

func (x *DeleteCheckpointsRequest) Reset() {
	*x = DeleteCheckpointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCheckpointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCheckpointsRequest) ProtoMessage() {}

func (x *DeleteCheckpointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCheckpointsRequest.ProtoReflect.Descriptor instead.
func (*DeleteCheckpointsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCheckpointsRequest) GetCheckpointUuids() []string {
	if x != nil {
		return x.CheckpointUuids
	}
	return nil
}

// Response to DeleteCheckpointsRequest
type DeleteCheckpointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCheckpointsResponse) Reset() {
	*x = DeleteCheckpointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCheckpointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCheckpointsResponse) ProtoMessage() {}

func (x *DeleteCheckpointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCheckpointsResponse.ProtoReflect.Descriptor instead.
func (*DeleteCheckpointsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{9}
}

// Request for all metrics related to a given checkpoint
type GetTrialMetricsByCheckpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of the checkpoint.
	CheckpointUuid string `protobuf:"bytes,1,opt,name=checkpoint_uuid,json=checkpointUuid,proto3" json:"checkpoint_uuid,omitempty"`
	// Type of the TrialSourceInfo
	TrialSourceInfoType *trialv1.TrialSourceInfoType `protobuf:"varint,2,opt,name=trial_source_info_type,json=trialSourceInfoType,proto3,enum=determined.trial.v1.TrialSourceInfoType,oneof" json:"trial_source_info_type,omitempty"`
	// Metric Group string ("training", "validation", or anything else) (nil means
	// all groups)
	MetricGroup *string `protobuf:"bytes,3,opt,name=metric_group,json=metricGroup,proto3,oneof" json:"metric_group,omitempty"`
}

func (x *GetTrialMetricsByCheckpointRequest) Reset() {
	*x = GetTrialMetricsByCheckpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrialMetricsByCheckpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrialMetricsByCheckpointRequest) ProtoMessage() {}

func (x *GetTrialMetricsByCheckpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrialMetricsByCheckpointRequest.ProtoReflect.Descriptor instead.
func (*GetTrialMetricsByCheckpointRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{10}
}

func (x *GetTrialMetricsByCheckpointRequest) GetCheckpointUuid() string {
	if x != nil {
		return x.CheckpointUuid
	}
	return ""
}

func (x *GetTrialMetricsByCheckpointRequest) GetTrialSourceInfoType() trialv1.TrialSourceInfoType {
	if x != nil && x.TrialSourceInfoType != nil {
		return *x.TrialSourceInfoType
	}
	return trialv1.TrialSourceInfoType_TRIAL_SOURCE_INFO_TYPE_UNSPECIFIED
}

func (x *GetTrialMetricsByCheckpointRequest) GetMetricGroup() string {
	if x != nil && x.MetricGroup != nil {
		return *x.MetricGroup
	}
	return ""
}

// Response for all metrics related to a given checkpoint
type GetTrialMetricsByCheckpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the related trials and their metrics
	Metrics []*trialv1.MetricsReport `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *GetTrialMetricsByCheckpointResponse) Reset() {
	*x = GetTrialMetricsByCheckpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_checkpoint_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrialMetricsByCheckpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrialMetricsByCheckpointResponse) ProtoMessage() {}

func (x *GetTrialMetricsByCheckpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_checkpoint_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrialMetricsByCheckpointResponse.ProtoReflect.Descriptor instead.
func (*GetTrialMetricsByCheckpointResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_checkpoint_proto_rawDescGZIP(), []int{11}
}

func (x *GetTrialMetricsByCheckpointResponse) GetMetrics() []*trialv1.MetricsReport {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_determined_api_v1_checkpoint_proto protoreflect.FileDescriptor

var file_determined_api_v1_checkpoint_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x64, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x74,
	0x72, 0x69, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x75,
	0x69, 0x64, 0x22, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x3a, 0x12, 0x92, 0x41, 0x0f, 0x0a, 0x0d, 0xd2, 0x01, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x65, 0x0a, 0x1d, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x1e,
	0x50, 0x6f, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x1d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x67, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x73, 0x3a, 0x2b, 0x92, 0x41,
	0x28, 0x0a, 0x26, 0xd2, 0x01, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x73, 0xd2, 0x01, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x73, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7b, 0x0a, 0x17, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x3a, 0x13, 0x92, 0x41, 0x10, 0x0a, 0x0e, 0xd2, 0x01, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x73, 0x3a, 0x18, 0x92, 0x41, 0x15,
	0x0a, 0x13, 0xd2, 0x01, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x62, 0x0a, 0x16, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x13,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01, 0x3a, 0x17,
	0x92, 0x41, 0x14, 0x0a, 0x12, 0xd2, 0x01, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x74, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3a, 0x0f, 0x92, 0x41, 0x0c, 0x0a, 0x0a, 0xd2,
	0x01, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_api_v1_checkpoint_proto_rawDescOnce sync.Once
	file_determined_api_v1_checkpoint_proto_rawDescData = file_determined_api_v1_checkpoint_proto_rawDesc
)

func file_determined_api_v1_checkpoint_proto_rawDescGZIP() []byte {
	file_determined_api_v1_checkpoint_proto_rawDescOnce.Do(func() {
		file_determined_api_v1_checkpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_api_v1_checkpoint_proto_rawDescData)
	})
	return file_determined_api_v1_checkpoint_proto_rawDescData
}

var file_determined_api_v1_checkpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_determined_api_v1_checkpoint_proto_goTypes = []interface{}{
	(*GetCheckpointRequest)(nil),                // 0: determined.api.v1.GetCheckpointRequest
	(*GetCheckpointResponse)(nil),               // 1: determined.api.v1.GetCheckpointResponse
	(*PostCheckpointMetadataRequest)(nil),       // 2: determined.api.v1.PostCheckpointMetadataRequest
	(*PostCheckpointMetadataResponse)(nil),      // 3: determined.api.v1.PostCheckpointMetadataResponse
	(*CheckpointsRemoveFilesRequest)(nil),       // 4: determined.api.v1.CheckpointsRemoveFilesRequest
	(*CheckpointsRemoveFilesResponse)(nil),      // 5: determined.api.v1.CheckpointsRemoveFilesResponse
	(*PatchCheckpointsRequest)(nil),             // 6: determined.api.v1.PatchCheckpointsRequest
	(*PatchCheckpointsResponse)(nil),            // 7: determined.api.v1.PatchCheckpointsResponse
	(*DeleteCheckpointsRequest)(nil),            // 8: determined.api.v1.DeleteCheckpointsRequest
	(*DeleteCheckpointsResponse)(nil),           // 9: determined.api.v1.DeleteCheckpointsResponse
	(*GetTrialMetricsByCheckpointRequest)(nil),  // 10: determined.api.v1.GetTrialMetricsByCheckpointRequest
	(*GetTrialMetricsByCheckpointResponse)(nil), // 11: determined.api.v1.GetTrialMetricsByCheckpointResponse
	(*checkpointv1.Checkpoint)(nil),             // 12: determined.checkpoint.v1.Checkpoint
	(*checkpointv1.PatchCheckpoint)(nil),        // 13: determined.checkpoint.v1.PatchCheckpoint
	(trialv1.TrialSourceInfoType)(0),            // 14: determined.trial.v1.TrialSourceInfoType
	(*trialv1.MetricsReport)(nil),               // 15: determined.trial.v1.MetricsReport
}
var file_determined_api_v1_checkpoint_proto_depIdxs = []int32{
	12, // 0: determined.api.v1.GetCheckpointResponse.checkpoint:type_name -> determined.checkpoint.v1.Checkpoint
	12, // 1: determined.api.v1.PostCheckpointMetadataRequest.checkpoint:type_name -> determined.checkpoint.v1.Checkpoint
	12, // 2: determined.api.v1.PostCheckpointMetadataResponse.checkpoint:type_name -> determined.checkpoint.v1.Checkpoint
	13, // 3: determined.api.v1.PatchCheckpointsRequest.checkpoints:type_name -> determined.checkpoint.v1.PatchCheckpoint
	14, // 4: determined.api.v1.GetTrialMetricsByCheckpointRequest.trial_source_info_type:type_name -> determined.trial.v1.TrialSourceInfoType
	15, // 5: determined.api.v1.GetTrialMetricsByCheckpointResponse.metrics:type_name -> determined.trial.v1.MetricsReport
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_determined_api_v1_checkpoint_proto_init() }
func file_determined_api_v1_checkpoint_proto_init() {
	if File_determined_api_v1_checkpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_api_v1_checkpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCheckpointRequest); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCheckpointResponse); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCheckpointMetadataRequest); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCheckpointMetadataResponse); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointsRemoveFilesRequest); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointsRemoveFilesResponse); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchCheckpointsRequest); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchCheckpointsResponse); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCheckpointsRequest); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCheckpointsResponse); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrialMetricsByCheckpointRequest); i {
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
		file_determined_api_v1_checkpoint_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrialMetricsByCheckpointResponse); i {
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
	file_determined_api_v1_checkpoint_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_determined_api_v1_checkpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_api_v1_checkpoint_proto_goTypes,
		DependencyIndexes: file_determined_api_v1_checkpoint_proto_depIdxs,
		MessageInfos:      file_determined_api_v1_checkpoint_proto_msgTypes,
	}.Build()
	File_determined_api_v1_checkpoint_proto = out.File
	file_determined_api_v1_checkpoint_proto_rawDesc = nil
	file_determined_api_v1_checkpoint_proto_goTypes = nil
	file_determined_api_v1_checkpoint_proto_depIdxs = nil
}
