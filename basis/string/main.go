package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
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
	// 取第3个字符串
	fmt.Printf("%q s1[2] \n", s1[2]) // error
	fmt.Printf("%q \n", []rune(s1)[2])
	// 长度
	fmt.Printf("len(s1) = %d \n", len(s1))                                       // 6
	fmt.Printf("len(s1) = %d \n", len([]rune(s1)))                               // 3
	fmt.Printf("utf8.RuneCountInString(s1) = %d \n", utf8.RuneCountInString(s1)) // 3

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

	fmt.Printf("string(65) = %q \n", string(65))       // "A"
	fmt.Printf("string(65) = %q \n", strconv.Itoa(65)) // 65

	// 特殊表示
	fmt.Println("\xe6\x96\xb0世界" == "新世界")
	fmt.Println('\u65b0' == '新')
	// unicode转utf-8
	var ss rune = '\xfe'
	var p = make([]byte, 4)
	size := utf8.EncodeRune(p, ss)
	fmt.Println(p[:size])

	// 字符串反转
	originStr := "你好abc啊哈哈a"
	orune := []rune(originStr)
	for i, j := 0, len(orune)-1; i < j; i, j = i+1, j-1 {
		orune[i], orune[j] = orune[j], orune[i]
	}
	fmt.Println(string(orune))

}
