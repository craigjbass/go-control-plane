// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdsDummy) Reset()         { *m = CdsDummy{} }
func (m *CdsDummy) String() string { return proto.CompactTextString(m) }
func (*CdsDummy) ProtoMessage()    {}
func (*CdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e73f50fbb1daa302, []int{0}
}

func (m *CdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdsDummy.Unmarshal(m, b)
}
func (m *CdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdsDummy.Marshal(b, m, deterministic)
}
func (m *CdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdsDummy.Merge(m, src)
}
func (m *CdsDummy) XXX_Size() int {
	return xxx_messageInfo_CdsDummy.Size(m)
}
func (m *CdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_CdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_CdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CdsDummy)(nil), "envoy.api.v2.CdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/cds.proto", fileDescriptor_e73f50fbb1daa302) }

var fileDescriptor_e73f50fbb1daa302 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xff, 0xe9, 0x9f, 0x2f, 0x59, 0x94, 0x21, 0x42, 0x05, 0x59, 0xa5, 0x45, 0x05, 0x09,
	0xc4, 0x90, 0xa0, 0x76, 0xeb, 0xd8, 0x56, 0xcc, 0x11, 0x1d, 0x58, 0x31, 0xc9, 0x55, 0x6b, 0x29,
	0xc9, 0x35, 0xb6, 0x13, 0x91, 0x95, 0x09, 0xb1, 0x30, 0x20, 0x21, 0x1e, 0x80, 0x37, 0xe2, 0x15,
	0x58, 0x59, 0x60, 0x47, 0xa8, 0xf9, 0x28, 0x18, 0x04, 0x13, 0x6b, 0x7e, 0xe7, 0x9c, 0x1c, 0xdf,
	0x43, 0x1a, 0x10, 0xa7, 0x98, 0xb9, 0x4c, 0x70, 0x37, 0xed, 0xba, 0x7e, 0xa0, 0x1c, 0x21, 0x51,
	0xa3, 0xbd, 0x9a, 0x7f, 0x77, 0x98, 0xe0, 0x4e, 0xda, 0xa5, 0x4d, 0x43, 0x15, 0x70, 0xe5, 0x63,
	0x0a, 0x32, 0x2b, 0xb4, 0xb4, 0x39, 0x41, 0x9c, 0x84, 0x90, 0x63, 0x16, 0xc7, 0xa8, 0x99, 0xe6,
	0x18, 0x97, 0x49, 0x74, 0xbb, 0xf4, 0x7e, 0x00, 0x57, 0x82, 0xc2, 0x44, 0xfa, 0x50, 0x2a, 0x5a,
	0x49, 0x20, 0x98, 0x21, 0x88, 0xf8, 0x44, 0x32, 0x5d, 0x71, 0x6a, 0x76, 0x0c, 0x13, 0xa5, 0x41,
	0x16, 0xac, 0x43, 0xc8, 0xca, 0x30, 0x50, 0xa3, 0x24, 0x8a, 0xb2, 0xee, 0x73, 0x8d, 0x6c, 0x0c,
	0x0b, 0x3a, 0xaa, 0x2a, 0x8e, 0x41, 0xa6, 0xdc, 0x07, 0xfb, 0x84, 0xac, 0x8d, 0xb5, 0x04, 0x16,
	0x95, 0x02, 0x65, 0xb7, 0x9c, 0xcf, 0x4f, 0x74, 0xe6, 0x8e, 0x63, 0x38, 0x4f, 0x40, 0x69, 0xda,
	0xfe, 0x91, 0x2b, 0x81, 0xb1, 0x82, 0xce, 0xbf, 0x7d, 0xeb, 0xd0, 0xb2, 0x4f, 0x49, 0x7d, 0x04,
	0xa1, 0x66, 0xf3, 0xdc, 0x9d, 0x2f, 0xbe, 0x19, 0xfc, 0x16, 0xbe, 0xfb, 0xbb, 0xc8, 0xf8, 0x43,
	0x46, 0xea, 0x47, 0xa0, 0xfd, 0xe9, 0xdf, 0x35, 0xdf, 0xbb, 0x7c, 0x7c, 0xba, 0xad, 0x6d, 0x76,
	0x1a, 0xc6, 0x9a, 0xfd, 0xf2, 0xb2, 0x2a, 0xa7, 0xff, 0xfb, 0xd6, 0x01, 0x6d, 0x5e, 0x3f, 0xdc,
	0xbd, 0x2e, 0x37, 0xc8, 0xba, 0x11, 0x58, 0x16, 0x19, 0x78, 0x2f, 0xf7, 0x6f, 0x37, 0x8b, 0x6d,
	0x7b, 0xab, 0xa0, 0xaa, 0x38, 0xb5, 0x53, 0x0d, 0x94, 0xf6, 0x58, 0x28, 0xa6, 0x8c, 0x50, 0x8e,
	0x45, 0x21, 0x21, 0xf1, 0x22, 0x33, 0xba, 0x0d, 0x66, 0xe3, 0x79, 0xb3, 0x21, 0x3d, 0xeb, 0xca,
	0xb2, 0xbc, 0x85, 0xb3, 0xa5, 0x7c, 0xd6, 0xde, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xf9,
	0x75, 0x8a, 0x98, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterDiscoveryServiceClient is the client API for ClusterDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterDiscoveryServiceClient interface {
	StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error)
	DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error)
	FetchClusters(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type clusterDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterDiscoveryServiceClient(cc *grpc.ClientConn) ClusterDiscoveryServiceClient {
	return &clusterDiscoveryServiceClient{cc}
}

func (c *clusterDiscoveryServiceClient) StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ClusterDiscoveryService/StreamClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceStreamClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_StreamClustersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceStreamClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceStreamClustersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.ClusterDiscoveryService/DeltaClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceDeltaClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_DeltaClustersClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceDeltaClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) FetchClusters(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ClusterDiscoveryService/FetchClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterDiscoveryServiceServer is the server API for ClusterDiscoveryService service.
type ClusterDiscoveryServiceServer interface {
	StreamClusters(ClusterDiscoveryService_StreamClustersServer) error
	DeltaClusters(ClusterDiscoveryService_DeltaClustersServer) error
	FetchClusters(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedClusterDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterDiscoveryServiceServer struct {
}

func (*UnimplementedClusterDiscoveryServiceServer) StreamClusters(srv ClusterDiscoveryService_StreamClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) DeltaClusters(srv ClusterDiscoveryService_DeltaClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) FetchClusters(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClusters not implemented")
}

func RegisterClusterDiscoveryServiceServer(s *grpc.Server, srv ClusterDiscoveryServiceServer) {
	s.RegisterService(&_ClusterDiscoveryService_serviceDesc, srv)
}

func _ClusterDiscoveryService_StreamClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).StreamClusters(&clusterDiscoveryServiceStreamClustersServer{stream})
}

type ClusterDiscoveryService_StreamClustersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceStreamClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceStreamClustersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_DeltaClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).DeltaClusters(&clusterDiscoveryServiceDeltaClustersServer{stream})
}

type ClusterDiscoveryService_DeltaClustersServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceDeltaClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_FetchClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ClusterDiscoveryService/FetchClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ClusterDiscoveryService",
	HandlerType: (*ClusterDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClusters",
			Handler:    _ClusterDiscoveryService_FetchClusters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClusters",
			Handler:       _ClusterDiscoveryService_StreamClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaClusters",
			Handler:       _ClusterDiscoveryService_DeltaClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/cds.proto",
}
