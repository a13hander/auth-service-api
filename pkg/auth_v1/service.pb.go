// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: service.proto

package auth_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_USER  Role = 0
	Role_ADMIN Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "USER",
		1: "ADMIN",
	}
	Role_value = map[string]int32{
		"USER":  0,
		"ADMIN": 1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Role     Role   `protobuf:"varint,3,opt,name=role,proto3,enum=auth_v1.Role" json:"role,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_USER
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info      *UserInfo              `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User            *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password        string    `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm string    `protobuf:"bytes,3,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateRequest) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60,
	0x01, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18,
	0x64, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x78,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x08, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x20, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x39, 0x0a,
	0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x08,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x20, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x2a, 0x1b, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x32, 0x9b, 0x01, 0x0a, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x56, 0x31, 0x12, 0x5a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x31, 0x33, 0x68, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x61, 0x70,
	0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(Role)(0),                     // 0: auth_v1.Role
	(*UserInfo)(nil),              // 1: auth_v1.UserInfo
	(*User)(nil),                  // 2: auth_v1.User
	(*CreateRequest)(nil),         // 3: auth_v1.CreateRequest
	(*CreateResponse)(nil),        // 4: auth_v1.CreateResponse
	(*ListResponse)(nil),          // 5: auth_v1.ListResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	0, // 0: auth_v1.UserInfo.role:type_name -> auth_v1.Role
	1, // 1: auth_v1.User.info:type_name -> auth_v1.UserInfo
	6, // 2: auth_v1.User.created_at:type_name -> google.protobuf.Timestamp
	1, // 3: auth_v1.CreateRequest.user:type_name -> auth_v1.UserInfo
	2, // 4: auth_v1.ListResponse.user:type_name -> auth_v1.User
	3, // 5: auth_v1.AuthV1.Create:input_type -> auth_v1.CreateRequest
	7, // 6: auth_v1.AuthV1.List:input_type -> google.protobuf.Empty
	4, // 7: auth_v1.AuthV1.Create:output_type -> auth_v1.CreateResponse
	5, // 8: auth_v1.AuthV1.List:output_type -> auth_v1.ListResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
