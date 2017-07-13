// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compass.proto

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	compass.proto

It has these top-level messages:
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hapi_services_tiller "k8s.io/helm/pkg/proto/hapi/services"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CompassService service

type CompassServiceClient interface {
	// ListReleases retrieves release history.
	// TODO: Allow filtering the set of releases by
	// release status. By default, ListAllReleases returns the releases who
	// current status is "Active".
	ListReleases(ctx context.Context, in *hapi_services_tiller.ListReleasesRequest, opts ...grpc.CallOption) (*hapi_services_tiller.ListReleasesResponse, error)
	// GetReleasesStatus retrieves status information for the specified release.
	GetReleaseStatus(ctx context.Context, in *hapi_services_tiller.GetReleaseStatusRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetReleaseStatusResponse, error)
	// // GetReleaseContent retrieves the release content (chart + value) for the specified release.
	GetReleaseContent(ctx context.Context, in *hapi_services_tiller.GetReleaseContentRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetReleaseContentResponse, error)
	// // UpdateRelease updates release content.
	UpdateRelease(ctx context.Context, in *hapi_services_tiller.UpdateReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.UpdateReleaseResponse, error)
	// // InstallRelease requests installation of a chart as a new release.
	InstallRelease(ctx context.Context, in *hapi_services_tiller.InstallReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.InstallReleaseResponse, error)
	// // UninstallRelease requests deletion of a named release.
	UninstallRelease(ctx context.Context, in *hapi_services_tiller.UninstallReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.UninstallReleaseResponse, error)
	// // GetVersion returns the current version of the server.
	GetVersion(ctx context.Context, in *hapi_services_tiller.GetVersionRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetVersionResponse, error)
	// // RollbackRelease rolls back a release to a previous version.
	RollbackRelease(ctx context.Context, in *hapi_services_tiller.RollbackReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.RollbackReleaseResponse, error)
	// // ReleaseHistory retrieves a releasse's history.
	GetHistory(ctx context.Context, in *hapi_services_tiller.GetHistoryRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetHistoryResponse, error)
	// // RunReleaseTest executes the tests defined of a named release
	RunReleaseTest(ctx context.Context, in *hapi_services_tiller.TestReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.TestReleaseResponse, error)
}

type compassServiceClient struct {
	cc *grpc.ClientConn
}

func NewCompassServiceClient(cc *grpc.ClientConn) CompassServiceClient {
	return &compassServiceClient{cc}
}

