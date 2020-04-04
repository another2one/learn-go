package main

import "fmt"

// range 遍历的是值拷贝的副本

func main() {

	// 1 slice 底层对slice做一个拷贝，记录此时长度，然后for循环到其长度
	//s := []int{5, 6, 7}
	s := make([]int, 3, 8)
	s[0] = 5
	s[1] = 6
	s[2] = 7
	for i, v := range s {
		if i == 0 {
			s[0] = 0
			s[1] = 1
			s[2] = 3
		}
		fmt.Printf("v = %d \n", v)
		s[i] = v + 100 // v变了
	}
	fmt.Printf("s = %v \n", s) // 105 101 103
	for _, v := range s {
		s = append(s, v)
	}
	fmt.Printf("s = %v \n", s) // 105 101 103 105 101 103 s为值副本

	// 2 array
	a := [...]int{5, 6, 7}
	for i, v := range a {
		if i == 0 {
			a[0] = 0
			a[1] = 1
			a[2] = 3
		}
		fmt.Printf("v = %d \n", v)
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
		fmt.Printf("i, v = %d, %d \n", i, v)
		m[i] = v + 100 // v 会变
	}
	fmt.Printf("m = %v \n", m)  // 结果不一定，因为map是无序的，如果按顺序就是 105 101 3
	for j := 0; j < 1000; j++ { // TODO: 测试的重要性，只要多次不同环境运行才能发现很多问题
		m = map[int]int{
			2: 7,
			0: 5,
			1: 6,
		}
		for i, v := range m {
			//fmt.Printf("i, v = %d, %d \n", i, v)
			m[i+1] = v + 100
		}
		fmt.Printf("m = %v \n", m) // 可以正常退出 原因 map 无序 实际结果也不定
	}

}
