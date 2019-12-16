package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	var s string
	var i int
	var f float64
	var b bool

	fmt.Printf("s = %v \t s.size = %d \n", s, unsafe.Sizeof(s));
	fmt.Printf("i = %v \t i.size = %d \n", i, unsafe.Sizeof(i));
	fmt.Printf("f = %v \t f.size = %d \n", f, unsafe.Sizeof(f));
	fmt.Printf("b = %v \t b.size = %d \n", b, unsafe.Sizeof(b));
}