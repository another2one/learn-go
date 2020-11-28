package main

import (
	"fmt"
	"time"
)

func main() {

	//s1 := []int{1, 2, 4}
	s1 := make(map[int]int)
	s1[1] = 1
	for i, v := range s1 {
		s1[i+1] = v + 1
	}
	fmt.Println(s1)

	for _, v := range s1 {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second * 1)
}
