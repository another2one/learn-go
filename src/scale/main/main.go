package main

import(
	"fmt"
)

func main() {

	var i int = 5
	var j int = 10

	// %b 二进制输出
	fmt.Printf("i = %b \n", i)

	// 0开头为8进制
	i = 011
	fmt.Printf("i = %d \n", i)
	// 0x开头为16进制
	i = 0x11
	fmt.Printf("i = %d \n", i)

	// 内存中都是二进制，可以直接加
	fmt.Printf("i + j = %d \n", i + j)
}