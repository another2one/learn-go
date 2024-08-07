package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func do(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("cancel receive")
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("cancel send")
		cancel()
	}()
	do(ctx)
}
