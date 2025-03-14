// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: feed_server.proto

package pbfs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Upstream_Handshake_FullMethodName            = "/pbfs.Upstream/Handshake"
	Upstream_Messaging_FullMethodName            = "/pbfs.Upstream/Messaging"
	Upstream_Watch_FullMethodName                = "/pbfs.Upstream/Watch"
	Upstream_PullAppFileMeta_FullMethodName      = "/pbfs.Upstream/PullAppFileMeta"
	Upstream_GetDownloadURL_FullMethodName       = "/pbfs.Upstream/GetDownloadURL"
	Upstream_PullKvMeta_FullMethodName           = "/pbfs.Upstream/PullKvMeta"
	Upstream_GetKvValue_FullMethodName           = "/pbfs.Upstream/GetKvValue"
	Upstream_ListApps_FullMethodName             = "/pbfs.Upstream/ListApps"
	Upstream_AsyncDownload_FullMethodName        = "/pbfs.Upstream/AsyncDownload"
	Upstream_AsyncDownloadStatus_FullMethodName  = "/pbfs.Upstream/AsyncDownloadStatus"
	Upstream_GetSingleKvValue_FullMethodName     = "/pbfs.Upstream/GetSingleKvValue"
	Upstream_GetSingleKvMeta_FullMethodName      = "/pbfs.Upstream/GetSingleKvMeta"
	Upstream_GetSingleFileContent_FullMethodName = "/pbfs.Upstream/GetSingleFileContent"
)

// UpstreamClient is the client API for Upstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpstreamClient interface {
	// APIs for sidecar.
	Handshake(ctx context.Context, in *HandshakeMessage, opts ...grpc.CallOption) (*HandshakeResp, error)
	Messaging(ctx context.Context, in *MessagingMeta, opts ...grpc.CallOption) (*MessagingResp, error)
	Watch(ctx context.Context, in *SideWatchMeta, opts ...grpc.CallOption) (Upstream_WatchClient, error)
	PullAppFileMeta(ctx context.Context, in *PullAppFileMetaReq, opts ...grpc.CallOption) (*PullAppFileMetaResp, error)
	GetDownloadURL(ctx context.Context, in *GetDownloadURLReq, opts ...grpc.CallOption) (*GetDownloadURLResp, error)
	PullKvMeta(ctx context.Context, in *PullKvMetaReq, opts ...grpc.CallOption) (*PullKvMetaResp, error)
	GetKvValue(ctx context.Context, in *GetKvValueReq, opts ...grpc.CallOption) (*GetKvValueResp, error)
	ListApps(ctx context.Context, in *ListAppsReq, opts ...grpc.CallOption) (*ListAppsResp, error)
	AsyncDownload(ctx context.Context, in *AsyncDownloadReq, opts ...grpc.CallOption) (*AsyncDownloadResp, error)
	AsyncDownloadStatus(ctx context.Context, in *AsyncDownloadStatusReq, opts ...grpc.CallOption) (*AsyncDownloadStatusResp, error)
	// 获取单个kv实例
	GetSingleKvValue(ctx context.Context, in *GetSingleKvValueReq, opts ...grpc.CallOption) (*GetSingleKvValueResp, error)
	GetSingleKvMeta(ctx context.Context, in *GetSingleKvValueReq, opts ...grpc.CallOption) (*GetSingleKvMetaResp, error)
	// 获取单个文件配置项
	GetSingleFileContent(ctx context.Context, in *GetSingleFileContentReq, opts ...grpc.CallOption) (Upstream_GetSingleFileContentClient, error)
}

type upstreamClient struct {
	cc grpc.ClientConnInterface
}

func NewUpstreamClient(cc grpc.ClientConnInterface) UpstreamClient {
	return &upstreamClient{cc}
}

