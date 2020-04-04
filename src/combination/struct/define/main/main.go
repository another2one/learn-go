package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int
	Sex  bool
}

func main() {

	// 第一种
	stu1 := Student{
		"lizhi",
		20,
		true,
	}
	fmt.Printf("%+v \n", stu1)

	// 第二种
	var stu2 Student
	stu2.Name = "lizhi"
	stu2.Age = 30
	fmt.Printf("%+v \n", stu2)

	// 第三种 注意：返回的是指针
	var stu3 *Student = new(Student)
	(*stu3).Name = "666" // 标准写法
	stu3.Name = "lizhi"  // 简写，底层转为上面写法
	stu3.Age = 30
	fmt.Printf("%+v \n", stu3)

	// 第四种
	var stu4 *Student = &Student{}
	(*stu4).Name = "666" // 标准写法
	stu4.Name = "lizhi"  // 简写，底层转为上面写法
	stu4.Age = 30
	fmt.Printf("%+v \n", stu4)

	// ############################ 注意事项

	// 值类型 默认值拷贝
	stu11 := stu1
	stu1.Name = "66"
	fmt.Printf("%+v \n", stu11)

	// 内存连续
	fmt.Printf("%p, %p, %p \n", &stu11.Name, &stu11.Age, &stu11.Sex)
	// tag 标签
	ss, _ := json.Marshal(stu11)
	fmt.Printf("%+v \n", string(ss))

	// 类型互转 相同字段名和类型的可以转换 但不能直接赋值
	type t1 struct{ Name string }
	type t2 struct{ Name string }
	test1 := t1{"lizhi"}
	var test2 t2
	// test2 = test1 // 错误 类型不同
	test2 = t2(test1)
	fmt.Printf("%+v \n", test2)
}
