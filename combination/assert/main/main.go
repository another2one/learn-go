package main

import (
	"fmt"
)

// 由于接口没有具体类型，所以赋值时需要类型转换

type Point struct {
	x, y int
}
func (this Point) Shit (num int) {
	fmt.Println("吃了", num , "坨屎")
}

type Test struct {}

func (this Test) Shit (num int) {
	fmt.Println("拉了", num , "坨屎")
}

type Inter interface {
	Shit(int)
}

func main() {
	var a Inter
	var b Point
	var c = Test {}
	b = Point{3, 4}
	a = &b // a指向b的地址
	a.Shit(3)

	a = c // a指向c
	a.Shit(3)
	c, ok := a.(Test)
	
	// b = a // cannot use a (type Inter) as type Point in assignment: need type assertion
	b, ok = a.(Point) // 类型断言 判断a是否指向Point类型变量，是则转换为Point类型的变量，否则报错
	if !ok {
		fmt.Println("转换失误")
	}
	fmt.Printf("%+v \n", b)

	// inters := make([]Inter, 2, 2)
	// inters[0] = Point{1, 2}
	// inters[1] = Test{}
	// for _, value := range inters {
	// 	a = value
	// 	if _, ok := a.(Point); ok{
	// 		fmt.Println("I am Point")
	// 	}else if value, ok := a.(Test); ok {
	// 		value.Shit(6)
	// 	}else {
	// 		fmt.Println("error")
	// 	}
	// }
}