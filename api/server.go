package main

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/syumai/shop.syum.ai/api/gen/shop/v1"
	"github.com/syumai/shop.syum.ai/api/gen/shop/v1/shopv1connect"
)

type ShopServer struct{}

var _ shopv1connect.ShopServiceHandler = &ShopServer{}

func (s *ShopServer) Echo(
	_ context.Context,
	req *connect.Request[v1.EchoRequest],
) (*connect.Response[v1.EchoResponse], error) {
	res := connect.NewResponse(&v1.EchoResponse{
		Msg: req.Msg,
	})
	return res, nil
}
