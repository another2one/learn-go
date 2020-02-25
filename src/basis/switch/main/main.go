package main

import(
	"fmt"
)

func main() {
	var i int
	fmt.Println("请输入星期：")
	fmt.Scanln(&i)
	switch i {
		case 1, 2, 3, 4, 5, 6, 7:
			fmt.Printf("星期 %d \n", i)
		default :
			fmt.Println("星期格式错误")
	}
	
	// type Integer interface {}
	// var integer Integer = i
	// 简写
	var integer interface {} = i

	// type用法
	switch integer.(type) {
		case int :
			fmt.Println("int")
		case float64 :
			fmt.Println("float")
		default :
			fmt.Println("not found")
	}
}