package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PasswordCheckInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password          string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	EncryptedPassword string `protobuf:"bytes,2,opt,name=encryptedPassword,proto3" json:"encryptedPassword,omitempty"`
}

func (x *PasswordCheckInfo) Reset() {
	*x = PasswordCheckInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordCheckInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordCheckInfo) ProtoMessage() {}

func (x *PasswordCheckInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordCheckInfo.ProtoReflect.Descriptor instead.
func (*PasswordCheckInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordCheckInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PasswordCheckInfo) GetEncryptedPassword() string {
	if x != nil {
		return x.EncryptedPassword
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pn    uint32 `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	PSize uint32 `protobuf:"varint,2,opt,name=pSize,proto3" json:"pSize,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *PageInfo) GetPn() uint32 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *PageInfo) GetPSize() uint32 {
	if x != nil {
		return x.PSize
	}
	return 0
}

type MobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *MobileRequest) Reset() {
	*x = MobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileRequest) ProtoMessage() {}

func (x *MobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileRequest.ProtoReflect.Descriptor instead.
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *MobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *IdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	PassWord string `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *CreateUserInfo) Reset() {
	*x = CreateUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserInfo) ProtoMessage() {}

func (x *CreateUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserInfo.ProtoReflect.Descriptor instead.
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CreateUserInfo) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *CreateUserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type UpdateUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Gender   string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDay uint64 `protobuf:"varint,4,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
}

func (x *UpdateUserInfo) Reset() {
	*x = UpdateUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfo) ProtoMessage() {}

func (x *UpdateUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfo.ProtoReflect.Descriptor instead.
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UpdateUserInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateUserInfo) GetBirthDay() uint64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PassWord string `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	BirthDay uint64 `protobuf:"varint,5,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Gender   string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Role     int32  `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *UserInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfoResponse) GetBirthDay() uint64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

func (x *UserInfoResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfoResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type UserListResonse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserListResonse) Reset() {
	*x = UserListResonse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResonse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResonse) ProtoMessage() {}

func (x *UserListResonse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResonse.ProtoReflect.Descriptor instead.
func (*UserListResonse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserListResonse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserListResonse) GetData() []*UserInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x11, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x70, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x70, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x27, 0x0a, 0x0d, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x1b,
	0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x70, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x22,
	0xba, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4e, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb4, 0x02, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x10, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33,
	0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_proto_goTypes = []interface{}{
	(*PasswordCheckInfo)(nil), // 0: PasswordCheckInfo
	(*CheckResponse)(nil),     // 1: CheckResponse
	(*PageInfo)(nil),          // 2: PageInfo
	(*MobileRequest)(nil),     // 3: MobileRequest
	(*IdRequest)(nil),         // 4: IdRequest
	(*CreateUserInfo)(nil),    // 5: CreateUserInfo
	(*UpdateUserInfo)(nil),    // 6: UpdateUserInfo
	(*UserInfoResponse)(nil),  // 7: UserInfoResponse
	(*UserListResonse)(nil),   // 8: UserListResonse
	(*empty.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	7, // 0: UserListResonse.data:type_name -> UserInfoResponse
	2, // 1: User.GetUserList:input_type -> PageInfo
	3, // 2: User.GetUserByMobile:input_type -> MobileRequest
	4, // 3: User.GetUserById:input_type -> IdRequest
	5, // 4: User.CreateUser:input_type -> CreateUserInfo
	6, // 5: User.UpdateUser:input_type -> UpdateUserInfo
	0, // 6: User.CheckPassWord:input_type -> PasswordCheckInfo
	8, // 7: User.GetUserList:output_type -> UserListResonse
	7, // 8: User.GetUserByMobile:output_type -> UserInfoResponse
	7, // 9: User.GetUserById:output_type -> UserInfoResponse
	7, // 10: User.CreateUser:output_type -> UserInfoResponse
	9, // 11: User.UpdateUser:output_type -> google.protobuf.Empty
	1, // 12: User.CheckPassWord:output_type -> CheckResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordCheckInfo); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserInfo); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfo); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResonse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResonse, error)
	GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckPassWord(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResonse, error) {
	out := new(UserListResonse)
	err := c.cc.Invoke(ctx, "/User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassWord(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/User/CheckPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListResonse, error)
	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*empty.Empty, error)
	CheckPassWord(context.Context, *PasswordCheckInfo) (*CheckResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserList(context.Context, *PageInfo) (*UserListResonse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedUserServer) GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (*UnimplementedUserServer) GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServer) CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServer) CheckPassWord(context.Context, *PasswordCheckInfo) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordCheckInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CheckPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassWord(ctx, req.(*PasswordCheckInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _User_GetUserByMobile_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassWord",
			Handler:    _User_CheckPassWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
