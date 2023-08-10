package main

import "learn-go/combination/slice/utils"

// go 切片作为函数参数
func main() {
	s1 := []int{0}
	utils.PrintSlice(s1, "start s1")
	s2 := f1(s1)
	utils.PrintSlice(s1, "after f1 s1")
	utils.PrintSlice(s2, "after f1 s2")
	s3 := f2(s1)
	utils.PrintSlice(s1, "after f2 s1")
	utils.PrintSlice(s3, "after f3 s3")
	s4 := f3(s1)
	utils.PrintSlice(s1, "after f3 s1")
	utils.PrintSlice(s4, "after f3 s4")
}

func f1(s1 []int) []int {
	if len(s1) == 0 {
		return s1
	}
	s1[0] = 9
	return s1
}

func f2(s1 []int) []int {
	if len(s1) == 0 {
		return s1
	}
	s1 = append(s1, 9)
	return s1
}

func f3(s1 []int) (s3 []int) {
	if len(s1) == 0 {
		return s1
	}
	s3 = s1
	return
}
