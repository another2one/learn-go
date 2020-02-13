package main

import (
	"fmt"
	"test/test"
	"strings"
)

func init(){
	fmt.Println("init")
}

func swap (s1, s2 *int){
	*s1, *s2 = *s2, *s1
}

func makeSuffix(suffix string) func(string)string {
	return func (str string) string {
		if !strings.HasSuffix(str, suffix) {
			str += suffix
		}
		return str
	}
}

func init(){
	defer fmt.Println("ok3")
}

func main() {
	ss = 22
	n1 := 1
	defer fmt.Println(n1)
	s1, s2 := 1, 2
	fmt.Println(s1, s2)
	swap(&s1, &s2)
	fmt.Println(s1, s2)
	test.Test(565656)
	n1 = 8
	defer fmt.Println(n1)
	formatSuffix := makeSuffix(".jpg")
	fmt.Printf("type is %T \n", formatSuffix)
	fmt.Println(formatSuffix("66"))
	fmt.Println(formatSuffix("66.jpg"))
}

