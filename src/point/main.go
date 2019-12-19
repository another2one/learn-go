package main

import (
	"fmt"
)

func main()  {
	
	var i int = 65

	// 指针定义 取地址 &
	var p *int = &i
	// var p *int  // 空指针值为nil
	print(p)
	// 指针取值 *
	print(*p)

	*p = 66
	print(i)

	var _ int = 64
}


func print(val interface{}) {
	fmt.Printf("val's type is %T, val = %v \n", val, val)
}