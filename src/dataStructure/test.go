package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ArrayQueue struct {
	Push    int
	Pop     int
	MaxSize int
	Slice   []int
}

var lock sync.Mutex

var (
	GoRoutingNum = 3
)

func (aq *ArrayQueue) add(i int) error {
	if aq.Pop-aq.Push >= aq.MaxSize-1 {
		return errors.New("队列已满")
	} else {
		aq.Slice[aq.Push%aq.MaxSize] = i
		aq.Push++
	}
	return nil
}

func (aq *ArrayQueue) get() (int, error) {
	if aq.Pop == aq.Push {
		return 0, errors.New("队列为空")
	} else {
		res := aq.Slice[aq.Pop%aq.MaxSize]
		aq.Pop++
		return res, nil
	}

}

type lineQueue struct {
	Num  int
	Next *lineQueue
}

func push(head, new *lineQueue) {
	getLast(head).Next = new
}

func getLast(head *lineQueue) *lineQueue {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	return temp
}

func delete(head, line *lineQueue) error {
	if head == line {
		return errors.New("队列为空")
	}
	temp := head
	for {
		if temp.Next == line {
			break
		}
		temp = temp.Next
	}
	temp.Next = line.Next
	return nil
}

func pop(head *lineQueue) (*lineQueue, error) {
	return getLast(head), delete(head, getLast(head))
}

func main() {

	// 数组队列实现
	fmt.Printf("\n 数组队列 \n")

	aq := ArrayQueue{
		Push:    0,
		Pop:     0,
		MaxSize: 8,
		Slice:   make([]int, 8),
	}
	exitChan := make(chan bool)
	go func() {
		i := 1
		for i < 10 {
			aq.add(i)
			i++
			time.Sleep(time.Second * 1)
		}
		exitChan <- true
	}()
	for i := 1; i < GoRoutingNum+1; i++ {
		go func(i int) {
			for {
				lock.Lock()
				res, err := aq.get()
				lock.Unlock()
				if err == nil {
					fmt.Printf("%d 号协程取出了 %d \n", i, res)
					time.Sleep(time.Second * 1)
				}
			}
		}(i)
	}
	<-exitChan

	// 链表实现
	fmt.Printf("\n 链表实现 \n")

	head := &lineQueue{}

	go func() {
		i := 1
		for i < 10 {
			push(head, &lineQueue{
				Num:  i,
				Next: nil,
			})
			i++
			time.Sleep(time.Second * 1)
		}
		exitChan <- true
	}()
	for i := 1; i < GoRoutingNum+1; i++ {
		go func(i int) {
			for {
				lock.Lock()
				res, err := pop(head)
				lock.Unlock()
				if err == nil {
					fmt.Printf("%d 号协程取出了 %d \n", i, res.Num)
					time.Sleep(time.Second * 1)
				}
			}
		}(i)
	}
	<-exitChan

	// 通道
	fmt.Printf("\n 通道 \n")

	queueChan := make(chan int, 4)
	go func() {
		i := 1
		for i < 10 {
			queueChan <- i
			i++
			time.Sleep(time.Second * 1)
		}
		exitChan <- true
	}()
	for i := 1; i < GoRoutingNum+1; i++ {
		go func(i int) {
			for {
				fmt.Printf("%d 号协程取出了 %d \n", i, <-queueChan)
			}
		}(i)
	}
	<-exitChan
}
