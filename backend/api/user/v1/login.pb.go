// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.6
// source: user/v1/login.proto

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

type Type int32

const (
	Type_TYPE_UNSET     Type = 0
	Type_TYPE_ACCOUNT   Type = 1
	Type_TYPE_EMAIL     Type = 2
	Type_TYPE_PHONE     Type = 3
	Type_TYPE_GITHUB    Type = 4
	Type_TYPE_MICROSOFT Type = 5
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "TYPE_ACCOUNT",
		2: "TYPE_EMAIL",
		3: "TYPE_PHONE",
		4: "TYPE_GITHUB",
		5: "TYPE_MICROSOFT",
	}
	Type_value = map[string]int32{
		"TYPE_UNSET":     0,
		"TYPE_ACCOUNT":   1,
		"TYPE_EMAIL":     2,
		"TYPE_PHONE":     3,
		"TYPE_GITHUB":    4,
		"TYPE_MICROSOFT": 5,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_user_v1_login_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_user_v1_login_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{0}
}

type Method int32

const (
	Method_METHOD_UNSET    Method = 0
	Method_METHOD_PASSWORD Method = 1
	Method_METHOD_CODE     Method = 2
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "METHOD_UNSET",
		1: "METHOD_PASSWORD",
		2: "METHOD_CODE",
	}
	Method_value = map[string]int32{
		"METHOD_UNSET":    0,
		"METHOD_PASSWORD": 1,
		"METHOD_CODE":     2,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_user_v1_login_proto_enumTypes[1].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_user_v1_login_proto_enumTypes[1]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{1}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Type   `protobuf:"varint,1,opt,name=type,proto3,enum=api.user.v1.Type" json:"type,omitempty"`
	Method Method `protobuf:"varint,2,opt,name=method,proto3,enum=api.user.v1.Method" json:"method,omitempty"`
	Unique string `protobuf:"bytes,3,opt,name=unique,proto3" json:"unique,omitempty"`
	Secret string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSET
}

func (x *LoginRequest) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_METHOD_UNSET
}

func (x *LoginRequest) GetUnique() string {
	if x != nil {
		return x.Unique
	}
	return ""
}

func (x *LoginRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type OAuthLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType    string  `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn    int32   `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	Scope        string  `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	AccessToken  string  `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken *string `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3,oneof" json:"refresh_token,omitempty"`
	IdToken      *string `protobuf:"bytes,6,opt,name=id_token,json=idToken,proto3,oneof" json:"id_token,omitempty"`
}

func (x *OAuthLoginReply) Reset() {
	*x = OAuthLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginReply) ProtoMessage() {}

func (x *OAuthLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthLoginReply.ProtoReflect.Descriptor instead.
func (*OAuthLoginReply) Descriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthLoginReply) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *OAuthLoginReply) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *OAuthLoginReply) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *OAuthLoginReply) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *OAuthLoginReply) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

