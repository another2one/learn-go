package main

import "learn-go/combination/slice/utils"

func main() {

	// 切片为引用类型
	//cap := 4
	cap := 3
	slice1 := make([]int, 3, cap)
	slice2 := slice1
	slice2[0] = 1
	utils.PrintSlice(slice1, "slice1")
	utils.PrintSlice(slice2, "slice2")

	// 切片底层指向数组，容量不够时会指向新数组并复制内容
	slice3 := append(slice1, 4)
	slice2[0] = 2
	utils.PrintSlice(slice1, "slice1")
	utils.PrintSlice(slice2, "slice2")
	utils.PrintSlice(slice3, "slice3")

	// copy 方法可以复制值
	slice11 := []int{3, 3, 3}
	copy(slice11, slice3)
	utils.PrintSlice(slice11, "slice11")
}
