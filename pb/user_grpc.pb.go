// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	ClientSignup(ctx context.Context, in *ClientSignUpRequest, opts ...grpc.CallOption) (*ClientSignUpResponse, error)
	FreelancerSignup(ctx context.Context, in *FreelancerSignUpRequest, opts ...grpc.CallOption) (*FreelancerSignUpResponse, error)
	ClientLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*ClientSignUpResponse, error)
	FreelancerLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*FreelancerSignUpResponse, error)
	AdminLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*ClientSignUpResponse, error)
	ClientCreateProfile(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FreelancerCreateProfile(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllCategoryClient, error)
	GetCategoryById(ctx context.Context, in *GetCategoryByIdRequest, opts ...grpc.CallOption) (*UpdateCategoryRequest, error)
	AdminAddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AdminUpdateSkill(ctx context.Context, in *SkillResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllSkills(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllSkillsClient, error)
	ClientAddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ClientUpdateAddress(ctx context.Context, in *AddressResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ClientGetAddress(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*AddressResponse, error)
	FreelancerAddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FreelancerUpdateAddress(ctx context.Context, in *AddressResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FreelancerGetAddress(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*AddressResponse, error)
	BlockClient(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnBlockClient(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BlockFreelancer(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnBlockFreelancer(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ClientSignup(ctx context.Context, in *ClientSignUpRequest, opts ...grpc.CallOption) (*ClientSignUpResponse, error) {
	out := new(ClientSignUpResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ClientSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelancerSignup(ctx context.Context, in *FreelancerSignUpRequest, opts ...grpc.CallOption) (*FreelancerSignUpResponse, error) {
	out := new(FreelancerSignUpResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FreelancerSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ClientLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*ClientSignUpResponse, error) {
	out := new(ClientSignUpResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ClientLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelancerLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*FreelancerSignUpResponse, error) {
	out := new(FreelancerSignUpResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FreelancerLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AdminLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*ClientSignUpResponse, error) {
	out := new(ClientSignUpResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/AdminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ClientCreateProfile(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/ClientCreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelancerCreateProfile(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/FreelancerCreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllCategoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/user.UserService/GetAllCategory", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetAllCategoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetAllCategoryClient interface {
	Recv() (*UpdateCategoryRequest, error)
	grpc.ClientStream
}

type userServiceGetAllCategoryClient struct {
	grpc.ClientStream
}

func (x *userServiceGetAllCategoryClient) Recv() (*UpdateCategoryRequest, error) {
	m := new(UpdateCategoryRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetCategoryById(ctx context.Context, in *GetCategoryByIdRequest, opts ...grpc.CallOption) (*UpdateCategoryRequest, error) {
	out := new(UpdateCategoryRequest)
	err := c.cc.Invoke(ctx, "/user.UserService/GetCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AdminAddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/AdminAddSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AdminUpdateSkill(ctx context.Context, in *SkillResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/AdminUpdateSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllSkills(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserService_GetAllSkillsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/user.UserService/GetAllSkills", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetAllSkillsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetAllSkillsClient interface {
	Recv() (*SkillResponse, error)
	grpc.ClientStream
}

type userServiceGetAllSkillsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetAllSkillsClient) Recv() (*SkillResponse, error) {
	m := new(SkillResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) ClientAddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/ClientAddAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ClientUpdateAddress(ctx context.Context, in *AddressResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/ClientUpdateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ClientGetAddress(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*AddressResponse, error) {
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ClientGetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelancerAddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/FreelancerAddAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelancerUpdateAddress(ctx context.Context, in *AddressResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/FreelancerUpdateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelancerGetAddress(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*AddressResponse, error) {
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FreelancerGetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BlockClient(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/BlockClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnBlockClient(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/UnBlockClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BlockFreelancer(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/BlockFreelancer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnBlockFreelancer(ctx context.Context, in *GetUserById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/UnBlockFreelancer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	ClientSignup(context.Context, *ClientSignUpRequest) (*ClientSignUpResponse, error)
	FreelancerSignup(context.Context, *FreelancerSignUpRequest) (*FreelancerSignUpResponse, error)
	ClientLogin(context.Context, *LoginRequest) (*ClientSignUpResponse, error)
	FreelancerLogin(context.Context, *LoginRequest) (*FreelancerSignUpResponse, error)
	AdminLogin(context.Context, *LoginRequest) (*ClientSignUpResponse, error)
	ClientCreateProfile(context.Context, *GetUserById) (*emptypb.Empty, error)
	FreelancerCreateProfile(context.Context, *GetUserById) (*emptypb.Empty, error)
	AddCategory(context.Context, *AddCategoryRequest) (*emptypb.Empty, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*emptypb.Empty, error)
	GetAllCategory(*emptypb.Empty, UserService_GetAllCategoryServer) error
	GetCategoryById(context.Context, *GetCategoryByIdRequest) (*UpdateCategoryRequest, error)
	AdminAddSkill(context.Context, *AddSkillRequest) (*emptypb.Empty, error)
	AdminUpdateSkill(context.Context, *SkillResponse) (*emptypb.Empty, error)
	GetAllSkills(*emptypb.Empty, UserService_GetAllSkillsServer) error
	ClientAddAddress(context.Context, *AddAddressRequest) (*emptypb.Empty, error)
	ClientUpdateAddress(context.Context, *AddressResponse) (*emptypb.Empty, error)
	ClientGetAddress(context.Context, *GetUserById) (*AddressResponse, error)
	FreelancerAddAddress(context.Context, *AddAddressRequest) (*emptypb.Empty, error)
	FreelancerUpdateAddress(context.Context, *AddressResponse) (*emptypb.Empty, error)
	FreelancerGetAddress(context.Context, *GetUserById) (*AddressResponse, error)
	BlockClient(context.Context, *GetUserById) (*emptypb.Empty, error)
	UnBlockClient(context.Context, *GetUserById) (*emptypb.Empty, error)
	BlockFreelancer(context.Context, *GetUserById) (*emptypb.Empty, error)
	UnBlockFreelancer(context.Context, *GetUserById) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) ClientSignup(context.Context, *ClientSignUpRequest) (*ClientSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientSignup not implemented")
}
func (UnimplementedUserServiceServer) FreelancerSignup(context.Context, *FreelancerSignUpRequest) (*FreelancerSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelancerSignup not implemented")
}
func (UnimplementedUserServiceServer) ClientLogin(context.Context, *LoginRequest) (*ClientSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientLogin not implemented")
}
func (UnimplementedUserServiceServer) FreelancerLogin(context.Context, *LoginRequest) (*FreelancerSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelancerLogin not implemented")
}
func (UnimplementedUserServiceServer) AdminLogin(context.Context, *LoginRequest) (*ClientSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedUserServiceServer) ClientCreateProfile(context.Context, *GetUserById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientCreateProfile not implemented")
}
func (UnimplementedUserServiceServer) FreelancerCreateProfile(context.Context, *GetUserById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelancerCreateProfile not implemented")
}
func (UnimplementedUserServiceServer) AddCategory(context.Context, *AddCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedUserServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedUserServiceServer) GetAllCategory(*emptypb.Empty, UserService_GetAllCategoryServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllCategory not implemented")
}
func (UnimplementedUserServiceServer) GetCategoryById(context.Context, *GetCategoryByIdRequest) (*UpdateCategoryRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
func (UnimplementedUserServiceServer) AdminAddSkill(context.Context, *AddSkillRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminAddSkill not implemented")
}
func (UnimplementedUserServiceServer) AdminUpdateSkill(context.Context, *SkillResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateSkill not implemented")
}
func (UnimplementedUserServiceServer) GetAllSkills(*emptypb.Empty, UserService_GetAllSkillsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllSkills not implemented")
}
func (UnimplementedUserServiceServer) ClientAddAddress(context.Context, *AddAddressRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientAddAddress not implemented")
}
func (UnimplementedUserServiceServer) ClientUpdateAddress(context.Context, *AddressResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUpdateAddress not implemented")
}
func (UnimplementedUserServiceServer) ClientGetAddress(context.Context, *GetUserById) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientGetAddress not implemented")
}
func (UnimplementedUserServiceServer) FreelancerAddAddress(context.Context, *AddAddressRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelancerAddAddress not implemented")
}
func (UnimplementedUserServiceServer) FreelancerUpdateAddress(context.Context, *AddressResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelancerUpdateAddress not implemented")
}
func (UnimplementedUserServiceServer) FreelancerGetAddress(context.Context, *GetUserById) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelancerGetAddress not implemented")
}
func (UnimplementedUserServiceServer) BlockClient(context.Context, *GetUserById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockClient not implemented")
}
func (UnimplementedUserServiceServer) UnBlockClient(context.Context, *GetUserById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockClient not implemented")
}
func (UnimplementedUserServiceServer) BlockFreelancer(context.Context, *GetUserById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockFreelancer not implemented")
}
func (UnimplementedUserServiceServer) UnBlockFreelancer(context.Context, *GetUserById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockFreelancer not implemented")
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

func _UserService_ClientSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClientSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ClientSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClientSignup(ctx, req.(*ClientSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelancerSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreelancerSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelancerSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FreelancerSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelancerSignup(ctx, req.(*FreelancerSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ClientLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClientLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ClientLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClientLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelancerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelancerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FreelancerLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelancerLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AdminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AdminLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ClientCreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClientCreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ClientCreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClientCreateProfile(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelancerCreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelancerCreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FreelancerCreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelancerCreateProfile(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllCategory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetAllCategory(m, &userServiceGetAllCategoryServer{stream})
}

type UserService_GetAllCategoryServer interface {
	Send(*UpdateCategoryRequest) error
	grpc.ServerStream
}

type userServiceGetAllCategoryServer struct {
	grpc.ServerStream
}

func (x *userServiceGetAllCategoryServer) Send(m *UpdateCategoryRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_GetCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCategoryById(ctx, req.(*GetCategoryByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AdminAddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AdminAddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AdminAddSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AdminAddSkill(ctx, req.(*AddSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AdminUpdateSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkillResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AdminUpdateSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AdminUpdateSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AdminUpdateSkill(ctx, req.(*SkillResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllSkills_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetAllSkills(m, &userServiceGetAllSkillsServer{stream})
}

type UserService_GetAllSkillsServer interface {
	Send(*SkillResponse) error
	grpc.ServerStream
}

type userServiceGetAllSkillsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetAllSkillsServer) Send(m *SkillResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_ClientAddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClientAddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ClientAddAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClientAddAddress(ctx, req.(*AddAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ClientUpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClientUpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ClientUpdateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClientUpdateAddress(ctx, req.(*AddressResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ClientGetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ClientGetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ClientGetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ClientGetAddress(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelancerAddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelancerAddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FreelancerAddAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelancerAddAddress(ctx, req.(*AddAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelancerUpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelancerUpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FreelancerUpdateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelancerUpdateAddress(ctx, req.(*AddressResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelancerGetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelancerGetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FreelancerGetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelancerGetAddress(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BlockClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BlockClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/BlockClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BlockClient(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnBlockClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnBlockClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UnBlockClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnBlockClient(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BlockFreelancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BlockFreelancer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/BlockFreelancer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BlockFreelancer(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnBlockFreelancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnBlockFreelancer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UnBlockFreelancer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnBlockFreelancer(ctx, req.(*GetUserById))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientSignup",
			Handler:    _UserService_ClientSignup_Handler,
		},
		{
			MethodName: "FreelancerSignup",
			Handler:    _UserService_FreelancerSignup_Handler,
		},
		{
			MethodName: "ClientLogin",
			Handler:    _UserService_ClientLogin_Handler,
		},
		{
			MethodName: "FreelancerLogin",
			Handler:    _UserService_FreelancerLogin_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _UserService_AdminLogin_Handler,
		},
		{
			MethodName: "ClientCreateProfile",
			Handler:    _UserService_ClientCreateProfile_Handler,
		},
		{
			MethodName: "FreelancerCreateProfile",
			Handler:    _UserService_FreelancerCreateProfile_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _UserService_AddCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _UserService_UpdateCategory_Handler,
		},
		{
			MethodName: "GetCategoryById",
			Handler:    _UserService_GetCategoryById_Handler,
		},
		{
			MethodName: "AdminAddSkill",
			Handler:    _UserService_AdminAddSkill_Handler,
		},
		{
			MethodName: "AdminUpdateSkill",
			Handler:    _UserService_AdminUpdateSkill_Handler,
		},
		{
			MethodName: "ClientAddAddress",
			Handler:    _UserService_ClientAddAddress_Handler,
		},
		{
			MethodName: "ClientUpdateAddress",
			Handler:    _UserService_ClientUpdateAddress_Handler,
		},
		{
			MethodName: "ClientGetAddress",
			Handler:    _UserService_ClientGetAddress_Handler,
		},
		{
			MethodName: "FreelancerAddAddress",
			Handler:    _UserService_FreelancerAddAddress_Handler,
		},
		{
			MethodName: "FreelancerUpdateAddress",
			Handler:    _UserService_FreelancerUpdateAddress_Handler,
		},
		{
			MethodName: "FreelancerGetAddress",
			Handler:    _UserService_FreelancerGetAddress_Handler,
		},
		{
			MethodName: "BlockClient",
			Handler:    _UserService_BlockClient_Handler,
		},
		{
			MethodName: "UnBlockClient",
			Handler:    _UserService_UnBlockClient_Handler,
		},
		{
			MethodName: "BlockFreelancer",
			Handler:    _UserService_BlockFreelancer_Handler,
		},
		{
			MethodName: "UnBlockFreelancer",
			Handler:    _UserService_UnBlockFreelancer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllCategory",
			Handler:       _UserService_GetAllCategory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllSkills",
			Handler:       _UserService_GetAllSkills_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}
