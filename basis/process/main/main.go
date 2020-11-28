package main

import(
	"fmt"
)

func main() {

	var age int

	// 单分支
	age = 101
	if age > 100 {
		fmt.Println("age > 100")
	}

	if age = 102; age > 100 {
		fmt.Println("age > 100")
	}

	// 双分支
	if age > 10 {
		fmt.Println("age > 10")
	} else {
		fmt.Println("age <= 10")
	}

	// 多分支
	if age > 10 {
		fmt.Println("age > 10")
	} else if age == 10 {
		fmt.Println("age = 10")
	} else {
		fmt.Println("age < 10")
	}

	var i1 int = 2020
	// var i2 int32 = 956
	
	if (i1 %4  == 0 && i1 % 100 > 0) || i1 % 400 == 0 {
		fmt.Printf("%d年是闰年", i1)
	}

}