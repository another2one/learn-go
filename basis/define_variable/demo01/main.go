package main

import (
	"fmt"
	"strconv"
)

var m int64

func main() {
	// 注意 此处不会修改外层m， main作用域内会重新生成 m
	// 这样才会修改
	// var err error
	// m, err = strconv.ParseInt("66", 10, 64)
	m, err := strconv.ParseInt("66", 10, 64)
	if err != nil {
		fmt.Println("parse error:", err)
	}
	fmt.Printf("main m = %d \n", m)

	// 这里也是 m 作用域在 if 里面
	if m, err1 := strconv.ParseInt("611", 10, 64); err1 != nil {
		fmt.Println("parse error:", err1)
		fmt.Printf("main m = %d \n", m)
	}
	fmt.Printf("main m = %d \n", m)

	a()
}

func a() {
	fmt.Printf("func m = %d \n", m)
}
