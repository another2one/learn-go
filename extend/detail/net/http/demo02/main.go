package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 平滑停止 - 监听程序退出信号，调用 Server.Shutdown 等程序执行完再退出
func main() {
	s := http.NewServeMux()
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		log.Println(w, "Hello world!")
	})
	server := &http.Server{
		Addr:    ":8090",
		Handler: s,
	}
	go server.ListenAndServe()

	listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		log.Println("notify sigs")
		httpSrv.Shutdown(ctx)
		log.Println("http shutdown")
	}
}
