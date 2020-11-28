package utils

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

var arrayLock sync.Mutex

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

func ArrayQueueRun(windowNum, personNum int) {

	aq := ArrayQueue{
		Push:    0,
		Pop:     0,
		MaxSize: 10,
		Slice:   make([]int, 10),
	}
	exitChan := make(chan bool)

	// 排队
	go func() {
		i := 1
		for i < personNum {
			aq.add(i)
			i++
			time.Sleep(time.Millisecond * 500)
		}
		exitChan <- true
	}()

	// 处理队列
	for i := 1; i <= windowNum; i++ {
		go func(i int) {
			for {
				arrayLock.Lock()
				res, err := aq.get()
				arrayLock.Unlock()
				if err == nil {
					fmt.Printf("%d 号窗口取出了 %d \n", i, res)
					time.Sleep(time.Second * 1)
				}
			}
		}(i)
	}
	<-exitChan
}
