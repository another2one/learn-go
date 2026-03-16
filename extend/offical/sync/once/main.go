package main

import (
	"fmt"
	"sync"
)

// once 保证程序只允许一遍
var once sync.Once

func main() {

	onceFunc := func() {
		fmt.Println("only once")
	}

	done := make(chan bool)
	for i := range 10 {
		go func(i int) {
			once.Do(onceFunc) // 保证只运行一遍
			done <- true
		}(i)
	}
	once.Do(onceFunc)
	for range 10 {
		<-done
	}
}
