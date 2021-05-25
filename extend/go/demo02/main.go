// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

// 发射倒计时 10 9 8 ... 按任意键停止发射
func main() {
	// 1 select 每次判断是否接收到abort信号
	// lanch1()
	// 2 cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go lanch2(ctx)
	os.Stdin.Read(make([]byte, 1))
}

func lanch1() {
	fmt.Println("Commencing countdown.")
	tick := time.NewTicker(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-abort:
			fmt.Println("stop ...")
			tick.Stop()
			return
		case <-tick.C:
			fmt.Println(countdown)
		}
	}
}

func lanch2(ctx context.Context) {
	fmt.Println("Commencing countdown.")
	tick := time.NewTicker(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ctx.Done():
			fmt.Println("stop ...")
			tick.Stop()
			return
		case <-tick.C:
			fmt.Println(countdown)
		}
	}
}
