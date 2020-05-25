package main

import (
	"fmt"
)

type Student struct {
	Name  string
	score int
}

type Animal struct {
	Name  string
	Wight float64
}

type Person struct {
	Animal
	Name string
}

type En struct {
	enscore int
}

func (p *Student) testing() {
	fmt.Println("I am testing ... ", p.Name)
}

func (p *Univercity) testing() {
	fmt.Println("I am Univercity ... ", p.Name)
}

type Univercity struct {
	Student    // 继承
	*Person    // 继承
	en      En // 组合
	Name    string
	Id      int
}

func main() {
	// 封装：隐藏细节，提供接口供外界调用，可以保证安全性同时提升操作简便
	// 继承：提高代码复用性，管理性
	college1 := Univercity{
		Student{},
		&Person{Name: "lizhi"},
		En{66},
		"lizhi",
		66,
	}
	college1.Name = "Univercity"
	college1.Student.Name = "Student"
	fmt.Println(college1.Person.Name)
	college1.testing()
	college1.en.enscore = 600
	college1.Student.testing()
	fmt.Printf("%+v \n", college1.Wight)
}
