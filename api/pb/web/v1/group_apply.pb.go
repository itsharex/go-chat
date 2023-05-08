// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: web/v1/group_apply.proto

package web

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type GroupApplyCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int32  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" binding:"required"`
	Remark  string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty" binding:"required"`
}

func (x *GroupApplyCreateRequest) Reset() {
	*x = GroupApplyCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyCreateRequest) ProtoMessage() {}

func (x *GroupApplyCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyCreateRequest.ProtoReflect.Descriptor instead.
func (*GroupApplyCreateRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{0}
}

func (x *GroupApplyCreateRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupApplyCreateRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type GroupApplyCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupApplyCreateResponse) Reset() {
	*x = GroupApplyCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyCreateResponse) ProtoMessage() {}

func (x *GroupApplyCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyCreateResponse.ProtoReflect.Descriptor instead.
func (*GroupApplyCreateResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{1}
}

type GroupApplyDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyId int32 `protobuf:"varint,1,opt,name=apply_id,json=applyId,proto3" json:"apply_id,omitempty" binding:"required"`
}

func (x *GroupApplyDeleteRequest) Reset() {
	*x = GroupApplyDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyDeleteRequest) ProtoMessage() {}

func (x *GroupApplyDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyDeleteRequest.ProtoReflect.Descriptor instead.
func (*GroupApplyDeleteRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{2}
}

func (x *GroupApplyDeleteRequest) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

type GroupApplyDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupApplyDeleteResponse) Reset() {
	*x = GroupApplyDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyDeleteResponse) ProtoMessage() {}

func (x *GroupApplyDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyDeleteResponse.ProtoReflect.Descriptor instead.
func (*GroupApplyDeleteResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{3}
}

type GroupApplyAgreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyId int32 `protobuf:"varint,1,opt,name=apply_id,json=applyId,proto3" json:"apply_id,omitempty" binding:"required"`
}

func (x *GroupApplyAgreeRequest) Reset() {
	*x = GroupApplyAgreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyAgreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyAgreeRequest) ProtoMessage() {}

func (x *GroupApplyAgreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyAgreeRequest.ProtoReflect.Descriptor instead.
func (*GroupApplyAgreeRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{4}
}

func (x *GroupApplyAgreeRequest) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

type GroupApplyAgreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupApplyAgreeResponse) Reset() {
	*x = GroupApplyAgreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyAgreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyAgreeResponse) ProtoMessage() {}

func (x *GroupApplyAgreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyAgreeResponse.ProtoReflect.Descriptor instead.
func (*GroupApplyAgreeResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{5}
}

type GroupApplyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" binding:"required"`
}

func (x *GroupApplyListRequest) Reset() {
	*x = GroupApplyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyListRequest) ProtoMessage() {}

func (x *GroupApplyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyListRequest.ProtoReflect.Descriptor instead.
func (*GroupApplyListRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{6}
}

func (x *GroupApplyListRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GroupApplyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GroupApplyListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GroupApplyListResponse) Reset() {
	*x = GroupApplyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyListResponse) ProtoMessage() {}

func (x *GroupApplyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyListResponse.ProtoReflect.Descriptor instead.
func (*GroupApplyListResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{7}
}

func (x *GroupApplyListResponse) GetItems() []*GroupApplyListResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GroupApplyListResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId   int32  `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Remark    string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Avatar    string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Nickname  string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GroupApplyListResponse_Item) Reset() {
	*x = GroupApplyListResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_group_apply_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupApplyListResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupApplyListResponse_Item) ProtoMessage() {}

func (x *GroupApplyListResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_group_apply_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupApplyListResponse_Item.ProtoReflect.Descriptor instead.
func (*GroupApplyListResponse_Item) Descriptor() ([]byte, []int) {
	return file_web_v1_group_apply_proto_rawDescGZIP(), []int{7, 0}
}

func (x *GroupApplyListResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupApplyListResponse_Item) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GroupApplyListResponse_Item) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupApplyListResponse_Item) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *GroupApplyListResponse_Item) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GroupApplyListResponse_Item) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GroupApplyListResponse_Item) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_web_v1_group_apply_proto protoreflect.FileDescriptor

var file_web_v1_group_apply_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x77, 0x65, 0x62, 0x1a,
	0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x17, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4d, 0x0a, 0x17, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x22,
	0x1a, 0x0a, 0x18, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x16, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x15, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x88, 0x02, 0x0a, 0x16, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0xb5, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a,
	0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_web_v1_group_apply_proto_rawDescOnce sync.Once
	file_web_v1_group_apply_proto_rawDescData = file_web_v1_group_apply_proto_rawDesc
)

func file_web_v1_group_apply_proto_rawDescGZIP() []byte {
	file_web_v1_group_apply_proto_rawDescOnce.Do(func() {
		file_web_v1_group_apply_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_v1_group_apply_proto_rawDescData)
	})
	return file_web_v1_group_apply_proto_rawDescData
}

var file_web_v1_group_apply_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_web_v1_group_apply_proto_goTypes = []interface{}{
	(*GroupApplyCreateRequest)(nil),     // 0: web.GroupApplyCreateRequest
	(*GroupApplyCreateResponse)(nil),    // 1: web.GroupApplyCreateResponse
	(*GroupApplyDeleteRequest)(nil),     // 2: web.GroupApplyDeleteRequest
	(*GroupApplyDeleteResponse)(nil),    // 3: web.GroupApplyDeleteResponse
	(*GroupApplyAgreeRequest)(nil),      // 4: web.GroupApplyAgreeRequest
	(*GroupApplyAgreeResponse)(nil),     // 5: web.GroupApplyAgreeResponse
	(*GroupApplyListRequest)(nil),       // 6: web.GroupApplyListRequest
	(*GroupApplyListResponse)(nil),      // 7: web.GroupApplyListResponse
	(*GroupApplyListResponse_Item)(nil), // 8: web.GroupApplyListResponse.Item
}
var file_web_v1_group_apply_proto_depIdxs = []int32{
	8, // 0: web.GroupApplyListResponse.items:type_name -> web.GroupApplyListResponse.Item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_v1_group_apply_proto_init() }
func file_web_v1_group_apply_proto_init() {
	if File_web_v1_group_apply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_v1_group_apply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyCreateRequest); i {
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
		file_web_v1_group_apply_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyCreateResponse); i {
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
		file_web_v1_group_apply_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyDeleteRequest); i {
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
		file_web_v1_group_apply_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyDeleteResponse); i {
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
		file_web_v1_group_apply_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyAgreeRequest); i {
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
		file_web_v1_group_apply_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyAgreeResponse); i {
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
		file_web_v1_group_apply_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyListRequest); i {
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
		file_web_v1_group_apply_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyListResponse); i {
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
		file_web_v1_group_apply_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupApplyListResponse_Item); i {
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
			RawDescriptor: file_web_v1_group_apply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_v1_group_apply_proto_goTypes,
		DependencyIndexes: file_web_v1_group_apply_proto_depIdxs,
		MessageInfos:      file_web_v1_group_apply_proto_msgTypes,
	}.Build()
	File_web_v1_group_apply_proto = out.File
	file_web_v1_group_apply_proto_rawDesc = nil
	file_web_v1_group_apply_proto_goTypes = nil
	file_web_v1_group_apply_proto_depIdxs = nil
}
