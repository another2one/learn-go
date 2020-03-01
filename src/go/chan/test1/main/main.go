package main

import (
	"fmt"
	"time"
	"sync"
)

const (
	NUMBER_NUM = 1000 // 数字个数
	GOROUTING_NUM = 8 // 协程个数
)

var wg sync.WaitGroup

//  写完关闭，读完通知结束

var count int

func writeData (intChan chan int) {
	for i := 1; i <= NUMBER_NUM; i++ {
		intChan<- i
	}
	close(intChan) // 写完以后关闭管道
}

func addToNum (n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum+= i
	}
	return sum
}


func readData (intChan chan int,resChan chan  map[int]int, exitChan chan bool) {
	defer wg.Done()
	for {
		v, ok := <-intChan // 取完关闭的管道数据后返回false, 并通知结束
		if !ok {
			exitChan<- true
			break
		}else{
			// time.Sleep(time.Millisecond*100)
			resChan<- map[int]int{v:addToNum(v)}
		}
	}
	fmt.Println("readData end ...")
}

func main() {

	startTime := time.Now()
	var intChan chan int
	intChan = make(chan int, 1000)
	resChan := make(chan map[int]int, NUMBER_NUM)
	exitChan := make(chan bool, GOROUTING_NUM)

	go writeData(intChan)

	wg.Add(GOROUTING_NUM)
	for i := 0; i < GOROUTING_NUM; i++ {
		go readData(intChan, resChan, exitChan)
	}

	// 第一种：通过接收每个写入协程的结束信号，关闭结果通道，不会阻塞读取
	// go func () {
	// 	for i := 0; i < GOROUTING_NUM; i++ {
	// 		<-exitChan
	// 	}
	// 	close(resChan)
	// }()	

	// 第二种：通过等待着，阻塞等待所有写入协程组完成，然后操作结果信道
	wg.Wait()
	close(resChan)

	for _ = range resChan {
		// fmt.Printf("%+v \n", v)
	}
	
	endTime := time.Now()
	usedTime := endTime.Sub(startTime)
	fmt.Println("process used " , usedTime)

	// 定时器实现
	 // 声明一个退出用的通道
	 exit := make(chan int)
	 // 打印开始
	 fmt.Println("start")
	 // 过1秒后, 调用匿名函数
	 time.AfterFunc(time.Second, func() {
		 // 1秒后, 打印结果
		 fmt.Println("one second after")
		 // 通知main()的goroutine已经结束
		 exit <- 0
	 })
	 // 等待结束
	 <-exit
}