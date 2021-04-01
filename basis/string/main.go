package main

import (
	"fmt"
)

/*
src/runtime/string.go:stringStruct
type stringStruct struct {
	str unsafe.Pointer // 字符串首地址，只读
	len int // 长度
}
*/

func main() {

	var s1 string = "li李"
	fmt.Printf("s1 len = %d, address = %p \n", len(s1), &s1)
	// 字符串赋值后不能修改
	// 因为string通常指向字符串字面量，而字符串字面量存储位置是只读段，而不是堆或栈上，所以才有了string不可修改的约定
	// s1[1] = 's' // error
	// 修改
	s2 := []rune(s1)
	s2[1] = 'z'
	s2 = append(s2, []rune("源码包src/runtime/string.go:stringStruct定义了string的数据结构")...)
	s1 = string(s2)
	fmt.Printf("s1=%q len = %d, address = %p \n", s1, len(s1), &s1)

	// "" 及 ``
	fmt.Println("my \t name is \" lizhi \"")
	fmt.Println(`my \t name is " lizhi "`)

	// 拼接
	s1 += "a" + // 可以换行， 注意加号在末尾
		"ss"

	fmt.Printf("s1 type is %T, s1 size is %d \n", s1, len(s1))
	// fmt.Println("c1 =", c1, "\t c2 =", c2);
	// fmt.Printf("c1 = %c \t c2 = %c \n", c1, c2);

}
