// 暂时测试
package main

import (
	"fmt"
	"time"
)

func main() {

	ch := 'b'
	fmt.Println(ch / 2.0)

	//s1 := []int{1, 2, 4}
	s1 := make(map[int]int, 6)
	s1[1] = 1
	printAddress(s1)
	for i, v := range s1 {
		printAddress(s1)
		s1[i+1] = v + 1
	}
	fmt.Println(s1)

	for _, v := range []int{1, 2, 4} {
		fmt.Printf("v address is %v \n", &v)
		go func(v int) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(time.Microsecond * 100)
}

func printAddress(v interface{}) {
	fmt.Printf("v address is %p \n", &v)
}
