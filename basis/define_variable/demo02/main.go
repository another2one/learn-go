package main

import "fmt"

func main() {
	a := make([]int, 20)
	fmt.Printf("a = %v \n", a)
	a = []int{17, 18, 19, 20}
	fmt.Printf("a = %v \n", a)
	b := a[15:16] // panic 数组越界 a重新分配了 与之前make没有关系了
	fmt.Printf("b = %v \n", b)
}