func (c *compassServiceClient) ListReleases(ctx context.Context, in *hapi_services_tiller.ListReleasesRequest, opts ...grpc.CallOption) (*hapi_services_tiller.ListReleasesResponse, error) {
	out := new(hapi_services_tiller.ListReleasesResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/ListReleases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) GetReleaseStatus(ctx context.Context, in *hapi_services_tiller.GetReleaseStatusRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetReleaseStatusResponse, error) {
	out := new(hapi_services_tiller.GetReleaseStatusResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/GetReleaseStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) GetReleaseContent(ctx context.Context, in *hapi_services_tiller.GetReleaseContentRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetReleaseContentResponse, error) {
	out := new(hapi_services_tiller.GetReleaseContentResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/GetReleaseContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) UpdateRelease(ctx context.Context, in *hapi_services_tiller.UpdateReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.UpdateReleaseResponse, error) {
	out := new(hapi_services_tiller.UpdateReleaseResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/UpdateRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) InstallRelease(ctx context.Context, in *hapi_services_tiller.InstallReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.InstallReleaseResponse, error) {
	out := new(hapi_services_tiller.InstallReleaseResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/InstallRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) UninstallRelease(ctx context.Context, in *hapi_services_tiller.UninstallReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.UninstallReleaseResponse, error) {
	out := new(hapi_services_tiller.UninstallReleaseResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/UninstallRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) GetVersion(ctx context.Context, in *hapi_services_tiller.GetVersionRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetVersionResponse, error) {
	out := new(hapi_services_tiller.GetVersionResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) RollbackRelease(ctx context.Context, in *hapi_services_tiller.RollbackReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.RollbackReleaseResponse, error) {
	out := new(hapi_services_tiller.RollbackReleaseResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/RollbackRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) GetHistory(ctx context.Context, in *hapi_services_tiller.GetHistoryRequest, opts ...grpc.CallOption) (*hapi_services_tiller.GetHistoryResponse, error) {
	out := new(hapi_services_tiller.GetHistoryResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/GetHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compassServiceClient) RunReleaseTest(ctx context.Context, in *hapi_services_tiller.TestReleaseRequest, opts ...grpc.CallOption) (*hapi_services_tiller.TestReleaseResponse, error) {
	out := new(hapi_services_tiller.TestReleaseResponse)
	err := grpc.Invoke(ctx, "/services.CompassService/RunReleaseTest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CompassService service

type CompassServiceServer interface {
	// ListReleases retrieves release history.
	// TODO: Allow filtering the set of releases by
	// release status. By default, ListAllReleases returns the releases who
	// current status is "Active".
	ListReleases(context.Context, *hapi_services_tiller.ListReleasesRequest) (*hapi_services_tiller.ListReleasesResponse, error)
	// GetReleasesStatus retrieves status information for the specified release.
	GetReleaseStatus(context.Context, *hapi_services_tiller.GetReleaseStatusRequest) (*hapi_services_tiller.GetReleaseStatusResponse, error)
	// // GetReleaseContent retrieves the release content (chart + value) for the specified release.
	GetReleaseContent(context.Context, *hapi_services_tiller.GetReleaseContentRequest) (*hapi_services_tiller.GetReleaseContentResponse, error)
	// // UpdateRelease updates release content.
	UpdateRelease(context.Context, *hapi_services_tiller.UpdateReleaseRequest) (*hapi_services_tiller.UpdateReleaseResponse, error)
	// // InstallRelease requests installation of a chart as a new release.
	InstallRelease(context.Context, *hapi_services_tiller.InstallReleaseRequest) (*hapi_services_tiller.InstallReleaseResponse, error)
	// // UninstallRelease requests deletion of a named release.
	UninstallRelease(context.Context, *hapi_services_tiller.UninstallReleaseRequest) (*hapi_services_tiller.UninstallReleaseResponse, error)
	// // GetVersion returns the current version of the server.
	GetVersion(context.Context, *hapi_services_tiller.GetVersionRequest) (*hapi_services_tiller.GetVersionResponse, error)
	// // RollbackRelease rolls back a release to a previous version.
	RollbackRelease(context.Context, *hapi_services_tiller.RollbackReleaseRequest) (*hapi_services_tiller.RollbackReleaseResponse, error)
	// // ReleaseHistory retrieves a releasse's history.
	GetHistory(context.Context, *hapi_services_tiller.GetHistoryRequest) (*hapi_services_tiller.GetHistoryResponse, error)
	// // RunReleaseTest executes the tests defined of a named release
	RunReleaseTest(context.Context, *hapi_services_tiller.TestReleaseRequest) (*hapi_services_tiller.TestReleaseResponse, error)
}

func RegisterCompassServiceServer(s *grpc.Server, srv CompassServiceServer) {
	s.RegisterService(&_CompassService_serviceDesc, srv)
}

func _CompassService_ListReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.ListReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).ListReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/ListReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).ListReleases(ctx, req.(*hapi_services_tiller.ListReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_GetReleaseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.GetReleaseStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).GetReleaseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/GetReleaseStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).GetReleaseStatus(ctx, req.(*hapi_services_tiller.GetReleaseStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_GetReleaseContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.GetReleaseContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).GetReleaseContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/GetReleaseContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).GetReleaseContent(ctx, req.(*hapi_services_tiller.GetReleaseContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_UpdateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.UpdateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).UpdateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/UpdateRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).UpdateRelease(ctx, req.(*hapi_services_tiller.UpdateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_InstallRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.InstallReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).InstallRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/InstallRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).InstallRelease(ctx, req.(*hapi_services_tiller.InstallReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_UninstallRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.UninstallReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).UninstallRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/UninstallRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).UninstallRelease(ctx, req.(*hapi_services_tiller.UninstallReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).GetVersion(ctx, req.(*hapi_services_tiller.GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_RollbackRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.RollbackReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).RollbackRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/RollbackRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).RollbackRelease(ctx, req.(*hapi_services_tiller.RollbackReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.GetHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).GetHistory(ctx, req.(*hapi_services_tiller.GetHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompassService_RunReleaseTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hapi_services_tiller.TestReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompassServiceServer).RunReleaseTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.CompassService/RunReleaseTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompassServiceServer).RunReleaseTest(ctx, req.(*hapi_services_tiller.TestReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompassService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.CompassService",
	HandlerType: (*CompassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListReleases",
			Handler:    _CompassService_ListReleases_Handler,
		},
		{
			MethodName: "GetReleaseStatus",
			Handler:    _CompassService_GetReleaseStatus_Handler,
		},
		{
			MethodName: "GetReleaseContent",
			Handler:    _CompassService_GetReleaseContent_Handler,
		},
		{
			MethodName: "UpdateRelease",
			Handler:    _CompassService_UpdateRelease_Handler,
		},
		{
			MethodName: "InstallRelease",
			Handler:    _CompassService_InstallRelease_Handler,
		},
		{
			MethodName: "UninstallRelease",
			Handler:    _CompassService_UninstallRelease_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _CompassService_GetVersion_Handler,
		},
		{
			MethodName: "RollbackRelease",
			Handler:    _CompassService_RollbackRelease_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _CompassService_GetHistory_Handler,
		},
		{
			MethodName: "RunReleaseTest",
			Handler:    _CompassService_RunReleaseTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compass.proto",
}

func init() { proto.RegisterFile("compass.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0x3d, 0x19, 0x32, 0x11, 0xd4, 0x1e, 0x39, 0x7a, 0x51, 0x54, 0x96, 0x44, 0xdf, 0x40,
	0x0e, 0x68, 0xe2, 0x69, 0x11, 0x0f, 0xde, 0x0a, 0x4e, 0xb0, 0x5a, 0xda, 0x75, 0x67, 0x20, 0xf1,
	0x49, 0x7d, 0x1d, 0x83, 0xec, 0x54, 0x58, 0xad, 0xdb, 0xe3, 0x66, 0xbe, 0x7f, 0xbe, 0xfc, 0x3b,
	0x49, 0xa1, 0x3d, 0xf3, 0x8b, 0x42, 0x13, 0x65, 0x45, 0xe9, 0xd9, 0xab, 0x16, 0x61, 0xb9, 0x32,
	0x33, 0xa4, 0x6e, 0xf7, 0x45, 0x17, 0x66, 0x20, 0x9f, 0x03, 0x36, 0xd6, 0x62, 0xb9, 0xa1, 0xae,
	0x3e, 0x5b, 0xd0, 0x19, 0x6e, 0x72, 0xe3, 0x0d, 0xa0, 0xe6, 0x70, 0x70, 0x6f, 0x88, 0x73, 0xb4,
	0xa8, 0x09, 0x49, 0xf5, 0xb2, 0x75, 0x3e, 0x93, 0x7c, 0x56, 0xe5, 0xb7, 0x99, 0x1c, 0xdf, 0x97,
	0x48, 0xdc, 0x3d, 0x4f, 0x41, 0xa9, 0xf0, 0x8e, 0xf0, 0x64, 0x4f, 0x11, 0x1c, 0x8d, 0x50, 0x06,
	0x63, 0xd6, 0xbc, 0x24, 0xd5, 0xff, 0x7b, 0x43, 0x9d, 0x13, 0x61, 0x96, 0x8a, 0x07, 0xe9, 0x0a,
	0x8e, 0x7f, 0xa6, 0x43, 0xef, 0x18, 0x1d, 0xab, 0xc6, 0x35, 0x15, 0x28, 0xda, 0x41, 0x32, 0x1f,
	0xbc, 0xaf, 0xd0, 0x9e, 0x14, 0xcf, 0x9a, 0xb1, 0x22, 0x54, 0xe4, 0x5f, 0xed, 0x40, 0xe2, 0xbb,
	0x48, 0x62, 0x83, 0x6b, 0x01, 0x9d, 0x3b, 0x47, 0xac, 0xad, 0x15, 0x59, 0x64, 0xc1, 0x2e, 0x25,
	0xb6, 0xcb, 0x34, 0x78, 0xfb, 0x8e, 0x13, 0x67, 0x76, 0x85, 0x91, 0x3b, 0xd6, 0xb9, 0x86, 0x3b,
	0xfe, 0xc6, 0x83, 0x54, 0x03, 0x8c, 0x90, 0x1f, 0xb1, 0x24, 0xe3, 0x9d, 0x3a, 0x8d, 0x1e, 0xa4,
	0x22, 0x44, 0x74, 0xd6, 0x0c, 0x06, 0x45, 0x01, 0x87, 0xb9, 0xb7, 0x76, 0xaa, 0x67, 0x6f, 0x52,
	0x2b, 0xf2, 0x6b, 0x6a, 0x98, 0xc8, 0xfa, 0x89, 0x74, 0xad, 0xd4, 0xad, 0x21, 0xf6, 0xe5, 0xc7,
	0x3f, 0xa5, 0x2a, 0xa2, 0xb9, 0x54, 0x00, 0x83, 0x62, 0x0e, 0x9d, 0x7c, 0xe9, 0x2a, 0xf5, 0x03,
	0x12, 0xab, 0x48, 0x7a, 0x3d, 0xab, 0xf5, 0xe9, 0x25, 0x90, 0x22, 0xba, 0x81, 0xa7, 0xf0, 0x02,
	0x4d, 0xf7, 0xbf, 0x1f, 0x9b, 0xeb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x19, 0x02, 0xba,
	0xa3, 0x04, 0x00, 0x00,
}