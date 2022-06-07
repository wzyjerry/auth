// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.6
// source: user/v1/register.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` // 账户名(*)
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // 密码
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"` // 昵称
}

func (x *RegisterAccountRequest) Reset() {
	*x = RegisterAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAccountRequest) ProtoMessage() {}

func (x *RegisterAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAccountRequest.ProtoReflect.Descriptor instead.
func (*RegisterAccountRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_register_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAccountRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterAccountRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 用户ID
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_user_v1_register_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RegisterPreEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // 邮箱(*)
}

func (x *RegisterPreEmailRequest) Reset() {
	*x = RegisterPreEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPreEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPreEmailRequest) ProtoMessage() {}

func (x *RegisterPreEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPreEmailRequest.ProtoReflect.Descriptor instead.
func (*RegisterPreEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_register_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterPreEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RegisterEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`       // 邮箱(*)
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // 密码
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"` // 昵称
	Code     string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`         // 验证码
}

func (x *RegisterEmailRequest) Reset() {
	*x = RegisterEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterEmailRequest) ProtoMessage() {}

func (x *RegisterEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterEmailRequest.ProtoReflect.Descriptor instead.
func (*RegisterEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_register_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterEmailRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterEmailRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterEmailRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RegisterPrePhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"` // 手机号(*)
}

func (x *RegisterPrePhoneRequest) Reset() {
	*x = RegisterPrePhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_register_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPrePhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPrePhoneRequest) ProtoMessage() {}

func (x *RegisterPrePhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_register_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPrePhoneRequest.ProtoReflect.Descriptor instead.
func (*RegisterPrePhoneRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_register_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterPrePhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type RegisterPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`       // 手机号(*)
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // 密码
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"` // 昵称
	Code     string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`         // 验证码
}

func (x *RegisterPhoneRequest) Reset() {
	*x = RegisterPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_register_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPhoneRequest) ProtoMessage() {}

func (x *RegisterPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_register_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPhoneRequest.ProtoReflect.Descriptor instead.
func (*RegisterPhoneRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_register_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterPhoneRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterPhoneRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterPhoneRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_user_v1_register_proto protoreflect.FileDescriptor

var file_user_v1_register_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x9c,
	0x01, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x49, 0x0a,
	0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15, 0x72, 0x13, 0x32, 0x11,
	0x5e, 0x5c, 0x2b, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x5c, 0x64, 0x7b, 0x31, 0x2c, 0x31, 0x34, 0x7d,
	0x24, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xb8, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x70, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x3a,
	0x01, 0x2a, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_register_proto_rawDescOnce sync.Once
	file_user_v1_register_proto_rawDescData = file_user_v1_register_proto_rawDesc
)

func file_user_v1_register_proto_rawDescGZIP() []byte {
	file_user_v1_register_proto_rawDescOnce.Do(func() {
		file_user_v1_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_register_proto_rawDescData)
	})
	return file_user_v1_register_proto_rawDescData
}

var file_user_v1_register_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_v1_register_proto_goTypes = []interface{}{
	(*RegisterAccountRequest)(nil),  // 0: api.user.v1.RegisterAccountRequest
	(*RegisterReply)(nil),           // 1: api.user.v1.RegisterReply
	(*RegisterPreEmailRequest)(nil), // 2: api.user.v1.RegisterPreEmailRequest
	(*RegisterEmailRequest)(nil),    // 3: api.user.v1.RegisterEmailRequest
	(*RegisterPrePhoneRequest)(nil), // 4: api.user.v1.RegisterPrePhoneRequest
	(*RegisterPhoneRequest)(nil),    // 5: api.user.v1.RegisterPhoneRequest
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_user_v1_register_proto_depIdxs = []int32{
	0, // 0: api.user.v1.Register.Account:input_type -> api.user.v1.RegisterAccountRequest
	2, // 1: api.user.v1.Register.PreEmail:input_type -> api.user.v1.RegisterPreEmailRequest
	3, // 2: api.user.v1.Register.Email:input_type -> api.user.v1.RegisterEmailRequest
	4, // 3: api.user.v1.Register.PrePhone:input_type -> api.user.v1.RegisterPrePhoneRequest
	5, // 4: api.user.v1.Register.Phone:input_type -> api.user.v1.RegisterPhoneRequest
	1, // 5: api.user.v1.Register.Account:output_type -> api.user.v1.RegisterReply
	6, // 6: api.user.v1.Register.PreEmail:output_type -> google.protobuf.Empty
	1, // 7: api.user.v1.Register.Email:output_type -> api.user.v1.RegisterReply
	6, // 8: api.user.v1.Register.PrePhone:output_type -> google.protobuf.Empty
	1, // 9: api.user.v1.Register.Phone:output_type -> api.user.v1.RegisterReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_v1_register_proto_init() }
func file_user_v1_register_proto_init() {
	if File_user_v1_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAccountRequest); i {
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
		file_user_v1_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_user_v1_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPreEmailRequest); i {
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
		file_user_v1_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterEmailRequest); i {
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
		file_user_v1_register_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPrePhoneRequest); i {
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
		file_user_v1_register_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPhoneRequest); i {
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
			RawDescriptor: file_user_v1_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_register_proto_goTypes,
		DependencyIndexes: file_user_v1_register_proto_depIdxs,
		MessageInfos:      file_user_v1_register_proto_msgTypes,
	}.Build()
	File_user_v1_register_proto = out.File
	file_user_v1_register_proto_rawDesc = nil
	file_user_v1_register_proto_goTypes = nil
	file_user_v1_register_proto_depIdxs = nil
}