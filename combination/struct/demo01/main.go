package main

import "fmt"

type Student struct{}

func (s *Student) look() {
	fmt.Println("look")
}

func (s Student) see(str string) {
	fmt.Println("see ", str)
}

func main() {

	// 对于地址参数的方法，只能通过实体调用
	s := Student{}
	s.look()
	// 等价于
	(&s).look() // 程序内部转化
	//Student.look(&s) // 错误 内部转为了 (&Student).look(&s)
	//Student.look(&Student{}) // 错误 内部转为了 (&Student).look(&Student{})

	// 对于传值调用的方法，可以传入结构体参数调用
	Student.see(Student{}, "666")
	s.see("666")
}
