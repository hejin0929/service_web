// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/auth/v1/auth.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUsersRequest) Reset() {
	*x = LoginUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUsersRequest) ProtoMessage() {}

func (x *LoginUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUsersRequest.ProtoReflect.Descriptor instead.
func (*LoginUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginUsersRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginUsersRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool       `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 返回是否成功
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // 返回信息
	Data    *LoginData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginUsersReply) Reset() {
	*x = LoginUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUsersReply) ProtoMessage() {}

func (x *LoginUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUsersReply.ProtoReflect.Descriptor instead.
func (*LoginUsersReply) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginUsersReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginUsersReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginUsersReply) GetData() *LoginData {
	if x != nil {
		return x.Data
	}
	return nil
}

type LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Uid          string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *LoginData) Reset() {
	*x = LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginData) ProtoMessage() {}

func (x *LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginData.ProtoReflect.Descriptor instead.
func (*LoginData) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginData) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginData) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ExitUsersLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ExitUsersLoginRequest) Reset() {
	*x = ExitUsersLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitUsersLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitUsersLoginRequest) ProtoMessage() {}

func (x *ExitUsersLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitUsersLoginRequest.ProtoReflect.Descriptor instead.
func (*ExitUsersLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ExitUsersLoginRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ExitUsersLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 返回是否成功
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // 返回信息
}

func (x *ExitUsersLoginReply) Reset() {
	*x = ExitUsersLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitUsersLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitUsersLoginReply) ProtoMessage() {}

func (x *ExitUsersLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitUsersLoginReply.ProtoReflect.Descriptor instead.
func (*ExitUsersLoginReply) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ExitUsersLoginReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ExitUsersLoginReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PatchUsersLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *PatchUsersLoginRequest) Reset() {
	*x = PatchUsersLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUsersLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUsersLoginRequest) ProtoMessage() {}

func (x *PatchUsersLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUsersLoginRequest.ProtoReflect.Descriptor instead.
func (*PatchUsersLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *PatchUsersLoginRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type PatchUsersLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 返回是否成功
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // 返回信息
}

func (x *PatchUsersLoginReply) Reset() {
	*x = PatchUsersLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUsersLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUsersLoginReply) ProtoMessage() {}

func (x *PatchUsersLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUsersLoginReply.ProtoReflect.Descriptor instead.
func (*PatchUsersLoginReply) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *PatchUsersLoginReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PatchUsersLoginReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PatchPasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 返回是否成功
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // 返回信息
}

func (x *PatchPasswordReply) Reset() {
	*x = PatchPasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchPasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchPasswordReply) ProtoMessage() {}

func (x *PatchPasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchPasswordReply.ProtoReflect.Descriptor instead.
func (*PatchPasswordReply) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *PatchPasswordReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PatchPasswordReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AuthLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthLoginRequest) Reset() {
	*x = AuthLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginRequest) ProtoMessage() {}

func (x *AuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *AuthLoginRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthLoginReply) Reset() {
	*x = AuthLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginReply) ProtoMessage() {}

func (x *AuthLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginReply.ProtoReflect.Descriptor instead.
func (*AuthLoginReply) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *AuthLoginReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthLoginReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PatchUsersLoginReplyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PatchUsersLoginReplyData) Reset() {
	*x = PatchUsersLoginReplyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUsersLoginReplyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUsersLoginReplyData) ProtoMessage() {}

func (x *PatchUsersLoginReplyData) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUsersLoginReplyData.ProtoReflect.Descriptor instead.
func (*PatchUsersLoginReplyData) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{6, 0}
}

func (x *PatchUsersLoginReplyData) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PatchUsersLoginReplyData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_api_auth_v1_auth_proto protoreflect.FileDescriptor

var file_api_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x11,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x71, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x29, 0x0a,
	0x15, 0x45, 0x78, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x45, 0x78, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x16, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x7a, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48,
	0x0a, 0x12, 0x50, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x44, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9a, 0x03, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x64, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6d, 0x0a, 0x0e, 0x45, 0x78, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x73, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x32, 0x0d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x47, 0x0a, 0x09,
	0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x20, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_auth_v1_auth_proto_rawDescOnce sync.Once
	file_api_auth_v1_auth_proto_rawDescData = file_api_auth_v1_auth_proto_rawDesc
)

func file_api_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_api_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_api_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_v1_auth_proto_rawDescData)
	})
	return file_api_auth_v1_auth_proto_rawDescData
}

var file_api_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_auth_v1_auth_proto_goTypes = []interface{}{
	(*LoginUsersRequest)(nil),        // 0: api.auth.v1.LoginUsersRequest
	(*LoginUsersReply)(nil),          // 1: api.auth.v1.LoginUsersReply
	(*LoginData)(nil),                // 2: api.auth.v1.LoginData
	(*ExitUsersLoginRequest)(nil),    // 3: api.auth.v1.ExitUsersLoginRequest
	(*ExitUsersLoginReply)(nil),      // 4: api.auth.v1.ExitUsersLoginReply
	(*PatchUsersLoginRequest)(nil),   // 5: api.auth.v1.PatchUsersLoginRequest
	(*PatchUsersLoginReply)(nil),     // 6: api.auth.v1.PatchUsersLoginReply
	(*PatchPasswordReply)(nil),       // 7: api.auth.v1.PatchPasswordReply
	(*AuthLoginRequest)(nil),         // 8: api.auth.v1.AuthLoginRequest
	(*AuthLoginReply)(nil),           // 9: api.auth.v1.AuthLoginReply
	(*PatchUsersLoginReplyData)(nil), // 10: api.auth.v1.PatchUsersLoginReply.data
}
var file_api_auth_v1_auth_proto_depIdxs = []int32{
	2, // 0: api.auth.v1.LoginUsersReply.data:type_name -> api.auth.v1.LoginData
	0, // 1: api.auth.v1.Users.LoginUsers:input_type -> api.auth.v1.LoginUsersRequest
	3, // 2: api.auth.v1.Users.ExitUsersLogin:input_type -> api.auth.v1.ExitUsersLoginRequest
	5, // 3: api.auth.v1.Users.PatchUsersLogin:input_type -> api.auth.v1.PatchUsersLoginRequest
	8, // 4: api.auth.v1.Users.AuthLogin:input_type -> api.auth.v1.AuthLoginRequest
	1, // 5: api.auth.v1.Users.LoginUsers:output_type -> api.auth.v1.LoginUsersReply
	4, // 6: api.auth.v1.Users.ExitUsersLogin:output_type -> api.auth.v1.ExitUsersLoginReply
	6, // 7: api.auth.v1.Users.PatchUsersLogin:output_type -> api.auth.v1.PatchUsersLoginReply
	9, // 8: api.auth.v1.Users.AuthLogin:output_type -> api.auth.v1.AuthLoginReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_auth_v1_auth_proto_init() }
func file_api_auth_v1_auth_proto_init() {
	if File_api_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUsersRequest); i {
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
		file_api_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUsersReply); i {
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
		file_api_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginData); i {
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
		file_api_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitUsersLoginRequest); i {
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
		file_api_auth_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitUsersLoginReply); i {
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
		file_api_auth_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchUsersLoginRequest); i {
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
		file_api_auth_v1_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchUsersLoginReply); i {
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
		file_api_auth_v1_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchPasswordReply); i {
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
		file_api_auth_v1_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLoginRequest); i {
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
		file_api_auth_v1_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLoginReply); i {
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
		file_api_auth_v1_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchUsersLoginReplyData); i {
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
			RawDescriptor: file_api_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_api_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_api_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_api_auth_v1_auth_proto = out.File
	file_api_auth_v1_auth_proto_rawDesc = nil
	file_api_auth_v1_auth_proto_goTypes = nil
	file_api_auth_v1_auth_proto_depIdxs = nil
}
