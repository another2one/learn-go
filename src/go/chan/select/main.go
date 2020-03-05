package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string, 10)
	go func () {
		for i := 0; i < 10 ; i++ {
			chan1<- "chan1: hello" + fmt.Sprintf("%d", i)
		}
	}()

	chan2 := make(chan string, 10)
	go func () {
		for i := 0; i < 10 ; i++ {
			chan2<- "chan2: hello" + fmt.Sprintf("%d", i)
		}
	}()

	// 没有select时无法同时处理两个没关闭的管道数据,只能顺序处理

	// for i := 0; i < 10 ; i++ {
	// 	fmt.Println(<-chan1)
	// 	fmt.Println(<-chan2)
	// }

	var s string
	main:
	for {
		select {
		case s = <-chan1:
			fmt.Println(s)
		case s = <-chan2:
			fmt.Println(s)
		case <-time.After(time.Second):
			fmt.Println("time out")
			break main
		}
	}
}