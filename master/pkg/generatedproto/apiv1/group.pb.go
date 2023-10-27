// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/api/v1/group.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	groupv1 "github.com/determined-ai/determined/master/pkg/generatedproto/groupv1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetGroupRequest is the body of the request for the call
// to get a group by id.
type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the group to return.
	GroupId int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// GetGroupResponse is the body of the response for the call
// to get a group by id.
type GetGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group info
	Group *groupv1.GroupDetails `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetGroupResponse) Reset() {
	*x = GetGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResponse) ProtoMessage() {}

func (x *GetGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResponse.ProtoReflect.Descriptor instead.
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupResponse) GetGroup() *groupv1.GroupDetails {
	if x != nil {
		return x.Group
	}
	return nil
}

// GetGroupsRequest is the body of the request for the call
// to search for groups.
type GetGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the user to use to find groups to which the user belongs.
	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The group name to use when searching.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Skip the number of groups before returning results. Negative values
	// denote number of groups to skip from the end before returning results.
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of groups. Required and must be must be <= 500.
	Limit int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetGroupsRequest) Reset() {
	*x = GetGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsRequest) ProtoMessage() {}

func (x *GetGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetGroupsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetGroupsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetGroupsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetGroupsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// GetGroupsResponse is the body of the response for the call
// to search for groups.
type GetGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The found groups
	Groups []*groupv1.GroupSearchResult `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	// Pagination information of the full dataset.
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetGroupsResponse) Reset() {
	*x = GetGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsResponse) ProtoMessage() {}

func (x *GetGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetGroupsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupsResponse) GetGroups() []*groupv1.GroupSearchResult {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GetGroupsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// UpdateGroupRequest is the body of the request for the call
// to update a group and its members.
type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the group
	GroupId int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The name of the group
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The user ids of users to add to the group
	AddUsers []int32 `protobuf:"varint,3,rep,packed,name=add_users,json=addUsers,proto3" json:"add_users,omitempty"`
	// The user ids of users to delete from the group
	RemoveUsers []int32 `protobuf:"varint,4,rep,packed,name=remove_users,json=removeUsers,proto3" json:"remove_users,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGroupRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UpdateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGroupRequest) GetAddUsers() []int32 {
	if x != nil {
		return x.AddUsers
	}
	return nil
}

func (x *UpdateGroupRequest) GetRemoveUsers() []int32 {
	if x != nil {
		return x.RemoveUsers
	}
	return nil
}

// CreateGroupResponse is the body of the response for the call
// to update a group and its members.
type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Info about the group after the update succeeded.
	Group *groupv1.GroupDetails `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGroupResponse) GetGroup() *groupv1.GroupDetails {
	if x != nil {
		return x.Group
	}
	return nil
}

// UpdateGroupResponse is the body of the response for the call
// to update a group and its members.
type UpdateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Info about the group after the update succeeded.
	Group *groupv1.GroupDetails `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupResponse) Reset() {
	*x = UpdateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupResponse) ProtoMessage() {}

func (x *UpdateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGroupResponse) GetGroup() *groupv1.GroupDetails {
	if x != nil {
		return x.Group
	}
	return nil
}

// CreateGroupRequest is the body of the request for the call
// to create a group.
type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name the new group should have
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The ids of users that should be added to the new group
	AddUsers []int32 `protobuf:"varint,2,rep,packed,name=add_users,json=addUsers,proto3" json:"add_users,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{7}
}

func (x *CreateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGroupRequest) GetAddUsers() []int32 {
	if x != nil {
		return x.AddUsers
	}
	return nil
}

// DeleteGroupRequest is the body of the request for the call
// to delete a group.
type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the group that should be deleted.
	GroupId int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGroupRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// DeleteGroupResponse is the body of the response for the call
// to delete a group.
type DeleteGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGroupResponse) Reset() {
	*x = DeleteGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResponse) ProtoMessage() {}

func (x *DeleteGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{9}
}

