package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/syumai/shop.syum.ai/api"
	"github.com/syumai/shop.syum.ai/api/gen/shop/v1/shopv1connect"
)

type testFunc func(t *testing.T, ctx context.Context, client shopv1connect.ShopServiceClient)

func RunTest(t *testing.T, testFunc testFunc) {
	mux := http.NewServeMux()
	mux.Handle(shopv1connect.NewShopServiceHandler(
		&api.ShopServer{},
	))
	server := httptest.NewServer(mux)
	defer server.Close()

	client := shopv1connect.NewShopServiceClient(
		server.Client(),
		server.URL,
	)

	testFunc(t, context.Background(), client)
}
