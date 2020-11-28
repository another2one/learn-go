package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main()  {

	var s string = "li"
	var i int = 1
	var f float64 = 1.1
	var b bool = false

	fmt.Printf("s = %v, s = %T, \t s.size = %d \n", s, s, unsafe.Sizeof(s));
	fmt.Printf("i = %v, i = %T, \t i.size = %d \n", i, i, unsafe.Sizeof(i));
	fmt.Printf("f = %v, f = %T, \t f.size = %d \n", f, f, unsafe.Sizeof(f));
	fmt.Printf("b = %v, b = %T \t b.size = %d \n", b, b, unsafe.Sizeof(b));


	// ----------------------  数字间的转换
	fmt.Println("----------------------  数字间的转换  type(i)")

	var i8 int8 = 90
	var i32 int32

	// i32 = i8 + 20 // 不能将int8类型赋给int32
	i32 = int32(i8 + 20)
	i32 = int32(i8) + 20

	i8 = int8(i32) + 127 // 编译能通过，但是结果溢出
	// i8 = int8(i32) + 128   // 编译不能通过
	fmt.Println("i32 = ", i32)
	fmt.Println("i8 = ", i8)

	// -----------------------  数字转字符串
	fmt.Println("----------------------  数字转字符串 fmt.Sprintf strconv.FormatInt")

	var sint int = 500
	var sfloat float64 = 500.5
	var sbool bool = true
	var schar byte = 85

	// 第一种， 使用 fmt.Sprintf 转换
	printStr(fmt.Sprintf("%d", sint))
	printStr(fmt.Sprintf("%.8f", sfloat))
	printStr(fmt.Sprintf("%t", sbool))
	printStr(fmt.Sprintf("%c", schar))

	// 第二种， 使用 strconv 转换
	printStr(strconv.FormatInt(int64(sint), 10)) // 10: 进制
	printStr(strconv.Itoa(sint))
	printStr(strconv.FormatFloat(sfloat, 'f', 8, 64)) // f:格式 8:保留小数位 64:float64
	printStr(strconv.FormatBool(sbool))

	// -----------------------  字符串转数字
	fmt.Println("----------------------  字符串转数字串")
	bs, _ := strconv.ParseBool("true")
	printOther(bs)

	num, _ := strconv.ParseInt("12346", 10, 64)
	printOther(num)

	str2float, _ := strconv.ParseFloat("12346.8967", 64)
	printOther(str2float)
}

func printStr(str interface{}) {
	fmt.Printf("str's type is %T, str = %q \n", str, str)
}

func printOther(other interface{}) {
	fmt.Printf("str's type is %T, str = %v \n", other, other)
}