// Add and remove multiple users from multiple groups.
type AssignMultipleGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user ids of users to edit group associations.
	UserIds []int32 `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	// The ids of groups to associate with users.
	AddGroups []int32 `protobuf:"varint,2,rep,packed,name=add_groups,json=addGroups,proto3" json:"add_groups,omitempty"`
	// The ids of groups to disassociate from users.
	RemoveGroups []int32 `protobuf:"varint,3,rep,packed,name=remove_groups,json=removeGroups,proto3" json:"remove_groups,omitempty"`
}

func (x *AssignMultipleGroupsRequest) Reset() {
	*x = AssignMultipleGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignMultipleGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignMultipleGroupsRequest) ProtoMessage() {}

func (x *AssignMultipleGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignMultipleGroupsRequest.ProtoReflect.Descriptor instead.
func (*AssignMultipleGroupsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{10}
}

func (x *AssignMultipleGroupsRequest) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *AssignMultipleGroupsRequest) GetAddGroups() []int32 {
	if x != nil {
		return x.AddGroups
	}
	return nil
}

func (x *AssignMultipleGroupsRequest) GetRemoveGroups() []int32 {
	if x != nil {
		return x.RemoveGroups
	}
	return nil
}

// Response to AssignMultipleGroupsRequest.
type AssignMultipleGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignMultipleGroupsResponse) Reset() {
	*x = AssignMultipleGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_group_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignMultipleGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignMultipleGroupsResponse) ProtoMessage() {}

func (x *AssignMultipleGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_group_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignMultipleGroupsResponse.ProtoReflect.Descriptor instead.
func (*AssignMultipleGroupsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_group_proto_rawDescGZIP(), []int{11}
}

var File_determined_api_v1_group_proto protoreflect.FileDescriptor

var file_determined_api_v1_group_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x1a, 0x22, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x3a, 0x0d, 0x92, 0x41, 0x0a, 0x0a, 0x08, 0xd2, 0x01, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0x7c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x3a, 0x0d, 0x92, 0x41, 0x0a, 0x0a, 0x08, 0xd2, 0x01, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x92, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64,
	0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a,
	0x0b, 0xd2, 0x01, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x0d, 0x92, 0x41,
	0x0a, 0x0a, 0x08, 0xd2, 0x01, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x5d, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x0d, 0x92, 0x41, 0x0a,
	0x0a, 0x08, 0xd2, 0x01, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x53, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x61, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x3a, 0x0c, 0x92, 0x41, 0x09, 0x0a, 0x07, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x41, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2, 0x01, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x1b, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x61, 0x64, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x3a, 0x2d, 0x92, 0x41, 0x2a, 0x0a, 0x28, 0xd2,
	0x01, 0x0a, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0xd2, 0x01, 0x0d, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0xd2, 0x01, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_api_v1_group_proto_rawDescOnce sync.Once
	file_determined_api_v1_group_proto_rawDescData = file_determined_api_v1_group_proto_rawDesc
)

func file_determined_api_v1_group_proto_rawDescGZIP() []byte {
	file_determined_api_v1_group_proto_rawDescOnce.Do(func() {
		file_determined_api_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_api_v1_group_proto_rawDescData)
	})
	return file_determined_api_v1_group_proto_rawDescData
}

var file_determined_api_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_determined_api_v1_group_proto_goTypes = []interface{}{
	(*GetGroupRequest)(nil),              // 0: determined.api.v1.GetGroupRequest
	(*GetGroupResponse)(nil),             // 1: determined.api.v1.GetGroupResponse
	(*GetGroupsRequest)(nil),             // 2: determined.api.v1.GetGroupsRequest
	(*GetGroupsResponse)(nil),            // 3: determined.api.v1.GetGroupsResponse
	(*UpdateGroupRequest)(nil),           // 4: determined.api.v1.UpdateGroupRequest
	(*CreateGroupResponse)(nil),          // 5: determined.api.v1.CreateGroupResponse
	(*UpdateGroupResponse)(nil),          // 6: determined.api.v1.UpdateGroupResponse
	(*CreateGroupRequest)(nil),           // 7: determined.api.v1.CreateGroupRequest
	(*DeleteGroupRequest)(nil),           // 8: determined.api.v1.DeleteGroupRequest
	(*DeleteGroupResponse)(nil),          // 9: determined.api.v1.DeleteGroupResponse
	(*AssignMultipleGroupsRequest)(nil),  // 10: determined.api.v1.AssignMultipleGroupsRequest
	(*AssignMultipleGroupsResponse)(nil), // 11: determined.api.v1.AssignMultipleGroupsResponse
	(*groupv1.GroupDetails)(nil),         // 12: determined.group.v1.GroupDetails
	(*groupv1.GroupSearchResult)(nil),    // 13: determined.group.v1.GroupSearchResult
	(*Pagination)(nil),                   // 14: determined.api.v1.Pagination
}
var file_determined_api_v1_group_proto_depIdxs = []int32{
	12, // 0: determined.api.v1.GetGroupResponse.group:type_name -> determined.group.v1.GroupDetails
	13, // 1: determined.api.v1.GetGroupsResponse.groups:type_name -> determined.group.v1.GroupSearchResult
	14, // 2: determined.api.v1.GetGroupsResponse.pagination:type_name -> determined.api.v1.Pagination
	12, // 3: determined.api.v1.CreateGroupResponse.group:type_name -> determined.group.v1.GroupDetails
	12, // 4: determined.api.v1.UpdateGroupResponse.group:type_name -> determined.group.v1.GroupDetails
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_determined_api_v1_group_proto_init() }
func file_determined_api_v1_group_proto_init() {
	if File_determined_api_v1_group_proto != nil {
		return
	}
	file_determined_api_v1_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_determined_api_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_determined_api_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResponse); i {
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
		file_determined_api_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsRequest); i {
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
		file_determined_api_v1_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsResponse); i {
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
		file_determined_api_v1_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_determined_api_v1_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_determined_api_v1_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupResponse); i {
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
		file_determined_api_v1_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_determined_api_v1_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_determined_api_v1_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupResponse); i {
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
		file_determined_api_v1_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignMultipleGroupsRequest); i {
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
		file_determined_api_v1_group_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignMultipleGroupsResponse); i {
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
			RawDescriptor: file_determined_api_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_api_v1_group_proto_goTypes,
		DependencyIndexes: file_determined_api_v1_group_proto_depIdxs,
		MessageInfos:      file_determined_api_v1_group_proto_msgTypes,
	}.Build()
	File_determined_api_v1_group_proto = out.File
	file_determined_api_v1_group_proto_rawDesc = nil
	file_determined_api_v1_group_proto_goTypes = nil
	file_determined_api_v1_group_proto_depIdxs = nil
}
