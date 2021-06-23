package main

import (
	"fmt"
)

// range 遍历的值是拷贝的副本

func main() {

	// 1 slice 底层对slice做一个拷贝，记录此时长度，然后for循环到其长度
	// s := []int{5, 6, 7}
	s := make([]int, 3, 6)
	fmt.Printf("s len = %d, cap = %d address = %p \n", len(s), cap(s), s)
	// s = []int{5, 6, 7} // 地址和容量变化
	s[0], s[1], s[2] = 5, 6, 7
	fmt.Printf("s len = %d, cap = %d address = %p \n", len(s), cap(s), s)
	for i, v := range s {
		if i == 0 {
			s[0] = 0
			s[1] = 1
			s[2] = 3
		}
		s[i] = v + 100
	}
	fmt.Printf("s = %v \n", s) // 105 101 103
	s = append(s, s...)
	fmt.Printf("s len = %d, cap = %d address = %p \n", len(s), cap(s), s)
	fmt.Printf("s = %v \n", s) // 105 101 103 105 101 103

	// 2 array
	a := [...]int{5, 6, 7}
	for i, v := range a {
		if i == 0 {
			a[0] = 0
			a[1] = 1
			a[2] = 3
		}
		a[i] = v + 100 // v 不会变
	}
	fmt.Printf("a = %v \n", a) // 105 106 107

	// 3 map hashMap结构体的指针拷贝， 但是map是无序的
	m := map[int]int{
		2: 7,
		0: 5,
		1: 6,
	}
	for i, v := range m {
		if i == 0 {
			m[0] = 0
			m[1] = 1
			m[2] = 3
		}
		m[i] = v + 100 // v 会变
	}
	fmt.Printf("m = %v \n", m) // 结果不一定

	// for j := 0; j < 1000; j++ { // TODO: 测试的重要性，只要多次不同环境运行才能发现很多问题
	// 	m = map[int]int{
	// 		2: 7,
	// 		0: 5,
	// 		1: 6,
	// 	}
	// 	for i, v := range m {
	// 		//fmt.Printf("i, v = %d, %d \n", i, v)
	// 		m[i+1] = v + 100
	// 	}
	// 	fmt.Printf("m = %v \n", m) // 可以正常退出 原因 map 无序 实际结果也不定
	// }

	// 4 string
	data := "你是谁啊???\xfd\xfe\xff" // þ 的unicode为 \xfe
	// 有问题 string range时以rune形式，碰到非正常utf8时会转为0xfffd
	for _, v := range data {
		fmt.Printf("%#x \n", v)
	}
	fmt.Println([]byte(data))
}
