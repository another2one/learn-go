package main

import (
	"fmt"
	"time"
)

// defer 执行顺序
//  panic后程序将会立马停止 当前协程defer执行完并退出
//  当前协程指的是最初开启的协程 协程会开新栈， defer是沿着栈执行的，所以recover也只能恢复当前协程的
func main() {
	defer func() {
		fmt.Println("in main")
	}()

	go func() {
		defer func() {
			fmt.Println("in goroutine")
		}()
		// 1
		// in goroutine
		// panic 666
		// panic(666)

		// 2
		// p2
		// p1
		// in goroutine
		// panic 666
		// p1()

		// 3
		// p2
		// p1
		// panic 666
		// go p1()

		time.Sleep(time.Second * 1)
	}()

	// 4
	// p2
	// p1
	// in mian
	// panic 666
	p1()

	time.Sleep(time.Second * 1)
}

func p1() {
	defer func() {
		fmt.Println("in p1")
	}()
	p2()
}

func p2() {
	defer func() {
		fmt.Println("in p2")
	}()
	panic(666)
}
