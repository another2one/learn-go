package test

import (
	"fmt"
)

func init(){
	fmt.Println("init test")
}

func Test (n int){
	if n > 2 {
		n--
		Test(n)
	}else{
		fmt.Println("n = ", n)
	}
}