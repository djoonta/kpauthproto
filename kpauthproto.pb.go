// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: kpauthproto.proto

package kpauthproto

import (
	auth "github.com/djoonta/kpauthproto/auth"
	role "github.com/djoonta/kpauthproto/role"
	user "github.com/djoonta/kpauthproto/user"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_kpauthproto_proto protoreflect.FileDescriptor

var file_kpauthproto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x9b, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1d, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x20, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68,
	0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26,
	0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0xdf, 0x04, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64,
	0x49, 0x44, 0x12, 0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x52, 0x6f, 0x6c,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x6b, 0x70,
	0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74,
	0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6b,
	0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x32, 0xbd, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x49,
	0x44, 0x12, 0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1e, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6b, 0x70, 0x61, 0x75,
	0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5e, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x23, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x1f,
	0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61, 0x2f, 0x6b, 0x70, 0x61, 0x75, 0x74, 0x68,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_kpauthproto_proto_goTypes = []interface{}{
	(*auth.AuthLoginRequest)(nil),           // 0: kpauthproto.AuthLoginRequest
	(*auth.AuthRegisterRequest)(nil),        // 1: kpauthproto.AuthRegisterRequest
	(*auth.AuthForgotPasswordRequest)(nil),  // 2: kpauthproto.AuthForgotPasswordRequest
	(*role.RoleCreateRequest)(nil),          // 3: kpauthproto.RoleCreateRequest
	(*role.RoleDeleteRequest)(nil),          // 4: kpauthproto.RoleDeleteRequest
	(*role.RoleUpdateRequest)(nil),          // 5: kpauthproto.RoleUpdateRequest
	(*role.RoleFindIDRequest)(nil),          // 6: kpauthproto.RoleFindIDRequest
	(*role.RoleFindAllRequest)(nil),         // 7: kpauthproto.RoleFindAllRequest
	(*role.RoleAssignUserRequest)(nil),      // 8: kpauthproto.RoleAssignUserRequest
	(*role.RoleRevokeUserRequest)(nil),      // 9: kpauthproto.RoleRevokeUserRequest
	(*user.UserFindIDRequest)(nil),          // 10: kpauthproto.UserFindIDRequest
	(*user.UserDeleteRequest)(nil),          // 11: kpauthproto.UserDeleteRequest
	(*user.UserActivatedRequest)(nil),       // 12: kpauthproto.UserActivatedRequest
	(*user.UserDeactivatedRequest)(nil),     // 13: kpauthproto.UserDeactivatedRequest
	(*user.UserFindAllRequest)(nil),         // 14: kpauthproto.UserFindAllRequest
	(*auth.AuthLoginResponse)(nil),          // 15: kpauthproto.AuthLoginResponse
	(*auth.AuthRegisterResponse)(nil),       // 16: kpauthproto.AuthRegisterResponse
	(*auth.AuthForgotPasswordResponse)(nil), // 17: kpauthproto.AuthForgotPasswordResponse
	(*role.RoleCreateResponse)(nil),         // 18: kpauthproto.RoleCreateResponse
	(*role.RoleDeleteResponse)(nil),         // 19: kpauthproto.RoleDeleteResponse
	(*role.RoleUpdateResponse)(nil),         // 20: kpauthproto.RoleUpdateResponse
	(*role.RoleFindIDResponse)(nil),         // 21: kpauthproto.RoleFindIDResponse
	(*role.RoleFindAllResponse)(nil),        // 22: kpauthproto.RoleFindAllResponse
	(*role.RoleAssignUserResponse)(nil),     // 23: kpauthproto.RoleAssignUserResponse
	(*role.RoleRevokeUserResponse)(nil),     // 24: kpauthproto.RoleRevokeUserResponse
	(*user.UserFindIDResponse)(nil),         // 25: kpauthproto.UserFindIDResponse
	(*user.UserDeleteResponse)(nil),         // 26: kpauthproto.UserDeleteResponse
	(*user.UserActivatedResponse)(nil),      // 27: kpauthproto.UserActivatedResponse
	(*user.UserDeactivatedResponse)(nil),    // 28: kpauthproto.UserDeactivatedResponse
	(*user.UserFindAllResponse)(nil),        // 29: kpauthproto.UserFindAllResponse
}
var file_kpauthproto_proto_depIdxs = []int32{
	0,  // 0: kpauthproto.AuthService.AuthLogin:input_type -> kpauthproto.AuthLoginRequest
	1,  // 1: kpauthproto.AuthService.AuthRegister:input_type -> kpauthproto.AuthRegisterRequest
	2,  // 2: kpauthproto.AuthService.AuthForgotPassword:input_type -> kpauthproto.AuthForgotPasswordRequest
	3,  // 3: kpauthproto.RoleService.RoleCreate:input_type -> kpauthproto.RoleCreateRequest
	4,  // 4: kpauthproto.RoleService.RoleDelete:input_type -> kpauthproto.RoleDeleteRequest
	5,  // 5: kpauthproto.RoleService.RoleUpdate:input_type -> kpauthproto.RoleUpdateRequest
	6,  // 6: kpauthproto.RoleService.RoleFindID:input_type -> kpauthproto.RoleFindIDRequest
	7,  // 7: kpauthproto.RoleService.RoleFindAll:input_type -> kpauthproto.RoleFindAllRequest
	8,  // 8: kpauthproto.RoleService.RoleAssignUser:input_type -> kpauthproto.RoleAssignUserRequest
	9,  // 9: kpauthproto.RoleService.RoleRevokeUser:input_type -> kpauthproto.RoleRevokeUserRequest
	10, // 10: kpauthproto.UserService.UserFindID:input_type -> kpauthproto.UserFindIDRequest
	11, // 11: kpauthproto.UserService.UserDelete:input_type -> kpauthproto.UserDeleteRequest
	12, // 12: kpauthproto.UserService.UserActivated:input_type -> kpauthproto.UserActivatedRequest
	13, // 13: kpauthproto.UserService.UserDeactivated:input_type -> kpauthproto.UserDeactivatedRequest
	14, // 14: kpauthproto.UserService.UserFindAll:input_type -> kpauthproto.UserFindAllRequest
	15, // 15: kpauthproto.AuthService.AuthLogin:output_type -> kpauthproto.AuthLoginResponse
	16, // 16: kpauthproto.AuthService.AuthRegister:output_type -> kpauthproto.AuthRegisterResponse
	17, // 17: kpauthproto.AuthService.AuthForgotPassword:output_type -> kpauthproto.AuthForgotPasswordResponse
	18, // 18: kpauthproto.RoleService.RoleCreate:output_type -> kpauthproto.RoleCreateResponse
	19, // 19: kpauthproto.RoleService.RoleDelete:output_type -> kpauthproto.RoleDeleteResponse
	20, // 20: kpauthproto.RoleService.RoleUpdate:output_type -> kpauthproto.RoleUpdateResponse
	21, // 21: kpauthproto.RoleService.RoleFindID:output_type -> kpauthproto.RoleFindIDResponse
	22, // 22: kpauthproto.RoleService.RoleFindAll:output_type -> kpauthproto.RoleFindAllResponse
	23, // 23: kpauthproto.RoleService.RoleAssignUser:output_type -> kpauthproto.RoleAssignUserResponse
	24, // 24: kpauthproto.RoleService.RoleRevokeUser:output_type -> kpauthproto.RoleRevokeUserResponse
	25, // 25: kpauthproto.UserService.UserFindID:output_type -> kpauthproto.UserFindIDResponse
	26, // 26: kpauthproto.UserService.UserDelete:output_type -> kpauthproto.UserDeleteResponse
	27, // 27: kpauthproto.UserService.UserActivated:output_type -> kpauthproto.UserActivatedResponse
	28, // 28: kpauthproto.UserService.UserDeactivated:output_type -> kpauthproto.UserDeactivatedResponse
	29, // 29: kpauthproto.UserService.UserFindAll:output_type -> kpauthproto.UserFindAllResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_kpauthproto_proto_init() }
func file_kpauthproto_proto_init() {
	if File_kpauthproto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kpauthproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_kpauthproto_proto_goTypes,
		DependencyIndexes: file_kpauthproto_proto_depIdxs,
	}.Build()
	File_kpauthproto_proto = out.File
	file_kpauthproto_proto_rawDesc = nil
	file_kpauthproto_proto_goTypes = nil
	file_kpauthproto_proto_depIdxs = nil
}
