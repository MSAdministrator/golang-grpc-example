// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: example.proto

package pb

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
	ExampleEnrichment_ExampleEnrich_FullMethodName = "/pb.ExampleEnrichment/ExampleEnrich"
)

// ExampleEnrichmentClient is the client API for ExampleEnrichment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleEnrichmentClient interface {
	// ExampleEnrichment receives a batch of indicators and returns a batch of enriched indicators
	ExampleEnrich(ctx context.Context, in *ExampleEnrichmentRequest, opts ...grpc.CallOption) (*ExampleEnrichmentResponse, error)
}

type exampleEnrichmentClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleEnrichmentClient(cc grpc.ClientConnInterface) ExampleEnrichmentClient {
	return &exampleEnrichmentClient{cc}
}

func (c *exampleEnrichmentClient) ExampleEnrich(ctx context.Context, in *ExampleEnrichmentRequest, opts ...grpc.CallOption) (*ExampleEnrichmentResponse, error) {
	out := new(ExampleEnrichmentResponse)
	err := c.cc.Invoke(ctx, ExampleEnrichment_ExampleEnrich_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleEnrichmentServer is the server API for ExampleEnrichment service.
// All implementations must embed UnimplementedExampleEnrichmentServer
// for forward compatibility
type ExampleEnrichmentServer interface {
	// ExampleEnrichment receives a batch of indicators and returns a batch of enriched indicators
	ExampleEnrich(context.Context, *ExampleEnrichmentRequest) (*ExampleEnrichmentResponse, error)
	mustEmbedUnimplementedExampleEnrichmentServer()
}

// UnimplementedExampleEnrichmentServer must be embedded to have forward compatible implementations.
type UnimplementedExampleEnrichmentServer struct {
}

func (UnimplementedExampleEnrichmentServer) ExampleEnrich(context.Context, *ExampleEnrichmentRequest) (*ExampleEnrichmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleEnrich not implemented")
}
func (UnimplementedExampleEnrichmentServer) mustEmbedUnimplementedExampleEnrichmentServer() {}

// UnsafeExampleEnrichmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleEnrichmentServer will
// result in compilation errors.
type UnsafeExampleEnrichmentServer interface {
	mustEmbedUnimplementedExampleEnrichmentServer()
}

func RegisterExampleEnrichmentServer(s grpc.ServiceRegistrar, srv ExampleEnrichmentServer) {
	s.RegisterService(&ExampleEnrichment_ServiceDesc, srv)
}

func _ExampleEnrichment_ExampleEnrich_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleEnrichmentServer).ExampleEnrich(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleEnrichment_ExampleEnrich_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleEnrichmentServer).ExampleEnrich(ctx, req.(*ExampleEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleEnrichment_ServiceDesc is the grpc.ServiceDesc for ExampleEnrichment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleEnrichment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ExampleEnrichment",
	HandlerType: (*ExampleEnrichmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleEnrich",
			Handler:    _ExampleEnrichment_ExampleEnrich_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
