package main

import (
	"fmt"
)

// 外部变量定义
var (
	n1   = 1
	name = "lizhi" + "66"
)

func main() {

	fmt.Println("\n n1 = ", n1, "\n name = ", name)

	// -------------- 单个变量定义

	// 第一种
	var n1 float32 = 66 // 优先使用内部定义
	// 第二种
	var n2 = 32
	// 第三种
	n3 := 66

	fmt.Println("\n n1 = ", n1, "\n n2 = ", n2, "\n n3 = ", n3)

	// -------------- 多个变量定义

	// 第一种
	var s1, s2 int = 66, 77 // 优先使用内部定义
	// 第二种
	var s3, s4 = 88, 99
	// 第三种
	s5, s6, _ := 33, 44, 55

	fmt.Println("\n s1 = ", s1, "\n s2 = ", s2, "\n s3 = ", s3, "\n s4 = ", s4, "\n s5 = ", s5, "\n s6 = ", s6)
}
