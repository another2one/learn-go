package main

import "fmt"

// TODO: 汇编查看——到底是啥
func main() {
	s, q := a()
	fmt.Println(s, q)

	ss, _ := a()
	fmt.Println(ss)

	fmt.Println(a())
}

// ??? _难道只是一个特殊变量
func a() (int, _ int) {
	_ = 66
	return 1, 66
}

// ???
func a1(i, _ int) int {
	_ = 66
	return a
}
