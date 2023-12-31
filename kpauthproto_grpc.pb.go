// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: kpauthproto.proto

package kpauthproto

import (
	context "context"
	auth "github.com/djoonta/kpauthproto/auth"
	role "github.com/djoonta/kpauthproto/role"
	user "github.com/djoonta/kpauthproto/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	AuthLogin(ctx context.Context, in *auth.AuthLoginRequest, opts ...grpc.CallOption) (*auth.AuthLoginResponse, error)
	AuthRegister(ctx context.Context, in *auth.AuthRegisterRequest, opts ...grpc.CallOption) (*auth.AuthRegisterResponse, error)
	AuthForgotPassword(ctx context.Context, in *auth.AuthForgotPasswordRequest, opts ...grpc.CallOption) (*auth.AuthForgotPasswordResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AuthLogin(ctx context.Context, in *auth.AuthLoginRequest, opts ...grpc.CallOption) (*auth.AuthLoginResponse, error) {
	out := new(auth.AuthLoginResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.AuthService/AuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthRegister(ctx context.Context, in *auth.AuthRegisterRequest, opts ...grpc.CallOption) (*auth.AuthRegisterResponse, error) {
	out := new(auth.AuthRegisterResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.AuthService/AuthRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthForgotPassword(ctx context.Context, in *auth.AuthForgotPasswordRequest, opts ...grpc.CallOption) (*auth.AuthForgotPasswordResponse, error) {
	out := new(auth.AuthForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.AuthService/AuthForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	AuthLogin(context.Context, *auth.AuthLoginRequest) (*auth.AuthLoginResponse, error)
	AuthRegister(context.Context, *auth.AuthRegisterRequest) (*auth.AuthRegisterResponse, error)
	AuthForgotPassword(context.Context, *auth.AuthForgotPasswordRequest) (*auth.AuthForgotPasswordResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AuthLogin(context.Context, *auth.AuthLoginRequest) (*auth.AuthLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthLogin not implemented")
}
func (UnimplementedAuthServiceServer) AuthRegister(context.Context, *auth.AuthRegisterRequest) (*auth.AuthRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthRegister not implemented")
}
func (UnimplementedAuthServiceServer) AuthForgotPassword(context.Context, *auth.AuthForgotPasswordRequest) (*auth.AuthForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.AuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.AuthService/AuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthLogin(ctx, req.(*auth.AuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.AuthRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.AuthService/AuthRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthRegister(ctx, req.(*auth.AuthRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.AuthForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.AuthService/AuthForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthForgotPassword(ctx, req.(*auth.AuthForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kpauthproto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthLogin",
			Handler:    _AuthService_AuthLogin_Handler,
		},
		{
			MethodName: "AuthRegister",
			Handler:    _AuthService_AuthRegister_Handler,
		},
		{
			MethodName: "AuthForgotPassword",
			Handler:    _AuthService_AuthForgotPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kpauthproto.proto",
}

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	RoleCreate(ctx context.Context, in *role.RoleCreateRequest, opts ...grpc.CallOption) (*role.RoleCreateResponse, error)
	RoleDelete(ctx context.Context, in *role.RoleDeleteRequest, opts ...grpc.CallOption) (*role.RoleDeleteResponse, error)
	RoleUpdate(ctx context.Context, in *role.RoleUpdateRequest, opts ...grpc.CallOption) (*role.RoleUpdateResponse, error)
	RoleFindID(ctx context.Context, in *role.RoleFindIDRequest, opts ...grpc.CallOption) (*role.RoleFindIDResponse, error)
	RoleFindAll(ctx context.Context, in *role.RoleFindAllRequest, opts ...grpc.CallOption) (*role.RoleFindAllResponse, error)
	RoleAssignUser(ctx context.Context, in *role.RoleAssignUserRequest, opts ...grpc.CallOption) (*role.RoleAssignUserResponse, error)
	RoleRevokeUser(ctx context.Context, in *role.RoleRevokeUserRequest, opts ...grpc.CallOption) (*role.RoleRevokeUserResponse, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) RoleCreate(ctx context.Context, in *role.RoleCreateRequest, opts ...grpc.CallOption) (*role.RoleCreateResponse, error) {
	out := new(role.RoleCreateResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RoleDelete(ctx context.Context, in *role.RoleDeleteRequest, opts ...grpc.CallOption) (*role.RoleDeleteResponse, error) {
	out := new(role.RoleDeleteResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RoleUpdate(ctx context.Context, in *role.RoleUpdateRequest, opts ...grpc.CallOption) (*role.RoleUpdateResponse, error) {
	out := new(role.RoleUpdateResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RoleFindID(ctx context.Context, in *role.RoleFindIDRequest, opts ...grpc.CallOption) (*role.RoleFindIDResponse, error) {
	out := new(role.RoleFindIDResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleFindID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RoleFindAll(ctx context.Context, in *role.RoleFindAllRequest, opts ...grpc.CallOption) (*role.RoleFindAllResponse, error) {
	out := new(role.RoleFindAllResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleFindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RoleAssignUser(ctx context.Context, in *role.RoleAssignUserRequest, opts ...grpc.CallOption) (*role.RoleAssignUserResponse, error) {
	out := new(role.RoleAssignUserResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleAssignUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RoleRevokeUser(ctx context.Context, in *role.RoleRevokeUserRequest, opts ...grpc.CallOption) (*role.RoleRevokeUserResponse, error) {
	out := new(role.RoleRevokeUserResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.RoleService/RoleRevokeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	RoleCreate(context.Context, *role.RoleCreateRequest) (*role.RoleCreateResponse, error)
	RoleDelete(context.Context, *role.RoleDeleteRequest) (*role.RoleDeleteResponse, error)
	RoleUpdate(context.Context, *role.RoleUpdateRequest) (*role.RoleUpdateResponse, error)
	RoleFindID(context.Context, *role.RoleFindIDRequest) (*role.RoleFindIDResponse, error)
	RoleFindAll(context.Context, *role.RoleFindAllRequest) (*role.RoleFindAllResponse, error)
	RoleAssignUser(context.Context, *role.RoleAssignUserRequest) (*role.RoleAssignUserResponse, error)
	RoleRevokeUser(context.Context, *role.RoleRevokeUserRequest) (*role.RoleRevokeUserResponse, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) RoleCreate(context.Context, *role.RoleCreateRequest) (*role.RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleCreate not implemented")
}
func (UnimplementedRoleServiceServer) RoleDelete(context.Context, *role.RoleDeleteRequest) (*role.RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDelete not implemented")
}
func (UnimplementedRoleServiceServer) RoleUpdate(context.Context, *role.RoleUpdateRequest) (*role.RoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleUpdate not implemented")
}
func (UnimplementedRoleServiceServer) RoleFindID(context.Context, *role.RoleFindIDRequest) (*role.RoleFindIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleFindID not implemented")
}
func (UnimplementedRoleServiceServer) RoleFindAll(context.Context, *role.RoleFindAllRequest) (*role.RoleFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleFindAll not implemented")
}
func (UnimplementedRoleServiceServer) RoleAssignUser(context.Context, *role.RoleAssignUserRequest) (*role.RoleAssignUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAssignUser not implemented")
}
func (UnimplementedRoleServiceServer) RoleRevokeUser(context.Context, *role.RoleRevokeUserRequest) (*role.RoleRevokeUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleRevokeUser not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_RoleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleCreate(ctx, req.(*role.RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleDelete(ctx, req.(*role.RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RoleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleUpdate(ctx, req.(*role.RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RoleFindID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleFindIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleFindID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleFindID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleFindID(ctx, req.(*role.RoleFindIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RoleFindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleFindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleFindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleFindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleFindAll(ctx, req.(*role.RoleFindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RoleAssignUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleAssignUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleAssignUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleAssignUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleAssignUser(ctx, req.(*role.RoleAssignUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RoleRevokeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.RoleRevokeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RoleRevokeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.RoleService/RoleRevokeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RoleRevokeUser(ctx, req.(*role.RoleRevokeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kpauthproto.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoleCreate",
			Handler:    _RoleService_RoleCreate_Handler,
		},
		{
			MethodName: "RoleDelete",
			Handler:    _RoleService_RoleDelete_Handler,
		},
		{
			MethodName: "RoleUpdate",
			Handler:    _RoleService_RoleUpdate_Handler,
		},
		{
			MethodName: "RoleFindID",
			Handler:    _RoleService_RoleFindID_Handler,
		},
		{
			MethodName: "RoleFindAll",
			Handler:    _RoleService_RoleFindAll_Handler,
		},
		{
			MethodName: "RoleAssignUser",
			Handler:    _RoleService_RoleAssignUser_Handler,
		},
		{
			MethodName: "RoleRevokeUser",
			Handler:    _RoleService_RoleRevokeUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kpauthproto.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserFindID(ctx context.Context, in *user.UserFindIDRequest, opts ...grpc.CallOption) (*user.UserFindIDResponse, error)
	UserDelete(ctx context.Context, in *user.UserDeleteRequest, opts ...grpc.CallOption) (*user.UserDeleteResponse, error)
	UserActivated(ctx context.Context, in *user.UserActivatedRequest, opts ...grpc.CallOption) (*user.UserActivatedResponse, error)
	UserDeactivated(ctx context.Context, in *user.UserDeactivatedRequest, opts ...grpc.CallOption) (*user.UserDeactivatedResponse, error)
	UserFindAll(ctx context.Context, in *user.UserFindAllRequest, opts ...grpc.CallOption) (*user.UserFindAllResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserFindID(ctx context.Context, in *user.UserFindIDRequest, opts ...grpc.CallOption) (*user.UserFindIDResponse, error) {
	out := new(user.UserFindIDResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.UserService/UserFindID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserDelete(ctx context.Context, in *user.UserDeleteRequest, opts ...grpc.CallOption) (*user.UserDeleteResponse, error) {
	out := new(user.UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.UserService/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserActivated(ctx context.Context, in *user.UserActivatedRequest, opts ...grpc.CallOption) (*user.UserActivatedResponse, error) {
	out := new(user.UserActivatedResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.UserService/UserActivated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserDeactivated(ctx context.Context, in *user.UserDeactivatedRequest, opts ...grpc.CallOption) (*user.UserDeactivatedResponse, error) {
	out := new(user.UserDeactivatedResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.UserService/UserDeactivated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserFindAll(ctx context.Context, in *user.UserFindAllRequest, opts ...grpc.CallOption) (*user.UserFindAllResponse, error) {
	out := new(user.UserFindAllResponse)
	err := c.cc.Invoke(ctx, "/kpauthproto.UserService/UserFindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserFindID(context.Context, *user.UserFindIDRequest) (*user.UserFindIDResponse, error)
	UserDelete(context.Context, *user.UserDeleteRequest) (*user.UserDeleteResponse, error)
	UserActivated(context.Context, *user.UserActivatedRequest) (*user.UserActivatedResponse, error)
	UserDeactivated(context.Context, *user.UserDeactivatedRequest) (*user.UserDeactivatedResponse, error)
	UserFindAll(context.Context, *user.UserFindAllRequest) (*user.UserFindAllResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserFindID(context.Context, *user.UserFindIDRequest) (*user.UserFindIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFindID not implemented")
}
func (UnimplementedUserServiceServer) UserDelete(context.Context, *user.UserDeleteRequest) (*user.UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedUserServiceServer) UserActivated(context.Context, *user.UserActivatedRequest) (*user.UserActivatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserActivated not implemented")
}
func (UnimplementedUserServiceServer) UserDeactivated(context.Context, *user.UserDeactivatedRequest) (*user.UserDeactivatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDeactivated not implemented")
}
func (UnimplementedUserServiceServer) UserFindAll(context.Context, *user.UserFindAllRequest) (*user.UserFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFindAll not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserFindID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.UserFindIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserFindID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.UserService/UserFindID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserFindID(ctx, req.(*user.UserFindIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.UserService/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserDelete(ctx, req.(*user.UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserActivated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.UserActivatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserActivated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.UserService/UserActivated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserActivated(ctx, req.(*user.UserActivatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserDeactivated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.UserDeactivatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserDeactivated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.UserService/UserDeactivated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserDeactivated(ctx, req.(*user.UserDeactivatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserFindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.UserFindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserFindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kpauthproto.UserService/UserFindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserFindAll(ctx, req.(*user.UserFindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kpauthproto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserFindID",
			Handler:    _UserService_UserFindID_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _UserService_UserDelete_Handler,
		},
		{
			MethodName: "UserActivated",
			Handler:    _UserService_UserActivated_Handler,
		},
		{
			MethodName: "UserDeactivated",
			Handler:    _UserService_UserDeactivated_Handler,
		},
		{
			MethodName: "UserFindAll",
			Handler:    _UserService_UserFindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kpauthproto.proto",
}
