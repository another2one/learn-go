package main

import "learn-go/combination/slice/utils"

// 去重 利用map[string]struct{} 可实现无序set  reflect
// reflect.DeepEqual(a, b)

func main() {

	s1 := []int{2: 1}
	utils.PrintSlice(s1, "s1")
}
