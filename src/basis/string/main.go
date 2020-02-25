package main

import (
	"fmt"
)

func main()  {

	var s1 string = "li李"

	// 字符串赋值后不能修改
	// s1[1] = 's'
	// 修改
	s2 := []rune(s1)
	s2[1] = 's'
	s1 = string(s2)
	fmt.Printf("%q \n", s1)

	// "" 及 ``
	fmt.Println("my \t name is \" lizhi \"");
	fmt.Println(`my \t name is " lizhi "`);

	// 拼接
	s1 += "a" + // 可以换行， 当时加号在末尾
		"ss"

	fmt.Printf("s1 type is %T, s1 size is %d \n", s1, len(s1));
	// fmt.Println("c1 =", c1, "\t c2 =", c2);
	// fmt.Printf("c1 = %c \t c2 = %c \n", c1, c2);
}