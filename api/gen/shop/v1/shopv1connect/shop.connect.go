// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: shop/v1/shop.proto

package shopv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/syumai/shop.syum.ai/api/gen/shop/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ShopServiceName is the fully-qualified name of the ShopService service.
	ShopServiceName = "shop.v1.ShopService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ShopServiceEchoProcedure is the fully-qualified name of the ShopService's Echo RPC.
	ShopServiceEchoProcedure = "/shop.v1.ShopService/Echo"
	// ShopServiceCreateProductProcedure is the fully-qualified name of the ShopService's CreateProduct
	// RPC.
	ShopServiceCreateProductProcedure = "/shop.v1.ShopService/CreateProduct"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	shopServiceServiceDescriptor             = v1.File_shop_v1_shop_proto.Services().ByName("ShopService")
	shopServiceEchoMethodDescriptor          = shopServiceServiceDescriptor.Methods().ByName("Echo")
	shopServiceCreateProductMethodDescriptor = shopServiceServiceDescriptor.Methods().ByName("CreateProduct")
)

// ShopServiceClient is a client for the shop.v1.ShopService service.
type ShopServiceClient interface {
	Echo(context.Context, *connect.Request[v1.EchoRequest]) (*connect.Response[v1.EchoResponse], error)
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error)
}

// NewShopServiceClient constructs a client for the shop.v1.ShopService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewShopServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ShopServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &shopServiceClient{
		echo: connect.NewClient[v1.EchoRequest, v1.EchoResponse](
			httpClient,
			baseURL+ShopServiceEchoProcedure,
			connect.WithSchema(shopServiceEchoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProduct: connect.NewClient[v1.CreateProductRequest, v1.CreateProductResponse](
			httpClient,
			baseURL+ShopServiceCreateProductProcedure,
			connect.WithSchema(shopServiceCreateProductMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// shopServiceClient implements ShopServiceClient.
type shopServiceClient struct {
	echo          *connect.Client[v1.EchoRequest, v1.EchoResponse]
	createProduct *connect.Client[v1.CreateProductRequest, v1.CreateProductResponse]
}

// Echo calls shop.v1.ShopService.Echo.
func (c *shopServiceClient) Echo(ctx context.Context, req *connect.Request[v1.EchoRequest]) (*connect.Response[v1.EchoResponse], error) {
	return c.echo.CallUnary(ctx, req)
}

// CreateProduct calls shop.v1.ShopService.CreateProduct.
func (c *shopServiceClient) CreateProduct(ctx context.Context, req *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error) {
	return c.createProduct.CallUnary(ctx, req)
}

// ShopServiceHandler is an implementation of the shop.v1.ShopService service.
type ShopServiceHandler interface {
	Echo(context.Context, *connect.Request[v1.EchoRequest]) (*connect.Response[v1.EchoResponse], error)
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error)
}

// NewShopServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewShopServiceHandler(svc ShopServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	shopServiceEchoHandler := connect.NewUnaryHandler(
		ShopServiceEchoProcedure,
		svc.Echo,
		connect.WithSchema(shopServiceEchoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	shopServiceCreateProductHandler := connect.NewUnaryHandler(
		ShopServiceCreateProductProcedure,
		svc.CreateProduct,
		connect.WithSchema(shopServiceCreateProductMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/shop.v1.ShopService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ShopServiceEchoProcedure:
			shopServiceEchoHandler.ServeHTTP(w, r)
		case ShopServiceCreateProductProcedure:
			shopServiceCreateProductHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedShopServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedShopServiceHandler struct{}

func (UnimplementedShopServiceHandler) Echo(context.Context, *connect.Request[v1.EchoRequest]) (*connect.Response[v1.EchoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("shop.v1.ShopService.Echo is not implemented"))
}

func (UnimplementedShopServiceHandler) CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("shop.v1.ShopService.CreateProduct is not implemented"))
}
