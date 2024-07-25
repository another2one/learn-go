package main

import (
	"math/rand"
	"sync"
	"time"
)

// 多个发送者一个接收者如何优雅退出协程
func main() {
	rand.Seed(time.Now().Unix())
	maxDataNum := 1000000
	numSenders := 1000

	wg := sync.WaitGroup{}
	wg.Add(1)

	dataCh := make(chan int)

	go func() {
		defer wg.Done()
		for data := range dataCh {
			if data == numSenders {
				close(dataCh)
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// 多个发送者
	for numSenders < 0 {
		go func(workNum int) {
			dataCh <- rand.Intn(maxDataNum)
		}(numSenders)
		numSenders--
	}

	go wg.Wait()
}
