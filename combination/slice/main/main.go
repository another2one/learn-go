package main

import (
	"fmt"
	"learn-go/combination/slice/utils"
)

// 本质就是 结构体 struct{len cap *arr} 当容量不够时， 会创建一个新数组并指向他
func delete(s []int, index int) []int {
	length := len(s)
	if index < 0 {
		index = length + index
	}
	if index == 0 {
		return s[1:length]
	} else if index == length-1 {
		return s[0 : length-1]
	} else if index > 0 && index < length-1 {
		return append(s[0:index-1], s[index+1:length]...)
	} else {
		panic("index out of range")
	}
}

func search(s []int, find int) int {
	res := -1
	for index, val := range s {
		if val == find {
			return index
		}
	}
	return res
}

func main() {

	// ******** 定义
	// 1.
	slice1 := make([]int, 2, 5)
	utils.PrintSlice(slice1, "slice1")
	// 2
	slice2 := []int{1, 3, 5}
	utils.PrintSlice(slice2, "slice2")
	// 3
	var slice3 []int
	// slice3 = make([]int, 2, 5)
	// len 0 cap 0 此时无法使用 slice3[0] = 1 会报错：越界
	utils.PrintSlice(slice3, "slice3")
	// 4
	arr4 := [...]int{1, 3, 5, 7, 9}
	slice4 := arr4[:]
	utils.PrintSlice(slice4, "slice4")
	// 1, 3 中go底层会创建并维护数组(make方法)， 2, 4中直接显示指向数组

	// 增
	slice4 = append(slice4, 5)
	utils.PrintSlice(slice4, "slice4 add")

	// 删
	slice4 = delete(slice4, 2)
	utils.PrintSlice(slice4, "slice4 delete")
	slice4 = delete(slice4, 0)
	utils.PrintSlice(slice4, "slice4 delete")

	// 改
	slice4[0] = 9
	utils.PrintSlice(slice4, "slice4 update")

	// 查
	index := search(slice4, 9)
	if index == -1 {
		fmt.Println("not funnd 9 in slice4")
	} else {
		fmt.Println("index = ", index)
	}
}
