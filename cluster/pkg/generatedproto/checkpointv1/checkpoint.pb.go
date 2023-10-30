// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/checkpoint/v1/checkpoint.proto

package checkpointv1

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	commonv1 "github.com/determined-ai/determined/cluster/pkg/generatedproto/commonv1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The current state of the checkpoint.
type State int32

const (
	// The state of the checkpoint is unknown.
	State_STATE_UNSPECIFIED State = 0
	// The checkpoint is in an active state.
	State_STATE_ACTIVE State = 1
	// The checkpoint is persisted to checkpoint storage.
	State_STATE_COMPLETED State = 2
	// The checkpoint errored.
	State_STATE_ERROR State = 3
	// The checkpoint has been deleted.
	State_STATE_DELETED State = 4
	// The checkpoint has been partially deleted.
	State_STATE_PARTIALLY_DELETED State = 5
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_ACTIVE",
		2: "STATE_COMPLETED",
		3: "STATE_ERROR",
		4: "STATE_DELETED",
		5: "STATE_PARTIALLY_DELETED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED":       0,
		"STATE_ACTIVE":            1,
		"STATE_COMPLETED":         2,
		"STATE_ERROR":             3,
		"STATE_DELETED":           4,
		"STATE_PARTIALLY_DELETED": 5,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_determined_checkpoint_v1_checkpoint_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_determined_checkpoint_v1_checkpoint_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP(), []int{0}
}

// Sorts options for checkpoints by the given field.
type SortBy int32

const (
	// Returns checkpoints in an unsorted list.
	SortBy_SORT_BY_UNSPECIFIED SortBy = 0
	// Returns checkpoints sorted by UUID.
	SortBy_SORT_BY_UUID SortBy = 1
	// Returns checkpoints sorted by trial id.
	SortBy_SORT_BY_TRIAL_ID SortBy = 2
	// Returns checkpoints sorted by batch number.
	SortBy_SORT_BY_BATCH_NUMBER SortBy = 3
	// Returns checkpoints sorted by end time.
	SortBy_SORT_BY_END_TIME SortBy = 4
	// Returns checkpoints sorted by state.
	SortBy_SORT_BY_STATE SortBy = 5
	// Returns checkpoints sorted by the experiment's `searcher.metric`
	// configuration setting.
	SortBy_SORT_BY_SEARCHER_METRIC SortBy = 6
)

// Enum value maps for SortBy.
var (
	SortBy_name = map[int32]string{
		0: "SORT_BY_UNSPECIFIED",
		1: "SORT_BY_UUID",
		2: "SORT_BY_TRIAL_ID",
		3: "SORT_BY_BATCH_NUMBER",
		4: "SORT_BY_END_TIME",
		5: "SORT_BY_STATE",
		6: "SORT_BY_SEARCHER_METRIC",
	}
	SortBy_value = map[string]int32{
		"SORT_BY_UNSPECIFIED":     0,
		"SORT_BY_UUID":            1,
		"SORT_BY_TRIAL_ID":        2,
		"SORT_BY_BATCH_NUMBER":    3,
		"SORT_BY_END_TIME":        4,
		"SORT_BY_STATE":           5,
		"SORT_BY_SEARCHER_METRIC": 6,
	}
)

func (x SortBy) Enum() *SortBy {
	p := new(SortBy)
	*p = x
	return p
}

func (x SortBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortBy) Descriptor() protoreflect.EnumDescriptor {
	return file_determined_checkpoint_v1_checkpoint_proto_enumTypes[1].Descriptor()
}

func (SortBy) Type() protoreflect.EnumType {
	return &file_determined_checkpoint_v1_checkpoint_proto_enumTypes[1]
}

