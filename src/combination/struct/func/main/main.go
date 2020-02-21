package main

import (
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age int
	Sex bool
}

// 以方法中接收的参数为主，方法中为值，即使传递指针也会拷贝为值，指针则编译时自动改为指针调用
// 注意： 函数不会自动转化参数
func (s Student) test() {
	fmt.Println(s.Name)
	s.Name = "kk"
}

func test1(test *Student){
	fmt.Println(test)
}

type Integer int64

// 可以为
func (i Integer) test(){
	fmt.Println(i)
}

func (i Integer) String() string{
	str := fmt.Sprintf("%d", i)
	return str
}


func main() {

	var test2 Student
	test2.Name = "lizhi"
	(&test2).test()
	test1(&test2)
	fmt.Printf("%+v \n", test2)
	var i Integer
	i = 65
	i.test()
}