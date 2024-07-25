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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	do(ctx)
}
