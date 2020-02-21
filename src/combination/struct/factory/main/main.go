package main

import (
	"fmt"
	"combination/struct/factory/student"
)

func main() {
	stu := student.Create("lizhi")
	fmt.Printf("%+v \n", stu.GetName())
}