package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func testCause(seconds time.Duration) {
	ctx, cancel := context.WithCancelCause(context.Background())
	ctx, cancel1 := context.WithTimeoutCause(ctx, 3*time.Second, errors.New("timeout cause"))
	defer cancel1()
	go func() {
		time.Sleep(seconds)
		cancel(errors.New("cancel cause"))
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println(context.Cause(ctx))
			return
		default:
			log.Println("sleep")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("========= cancel first =============")
	testCause(2 * time.Second)
	log.Println("\n ========= timeout first =============")
	testCause(4 * time.Second)
}
