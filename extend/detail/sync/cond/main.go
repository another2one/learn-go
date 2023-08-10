package main

import (
	"fmt"
	"sync"
)

func main() {

	// 10个协程打印1-10，让其顺序显示
	var now = 0
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	for i := 0; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			for now != i {
				sendCond.Wait() // 解锁 - 阻塞并将当前程序放入队列等待信号 - 锁住并继续执行
			}
			fmt.Println(now)
			now++
			lock.Unlock()
			sendCond.Broadcast() // Signal只通知一个
		}(i)
	}
	lock.Lock()
	for now != 10 {
		fmt.Println("wait 10")
		sendCond.Wait()
	}
	lock.Unlock()
}
