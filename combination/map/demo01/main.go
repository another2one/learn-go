package main

// 注意：
// 1. map 的元素不可取址 因为可能发生扩容改变元素地址

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {

	m1 := map[string]Point{"foo": Point{1, 2}}
	// m1["foo"].x = 4 // 错误 无法地址偏移取值
	fmt.Println("m1[\"foo\"].x = ", m1["foo"].x)
	// 解决方法1：使用结构体指针
	m2 := map[string]*Point{"foo": &Point{1, 2}}
	m2["foo"].x = 4
	fmt.Println("m2[\"foo\"].x = ", m2["foo"].x)
	// 解决办法2：使用临时变量
	m3 := map[string]Point{"foo": Point{1, 2}}
	temp := m3["foo"]
	temp.x = 4
	m3["foo"] = temp
	fmt.Println("m3[\"foo\"].x = ", m3["foo"].x)

	// slice 可以
	s1 := make([]Point, 10)
	point := Point{4, 5}
	s1[0] = point
	s1[0].x = 6
	fmt.Println("s1[0].x = ", s1[0].x)
	fmt.Println("point = ", point)
	fmt.Printf("s1[0] type is %T, value is %+v \n", s1[0], s1[0])

	// map比较
	fmt.Printf("%t \n", equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv { // 直接y[k]会报错
			return false
		}
	}
	return true
}
