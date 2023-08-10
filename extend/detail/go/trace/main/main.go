package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			time.Sleep(50 * time.Millisecond)
			wg.Done()
		}(i)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	wg.Wait()

	//main
	fmt.Println("Hello World")
}
