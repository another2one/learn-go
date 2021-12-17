package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var _ Mux = (*mux)(nil)

type Mux interface {
}

type mux struct {
	name string
}

func (m mux) s() {
	fmt.Println("sssssssss")
}

func main() {
	// c 没有缓存的化将被释放
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// 正确用法
	go func() {
		// Block until a signal is received.
		time.Sleep(time.Second * 5)
	}()

	s := <-c
	fmt.Println("Got signal:", s)
}
