package main

import (
	"context"
	"log"
	"time"
)

func do(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		default:
			log.Println("sleep")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	// 这里一般是指定时间点时用 这里为了测试用的+2s
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	do(ctx)
}
