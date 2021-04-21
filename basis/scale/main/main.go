package main

import (
	"fmt"
)

func main() {

	var i int = 15
	var j int = 10

	// %b 二进制输出
	fmt.Printf("i = %b \n", i)
	fmt.Printf("i = %08b \n", i) // 8位
	// %o 8进制输出
	fmt.Printf("i = %o \n", i)
	fmt.Printf("i = %08o \n", i) // 8位
	// %x %x 16进制输出
	fmt.Printf("i = %X \n", i)
	fmt.Printf("i = %x \n", i)
	fmt.Printf("i = %#X \n", i)
	fmt.Printf("i = %#x \n", i)
	fmt.Printf("i = %08X \n", i) // 8位

	// 0开头为8进制
	i = 011
	fmt.Printf("i = %d \n", i)
	// 0x开头为16进制
	i = 0x11
	fmt.Printf("i = %d \n", i)

	// 内存中都是二进制，可以直接加
	fmt.Printf("i + j = %d \n", i+j)
}
