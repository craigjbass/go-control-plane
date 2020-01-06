// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/secret/v3alpha/sds.proto

package envoy_service_secret_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3alpha"
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

type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fedd920dfdba982, []int{0}
}

func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdsDummy.Unmarshal(m, b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
}
func (m *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(m, src)
}
func (m *SdsDummy) XXX_Size() int {
	return xxx_messageInfo_SdsDummy.Size(m)
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.secret.v3alpha.SdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/secret/v3alpha/sds.proto", fileDescriptor_3fedd920dfdba982)
}

var fileDescriptor_3fedd920dfdba982 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x4a, 0xe3, 0x40,
	0x1c, 0xc7, 0x77, 0x76, 0x97, 0xdd, 0x65, 0xe8, 0x5e, 0x72, 0xd8, 0x43, 0x28, 0xcb, 0x6e, 0x45,
	0x91, 0x0a, 0x33, 0xda, 0x62, 0x0b, 0x15, 0x2f, 0xa5, 0x78, 0x2e, 0xe6, 0x01, 0x64, 0x4c, 0x7e,
	0xb4, 0xc1, 0x74, 0x26, 0xce, 0x6f, 0x12, 0x1a, 0xf0, 0x54, 0x10, 0x8a, 0x57, 0xc1, 0x83, 0x47,
	0xf1, 0x55, 0x7c, 0x02, 0x5f, 0xc1, 0x37, 0xf0, 0x05, 0xa4, 0x99, 0x34, 0xb6, 0x8a, 0x4a, 0x2f,
	0x5e, 0x67, 0x3e, 0xbf, 0xef, 0x1f, 0xf8, 0xd2, 0x0d, 0x90, 0xa9, 0xca, 0x38, 0x82, 0x4e, 0x43,
	0x1f, 0x38, 0x82, 0xaf, 0xc1, 0xf0, 0xb4, 0x29, 0xa2, 0x78, 0x28, 0x38, 0x06, 0xc8, 0x62, 0xad,
	0x8c, 0x72, 0xaa, 0x39, 0xc7, 0x0a, 0x8e, 0x59, 0x8e, 0x15, 0x9c, 0xcb, 0x97, 0x55, 0x82, 0x10,
	0x7d, 0x95, 0x82, 0xce, 0x4a, 0xa1, 0xf2, 0xc5, 0xca, 0xb9, 0xd5, 0x81, 0x52, 0x83, 0x08, 0xb8,
	0x88, 0x43, 0x2e, 0xa4, 0x54, 0x46, 0x98, 0x50, 0xc9, 0xc2, 0xcc, 0xfd, 0x9f, 0x04, 0xb1, 0x58,
	0x7c, 0xe7, 0x29, 0x68, 0x0c, 0x95, 0x0c, 0xe5, 0xa0, 0x40, 0xfe, 0x59, 0xc7, 0x45, 0x46, 0x03,
	0xaa, 0x44, 0xfb, 0x60, 0x89, 0x5a, 0x8b, 0xfe, 0xf2, 0x02, 0xec, 0x25, 0xa3, 0x51, 0xd6, 0xa9,
	0x5f, 0xdf, 0x4d, 0xff, 0xae, 0xd3, 0xb5, 0xe5, 0x12, 0xcf, 0xa1, 0xd2, 0x06, 0x9b, 0xb3, 0x8d,
	0xc9, 0x77, 0xfa, 0xc7, 0xcb, 0xeb, 0xf5, 0xe6, 0xff, 0x9e, 0x3d, 0x70, 0xce, 0x09, 0xad, 0xf4,
	0x20, 0x32, 0xc2, 0xfe, 0xa3, 0xd3, 0x62, 0x6f, 0x2a, 0xda, 0xe2, 0x2c, 0xc7, 0x4b, 0xa1, 0x43,
	0x38, 0x4d, 0x00, 0x8d, 0xdb, 0x5e, 0xf9, 0x0e, 0x63, 0x25, 0x11, 0x6a, 0x5f, 0x36, 0xc9, 0x36,
	0x71, 0xce, 0xe8, 0x6f, 0xcf, 0x68, 0x10, 0xa3, 0x79, 0x8e, 0x9d, 0x8f, 0xf5, 0x5e, 0x46, 0x68,
	0xac, 0x72, 0xb2, 0xe4, 0x7e, 0x43, 0x68, 0xe5, 0x00, 0x8c, 0x3f, 0xfc, 0x64, 0xf7, 0xad, 0xc9,
	0xfd, 0xc3, 0xe5, 0xd7, 0x6a, 0xcd, 0x7d, 0xbd, 0xaa, 0x8e, 0xdd, 0x23, 0xe6, 0xc4, 0xb7, 0x0e,
	0xa9, 0xbb, 0xfb, 0x17, 0xb7, 0x57, 0x8f, 0x3f, 0xdb, 0x74, 0xd7, 0xfa, 0xc0, 0xd8, 0x80, 0xc4,
	0x7c, 0x62, 0x46, 0x0b, 0x89, 0xb1, 0xd2, 0xe6, 0x08, 0x95, 0x7f, 0x02, 0x06, 0x99, 0x89, 0xb0,
	0xb4, 0xb5, 0x9d, 0xba, 0x7b, 0xb4, 0x1e, 0x2a, 0x9b, 0x31, 0xd6, 0x6a, 0x9c, 0xb1, 0xf7, 0xe6,
	0xdf, 0x9d, 0x0d, 0xad, 0x3f, 0x1b, 0x5d, 0x9f, 0x4c, 0x09, 0x39, 0xfe, 0x91, 0x0f, 0xb0, 0xf9,
	0x14, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x34, 0xec, 0x13, 0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error)
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.secret.v3alpha.SecretDiscoveryService/DeltaSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceDeltaSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_DeltaSecretsClient interface {
	Send(*v3alpha.DeltaDiscoveryRequest) error
	Recv() (*v3alpha.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceDeltaSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Send(m *v3alpha.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Recv() (*v3alpha.DeltaDiscoveryResponse, error) {
	m := new(v3alpha.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[1], "/envoy.service.secret.v3alpha.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v3alpha.DiscoveryRequest) error
	Recv() (*v3alpha.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v3alpha.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v3alpha.DiscoveryResponse, error) {
	m := new(v3alpha.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error) {
	out := new(v3alpha.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.secret.v3alpha.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	DeltaSecrets(SecretDiscoveryService_DeltaSecretsServer) error
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error)
}

// UnimplementedSecretDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretDiscoveryServiceServer struct {
}

func (*UnimplementedSecretDiscoveryServiceServer) DeltaSecrets(srv SecretDiscoveryService_DeltaSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) StreamSecrets(srv SecretDiscoveryService_StreamSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) FetchSecrets(ctx context.Context, req *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSecrets not implemented")
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_DeltaSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).DeltaSecrets(&secretDiscoveryServiceDeltaSecretsServer{stream})
}

type SecretDiscoveryService_DeltaSecretsServer interface {
	Send(*v3alpha.DeltaDiscoveryResponse) error
	Recv() (*v3alpha.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceDeltaSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Send(m *v3alpha.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Recv() (*v3alpha.DeltaDiscoveryRequest, error) {
	m := new(v3alpha.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v3alpha.DiscoveryResponse) error
	Recv() (*v3alpha.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v3alpha.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v3alpha.DiscoveryRequest, error) {
	m := new(v3alpha.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.secret.v3alpha.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v3alpha.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.secret.v3alpha.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaSecrets",
			Handler:       _SecretDiscoveryService_DeltaSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/secret/v3alpha/sds.proto",
}