func (c *upstreamClient) Handshake(ctx context.Context, in *HandshakeMessage, opts ...grpc.CallOption) (*HandshakeResp, error) {
	out := new(HandshakeResp)
	err := c.cc.Invoke(ctx, Upstream_Handshake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) Messaging(ctx context.Context, in *MessagingMeta, opts ...grpc.CallOption) (*MessagingResp, error) {
	out := new(MessagingResp)
	err := c.cc.Invoke(ctx, Upstream_Messaging_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) Watch(ctx context.Context, in *SideWatchMeta, opts ...grpc.CallOption) (Upstream_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Upstream_ServiceDesc.Streams[0], Upstream_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Upstream_WatchClient interface {
	Recv() (*FeedWatchMessage, error)
	grpc.ClientStream
}

type upstreamWatchClient struct {
	grpc.ClientStream
}

func (x *upstreamWatchClient) Recv() (*FeedWatchMessage, error) {
	m := new(FeedWatchMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *upstreamClient) PullAppFileMeta(ctx context.Context, in *PullAppFileMetaReq, opts ...grpc.CallOption) (*PullAppFileMetaResp, error) {
	out := new(PullAppFileMetaResp)
	err := c.cc.Invoke(ctx, Upstream_PullAppFileMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) GetDownloadURL(ctx context.Context, in *GetDownloadURLReq, opts ...grpc.CallOption) (*GetDownloadURLResp, error) {
	out := new(GetDownloadURLResp)
	err := c.cc.Invoke(ctx, Upstream_GetDownloadURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) PullKvMeta(ctx context.Context, in *PullKvMetaReq, opts ...grpc.CallOption) (*PullKvMetaResp, error) {
	out := new(PullKvMetaResp)
	err := c.cc.Invoke(ctx, Upstream_PullKvMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) GetKvValue(ctx context.Context, in *GetKvValueReq, opts ...grpc.CallOption) (*GetKvValueResp, error) {
	out := new(GetKvValueResp)
	err := c.cc.Invoke(ctx, Upstream_GetKvValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) ListApps(ctx context.Context, in *ListAppsReq, opts ...grpc.CallOption) (*ListAppsResp, error) {
	out := new(ListAppsResp)
	err := c.cc.Invoke(ctx, Upstream_ListApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) AsyncDownload(ctx context.Context, in *AsyncDownloadReq, opts ...grpc.CallOption) (*AsyncDownloadResp, error) {
	out := new(AsyncDownloadResp)
	err := c.cc.Invoke(ctx, Upstream_AsyncDownload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) AsyncDownloadStatus(ctx context.Context, in *AsyncDownloadStatusReq, opts ...grpc.CallOption) (*AsyncDownloadStatusResp, error) {
	out := new(AsyncDownloadStatusResp)
	err := c.cc.Invoke(ctx, Upstream_AsyncDownloadStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) GetSingleKvValue(ctx context.Context, in *GetSingleKvValueReq, opts ...grpc.CallOption) (*GetSingleKvValueResp, error) {
	out := new(GetSingleKvValueResp)
	err := c.cc.Invoke(ctx, Upstream_GetSingleKvValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) GetSingleKvMeta(ctx context.Context, in *GetSingleKvValueReq, opts ...grpc.CallOption) (*GetSingleKvMetaResp, error) {
	out := new(GetSingleKvMetaResp)
	err := c.cc.Invoke(ctx, Upstream_GetSingleKvMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) GetSingleFileContent(ctx context.Context, in *GetSingleFileContentReq, opts ...grpc.CallOption) (Upstream_GetSingleFileContentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Upstream_ServiceDesc.Streams[1], Upstream_GetSingleFileContent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamGetSingleFileContentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Upstream_GetSingleFileContentClient interface {
	Recv() (*SingleFileChunk, error)
	grpc.ClientStream
}

type upstreamGetSingleFileContentClient struct {
	grpc.ClientStream
}

func (x *upstreamGetSingleFileContentClient) Recv() (*SingleFileChunk, error) {
	m := new(SingleFileChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpstreamServer is the server API for Upstream service.
// All implementations should embed UnimplementedUpstreamServer
// for forward compatibility
type UpstreamServer interface {
	// APIs for sidecar.
	Handshake(context.Context, *HandshakeMessage) (*HandshakeResp, error)
	Messaging(context.Context, *MessagingMeta) (*MessagingResp, error)
	Watch(*SideWatchMeta, Upstream_WatchServer) error
	PullAppFileMeta(context.Context, *PullAppFileMetaReq) (*PullAppFileMetaResp, error)
	GetDownloadURL(context.Context, *GetDownloadURLReq) (*GetDownloadURLResp, error)
	PullKvMeta(context.Context, *PullKvMetaReq) (*PullKvMetaResp, error)
	GetKvValue(context.Context, *GetKvValueReq) (*GetKvValueResp, error)
	ListApps(context.Context, *ListAppsReq) (*ListAppsResp, error)
	AsyncDownload(context.Context, *AsyncDownloadReq) (*AsyncDownloadResp, error)
	AsyncDownloadStatus(context.Context, *AsyncDownloadStatusReq) (*AsyncDownloadStatusResp, error)
	// 获取单个kv实例
	GetSingleKvValue(context.Context, *GetSingleKvValueReq) (*GetSingleKvValueResp, error)
	GetSingleKvMeta(context.Context, *GetSingleKvValueReq) (*GetSingleKvMetaResp, error)
	// 获取单个文件配置项
	GetSingleFileContent(*GetSingleFileContentReq, Upstream_GetSingleFileContentServer) error
}

// UnimplementedUpstreamServer should be embedded to have forward compatible implementations.
type UnimplementedUpstreamServer struct {
}

func (UnimplementedUpstreamServer) Handshake(context.Context, *HandshakeMessage) (*HandshakeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedUpstreamServer) Messaging(context.Context, *MessagingMeta) (*MessagingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Messaging not implemented")
}
func (UnimplementedUpstreamServer) Watch(*SideWatchMeta, Upstream_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedUpstreamServer) PullAppFileMeta(context.Context, *PullAppFileMetaReq) (*PullAppFileMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullAppFileMeta not implemented")
}
func (UnimplementedUpstreamServer) GetDownloadURL(context.Context, *GetDownloadURLReq) (*GetDownloadURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadURL not implemented")
}
func (UnimplementedUpstreamServer) PullKvMeta(context.Context, *PullKvMetaReq) (*PullKvMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullKvMeta not implemented")
}
func (UnimplementedUpstreamServer) GetKvValue(context.Context, *GetKvValueReq) (*GetKvValueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKvValue not implemented")
}
func (UnimplementedUpstreamServer) ListApps(context.Context, *ListAppsReq) (*ListAppsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApps not implemented")
}
func (UnimplementedUpstreamServer) AsyncDownload(context.Context, *AsyncDownloadReq) (*AsyncDownloadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsyncDownload not implemented")
}
func (UnimplementedUpstreamServer) AsyncDownloadStatus(context.Context, *AsyncDownloadStatusReq) (*AsyncDownloadStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsyncDownloadStatus not implemented")
}
func (UnimplementedUpstreamServer) GetSingleKvValue(context.Context, *GetSingleKvValueReq) (*GetSingleKvValueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleKvValue not implemented")
}
func (UnimplementedUpstreamServer) GetSingleKvMeta(context.Context, *GetSingleKvValueReq) (*GetSingleKvMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleKvMeta not implemented")
}
func (UnimplementedUpstreamServer) GetSingleFileContent(*GetSingleFileContentReq, Upstream_GetSingleFileContentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSingleFileContent not implemented")
}

// UnsafeUpstreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpstreamServer will
// result in compilation errors.
type UnsafeUpstreamServer interface {
	mustEmbedUnimplementedUpstreamServer()
}

func RegisterUpstreamServer(s grpc.ServiceRegistrar, srv UpstreamServer) {
	s.RegisterService(&Upstream_ServiceDesc, srv)
}

func _Upstream_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_Handshake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).Handshake(ctx, req.(*HandshakeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_Messaging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagingMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).Messaging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_Messaging_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).Messaging(ctx, req.(*MessagingMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SideWatchMeta)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpstreamServer).Watch(m, &upstreamWatchServer{stream})
}

type Upstream_WatchServer interface {
	Send(*FeedWatchMessage) error
	grpc.ServerStream
}

type upstreamWatchServer struct {
	grpc.ServerStream
}

func (x *upstreamWatchServer) Send(m *FeedWatchMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Upstream_PullAppFileMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullAppFileMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).PullAppFileMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_PullAppFileMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).PullAppFileMeta(ctx, req.(*PullAppFileMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_GetDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).GetDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_GetDownloadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).GetDownloadURL(ctx, req.(*GetDownloadURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_PullKvMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullKvMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).PullKvMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_PullKvMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).PullKvMeta(ctx, req.(*PullKvMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_GetKvValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKvValueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).GetKvValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_GetKvValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).GetKvValue(ctx, req.(*GetKvValueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_ListApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).ListApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_ListApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).ListApps(ctx, req.(*ListAppsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_AsyncDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsyncDownloadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).AsyncDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_AsyncDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).AsyncDownload(ctx, req.(*AsyncDownloadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_AsyncDownloadStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsyncDownloadStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).AsyncDownloadStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_AsyncDownloadStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).AsyncDownloadStatus(ctx, req.(*AsyncDownloadStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_GetSingleKvValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleKvValueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).GetSingleKvValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_GetSingleKvValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).GetSingleKvValue(ctx, req.(*GetSingleKvValueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_GetSingleKvMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleKvValueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).GetSingleKvMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_GetSingleKvMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).GetSingleKvMeta(ctx, req.(*GetSingleKvValueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_GetSingleFileContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSingleFileContentReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpstreamServer).GetSingleFileContent(m, &upstreamGetSingleFileContentServer{stream})
}

type Upstream_GetSingleFileContentServer interface {
	Send(*SingleFileChunk) error
	grpc.ServerStream
}

type upstreamGetSingleFileContentServer struct {
	grpc.ServerStream
}

func (x *upstreamGetSingleFileContentServer) Send(m *SingleFileChunk) error {
	return x.ServerStream.SendMsg(m)
}

// Upstream_ServiceDesc is the grpc.ServiceDesc for Upstream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upstream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbfs.Upstream",
	HandlerType: (*UpstreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Upstream_Handshake_Handler,
		},
		{
			MethodName: "Messaging",
			Handler:    _Upstream_Messaging_Handler,
		},
		{
			MethodName: "PullAppFileMeta",
			Handler:    _Upstream_PullAppFileMeta_Handler,
		},
		{
			MethodName: "GetDownloadURL",
			Handler:    _Upstream_GetDownloadURL_Handler,
		},
		{
			MethodName: "PullKvMeta",
			Handler:    _Upstream_PullKvMeta_Handler,
		},
		{
			MethodName: "GetKvValue",
			Handler:    _Upstream_GetKvValue_Handler,
		},
		{
			MethodName: "ListApps",
			Handler:    _Upstream_ListApps_Handler,
		},
		{
			MethodName: "AsyncDownload",
			Handler:    _Upstream_AsyncDownload_Handler,
		},
		{
			MethodName: "AsyncDownloadStatus",
			Handler:    _Upstream_AsyncDownloadStatus_Handler,
		},
		{
			MethodName: "GetSingleKvValue",
			Handler:    _Upstream_GetSingleKvValue_Handler,
		},
		{
			MethodName: "GetSingleKvMeta",
			Handler:    _Upstream_GetSingleKvMeta_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Upstream_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSingleFileContent",
			Handler:       _Upstream_GetSingleFileContent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "feed_server.proto",
}
