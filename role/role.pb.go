// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: role/role.proto

package role

import (
	pagination "github.com/djoonta/kpauthproto/pagination"
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

type RoleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RoleInfo) Reset() {
	*x = RoleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleInfo) ProtoMessage() {}

func (x *RoleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleInfo.ProtoReflect.Descriptor instead.
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RoleCreateRequest) Reset() {
	*x = RoleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateRequest) ProtoMessage() {}

func (x *RoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RoleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RoleCreateResponse) Reset() {
	*x = RoleCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateResponse) ProtoMessage() {}

func (x *RoleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateResponse.ProtoReflect.Descriptor instead.
func (*RoleCreateResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleCreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RoleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RoleUpdateRequest) Reset() {
	*x = RoleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateRequest) ProtoMessage() {}

func (x *RoleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{3}
}

func (x *RoleUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RoleUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RoleUpdateResponse) Reset() {
	*x = RoleUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateResponse) ProtoMessage() {}

func (x *RoleUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateResponse.ProtoReflect.Descriptor instead.
func (*RoleUpdateResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{4}
}

func (x *RoleUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RoleDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleDeleteRequest) Reset() {
	*x = RoleDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteRequest) ProtoMessage() {}

func (x *RoleDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteRequest.ProtoReflect.Descriptor instead.
func (*RoleDeleteRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{5}
}

func (x *RoleDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RoleDeleteResponse) Reset() {
	*x = RoleDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteResponse) ProtoMessage() {}

func (x *RoleDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteResponse.ProtoReflect.Descriptor instead.
func (*RoleDeleteResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{6}
}

func (x *RoleDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RoleFindIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleFindIDRequest) Reset() {
	*x = RoleFindIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindIDRequest) ProtoMessage() {}

func (x *RoleFindIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindIDRequest.ProtoReflect.Descriptor instead.
func (*RoleFindIDRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{7}
}

func (x *RoleFindIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleFindIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleInfo *RoleInfo `protobuf:"bytes,1,opt,name=role_info,json=roleInfo,proto3" json:"role_info,omitempty"`
}

func (x *RoleFindIDResponse) Reset() {
	*x = RoleFindIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindIDResponse) ProtoMessage() {}

func (x *RoleFindIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindIDResponse.ProtoReflect.Descriptor instead.
func (*RoleFindIDResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{8}
}

func (x *RoleFindIDResponse) GetRoleInfo() *RoleInfo {
	if x != nil {
		return x.RoleInfo
	}
	return nil
}

type RoleFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RoleFindAllRequest) Reset() {
	*x = RoleFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindAllRequest) ProtoMessage() {}

func (x *RoleFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindAllRequest.ProtoReflect.Descriptor instead.
func (*RoleFindAllRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{9}
}

func (x *RoleFindAllRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *RoleFindAllRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type RoleFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles      []*RoleInfo                `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	Pagination *pagination.PaginationInfo `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *RoleFindAllResponse) Reset() {
	*x = RoleFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindAllResponse) ProtoMessage() {}

func (x *RoleFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindAllResponse.ProtoReflect.Descriptor instead.
func (*RoleFindAllResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{10}
}

func (x *RoleFindAllResponse) GetRoles() []*RoleInfo {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *RoleFindAllResponse) GetPagination() *pagination.PaginationInfo {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type RoleAssignUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RoleAssignUserRequest) Reset() {
	*x = RoleAssignUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleAssignUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAssignUserRequest) ProtoMessage() {}

func (x *RoleAssignUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAssignUserRequest.ProtoReflect.Descriptor instead.
func (*RoleAssignUserRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{11}
}

func (x *RoleAssignUserRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleAssignUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RoleAssignUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RoleAssignUserResponse) Reset() {
	*x = RoleAssignUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleAssignUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAssignUserResponse) ProtoMessage() {}

func (x *RoleAssignUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAssignUserResponse.ProtoReflect.Descriptor instead.
func (*RoleAssignUserResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{12}
}

func (x *RoleAssignUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RoleRevokeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RoleRevokeUserRequest) Reset() {
	*x = RoleRevokeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleRevokeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRevokeUserRequest) ProtoMessage() {}

func (x *RoleRevokeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRevokeUserRequest.ProtoReflect.Descriptor instead.
func (*RoleRevokeUserRequest) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{13}
}

func (x *RoleRevokeUserRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleRevokeUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RoleRevokeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RoleRevokeUserResponse) Reset() {
	*x = RoleRevokeUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_role_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleRevokeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRevokeUserResponse) ProtoMessage() {}

func (x *RoleRevokeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_role_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRevokeUserResponse.ProtoReflect.Descriptor instead.
func (*RoleRevokeUserResponse) Descriptor() ([]byte, []int) {
	return file_role_role_proto_rawDescGZIP(), []int{14}
}

func (x *RoleRevokeUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_role_role_proto protoreflect.FileDescriptor

var file_role_role_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a,
	0x11, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65,
	0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a,
	0x12, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3e, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7f, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x65, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x15, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61, 0x2f, 0x6b, 0x70, 0x61,
	0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_role_proto_rawDescOnce sync.Once
	file_role_role_proto_rawDescData = file_role_role_proto_rawDesc
)

func file_role_role_proto_rawDescGZIP() []byte {
	file_role_role_proto_rawDescOnce.Do(func() {
		file_role_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_role_proto_rawDescData)
	})
	return file_role_role_proto_rawDescData
}

var file_role_role_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_role_role_proto_goTypes = []interface{}{
	(*RoleInfo)(nil),                  // 0: kpauthproto.RoleInfo
	(*RoleCreateRequest)(nil),         // 1: kpauthproto.RoleCreateRequest
	(*RoleCreateResponse)(nil),        // 2: kpauthproto.RoleCreateResponse
	(*RoleUpdateRequest)(nil),         // 3: kpauthproto.RoleUpdateRequest
	(*RoleUpdateResponse)(nil),        // 4: kpauthproto.RoleUpdateResponse
	(*RoleDeleteRequest)(nil),         // 5: kpauthproto.RoleDeleteRequest
	(*RoleDeleteResponse)(nil),        // 6: kpauthproto.RoleDeleteResponse
	(*RoleFindIDRequest)(nil),         // 7: kpauthproto.RoleFindIDRequest
	(*RoleFindIDResponse)(nil),        // 8: kpauthproto.RoleFindIDResponse
	(*RoleFindAllRequest)(nil),        // 9: kpauthproto.RoleFindAllRequest
	(*RoleFindAllResponse)(nil),       // 10: kpauthproto.RoleFindAllResponse
	(*RoleAssignUserRequest)(nil),     // 11: kpauthproto.RoleAssignUserRequest
	(*RoleAssignUserResponse)(nil),    // 12: kpauthproto.RoleAssignUserResponse
	(*RoleRevokeUserRequest)(nil),     // 13: kpauthproto.RoleRevokeUserRequest
	(*RoleRevokeUserResponse)(nil),    // 14: kpauthproto.RoleRevokeUserResponse
	(*pagination.PaginationInfo)(nil), // 15: kpauthproto.PaginationInfo
}
var file_role_role_proto_depIdxs = []int32{
	0,  // 0: kpauthproto.RoleFindIDResponse.role_info:type_name -> kpauthproto.RoleInfo
	0,  // 1: kpauthproto.RoleFindAllResponse.roles:type_name -> kpauthproto.RoleInfo
	15, // 2: kpauthproto.RoleFindAllResponse.pagination:type_name -> kpauthproto.PaginationInfo
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_role_role_proto_init() }
func file_role_role_proto_init() {
	if File_role_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleInfo); i {
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
		file_role_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreateRequest); i {
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
		file_role_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreateResponse); i {
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
		file_role_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateRequest); i {
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
		file_role_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateResponse); i {
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
		file_role_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteRequest); i {
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
		file_role_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteResponse); i {
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
		file_role_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindIDRequest); i {
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
		file_role_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindIDResponse); i {
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
		file_role_role_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindAllRequest); i {
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
		file_role_role_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindAllResponse); i {
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
		file_role_role_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleAssignUserRequest); i {
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
		file_role_role_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleAssignUserResponse); i {
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
		file_role_role_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleRevokeUserRequest); i {
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
		file_role_role_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleRevokeUserResponse); i {
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
			RawDescriptor: file_role_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_role_proto_goTypes,
		DependencyIndexes: file_role_role_proto_depIdxs,
		MessageInfos:      file_role_role_proto_msgTypes,
	}.Build()
	File_role_role_proto = out.File
	file_role_role_proto_rawDesc = nil
	file_role_role_proto_goTypes = nil
	file_role_role_proto_depIdxs = nil
}