func (x SortBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortBy.Descriptor instead.
func (SortBy) EnumDescriptor() ([]byte, []int) {
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP(), []int{1}
}

// CheckpointTrainingMetadata is specifically metadata about training.
type CheckpointTrainingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the trial that created this checkpoint.
	TrialId *wrappers.Int32Value `protobuf:"bytes,1,opt,name=trial_id,json=trialId,proto3" json:"trial_id,omitempty"`
	// The ID of the experiment that created this checkpoint.
	ExperimentId *wrappers.Int32Value `protobuf:"bytes,2,opt,name=experiment_id,json=experimentId,proto3" json:"experiment_id,omitempty"`
	// The configuration of the experiment that created this checkpoint.
	ExperimentConfig *_struct.Struct `protobuf:"bytes,3,opt,name=experiment_config,json=experimentConfig,proto3" json:"experiment_config,omitempty"`
	// Hyperparameter values for the trial that created this checkpoint.
	Hparams *_struct.Struct `protobuf:"bytes,4,opt,name=hparams,proto3" json:"hparams,omitempty"`
	// Training metrics reported at the same steps_completed as the checkpoint.
	TrainingMetrics *commonv1.Metrics `protobuf:"bytes,5,opt,name=training_metrics,json=trainingMetrics,proto3" json:"training_metrics,omitempty"`
	// Validation metrics reported at the same steps_completed as the checkpoint.
	ValidationMetrics *commonv1.Metrics `protobuf:"bytes,6,opt,name=validation_metrics,json=validationMetrics,proto3" json:"validation_metrics,omitempty"`
	// Searcher metric (as specified by the expconf) at the same steps_completed
	// of the checkpoint.
	SearcherMetric *wrappers.DoubleValue `protobuf:"bytes,17,opt,name=searcher_metric,json=searcherMetric,proto3" json:"searcher_metric,omitempty"`
}

func (x *CheckpointTrainingMetadata) Reset() {
	*x = CheckpointTrainingMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointTrainingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointTrainingMetadata) ProtoMessage() {}

func (x *CheckpointTrainingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointTrainingMetadata.ProtoReflect.Descriptor instead.
func (*CheckpointTrainingMetadata) Descriptor() ([]byte, []int) {
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP(), []int{0}
}

func (x *CheckpointTrainingMetadata) GetTrialId() *wrappers.Int32Value {
	if x != nil {
		return x.TrialId
	}
	return nil
}

func (x *CheckpointTrainingMetadata) GetExperimentId() *wrappers.Int32Value {
	if x != nil {
		return x.ExperimentId
	}
	return nil
}

func (x *CheckpointTrainingMetadata) GetExperimentConfig() *_struct.Struct {
	if x != nil {
		return x.ExperimentConfig
	}
	return nil
}

func (x *CheckpointTrainingMetadata) GetHparams() *_struct.Struct {
	if x != nil {
		return x.Hparams
	}
	return nil
}

func (x *CheckpointTrainingMetadata) GetTrainingMetrics() *commonv1.Metrics {
	if x != nil {
		return x.TrainingMetrics
	}
	return nil
}

func (x *CheckpointTrainingMetadata) GetValidationMetrics() *commonv1.Metrics {
	if x != nil {
		return x.ValidationMetrics
	}
	return nil
}

func (x *CheckpointTrainingMetadata) GetSearcherMetric() *wrappers.DoubleValue {
	if x != nil {
		return x.SearcherMetric
	}
	return nil
}

// Checkpoint a collection of files saved by a task.
type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the task which generated this checkpoint.
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// ID of the allocation which generated this checkpoint.
	AllocationId *string `protobuf:"bytes,2,opt,name=allocation_id,json=allocationId,proto3,oneof" json:"allocation_id,omitempty"`
	// UUID of the checkpoint.
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Timestamp when the checkpoint was reported.
	ReportTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=report_time,json=reportTime,proto3" json:"report_time,omitempty"`
	// Dictionary of file paths to file sizes in bytes of all files in the
	// checkpoint.
	Resources map[string]int64 `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// User defined metadata associated with the checkpoint.
	Metadata *_struct.Struct `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The state of the underlying checkpoint.
	State State `protobuf:"varint,7,opt,name=state,proto3,enum=determined.checkpoint.v1.State" json:"state,omitempty"`
	// Training-related data for this checkpoint.
	Training *CheckpointTrainingMetadata `protobuf:"bytes,8,opt,name=training,proto3" json:"training,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP(), []int{1}
}

func (x *Checkpoint) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Checkpoint) GetAllocationId() string {
	if x != nil && x.AllocationId != nil {
		return *x.AllocationId
	}
	return ""
}

func (x *Checkpoint) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Checkpoint) GetReportTime() *timestamp.Timestamp {
	if x != nil {
		return x.ReportTime
	}
	return nil
}

func (x *Checkpoint) GetResources() map[string]int64 {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Checkpoint) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Checkpoint) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

func (x *Checkpoint) GetTraining() *CheckpointTrainingMetadata {
	if x != nil {
		return x.Training
	}
	return nil
}

// Request to change checkpoint database information.
type PatchCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uuid of the checkpoint.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Dictionary of file paths to file sizes in bytes of all files in the
	// checkpoint. This won't update actual checkpoint files.
	// If len(resources) == 0 => the checkpoint is considered deleted
	// Otherwise if resources are updated the checkpoint is considered partially
	// deleted.
	Resources *PatchCheckpoint_OptionalResources `protobuf:"bytes,2,opt,name=resources,proto3,oneof" json:"resources,omitempty"`
}

func (x *PatchCheckpoint) Reset() {
	*x = PatchCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchCheckpoint) ProtoMessage() {}

func (x *PatchCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchCheckpoint.ProtoReflect.Descriptor instead.
func (*PatchCheckpoint) Descriptor() ([]byte, []int) {
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP(), []int{2}
}

func (x *PatchCheckpoint) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PatchCheckpoint) GetResources() *PatchCheckpoint_OptionalResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

// Gets around not being able to do "Optional map<string, int64>".
// Not ideal but this API is marked internal for now.
type PatchCheckpoint_OptionalResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resources.
	Resources map[string]int64 `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PatchCheckpoint_OptionalResources) Reset() {
	*x = PatchCheckpoint_OptionalResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchCheckpoint_OptionalResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchCheckpoint_OptionalResources) ProtoMessage() {}

