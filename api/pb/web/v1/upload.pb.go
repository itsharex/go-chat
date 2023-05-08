// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: web/v1/upload.proto

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

// 头像上传接口请求参数
type UploadAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadAvatarRequest) Reset() {
	*x = UploadAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarRequest) ProtoMessage() {}

func (x *UploadAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarRequest.ProtoReflect.Descriptor instead.
func (*UploadAvatarRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_upload_proto_rawDescGZIP(), []int{0}
}

// 头像上传接口响应参数
type UploadAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar string `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UploadAvatarResponse) Reset() {
	*x = UploadAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarResponse) ProtoMessage() {}

func (x *UploadAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarResponse.ProtoReflect.Descriptor instead.
func (*UploadAvatarResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadAvatarResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

// 批量上传文件初始化接口请求参数
type UploadInitiateMultipartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" binding:"required"`
	FileSize int64  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty" binding:"required"`
}

func (x *UploadInitiateMultipartRequest) Reset() {
	*x = UploadInitiateMultipartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInitiateMultipartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInitiateMultipartRequest) ProtoMessage() {}

func (x *UploadInitiateMultipartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInitiateMultipartRequest.ProtoReflect.Descriptor instead.
func (*UploadInitiateMultipartRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UploadInitiateMultipartRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadInitiateMultipartRequest) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

// 批量上传文件初始化接口响应参数
type UploadInitiateMultipartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId  string `protobuf:"bytes,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	SplitSize int32  `protobuf:"varint,2,opt,name=split_size,json=splitSize,proto3" json:"split_size,omitempty"`
}

func (x *UploadInitiateMultipartResponse) Reset() {
	*x = UploadInitiateMultipartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInitiateMultipartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInitiateMultipartResponse) ProtoMessage() {}

func (x *UploadInitiateMultipartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInitiateMultipartResponse.ProtoReflect.Descriptor instead.
func (*UploadInitiateMultipartResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UploadInitiateMultipartResponse) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *UploadInitiateMultipartResponse) GetSplitSize() int32 {
	if x != nil {
		return x.SplitSize
	}
	return 0
}

// 批量上传文件接口请求参数
type UploadMultipartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId   string `protobuf:"bytes,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty" form:"upload_id" binding:"required"`
	SplitIndex int32  `protobuf:"varint,2,opt,name=split_index,json=splitIndex,proto3" json:"split_index,omitempty" form:"split_index" binding:"min=0"`
	SplitNum   int32  `protobuf:"varint,3,opt,name=split_num,json=splitNum,proto3" json:"split_num,omitempty" form:"split_num" binding:"required,min=1"`
}

func (x *UploadMultipartRequest) Reset() {
	*x = UploadMultipartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_upload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadMultipartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMultipartRequest) ProtoMessage() {}

func (x *UploadMultipartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_upload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMultipartRequest.ProtoReflect.Descriptor instead.
func (*UploadMultipartRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_upload_proto_rawDescGZIP(), []int{4}
}

func (x *UploadMultipartRequest) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *UploadMultipartRequest) GetSplitIndex() int32 {
	if x != nil {
		return x.SplitIndex
	}
	return 0
}

func (x *UploadMultipartRequest) GetSplitNum() int32 {
	if x != nil {
		return x.SplitNum
	}
	return 0
}

// 批量上传文件接口请求参数
type UploadMultipartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId string `protobuf:"bytes,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	IsMerge  bool   `protobuf:"varint,2,opt,name=is_merge,json=isMerge,proto3" json:"is_merge,omitempty"`
}

func (x *UploadMultipartResponse) Reset() {
	*x = UploadMultipartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_upload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadMultipartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMultipartResponse) ProtoMessage() {}

func (x *UploadMultipartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_upload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMultipartResponse.ProtoReflect.Descriptor instead.
func (*UploadMultipartResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_upload_proto_rawDescGZIP(), []int{5}
}

func (x *UploadMultipartResponse) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *UploadMultipartResponse) GetIsMerge() bool {
	if x != nil {
		return x.IsMerge
	}
	return false
}

var File_web_v1_upload_proto protoreflect.FileDescriptor

var file_web_v1_upload_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x77, 0x65, 0x62, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x15, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84,
	0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5d, 0x0a, 0x1f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x45, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x28, 0x9a, 0x84, 0x9e, 0x03, 0x23, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x27, 0x9a, 0x84, 0x9e,
	0x03, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x6d, 0x69,
	0x6e, 0x3d, 0x30, 0x22, 0x52, 0x0a, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x4b, 0x0a, 0x09, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x2e, 0x9a, 0x84, 0x9e, 0x03, 0x29, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6d, 0x69, 0x6e,
	0x3d, 0x31, 0x22, 0x52, 0x08, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x51, 0x0a,
	0x17, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_v1_upload_proto_rawDescOnce sync.Once
	file_web_v1_upload_proto_rawDescData = file_web_v1_upload_proto_rawDesc
)

func file_web_v1_upload_proto_rawDescGZIP() []byte {
	file_web_v1_upload_proto_rawDescOnce.Do(func() {
		file_web_v1_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_v1_upload_proto_rawDescData)
	})
	return file_web_v1_upload_proto_rawDescData
}

var file_web_v1_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_web_v1_upload_proto_goTypes = []interface{}{
	(*UploadAvatarRequest)(nil),             // 0: web.UploadAvatarRequest
	(*UploadAvatarResponse)(nil),            // 1: web.UploadAvatarResponse
	(*UploadInitiateMultipartRequest)(nil),  // 2: web.UploadInitiateMultipartRequest
	(*UploadInitiateMultipartResponse)(nil), // 3: web.UploadInitiateMultipartResponse
	(*UploadMultipartRequest)(nil),          // 4: web.UploadMultipartRequest
	(*UploadMultipartResponse)(nil),         // 5: web.UploadMultipartResponse
}
var file_web_v1_upload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web_v1_upload_proto_init() }
func file_web_v1_upload_proto_init() {
	if File_web_v1_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_v1_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAvatarRequest); i {
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
		file_web_v1_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAvatarResponse); i {
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
		file_web_v1_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInitiateMultipartRequest); i {
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
		file_web_v1_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInitiateMultipartResponse); i {
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
		file_web_v1_upload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadMultipartRequest); i {
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
		file_web_v1_upload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadMultipartResponse); i {
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
			RawDescriptor: file_web_v1_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_v1_upload_proto_goTypes,
		DependencyIndexes: file_web_v1_upload_proto_depIdxs,
		MessageInfos:      file_web_v1_upload_proto_msgTypes,
	}.Build()
	File_web_v1_upload_proto = out.File
	file_web_v1_upload_proto_rawDesc = nil
	file_web_v1_upload_proto_goTypes = nil
	file_web_v1_upload_proto_depIdxs = nil
}
