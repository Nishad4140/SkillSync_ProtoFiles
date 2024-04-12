// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: project.proto

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

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	CreateGig(ctx context.Context, in *CreateGigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateGig(ctx context.Context, in *GigResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetGig(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*GigResponse, error)
	GetAllFreelancerGigs(ctx context.Context, in *GetByUserId, opts ...grpc.CallOption) (ProjectService_GetAllFreelancerGigsClient, error)
	ShowIntrest(ctx context.Context, in *IntrestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllIntrest(ctx context.Context, in *GetAllIntrestRequest, opts ...grpc.CallOption) (ProjectService_GetAllIntrestClient, error)
	AddPackageType(ctx context.Context, in *AddPackageTypeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EditPackageType(ctx context.Context, in *PackageTypeResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPackageType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ProjectService_GetPackageTypeClient, error)
	ClientAddRequest(ctx context.Context, in *AddClientGigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ClientUpdateRequest(ctx context.Context, in *ClientRequestResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetClientRequest(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*ClientRequestResponse, error)
	GetAllClientRequest(ctx context.Context, in *RequestFilterQuery, opts ...grpc.CallOption) (ProjectService_GetAllClientRequestClient, error)
	ClientIntrestAcknowledgment(ctx context.Context, in *IntrestAcknowledgmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllGigs(ctx context.Context, in *GigFilterQuery, opts ...grpc.CallOption) (ProjectService_GetAllGigsClient, error)
	GetAllClientRequestForFreelancers(ctx context.Context, in *RequestFilterQuery, opts ...grpc.CallOption) (ProjectService_GetAllClientRequestForFreelancersClient, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateGig(ctx context.Context, in *CreateGigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/CreateGig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateGig(ctx context.Context, in *GigResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/UpdateGig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetGig(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*GigResponse, error) {
	out := new(GigResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/GetGig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetAllFreelancerGigs(ctx context.Context, in *GetByUserId, opts ...grpc.CallOption) (ProjectService_GetAllFreelancerGigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[0], "/project.ProjectService/GetAllFreelancerGigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetAllFreelancerGigsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetAllFreelancerGigsClient interface {
	Recv() (*GigResponse, error)
	grpc.ClientStream
}

type projectServiceGetAllFreelancerGigsClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetAllFreelancerGigsClient) Recv() (*GigResponse, error) {
	m := new(GigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) ShowIntrest(ctx context.Context, in *IntrestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/ShowIntrest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetAllIntrest(ctx context.Context, in *GetAllIntrestRequest, opts ...grpc.CallOption) (ProjectService_GetAllIntrestClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[1], "/project.ProjectService/GetAllIntrest", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetAllIntrestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetAllIntrestClient interface {
	Recv() (*IntrestResponse, error)
	grpc.ClientStream
}

type projectServiceGetAllIntrestClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetAllIntrestClient) Recv() (*IntrestResponse, error) {
	m := new(IntrestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) AddPackageType(ctx context.Context, in *AddPackageTypeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/AddPackageType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) EditPackageType(ctx context.Context, in *PackageTypeResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/EditPackageType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetPackageType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ProjectService_GetPackageTypeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[2], "/project.ProjectService/GetPackageType", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetPackageTypeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetPackageTypeClient interface {
	Recv() (*PackageTypeResponse, error)
	grpc.ClientStream
}

type projectServiceGetPackageTypeClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetPackageTypeClient) Recv() (*PackageTypeResponse, error) {
	m := new(PackageTypeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) ClientAddRequest(ctx context.Context, in *AddClientGigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/ClientAddRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ClientUpdateRequest(ctx context.Context, in *ClientRequestResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/ClientUpdateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetClientRequest(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*ClientRequestResponse, error) {
	out := new(ClientRequestResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/GetClientRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetAllClientRequest(ctx context.Context, in *RequestFilterQuery, opts ...grpc.CallOption) (ProjectService_GetAllClientRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[3], "/project.ProjectService/GetAllClientRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetAllClientRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetAllClientRequestClient interface {
	Recv() (*ClientRequestResponse, error)
	grpc.ClientStream
}

type projectServiceGetAllClientRequestClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetAllClientRequestClient) Recv() (*ClientRequestResponse, error) {
	m := new(ClientRequestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) ClientIntrestAcknowledgment(ctx context.Context, in *IntrestAcknowledgmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/project.ProjectService/ClientIntrestAcknowledgment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetAllGigs(ctx context.Context, in *GigFilterQuery, opts ...grpc.CallOption) (ProjectService_GetAllGigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[4], "/project.ProjectService/GetAllGigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetAllGigsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetAllGigsClient interface {
	Recv() (*GigResponse, error)
	grpc.ClientStream
}

type projectServiceGetAllGigsClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetAllGigsClient) Recv() (*GigResponse, error) {
	m := new(GigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) GetAllClientRequestForFreelancers(ctx context.Context, in *RequestFilterQuery, opts ...grpc.CallOption) (ProjectService_GetAllClientRequestForFreelancersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[5], "/project.ProjectService/GetAllClientRequestForFreelancers", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceGetAllClientRequestForFreelancersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_GetAllClientRequestForFreelancersClient interface {
	Recv() (*ClientRequestResponse, error)
	grpc.ClientStream
}

type projectServiceGetAllClientRequestForFreelancersClient struct {
	grpc.ClientStream
}

func (x *projectServiceGetAllClientRequestForFreelancersClient) Recv() (*ClientRequestResponse, error) {
	m := new(ClientRequestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	CreateGig(context.Context, *CreateGigRequest) (*emptypb.Empty, error)
	UpdateGig(context.Context, *GigResponse) (*emptypb.Empty, error)
	GetGig(context.Context, *GetById) (*GigResponse, error)
	GetAllFreelancerGigs(*GetByUserId, ProjectService_GetAllFreelancerGigsServer) error
	ShowIntrest(context.Context, *IntrestRequest) (*emptypb.Empty, error)
	GetAllIntrest(*GetAllIntrestRequest, ProjectService_GetAllIntrestServer) error
	AddPackageType(context.Context, *AddPackageTypeRequest) (*emptypb.Empty, error)
	EditPackageType(context.Context, *PackageTypeResponse) (*emptypb.Empty, error)
	GetPackageType(*emptypb.Empty, ProjectService_GetPackageTypeServer) error
	ClientAddRequest(context.Context, *AddClientGigRequest) (*emptypb.Empty, error)
	ClientUpdateRequest(context.Context, *ClientRequestResponse) (*emptypb.Empty, error)
	GetClientRequest(context.Context, *GetById) (*ClientRequestResponse, error)
	GetAllClientRequest(*RequestFilterQuery, ProjectService_GetAllClientRequestServer) error
	ClientIntrestAcknowledgment(context.Context, *IntrestAcknowledgmentRequest) (*emptypb.Empty, error)
	GetAllGigs(*GigFilterQuery, ProjectService_GetAllGigsServer) error
	GetAllClientRequestForFreelancers(*RequestFilterQuery, ProjectService_GetAllClientRequestForFreelancersServer) error
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) CreateGig(context.Context, *CreateGigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGig not implemented")
}
func (UnimplementedProjectServiceServer) UpdateGig(context.Context, *GigResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGig not implemented")
}
func (UnimplementedProjectServiceServer) GetGig(context.Context, *GetById) (*GigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGig not implemented")
}
func (UnimplementedProjectServiceServer) GetAllFreelancerGigs(*GetByUserId, ProjectService_GetAllFreelancerGigsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllFreelancerGigs not implemented")
}
func (UnimplementedProjectServiceServer) ShowIntrest(context.Context, *IntrestRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowIntrest not implemented")
}
func (UnimplementedProjectServiceServer) GetAllIntrest(*GetAllIntrestRequest, ProjectService_GetAllIntrestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllIntrest not implemented")
}
func (UnimplementedProjectServiceServer) AddPackageType(context.Context, *AddPackageTypeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPackageType not implemented")
}
func (UnimplementedProjectServiceServer) EditPackageType(context.Context, *PackageTypeResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPackageType not implemented")
}
func (UnimplementedProjectServiceServer) GetPackageType(*emptypb.Empty, ProjectService_GetPackageTypeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPackageType not implemented")
}
func (UnimplementedProjectServiceServer) ClientAddRequest(context.Context, *AddClientGigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientAddRequest not implemented")
}
func (UnimplementedProjectServiceServer) ClientUpdateRequest(context.Context, *ClientRequestResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUpdateRequest not implemented")
}
func (UnimplementedProjectServiceServer) GetClientRequest(context.Context, *GetById) (*ClientRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientRequest not implemented")
}
func (UnimplementedProjectServiceServer) GetAllClientRequest(*RequestFilterQuery, ProjectService_GetAllClientRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllClientRequest not implemented")
}
func (UnimplementedProjectServiceServer) ClientIntrestAcknowledgment(context.Context, *IntrestAcknowledgmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientIntrestAcknowledgment not implemented")
}
func (UnimplementedProjectServiceServer) GetAllGigs(*GigFilterQuery, ProjectService_GetAllGigsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllGigs not implemented")
}
func (UnimplementedProjectServiceServer) GetAllClientRequestForFreelancers(*RequestFilterQuery, ProjectService_GetAllClientRequestForFreelancersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllClientRequestForFreelancers not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateGig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateGig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/CreateGig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateGig(ctx, req.(*CreateGigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateGig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GigResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateGig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/UpdateGig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateGig(ctx, req.(*GigResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetGig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetGig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/GetGig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetGig(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetAllFreelancerGigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetByUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetAllFreelancerGigs(m, &projectServiceGetAllFreelancerGigsServer{stream})
}

type ProjectService_GetAllFreelancerGigsServer interface {
	Send(*GigResponse) error
	grpc.ServerStream
}

type projectServiceGetAllFreelancerGigsServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetAllFreelancerGigsServer) Send(m *GigResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_ShowIntrest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ShowIntrest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/ShowIntrest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ShowIntrest(ctx, req.(*IntrestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetAllIntrest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllIntrestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetAllIntrest(m, &projectServiceGetAllIntrestServer{stream})
}

type ProjectService_GetAllIntrestServer interface {
	Send(*IntrestResponse) error
	grpc.ServerStream
}

type projectServiceGetAllIntrestServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetAllIntrestServer) Send(m *IntrestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_AddPackageType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPackageTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddPackageType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/AddPackageType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddPackageType(ctx, req.(*AddPackageTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_EditPackageType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageTypeResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).EditPackageType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/EditPackageType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).EditPackageType(ctx, req.(*PackageTypeResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetPackageType_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetPackageType(m, &projectServiceGetPackageTypeServer{stream})
}

type ProjectService_GetPackageTypeServer interface {
	Send(*PackageTypeResponse) error
	grpc.ServerStream
}

type projectServiceGetPackageTypeServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetPackageTypeServer) Send(m *PackageTypeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_ClientAddRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientGigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ClientAddRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/ClientAddRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ClientAddRequest(ctx, req.(*AddClientGigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ClientUpdateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequestResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ClientUpdateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/ClientUpdateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ClientUpdateRequest(ctx, req.(*ClientRequestResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetClientRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetClientRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/GetClientRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetClientRequest(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetAllClientRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilterQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetAllClientRequest(m, &projectServiceGetAllClientRequestServer{stream})
}

type ProjectService_GetAllClientRequestServer interface {
	Send(*ClientRequestResponse) error
	grpc.ServerStream
}

type projectServiceGetAllClientRequestServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetAllClientRequestServer) Send(m *ClientRequestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_ClientIntrestAcknowledgment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntrestAcknowledgmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ClientIntrestAcknowledgment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/ClientIntrestAcknowledgment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ClientIntrestAcknowledgment(ctx, req.(*IntrestAcknowledgmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetAllGigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GigFilterQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetAllGigs(m, &projectServiceGetAllGigsServer{stream})
}

type ProjectService_GetAllGigsServer interface {
	Send(*GigResponse) error
	grpc.ServerStream
}

type projectServiceGetAllGigsServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetAllGigsServer) Send(m *GigResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_GetAllClientRequestForFreelancers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestFilterQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).GetAllClientRequestForFreelancers(m, &projectServiceGetAllClientRequestForFreelancersServer{stream})
}

type ProjectService_GetAllClientRequestForFreelancersServer interface {
	Send(*ClientRequestResponse) error
	grpc.ServerStream
}

type projectServiceGetAllClientRequestForFreelancersServer struct {
	grpc.ServerStream
}

func (x *projectServiceGetAllClientRequestForFreelancersServer) Send(m *ClientRequestResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGig",
			Handler:    _ProjectService_CreateGig_Handler,
		},
		{
			MethodName: "UpdateGig",
			Handler:    _ProjectService_UpdateGig_Handler,
		},
		{
			MethodName: "GetGig",
			Handler:    _ProjectService_GetGig_Handler,
		},
		{
			MethodName: "ShowIntrest",
			Handler:    _ProjectService_ShowIntrest_Handler,
		},
		{
			MethodName: "AddPackageType",
			Handler:    _ProjectService_AddPackageType_Handler,
		},
		{
			MethodName: "EditPackageType",
			Handler:    _ProjectService_EditPackageType_Handler,
		},
		{
			MethodName: "ClientAddRequest",
			Handler:    _ProjectService_ClientAddRequest_Handler,
		},
		{
			MethodName: "ClientUpdateRequest",
			Handler:    _ProjectService_ClientUpdateRequest_Handler,
		},
		{
			MethodName: "GetClientRequest",
			Handler:    _ProjectService_GetClientRequest_Handler,
		},
		{
			MethodName: "ClientIntrestAcknowledgment",
			Handler:    _ProjectService_ClientIntrestAcknowledgment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllFreelancerGigs",
			Handler:       _ProjectService_GetAllFreelancerGigs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllIntrest",
			Handler:       _ProjectService_GetAllIntrest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPackageType",
			Handler:       _ProjectService_GetPackageType_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllClientRequest",
			Handler:       _ProjectService_GetAllClientRequest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllGigs",
			Handler:       _ProjectService_GetAllGigs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllClientRequestForFreelancers",
			Handler:       _ProjectService_GetAllClientRequestForFreelancers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "project.proto",
}
