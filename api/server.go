package api

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/syumai/shop.syum.ai/api/gen/shop/v1"
	"github.com/syumai/shop.syum.ai/api/gen/shop/v1/shopv1connect"
	"github.com/syumai/shop.syum.ai/api/model"
)

type ShopServer struct {
	queries *model.Queries
}

func NewShopServer(queries *model.Queries) *ShopServer {
	return &ShopServer{
		queries: queries,
	}
}

var _ shopv1connect.ShopServiceHandler = &ShopServer{}

func (s *ShopServer) Echo(
	_ context.Context,
	req *connect.Request[v1.EchoRequest],
) (*connect.Response[v1.EchoResponse], error) {
	res := connect.NewResponse(&v1.EchoResponse{
		Msg: req.Msg.Msg,
	})
	return res, nil
}

func (s *ShopServer) CreateProduct(
	_ context.Context,
	req *connect.Request[v1.CreateProductRequest],
) (*connect.Response[v1.CreateProductResponse], error) {

	// DB 呼び出しを描いていく

	res := connect.NewResponse(&v1.CreateProductResponse{
		Product: &v1.Product{
			Id:            1,
			Name:          "",
			Description:   "",
			ProductStatus: v1.ProductStatus_PRODUCT_STATUS_UNSPECIFIED,
			CreatedTime:   nil,
			UpdatedTime:   nil,
			ReservedTime:  nil,
		},
	})
	return res, nil
}
