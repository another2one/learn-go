package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	ch := make(chan int)
	go func() {
		for i := 1; i < 27; i++ {
			ch <- 1
			fmt.Println(i)
			ch <- 1
		}
		wg.Done()
	}()
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			<-ch
			fmt.Printf("%c \n", i)
			<-ch
		}
		wg.Done()
	}()
	wg.Wait()
}
