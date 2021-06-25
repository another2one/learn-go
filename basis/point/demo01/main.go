package main

// go语言中哪些类型的值不可以被取地址
//  不可变的，临时结果和不安全的
func main() {
	// 常量不可寻址
	// const con = 123
	// _ = &con

	// 字符串的字节元素
	// s := "str"
	// _ = &s[0]

	// 映射元素
	// m := map[int]string{1: "lizhi"}
	// _ = &m[1]

	// 基本类型值的字面量不可寻址
	// _ = &(12)
	// _ = &(123 + 456)

	// 字面量数组元素不可寻址 变量可以
	// _ = &([...]int{123}[0]) // 不可寻址
	_ = &([]int{123}[0]) // ok
	ar := [...]int{2, 3, 4}
	_ = &(ar[2])

	// 接口值的动态值
	type stu struct{ name string }
	stru := stu{"lizhi"}
	// _ = &(interface{})(stru)
	ints := interface{}(stru)
	_ = &ints

	// 结构体字面量成员
	// _ = &(struct{ name string }{"lizhi"}.name)
	_ = &(struct{ name string }{"lizhi"}) // 可以

	// 声明的包级别函数
	// 方法(用做函数值)
	// 表达式中间结果值
	// a, b := 1, 1
	// _ = &(a + b)

	// 数据通道接收操作
	// c1 := make(chan int, 1)
	// c1 <- 1
	// _ = &(<-c1)

	// 子字符串操作

	// 子切片操作
	// _ = &([]int{1, 2, 3}[1:2])

	// 加法、减法、乘法、以及除法等等。
	// 函数调用
	// 显式值转换
	// 各种操作，不包含指针解引用(dereference)操作，但是包含数据通道接收操作、子字符串操作、子切片操作，以及加法/减法/乘法/除法等等
}