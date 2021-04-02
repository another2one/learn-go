package main

import "learn-go/combination/slice/utils"

// slice可以指定索引 []int{2:1} 结果: [0, 0, 1]

func main() {

	s1 := []int{2: 1}
	utils.PrintSlice(s1, "s1")
}