func (x *PatchCheckpoint_OptionalResources) ProtoReflect() protoreflect.Message {
	mi := &file_determined_checkpoint_v1_checkpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchCheckpoint_OptionalResources.ProtoReflect.Descriptor instead.
func (*PatchCheckpoint_OptionalResources) Descriptor() ([]byte, []int) {
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PatchCheckpoint_OptionalResources) GetResources() map[string]int64 {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_determined_checkpoint_v1_checkpoint_proto protoreflect.FileDescriptor

var file_determined_checkpoint_v1_checkpoint_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x03, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0d,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x44,
	0x0a, 0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x10, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x31, 0x0a, 0x07, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x48, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x4c, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x11, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x45, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x3a, 0x05, 0x92, 0x41, 0x02, 0x0a, 0x00, 0x22, 0xb9, 0x04,
	0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x51, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x50, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a,
	0x36, 0x92, 0x41, 0x33, 0x0a, 0x31, 0xd2, 0x01, 0x04, 0x75, 0x75, 0x69, 0x64, 0xd2, 0x01, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xd2, 0x01, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xd2, 0x01, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0xd2,
	0x01, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0xdf, 0x02, 0x0a, 0x0f, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x5e, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65,
	0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x1a, 0xbb, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a,
	0x0c, 0x92, 0x41, 0x09, 0x0a, 0x07, 0xd2, 0x01, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2a, 0x86, 0x01, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x05, 0x2a, 0xa9, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x42, 0x59, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x42, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10, 0x06,
	0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_checkpoint_v1_checkpoint_proto_rawDescOnce sync.Once
	file_determined_checkpoint_v1_checkpoint_proto_rawDescData = file_determined_checkpoint_v1_checkpoint_proto_rawDesc
)

func file_determined_checkpoint_v1_checkpoint_proto_rawDescGZIP() []byte {
	file_determined_checkpoint_v1_checkpoint_proto_rawDescOnce.Do(func() {
		file_determined_checkpoint_v1_checkpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_checkpoint_v1_checkpoint_proto_rawDescData)
	})
	return file_determined_checkpoint_v1_checkpoint_proto_rawDescData
}