func (x *OAuthLoginReply) GetIdToken() string {
	if x != nil && x.IdToken != nil {
		return *x.IdToken
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginPrePhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *LoginPrePhoneRequest) Reset() {
	*x = LoginPrePhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPrePhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPrePhoneRequest) ProtoMessage() {}

func (x *LoginPrePhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPrePhoneRequest.ProtoReflect.Descriptor instead.
func (*LoginPrePhoneRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginPrePhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type LoginPreEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *LoginPreEmailRequest) Reset() {
	*x = LoginPreEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPreEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPreEmailRequest) ProtoMessage() {}

func (x *LoginPreEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPreEmailRequest.ProtoReflect.Descriptor instead.
func (*LoginPreEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{4}
}

func (x *LoginPreEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TrashReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken   bool   `protobuf:"varint,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	UserToken   bool   `protobuf:"varint,2,opt,name=userToken,proto3" json:"userToken,omitempty"`
	ClientToken bool   `protobuf:"varint,3,opt,name=clientToken,proto3" json:"clientToken,omitempty"`
	Sub         string `protobuf:"bytes,4,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *TrashReply) Reset() {
	*x = TrashReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrashReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrashReply) ProtoMessage() {}

func (x *TrashReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrashReply.ProtoReflect.Descriptor instead.
func (*TrashReply) Descriptor() ([]byte, []int) {
	return file_user_v1_login_proto_rawDescGZIP(), []int{5}
}

func (x *TrashReply) GetAuthToken() bool {
	if x != nil {
		return x.AuthToken
	}
	return false
}

func (x *TrashReply) GetUserToken() bool {
	if x != nil {
		return x.UserToken
	}
	return false
}

func (x *TrashReply) GetClientToken() bool {
	if x != nil {
		return x.ClientToken
	}
	return false
}

func (x *TrashReply) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

var File_user_v1_login_proto protoreflect.FileDescriptor

var file_user_v1_login_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xf1, 0x01, 0x0a, 0x0f, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x46, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15, 0x72, 0x13, 0x32, 0x11,
	0x5e, 0x5c, 0x2b, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x5c, 0x64, 0x7b, 0x31, 0x2c, 0x31, 0x34, 0x7d,
	0x24, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x35, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x50, 0x72, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x7c, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x2a, 0x6d, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x4f, 0x46, 0x54, 0x10, 0x05, 0x2a, 0x40, 0x0a, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44,
	0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x54, 0x48,
	0x4f, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x32, 0x95,
	0x03, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6a, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a,
	0x12, 0x5c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x56,
	0x0a, 0x05, 0x54, 0x72, 0x61, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x12, 0x14, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2f, 0x74, 0x72, 0x61, 0x73, 0x68, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_login_proto_rawDescOnce sync.Once
	file_user_v1_login_proto_rawDescData = file_user_v1_login_proto_rawDesc
)

func file_user_v1_login_proto_rawDescGZIP() []byte {
	file_user_v1_login_proto_rawDescOnce.Do(func() {
		file_user_v1_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_login_proto_rawDescData)
	})
	return file_user_v1_login_proto_rawDescData
}

var file_user_v1_login_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_v1_login_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_v1_login_proto_goTypes = []interface{}{
	(Type)(0),                    // 0: api.user.v1.Type
	(Method)(0),                  // 1: api.user.v1.Method
	(*LoginRequest)(nil),         // 2: api.user.v1.LoginRequest
	(*OAuthLoginReply)(nil),      // 3: api.user.v1.OAuthLoginReply
	(*LoginReply)(nil),           // 4: api.user.v1.LoginReply
	(*LoginPrePhoneRequest)(nil), // 5: api.user.v1.LoginPrePhoneRequest
	(*LoginPreEmailRequest)(nil), // 6: api.user.v1.LoginPreEmailRequest
	(*TrashReply)(nil),           // 7: api.user.v1.TrashReply
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
}
var file_user_v1_login_proto_depIdxs = []int32{
	0, // 0: api.user.v1.LoginRequest.type:type_name -> api.user.v1.Type
	1, // 1: api.user.v1.LoginRequest.method:type_name -> api.user.v1.Method
	5, // 2: api.user.v1.Login.PrePhone:input_type -> api.user.v1.LoginPrePhoneRequest
	6, // 3: api.user.v1.Login.PreEmail:input_type -> api.user.v1.LoginPreEmailRequest
	2, // 4: api.user.v1.Login.Login:input_type -> api.user.v1.LoginRequest
	8, // 5: api.user.v1.Login.Trash:input_type -> google.protobuf.Empty
	8, // 6: api.user.v1.Login.PrePhone:output_type -> google.protobuf.Empty
	8, // 7: api.user.v1.Login.PreEmail:output_type -> google.protobuf.Empty
	4, // 8: api.user.v1.Login.Login:output_type -> api.user.v1.LoginReply
	7, // 9: api.user.v1.Login.Trash:output_type -> api.user.v1.TrashReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_v1_login_proto_init() }
func file_user_v1_login_proto_init() {
	if File_user_v1_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_user_v1_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthLoginReply); i {
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
		file_user_v1_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_user_v1_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPrePhoneRequest); i {
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
		file_user_v1_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPreEmailRequest); i {
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
		file_user_v1_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrashReply); i {
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
	file_user_v1_login_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_v1_login_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_login_proto_goTypes,
		DependencyIndexes: file_user_v1_login_proto_depIdxs,
		EnumInfos:         file_user_v1_login_proto_enumTypes,
		MessageInfos:      file_user_v1_login_proto_msgTypes,
	}.Build()
	File_user_v1_login_proto = out.File
	file_user_v1_login_proto_rawDesc = nil
	file_user_v1_login_proto_goTypes = nil
	file_user_v1_login_proto_depIdxs = nil
}
