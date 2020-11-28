package utils

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var linkLock sync.Mutex

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

func LinkListRun(windowNum, personNum int) {

	// 链表实现
	exitChan := make(chan bool)
	head := &lineQueue{}

	// 排队
	go func() {
		i := 1
		for i < personNum {
			push(head, &lineQueue{
				Num:  i,
				Next: nil,
			})
			i++
			time.Sleep(time.Second * 1)
		}
		exitChan <- true
	}()

	// 处理队列
	for i := 1; i <= windowNum; i++ {
		go func(i int) {
			for {
				linkLock.Lock()
				res, err := pop(head)
				linkLock.Unlock()
				if err == nil {
					fmt.Printf("%d 号协程取出了 %d \n", i, res.Num)
					time.Sleep(time.Second * 1)
				}
			}
		}(i)
	}
	<-exitChan
}
