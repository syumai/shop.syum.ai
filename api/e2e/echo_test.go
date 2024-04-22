package e2e

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	v1 "github.com/syumai/shop.syum.ai/api/gen/shop/v1"
	"github.com/syumai/shop.syum.ai/api/gen/shop/v1/shopv1connect"
)

func TestEcho(t *testing.T) {
	t.Parallel()
	RunTest(t, func(t *testing.T, ctx context.Context, client shopv1connect.ShopServiceClient) {
		want := "Hello, world!"
		res, err := client.Echo(ctx, connect.NewRequest(&v1.EchoRequest{Msg: want}))
		if err != nil {
			t.Fatalf("want err: nil, got: %v", err)
		}
		got := res.Msg.GetMsg()
		if want != got {
			t.Fatalf("want: %s, got: %s", want, got)
		}
	})
}
