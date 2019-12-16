package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	var c1 byte = 'c';
	var c2 rune = '你';

	fmt.Printf("c1 type is %T, size is %d, c1 type is %T, c2 size is %d \n", c1, unsafe.Sizeof(c1), c2, unsafe.Sizeof(c2));
	fmt.Println("c1 =", c1, "\t c2 =", c2);
	fmt.Printf("c1 = %c \t c2 = %c \n", c1, c2);
}