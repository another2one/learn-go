package main

import (
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int
	Sex  bool
}

// 以方法中接收的参数为主，方法中为值，即使传递指针也会拷贝为值，指针则编译时自动改为指针调用
// 注意： 函数不会自动转化参数
func (s Student) test() {
	fmt.Println(s.Name)
	s.Name = "kk"
}

func test1(test *Student) {
	fmt.Println(test)
}

type Integer int64

// 可以为
func (i Integer) test() {
	fmt.Println(i)
}

func (i Integer) String() string {
	str := fmt.Sprintf("%d", i)
	return str
}

// TODO: ...参数 切片会改变原来数据，多参数只是传值
func joinString(v ...string) string {
	str := ""
	for _, s := range v {
		str += s
	}
	return str
}

func joinSliceString(v []string) {
	for i, s := range v {
		v[i] = s + s
	}
}

type strFunc func(v ...string) string

func strWalk(ss []string, f strFunc) []string {
	for i, v := range ss {
		ss[i] = f(v)
	}
	return ss
}

func strAddNumber(v ...string) string {
	return v[0] + fmt.Sprint(len(v[0]))
}

func main() {

	var test2 Student
	test2.Name = "lizhi"
	(&test2).test()
	test1(&test2)
	fmt.Printf("%+v \n", test2)
	var i Integer = 65
	i.test()

	// 多参数
	fmt.Printf("多参数函数测试：%s \n", joinString("我", "是", "谁"))
	// 切片修改值
	ss := []string{"我", "是", "谁"}
	fmt.Printf("切片修改值前：%v \n", ss)
	joinSliceString(ss)
	fmt.Printf("切片修改值后：%v \n", ss)

	// 函数类型及参数
	fmt.Printf("函数类型及参数测试：%s \n", strWalk([]string{"我", "是", "谁"}, strAddNumber))
}
