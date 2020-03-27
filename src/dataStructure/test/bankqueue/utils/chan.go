package utils

import (
	"fmt"
	"time"
)

func ChanRun(windowNum, personNum int) {

	queueChan := make(chan int, 4)
	exitChan := make(chan bool)

	// 排队
	go func() {
		i := 1
		for i < personNum {
			queueChan <- i
			i++
			time.Sleep(time.Second * 1)
		}
		exitChan <- true
	}()

	// 取号
	for i := 1; i <= windowNum; i++ {
		go func(i int) {
			for {
				fmt.Printf("%d 号协程取出了 %d \n", i, <-queueChan)
			}
		}(i)
	}
	<-exitChan
}
