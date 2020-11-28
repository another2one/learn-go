package main

import (
	"fmt"
)

func main() {

	var name string
	var age int
	var score float64
	var hasMarried bool

	// 第一种方式 Scanln
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年纪")
	// fmt.Scanln(&age)
	// fmt.Println("请输入分数")
	// fmt.Scanln(&score)
	// fmt.Println("请输入是否结婚")
	// fmt.Scanln(&hasMarried)

	// fmt.Printf("name: %q, age: %d, score: %.1f, hasMarried: %v", name, age, score, hasMarried)

	// 第二种 Scanf
	fmt.Println("请依次输入姓名，年纪，分数，婚否, 使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &hasMarried)
	fmt.Printf("name: %q, age: %d, score: %.1f, hasMarried: %v", name, age, score, hasMarried)
}
