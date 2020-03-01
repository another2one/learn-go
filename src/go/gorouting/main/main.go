package main

import (
	"fmt"
	"time"
	"strconv"
)

// go主线程(线程/进程)：一个go线程可以起很多个协程协程是轻量级线程
// 协程特点：
// 1. 有独立栈空间
// 2. 共享堆空间
// 3. 调度由用户控制
// 4. 协程是轻量级线程

// MPG 		M 操作系统主线程  P 协程执行需要的上下文  G 协程


// 主线程结束后其他线程也会结束

func test () {
	for i := 1; i < 5; i++ {
		fmt.Println(i, ": hello ,world")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	startTime := time.Now()
	go test()
	for i := 1; i < 3; i++ {
		fmt.Println(strconv.Itoa(i) + ": hello ,golang")
		time.Sleep(time.Second * 1)
	}
	endTime := time.Now()
	fmt.Println("process used " , endTime.Sub(startTime))
}