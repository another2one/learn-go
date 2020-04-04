package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

// 1-100 的阶乘 map 集合
// 没有管道时， 需要对并发读取的同一数据加锁， 同时为了防止主线程提前结束，需要加休眠

var (
	factorialMap = make(map[int]*big.Int)
	lock         sync.Mutex // 互斥锁
)

func GetFactorial(i int) {

	product := big.NewInt(1)

	for j := i; j > 1; j-- {
		s := big.NewInt(int64(j))
		product = product.Mul(product, s)
	}

	lock.Lock()
	factorialMap[i] = product
	lock.Unlock()
}

func main() {
	startTime := time.Now()

	for i := 1; i <= 100; i++ {
		go GetFactorial(i)
	}

	endTime := time.Now()

	time.Sleep(time.Second)

	lock.Lock()
	fmt.Printf("%+v \n", factorialMap)
	lock.Unlock()

	fmt.Println("process used ", endTime.Sub(startTime))
}
