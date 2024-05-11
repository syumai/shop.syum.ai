package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/syumai/shop.syum.ai/api"
	"github.com/syumai/shop.syum.ai/api/gen/shop/v1/shopv1connect"
)

type Server struct {
	srv *http.Server
}

func New(port string) *Server {
	svc := &api.ShopService{}
	mux := http.NewServeMux()
	path, handler := shopv1connect.NewShopServiceHandler(svc)
	mux.Handle(path, handler)
	addr := fmt.Sprintf(":%s", port)
	srv := http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}
	return &Server{
		srv: &srv,
	}
}

func (s *Server) Start() error {
	fmt.Printf("listening server on %s\n", s.srv.Addr)
	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
