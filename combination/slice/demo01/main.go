package main

import "fmt"

// 切片底层指向数组，容量不够时会指向新数组并复制内容
// 切片为引用类型
// copy 方法可以复制值

func printSlice(s []int, tag string) {
	fmt.Printf("%s: value = %v, len = %d,  容量 = %d \n", tag, s, len(s), cap(s))
}

func main() {

	//cap := 4
	cap := 3
	slice1 := make([]int, 3, cap)
	slice2 := slice1
	slice2[0] = 1
	printSlice(slice1, "slice1")
	printSlice(slice2, "slice2")

	slice3 := append(slice1, 4) // 如果 cap 不够，会创建新的数组（一般为原来2倍大小）
	slice2[0] = 2
	printSlice(slice1, "slice1")
	printSlice(slice2, "slice2")
	printSlice(slice3, "slice3")

	slice11 := []int{3, 3, 3}
	copy(slice11, slice1)
	printSlice(slice11, "slice11")
}
