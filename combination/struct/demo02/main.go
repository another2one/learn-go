package main

import "fmt"

type Student struct {
	name string
}

type Student1 struct {
	name string
}

// struct比较
func main() {

	// 1 同一 struct 不同实例
	s := Student{"lizhi"}
	s1 := Student{"lizhi"}
	fmt.Printf("同一 struct 不同实例: %t \n", s1 == s)
	fmt.Printf("同一 struct 不同实例: %t \n", &s1 == &s)
	fmt.Printf("同一 struct 不同实例: %t \n", Student{"lizhi"} == Student{"lizhi"})
	fmt.Printf("同一 struct 不同实例: %t \n", &Student{"lizhi"} == &Student{"lizhi"})

	// 2 不同 struct
	// s2 := Student{"lizhi"}
	// s3 := Student1{"lizhi"}
	// fmt.Printf("不同 struct 不同实例: %t \n", s2 == s3)
	// fmt.Printf("不同 struct 不同实例: %t \n", &s2 == &s3)
	// fmt.Printf("不同 struct 不同实例: %t \n", Student{"lizhi"} == Student1{"lizhi"})
	// fmt.Printf("不同 struct 不同实例: %t \n", &Student{"lizhi"} == &Student1{"lizhi"})

	// 3 不同 struct 实例
	// fmt.Printf("同一 struct 不同实例: %t \n", s2 == s3)
}
