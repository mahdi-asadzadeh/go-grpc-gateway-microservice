// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: pb/product.proto

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (ProductService_ListProductClient, error)
	DetailProduct(ctx context.Context, in *DetailProductRequest, opts ...grpc.CallOption) (*DetailProductResponse, error)
	ImagesProduct(ctx context.Context, in *ImagesProductRequest, opts ...grpc.CallOption) (ProductService_ImagesProductClient, error)
	UploadImagesProduct(ctx context.Context, opts ...grpc.CallOption) (ProductService_UploadImagesProductClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (ProductService_ListProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/pb.ProductService/ListProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceListProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_ListProductClient interface {
	Recv() (*ListProductResponse, error)
	grpc.ClientStream
}

type productServiceListProductClient struct {
	grpc.ClientStream
}

func (x *productServiceListProductClient) Recv() (*ListProductResponse, error) {
	m := new(ListProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) DetailProduct(ctx context.Context, in *DetailProductRequest, opts ...grpc.CallOption) (*DetailProductResponse, error) {
	out := new(DetailProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductService/DetailProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ImagesProduct(ctx context.Context, in *ImagesProductRequest, opts ...grpc.CallOption) (ProductService_ImagesProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[1], "/pb.ProductService/ImagesProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceImagesProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_ImagesProductClient interface {
	Recv() (*ImagesProductResponse, error)
	grpc.ClientStream
}

type productServiceImagesProductClient struct {
	grpc.ClientStream
}

func (x *productServiceImagesProductClient) Recv() (*ImagesProductResponse, error) {
	m := new(ImagesProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) UploadImagesProduct(ctx context.Context, opts ...grpc.CallOption) (ProductService_UploadImagesProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[2], "/pb.ProductService/UploadImagesProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceUploadImagesProductClient{stream}
	return x, nil
}

type ProductService_UploadImagesProductClient interface {
	Send(*UploadImagesProductRequest) error
	CloseAndRecv() (*UploadImagesProductResponse, error)
	grpc.ClientStream
}

type productServiceUploadImagesProductClient struct {
	grpc.ClientStream
}

func (x *productServiceUploadImagesProductClient) Send(m *UploadImagesProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceUploadImagesProductClient) CloseAndRecv() (*UploadImagesProductResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImagesProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	ListProduct(*ListProductRequest, ProductService_ListProductServer) error
	DetailProduct(context.Context, *DetailProductRequest) (*DetailProductResponse, error)
	ImagesProduct(*ImagesProductRequest, ProductService_ImagesProductServer) error
	UploadImagesProduct(ProductService_UploadImagesProductServer) error
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) ListProduct(*ListProductRequest, ProductService_ListProductServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedProductServiceServer) DetailProduct(context.Context, *DetailProductRequest) (*DetailProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailProduct not implemented")
}
func (UnimplementedProductServiceServer) ImagesProduct(*ImagesProductRequest, ProductService_ImagesProductServer) error {
	return status.Errorf(codes.Unimplemented, "method ImagesProduct not implemented")
}
func (UnimplementedProductServiceServer) UploadImagesProduct(ProductService_UploadImagesProductServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImagesProduct not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListProductRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).ListProduct(m, &productServiceListProductServer{stream})
}

type ProductService_ListProductServer interface {
	Send(*ListProductResponse) error
	grpc.ServerStream
}

type productServiceListProductServer struct {
	grpc.ServerStream
}

func (x *productServiceListProductServer) Send(m *ListProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_DetailProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DetailProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/DetailProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DetailProduct(ctx, req.(*DetailProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ImagesProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImagesProductRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).ImagesProduct(m, &productServiceImagesProductServer{stream})
}

type ProductService_ImagesProductServer interface {
	Send(*ImagesProductResponse) error
	grpc.ServerStream
}

type productServiceImagesProductServer struct {
	grpc.ServerStream
}

func (x *productServiceImagesProductServer) Send(m *ImagesProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_UploadImagesProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).UploadImagesProduct(&productServiceUploadImagesProductServer{stream})
}

type ProductService_UploadImagesProductServer interface {
	SendAndClose(*UploadImagesProductResponse) error
	Recv() (*UploadImagesProductRequest, error)
	grpc.ServerStream
}

type productServiceUploadImagesProductServer struct {
	grpc.ServerStream
}

func (x *productServiceUploadImagesProductServer) SendAndClose(m *UploadImagesProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceUploadImagesProductServer) Recv() (*UploadImagesProductRequest, error) {
	m := new(UploadImagesProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "DetailProduct",
			Handler:    _ProductService_DetailProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProduct",
			Handler:       _ProductService_ListProduct_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ImagesProduct",
			Handler:       _ProductService_ImagesProduct_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadImagesProduct",
			Handler:       _ProductService_UploadImagesProduct_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pb/product.proto",
}
