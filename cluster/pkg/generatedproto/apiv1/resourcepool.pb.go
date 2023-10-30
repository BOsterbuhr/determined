// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/api/v1/resourcepool.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	resourcepoolv1 "github.com/determined-ai/determined/cluster/pkg/generatedproto/resourcepoolv1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Get the list of resource pools from the cluster.
type GetResourcePoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Skip the number of resource pools before returning results. Negative values
	// denote number of resource pools to skip from the end before returning
	// results.
	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of resource pools. A value of 0 denotes no limit.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// Indicate whether or not to return unbound pools only.
	Unbound bool `protobuf:"varint,3,opt,name=unbound,proto3" json:"unbound,omitempty"`
}

func (x *GetResourcePoolsRequest) Reset() {
	*x = GetResourcePoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePoolsRequest) ProtoMessage() {}

func (x *GetResourcePoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePoolsRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePoolsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{0}
}

func (x *GetResourcePoolsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetResourcePoolsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetResourcePoolsRequest) GetUnbound() bool {
	if x != nil {
		return x.Unbound
	}
	return false
}

// Response to GetResourcePoolsRequest.
type GetResourcePoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of returned resource pools.
	ResourcePools []*resourcepoolv1.ResourcePool `protobuf:"bytes,1,rep,name=resource_pools,json=resourcePools,proto3" json:"resource_pools,omitempty"`
	// Pagination information of the full dataset.
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetResourcePoolsResponse) Reset() {
	*x = GetResourcePoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePoolsResponse) ProtoMessage() {}

func (x *GetResourcePoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePoolsResponse.ProtoReflect.Descriptor instead.
func (*GetResourcePoolsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourcePoolsResponse) GetResourcePools() []*resourcepoolv1.ResourcePool {
	if x != nil {
		return x.ResourcePools
	}
	return nil
}

func (x *GetResourcePoolsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// Bind a resource pool to workspaces
type BindRPToWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource pool name.
	ResourcePoolName string `protobuf:"bytes,1,opt,name=resource_pool_name,json=resourcePoolName,proto3" json:"resource_pool_name,omitempty"`
	// The workspace IDs to be bound to the resource pool.
	WorkspaceIds []int32 `protobuf:"varint,2,rep,packed,name=workspace_ids,json=workspaceIds,proto3" json:"workspace_ids,omitempty"`
	// The workspace names to be bound to the resource pool.
	WorkspaceNames []string `protobuf:"bytes,3,rep,name=workspace_names,json=workspaceNames,proto3" json:"workspace_names,omitempty"`
}

func (x *BindRPToWorkspaceRequest) Reset() {
	*x = BindRPToWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRPToWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRPToWorkspaceRequest) ProtoMessage() {}

func (x *BindRPToWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRPToWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*BindRPToWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{2}
}

func (x *BindRPToWorkspaceRequest) GetResourcePoolName() string {
	if x != nil {
		return x.ResourcePoolName
	}
	return ""
}

func (x *BindRPToWorkspaceRequest) GetWorkspaceIds() []int32 {
	if x != nil {
		return x.WorkspaceIds
	}
	return nil
}

func (x *BindRPToWorkspaceRequest) GetWorkspaceNames() []string {
	if x != nil {
		return x.WorkspaceNames
	}
	return nil
}

// Bind a resource pool to workspaces response.
type BindRPToWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindRPToWorkspaceResponse) Reset() {
	*x = BindRPToWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRPToWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRPToWorkspaceResponse) ProtoMessage() {}

func (x *BindRPToWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRPToWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*BindRPToWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{3}
}

// Unbind a resource pool to workspaces.
type UnbindRPFromWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource pool name.
	ResourcePoolName string `protobuf:"bytes,1,opt,name=resource_pool_name,json=resourcePoolName,proto3" json:"resource_pool_name,omitempty"`
	// The workspace IDs to be unbound.
	WorkspaceIds []int32 `protobuf:"varint,2,rep,packed,name=workspace_ids,json=workspaceIds,proto3" json:"workspace_ids,omitempty"`
	// The workspace names to be unbound.
	WorkspaceNames []string `protobuf:"bytes,3,rep,name=workspace_names,json=workspaceNames,proto3" json:"workspace_names,omitempty"`
}

func (x *UnbindRPFromWorkspaceRequest) Reset() {
	*x = UnbindRPFromWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindRPFromWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRPFromWorkspaceRequest) ProtoMessage() {}

func (x *UnbindRPFromWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRPFromWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*UnbindRPFromWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{4}
}

func (x *UnbindRPFromWorkspaceRequest) GetResourcePoolName() string {
	if x != nil {
		return x.ResourcePoolName
	}
	return ""
}

func (x *UnbindRPFromWorkspaceRequest) GetWorkspaceIds() []int32 {
	if x != nil {
		return x.WorkspaceIds
	}
	return nil
}

