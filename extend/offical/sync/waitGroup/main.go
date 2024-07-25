package main

import (
	"fmt"
	"sync"
)

func main() {

	// 10个协程打印1-10，让其顺序显示
	c := make(chan int, 1)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		c <- i
		go func() {
			fmt.Println(<-c)
			wg.Done()
		}()
	}
	wg.Wait()

	// 不带缓冲通道必须同时写入和读取通道，不然会报错
	c = make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			c <- i
		}(i)
		//c<-i // 错误，写入时还没有其他协程等待取出
		go func() {
			fmt.Println(<-c)
			wg.Done()
		}()
	}
	wg.Wait()
}
