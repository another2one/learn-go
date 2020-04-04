package main

import (
	"fmt"
	"time"
)

var (
	timeOutChan = make(chan bool, 1)
)

func main() {

	// 延时器
	fmt.Println("start")
	start := time.Now()
	go afterTimeExec(2*time.Second, func() {
		fmt.Println("exec")
	})
	<-timeOutChan

	fmt.Printf("use time: %ds \n", time.Now().Second()-start.Second())
	// 官方自带的
	select {
	case <-timeOutChan:
		fmt.Println("time chan write ok !!!")
	case <-time.After(1 * time.Second): // 超时
		fmt.Println("time out")

	}

	// 定时器 + 清除
	inter1 := make(chan bool, 1)
	go intervalTimeExec(1*time.Second, inter1, func() {
		fmt.Println("interval")
	})
	time.Sleep(6 * time.Second)
	clearInterval(inter1)
}

func afterTimeExec(after time.Duration, f func()) {
	time.Sleep(after)
	f()
	timeOutChan <- true
}

func intervalTimeExec(intervalTime time.Duration, interval chan bool, f func()) {
loop:
	for {
		select {
		case <-interval:
			break loop
		default:
			time.Sleep(intervalTime)
			f()
		}
	}
}

func clearInterval(interval chan bool) {
	interval <- true
}
