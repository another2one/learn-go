package main

// https://zhuanlan.zhihu.com/p/101623663
// 这个更完全一些
import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	Handler func(v ...interface{}) error
	Params  []interface{}
}

type Pool struct {
	capacity       uint64
	runningWorkers uint64
	totalWorkers   uint64
	status         int64 // 协程池状态
	maxWaitSecond  int64
	chTask         chan *Task
	sync.Mutex
}

const (
	RUNNING = 1
	STOPED  = 0
)

var (
	workNum = 0
)

func NewTask(handler func(v ...interface{}) error, params []interface{}) *Task {
	return &Task{
		Handler: handler,
		Params:  params,
	}
}

func NewPool(num uint64) *Pool {
	if num <= 0 {
		panic("workernum error: should > 0")
	}
	return &Pool{
		capacity:      num,
		status:        RUNNING,
		chTask:        make(chan *Task, num),
		maxWaitSecond: 5,
	}
}

func (p *Pool) SetMaxWaitSecond(senconds int64) {
	p.maxWaitSecond = senconds
}

// 开启协程池处理任务
func (p *Pool) run(workNum int) {
	atomic.AddUint64(&p.totalWorkers, 1)
	defer p.decrTotalNum()
	defer fmt.Printf("%d 号协程销毁 \n", workNum)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%d 号协程 panic: %s\n", workNum, r)
			p.decrRunningNum()
		}
	}()
	for {
		select {
		case w, ok := <-p.chTask:
			if !ok {
				return
			}
			atomic.AddUint64(&p.runningWorkers, 1)
			w.Handler(w.Params...)
			log.Printf("%d 号协程执行完毕 \n", workNum)
			p.decrRunningNum()
		case <-time.After(time.Duration(p.maxWaitSecond) * time.Second):
			// 销毁
			return
		}
	}
}

func (p *Pool) decrTotalNum() {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	atomic.SwapUint64(&p.totalWorkers, p.totalWorkers-1)
}

func (p *Pool) decrRunningNum() {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	atomic.SwapUint64(&p.runningWorkers, p.runningWorkers-1)
}

func (p *Pool) Put(t *Task) {
	if p.status == STOPED {
		return
	}
	// 检查没有超过容量以及没有空闲协程的话就新建
	if atomic.LoadUint64(&p.totalWorkers) < p.capacity && atomic.LoadUint64(&p.totalWorkers) <= atomic.LoadUint64(&p.runningWorkers) {
		go p.run(workNum)
		workNum++
	}
	p.chTask <- t
}

func (p *Pool) Close() {
	p.status = STOPED

	fmt.Println("协程数量", atomic.LoadUint64(&p.totalWorkers))
	// 等待所有任务执行完成再关闭
	for atomic.LoadUint64(&p.runningWorkers) > 0 {
		time.Sleep(time.Millisecond * 1)
	}
	// 等待所有协程关闭
	// for atomic.LoadUint64(&p.totalWorkers) > 0 {
	// 	fmt.Println(" -- 协程数量", atomic.LoadUint64(&p.totalWorkers))
	// 	time.Sleep(time.Millisecond * 1000)
	// }
	fmt.Println("全部执行完成了")
	// 执行完成
	close(p.chTask)
}

//主函数
func main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("消耗时间 %.3f \n", float64(time.Since(timeStart).Milliseconds())/1000)
	}()
	// 创建一个协程池
	p := NewPool(10)
	defer p.Close()
	for i := 0; i < 25; i++ {
		t := NewTask(func(v ...interface{}) error {
			switch v[0].(type) {
			case int:
				if v[0] == 3 {
					panic("not like 3")
				}
				fmt.Printf("%d 号任务执行完成 。。。 \n", v[0])
			}
			return nil
		}, []interface{}{i})
		p.Put(t)
	}
}
