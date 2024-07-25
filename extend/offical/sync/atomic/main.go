package main

import (
	"fmt"
	"sync/atomic"
)

// 可操作的数据类型: int32、int64、uint32、uint64、uintptr
// 可以做的原子操作有: 加法（add）、比较并交换（compare and swap，简称 CAS）、加载（load）、存储（store）和交换（swap）

func main() {

	// 10个协程打印1-10，让其顺序显示
	var count uint32
	trigger := func(i uint32, fn func()) {
		// 确保操作原子性 读取数据和加一操作
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})

	// CAS CompareAndSwap 实现自旋锁，不会休眠线程(即进入内核态线程), 线程一直处于用户态避免了状态切换的资源和时间浪费
	// 适合简单的操作，等待时间不长的
	for {
		if atomic.CompareAndSwapUint32(&count, 10, 2) {
			// do something ....
			break
		}
	}
}
