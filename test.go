package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const N = 1000

func main() {
	// 编程实现：使用add原子操作来并发地递增一个int32值。创建1000个新协程。每个新协程将整数n的值增加1。 原子操作保证这1000个新协程之间不会发生数据竞争，最终程序打印1000
	var sum int64 = 0

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&sum, 1)
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
