package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func do(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println(context.Cause(ctx))
			return
		default:
			log.Println("sleep")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	ctx, cancel := context.WithTimeoutCause(context.Background(), 2*time.Second, errors.New("cause timeout"))
	defer cancel()
	do(ctx)
}
