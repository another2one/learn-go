package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	server := http.Server{Addr: ":8081"}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("hello world"))
			if err != nil {
				fmt.Printf("http error: %v", err)
			}
		})
		fmt.Println("start server: 127.0.0.1:8081")
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Shutdown(context.Background())
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("exit with: %v \n", err)
	}
}