func (x *UnbindRPFromWorkspaceRequest) GetWorkspaceNames() []string {
	if x != nil {
		return x.WorkspaceNames
	}
	return nil
}

// Unbind a resource pool to workspaces response.
type UnbindRPFromWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnbindRPFromWorkspaceResponse) Reset() {
	*x = UnbindRPFromWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindRPFromWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRPFromWorkspaceResponse) ProtoMessage() {}

func (x *UnbindRPFromWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRPFromWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*UnbindRPFromWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{5}
}

// Overwrite and replace the workspaces bound to an RP request.
type OverwriteRPWorkspaceBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource pool name.
	ResourcePoolName string `protobuf:"bytes,1,opt,name=resource_pool_name,json=resourcePoolName,proto3" json:"resource_pool_name,omitempty"`
	// The new workspace IDs to bind to the resource_pool.
	WorkspaceIds []int32 `protobuf:"varint,2,rep,packed,name=workspace_ids,json=workspaceIds,proto3" json:"workspace_ids,omitempty"`
	// The new workspace names to bind to the resource_pool.
	WorkspaceNames []string `protobuf:"bytes,3,rep,name=workspace_names,json=workspaceNames,proto3" json:"workspace_names,omitempty"`
}

func (x *OverwriteRPWorkspaceBindingsRequest) Reset() {
	*x = OverwriteRPWorkspaceBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverwriteRPWorkspaceBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverwriteRPWorkspaceBindingsRequest) ProtoMessage() {}

func (x *OverwriteRPWorkspaceBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverwriteRPWorkspaceBindingsRequest.ProtoReflect.Descriptor instead.
func (*OverwriteRPWorkspaceBindingsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{6}
}

func (x *OverwriteRPWorkspaceBindingsRequest) GetResourcePoolName() string {
	if x != nil {
		return x.ResourcePoolName
	}
	return ""
}

func (x *OverwriteRPWorkspaceBindingsRequest) GetWorkspaceIds() []int32 {
	if x != nil {
		return x.WorkspaceIds
	}
	return nil
}

func (x *OverwriteRPWorkspaceBindingsRequest) GetWorkspaceNames() []string {
	if x != nil {
		return x.WorkspaceNames
	}
	return nil
}

// Overwrite and replace the workspaces bound to an RP response.
type OverwriteRPWorkspaceBindingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OverwriteRPWorkspaceBindingsResponse) Reset() {
	*x = OverwriteRPWorkspaceBindingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverwriteRPWorkspaceBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverwriteRPWorkspaceBindingsResponse) ProtoMessage() {}

func (x *OverwriteRPWorkspaceBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverwriteRPWorkspaceBindingsResponse.ProtoReflect.Descriptor instead.
func (*OverwriteRPWorkspaceBindingsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{7}
}

// List all the workspaces bound to the RP.
type ListWorkspacesBoundToRPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource pool name.
	ResourcePoolName string `protobuf:"bytes,1,opt,name=resource_pool_name,json=resourcePoolName,proto3" json:"resource_pool_name,omitempty"`
	// The offset to use with pagination
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// The maximum number of results to return
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListWorkspacesBoundToRPRequest) Reset() {
	*x = ListWorkspacesBoundToRPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspacesBoundToRPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspacesBoundToRPRequest) ProtoMessage() {}

func (x *ListWorkspacesBoundToRPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspacesBoundToRPRequest.ProtoReflect.Descriptor instead.
func (*ListWorkspacesBoundToRPRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{8}
}

func (x *ListWorkspacesBoundToRPRequest) GetResourcePoolName() string {
	if x != nil {
		return x.ResourcePoolName
	}
	return ""
}

func (x *ListWorkspacesBoundToRPRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListWorkspacesBoundToRPRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Response to ListWorkspacesBoundToRPRequest.
type ListWorkspacesBoundToRPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of workspace IDs.
	WorkspaceIds []int32 `protobuf:"varint,1,rep,packed,name=workspace_ids,json=workspaceIds,proto3" json:"workspace_ids,omitempty"`
	// Pagination information of the full dataset.
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListWorkspacesBoundToRPResponse) Reset() {
	*x = ListWorkspacesBoundToRPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspacesBoundToRPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspacesBoundToRPResponse) ProtoMessage() {}

func (x *ListWorkspacesBoundToRPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspacesBoundToRPResponse.ProtoReflect.Descriptor instead.
func (*ListWorkspacesBoundToRPResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{9}
}

func (x *ListWorkspacesBoundToRPResponse) GetWorkspaceIds() []int32 {
	if x != nil {
		return x.WorkspaceIds
	}
	return nil
}

func (x *ListWorkspacesBoundToRPResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_determined_api_v1_resourcepool_proto protoreflect.FileDescriptor

var file_determined_api_v1_resourcepool_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0xaa, 0x01,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb2, 0x01, 0x0a, 0x18, 0x42,
	0x69, 0x6e, 0x64, 0x52, 0x50, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a, 0x15, 0xd2, 0x01, 0x12, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x1b, 0x0a, 0x19, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x50, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb6, 0x01, 0x0a,
	0x1c, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a, 0x15,
	0xd2, 0x01, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52,
	0x50, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x23, 0x4f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x50, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a,
	0x15, 0xd2, 0x01, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x24, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x50, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa0,
	0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x6f, 0x52, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x22, 0x92,
	0x41, 0x1f, 0x0a, 0x1d, 0xd2, 0x01, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0xa3, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x6f, 0x52, 0x50, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x1c, 0x92, 0x41, 0x19, 0x0a, 0x17,
	0xd2, 0x01, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0xd2,
	0x01, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_api_v1_resourcepool_proto_rawDescOnce sync.Once
	file_determined_api_v1_resourcepool_proto_rawDescData = file_determined_api_v1_resourcepool_proto_rawDesc
)

func file_determined_api_v1_resourcepool_proto_rawDescGZIP() []byte {
	file_determined_api_v1_resourcepool_proto_rawDescOnce.Do(func() {
		file_determined_api_v1_resourcepool_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_api_v1_resourcepool_proto_rawDescData)
	})
	return file_determined_api_v1_resourcepool_proto_rawDescData
}

var file_determined_api_v1_resourcepool_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_determined_api_v1_resourcepool_proto_goTypes = []interface{}{
	(*GetResourcePoolsRequest)(nil),              // 0: determined.api.v1.GetResourcePoolsRequest
	(*GetResourcePoolsResponse)(nil),             // 1: determined.api.v1.GetResourcePoolsResponse
	(*BindRPToWorkspaceRequest)(nil),             // 2: determined.api.v1.BindRPToWorkspaceRequest
	(*BindRPToWorkspaceResponse)(nil),            // 3: determined.api.v1.BindRPToWorkspaceResponse
	(*UnbindRPFromWorkspaceRequest)(nil),         // 4: determined.api.v1.UnbindRPFromWorkspaceRequest
	(*UnbindRPFromWorkspaceResponse)(nil),        // 5: determined.api.v1.UnbindRPFromWorkspaceResponse
	(*OverwriteRPWorkspaceBindingsRequest)(nil),  // 6: determined.api.v1.OverwriteRPWorkspaceBindingsRequest
	(*OverwriteRPWorkspaceBindingsResponse)(nil), // 7: determined.api.v1.OverwriteRPWorkspaceBindingsResponse
	(*ListWorkspacesBoundToRPRequest)(nil),       // 8: determined.api.v1.ListWorkspacesBoundToRPRequest
	(*ListWorkspacesBoundToRPResponse)(nil),      // 9: determined.api.v1.ListWorkspacesBoundToRPResponse
	(*resourcepoolv1.ResourcePool)(nil),          // 10: determined.resourcepool.v1.ResourcePool
	(*Pagination)(nil),                           // 11: determined.api.v1.Pagination
}
var file_determined_api_v1_resourcepool_proto_depIdxs = []int32{
	10, // 0: determined.api.v1.GetResourcePoolsResponse.resource_pools:type_name -> determined.resourcepool.v1.ResourcePool
	11, // 1: determined.api.v1.GetResourcePoolsResponse.pagination:type_name -> determined.api.v1.Pagination
	11, // 2: determined.api.v1.ListWorkspacesBoundToRPResponse.pagination:type_name -> determined.api.v1.Pagination
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_determined_api_v1_resourcepool_proto_init() }
func file_determined_api_v1_resourcepool_proto_init() {
	if File_determined_api_v1_resourcepool_proto != nil {
		return
	}
	file_determined_api_v1_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_determined_api_v1_resourcepool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcePoolsRequest); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcePoolsResponse); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindRPToWorkspaceRequest); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindRPToWorkspaceResponse); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindRPFromWorkspaceRequest); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindRPFromWorkspaceResponse); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverwriteRPWorkspaceBindingsRequest); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverwriteRPWorkspaceBindingsResponse); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspacesBoundToRPRequest); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspacesBoundToRPResponse); i {
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
			RawDescriptor: file_determined_api_v1_resourcepool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_api_v1_resourcepool_proto_goTypes,
		DependencyIndexes: file_determined_api_v1_resourcepool_proto_depIdxs,
		MessageInfos:      file_determined_api_v1_resourcepool_proto_msgTypes,
	}.Build()
	File_determined_api_v1_resourcepool_proto = out.File
	file_determined_api_v1_resourcepool_proto_rawDesc = nil
	file_determined_api_v1_resourcepool_proto_goTypes = nil
	file_determined_api_v1_resourcepool_proto_depIdxs = nil
}
