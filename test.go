// 暂时测试
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

	for _, v := range []int{1, 2, 4} {
		fmt.Printf("v address is %v \n", &v)
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Microsecond * 100)
}
