package main

import (
	"fmt"
	"learn-go/combination/slice/utils"
	"math"
)

func editSlice(s []int) {
	s[0] = 1
	s = append(s, 1)
}

func editSlicePoint(s *[]int) {
	*s = append(*s, 1)
}

func editAll(a func(s []int), s ...int) (res []int) {
	res = s
	if cap(s) > 0 {
		a(s)
	}
	return
}

func edit1(a int, _ string) bool {
	return a > 10
}

// 1. 函数都是值传递，slice， map, chan为引用类型,也是作为值传递
// 2. 支持函数作为参数和返回值
// 3. 不支持参数默认值
// 4. 支持多参数(...)，必须传一个
// 5. 返回参数支持命名, return时自动返回（return可不带参数但必须return）
// 6. 参数支持_，但是不能再使用
// 7. 没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数签名 math.Sin
func main() {
	// 1
	s1 := make([]int, 1)
	utils.PrintSlice(s1, "s1 before:")
	editSlice(s1)
	utils.PrintSlice(s1, "s1 after:")

	s2 := make([]int, 1)
	utils.PrintSlice(s2, "s2 before:")
	editSlicePoint(&s2)
	utils.PrintSlice(s2, "s2 after:")

	// 2 4 5
	_ = editAll(editSlice, 1)

	// 6
	fmt.Println(edit1(10, "20"))

	// 7
	math.Sin(60)
}
