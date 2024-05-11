package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/syumai/shop.syum.ai/api/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s := server.New(port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(s.Start)
	eg.Go(func() error {
		<-egCtx.Done()
		return s.Stop()
	})

	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
