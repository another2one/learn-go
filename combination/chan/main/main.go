package main

import (
	"fmt"
	"time"
)

// 管道本质是一个线程安全的队列 (先进先出)
// 1. 引用类型
// 2. 必须make分配内存后使用
// 3. 管道内数据有类型
// 4. 存放数据不能超过其容量 deadLock
// 5. 没有协程的情况下，如果管道数据以取完，再取就会报错 deadLock
// 6. 管道关闭后不能再写入，但是可以取出剩余数据
// 7. 遍历没关闭管道时会报错
// 8. 只读只写为管道属性，数据类型本身还是一个
// 9. 通道满以后(或者无缓冲通道)写入会阻塞，如果协程都退出了就好报错

func test() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("666")
	}
}

func main() {

	var intChan chan float64
	intChan = make(chan float64, 4)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	intChan <- 4

	vs, ok := <-intChan // ok 为 true
	fmt.Printf("ok = %v, vs = %v \n", ok, vs)

	var intChan2 <-chan float64 = intChan // 只能读取的单向管道， 引用类型，所以指向同一地址
	fmt.Printf("intChan2 = %v, intChan = %v \n", intChan2, intChan)

	// intChan<- 5  // 管道满了不能再入
	close(intChan) // 关闭管道

	for value := range intChan2 {
		fmt.Printf("value = %v \n", value)
	}

	vv, ok := <-intChan
	if !ok {
		fmt.Printf("数据已取完 ... vv = %v \n", vv)
	}

	vv, ok = <-intChan
	if !ok {
		fmt.Printf("数据已取完 ... vv = %T \n", vv)
	}
	for i := 1; i < 10; i++ {
		<-intChan
	}

	intChan3 := make(chan int, 3)
	go test()

	for i := 0; i < 3; i++ {
		intChan3 <- i // 写满后，等待其他协程读取，没有其他协程则会报错
	}

}
