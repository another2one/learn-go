package main

import (
	"common"
	"fmt"
)

func main() {

	// 整型直接运算必须相同类型(int 和 int64也不行,浮点规则相同),且结果仍为此类型,不足会溢出
	// 整型和浮点运算后结果会变成浮点型
	// ++ -- 为独立语句(单行)

	// 算数运算符
	var i1 int8 = 89
	var i2 int8 = 89
	// i3 := i1 + i2 // 溢出
	i3 := int64(i1) + int64(i2)
	i4 := int64(i1) * int64(i2)
	common.PrintOther(i3)
	common.PrintOther(i4)
	common.PrintOther(10 / float32(4))
	common.PrintOther(10 / 2.5)
	common.PrintOther(10 - 4.5)
	common.PrintOther(10 + 4.5)
	common.PrintOther(10 * 4.5)
	common.PrintOther(-10 % 9)
	i5 := -10
	i6 := 10.1
	i5++ // 独立语言
	common.PrintOther(i5)
	// common.Print(i5++) // 不允许
	// i5 = i5++  // 不允许
	common.PrintOther(i6)
	fmt.Printf("还有%d个星期%d天放假 \n", 97/7, 97%7)
	fmt.Printf("温度为:%v", calc(130))

	var i7 int64 = 89
	// common.Print(i6 > i7) // 两个整型或俩个浮点型,位数不同不能比较
	common.PrintOther(89 > i7)

	a,b := 10,2
	a += b
	b = a - b
	a = a - b


	fmt.Printf("a = %d, b = %d", a, b)
}

func calc(temp float64) float64 {
	return 5.0/9*(temp-100)
}