var file_determined_checkpoint_v1_checkpoint_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_determined_checkpoint_v1_checkpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_determined_checkpoint_v1_checkpoint_proto_goTypes = []interface{}{
	(State)(0),                         // 0: determined.checkpoint.v1.State
	(SortBy)(0),                        // 1: determined.checkpoint.v1.SortBy
	(*CheckpointTrainingMetadata)(nil), // 2: determined.checkpoint.v1.CheckpointTrainingMetadata
	(*Checkpoint)(nil),                 // 3: determined.checkpoint.v1.Checkpoint
	(*PatchCheckpoint)(nil),            // 4: determined.checkpoint.v1.PatchCheckpoint
	nil,                                // 5: determined.checkpoint.v1.Checkpoint.ResourcesEntry
	(*PatchCheckpoint_OptionalResources)(nil), // 6: determined.checkpoint.v1.PatchCheckpoint.OptionalResources
	nil,                          // 7: determined.checkpoint.v1.PatchCheckpoint.OptionalResources.ResourcesEntry
	(*wrappers.Int32Value)(nil),  // 8: google.protobuf.Int32Value
	(*_struct.Struct)(nil),       // 9: google.protobuf.Struct
	(*commonv1.Metrics)(nil),     // 10: determined.common.v1.Metrics
	(*wrappers.DoubleValue)(nil), // 11: google.protobuf.DoubleValue
	(*timestamp.Timestamp)(nil),  // 12: google.protobuf.Timestamp
}
var file_determined_checkpoint_v1_checkpoint_proto_depIdxs = []int32{
	8,  // 0: determined.checkpoint.v1.CheckpointTrainingMetadata.trial_id:type_name -> google.protobuf.Int32Value
	8,  // 1: determined.checkpoint.v1.CheckpointTrainingMetadata.experiment_id:type_name -> google.protobuf.Int32Value
	9,  // 2: determined.checkpoint.v1.CheckpointTrainingMetadata.experiment_config:type_name -> google.protobuf.Struct
	9,  // 3: determined.checkpoint.v1.CheckpointTrainingMetadata.hparams:type_name -> google.protobuf.Struct
	10, // 4: determined.checkpoint.v1.CheckpointTrainingMetadata.training_metrics:type_name -> determined.common.v1.Metrics
	10, // 5: determined.checkpoint.v1.CheckpointTrainingMetadata.validation_metrics:type_name -> determined.common.v1.Metrics
	11, // 6: determined.checkpoint.v1.CheckpointTrainingMetadata.searcher_metric:type_name -> google.protobuf.DoubleValue
	12, // 7: determined.checkpoint.v1.Checkpoint.report_time:type_name -> google.protobuf.Timestamp
	5,  // 8: determined.checkpoint.v1.Checkpoint.resources:type_name -> determined.checkpoint.v1.Checkpoint.ResourcesEntry
	9,  // 9: determined.checkpoint.v1.Checkpoint.metadata:type_name -> google.protobuf.Struct
	0,  // 10: determined.checkpoint.v1.Checkpoint.state:type_name -> determined.checkpoint.v1.State
	2,  // 11: determined.checkpoint.v1.Checkpoint.training:type_name -> determined.checkpoint.v1.CheckpointTrainingMetadata
	6,  // 12: determined.checkpoint.v1.PatchCheckpoint.resources:type_name -> determined.checkpoint.v1.PatchCheckpoint.OptionalResources
	7,  // 13: determined.checkpoint.v1.PatchCheckpoint.OptionalResources.resources:type_name -> determined.checkpoint.v1.PatchCheckpoint.OptionalResources.ResourcesEntry
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_determined_checkpoint_v1_checkpoint_proto_init() }
func file_determined_checkpoint_v1_checkpoint_proto_init() {
	if File_determined_checkpoint_v1_checkpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_checkpoint_v1_checkpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointTrainingMetadata); i {
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
		file_determined_checkpoint_v1_checkpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
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
		file_determined_checkpoint_v1_checkpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchCheckpoint); i {
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
		file_determined_checkpoint_v1_checkpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchCheckpoint_OptionalResources); i {
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
	file_determined_checkpoint_v1_checkpoint_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_determined_checkpoint_v1_checkpoint_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_determined_checkpoint_v1_checkpoint_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_checkpoint_v1_checkpoint_proto_goTypes,
		DependencyIndexes: file_determined_checkpoint_v1_checkpoint_proto_depIdxs,
		EnumInfos:         file_determined_checkpoint_v1_checkpoint_proto_enumTypes,
		MessageInfos:      file_determined_checkpoint_v1_checkpoint_proto_msgTypes,
	}.Build()
	File_determined_checkpoint_v1_checkpoint_proto = out.File
	file_determined_checkpoint_v1_checkpoint_proto_rawDesc = nil
	file_determined_checkpoint_v1_checkpoint_proto_goTypes = nil
	file_determined_checkpoint_v1_checkpoint_proto_depIdxs = nil
}
