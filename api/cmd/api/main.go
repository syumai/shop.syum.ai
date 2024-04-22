package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/syumai/shop.syum.ai/api"
	"github.com/syumai/shop.syum.ai/api/gen/shop/v1/shopv1connect"
	"github.com/syumai/shop.syum.ai/api/model"
)

func main() {
	db, err := sql.Open("sqlite3", "file:db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	queries := model.New(db)
	srv := api.NewShopServer(queries)

	mux := http.NewServeMux()
	path, handler := shopv1connect.NewShopServiceHandler(srv)
	mux.Handle(path, handler)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	s := http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.Shutdown(ctx)
	}()

	fmt.Printf("listening server on %s\n", addr)
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
